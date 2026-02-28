#!/usr/bin/env python3
"""
Generate coverage documentation for Jamf Pro SDK services.

This script analyzes the codebase to generate comprehensive documentation
showing test coverage, examples, and API documentation for each service.
"""

import os
import re
import subprocess
from pathlib import Path
from datetime import datetime
from typing import Dict, List, Optional
from dataclasses import dataclass


@dataclass
class FunctionInfo:  # pylint: disable=too-many-instance-attributes
    """Information about a service function."""
    name: str
    endpoint: str
    doc_url: Optional[str]
    example_path: Optional[str]
    unit_test_status: str
    acceptance_test_status: str
    is_helper: bool = False
    description: str = ""


@dataclass
class ServiceInfo:
    """Information about a service."""
    name: str
    service_path: str
    api_type: str
    functions: List[FunctionInfo]
    total_functions: int = 0
    unit_tests_passing: int = 0
    acceptance_tests_passing: int = 0
    examples_available: int = 0
    docs_available: int = 0


class CoverageGenerator:
    """Generates coverage documentation for Jamf Pro SDK."""

    def __init__(self, repo_root: Path, jamf_version: Optional[str] = None):
        self.repo_root = repo_root
        self.jamf_version = jamf_version or self.get_jamf_version()
        self.services_base_path = repo_root / "jamfpro" / "services"
        self.examples_base_path = repo_root / "examples"
        self.output_path = repo_root / "docs" / "coverage"
    
    def get_jamf_version(self) -> str:
        env_version = os.environ.get('JAMF_PRO_VERSION')
        if env_version:
            return env_version
        
        print("  Attempting to fetch Jamf Pro version from live instance...")
        try:
            script_dir = Path(__file__).parent
            version_helper = script_dir / "get_jamf_version.go"
            
            result = subprocess.run(
                ["go", "run", str(version_helper)],
                cwd=self.repo_root,
                capture_output=True,
                text=True,
                timeout=15,
                check=False
            )
            
            if result.returncode == 0:
                version = result.stdout.strip()
                if version and version != "unknown":
                    print(f"  ✅ Detected Jamf Pro version: {version}")
                    return version
                print("  ℹ️  Using default version (no live instance available)")
            else:
                print(f"  ℹ️  Using default version (error: {result.stderr.strip()[:100]})")
        except (subprocess.TimeoutExpired, subprocess.SubprocessError, OSError) as e:
            print(f"  ℹ️  Using default version (exception: {str(e)[:100]})")
        
        return "unknown"
        
    def run(self):
        """Main entry point to generate all coverage documentation."""
        print("🔍 Analyzing Jamf Pro SDK services...")
        
        services = self.discover_services()
        print(f"📦 Found {len(services)} services")
        
        print("🧪 Running tests to get status...")
        test_results = self.run_tests()
        
        print("📝 Generating documentation...")
        for service in services:
            self.analyze_service(service, test_results)
            self.generate_service_doc(service)
        
        self.generate_index(services)
        
        print(f"✅ Documentation generated in {self.output_path}")
        
    def discover_services(self) -> List[ServiceInfo]:
        """Discover all services in the codebase."""
        services = []
        
        for api_type in ["jamf_pro_api", "classic_api"]:
            api_path = self.services_base_path / api_type
            if not api_path.exists():
                continue
            
            for service_dir in sorted(api_path.iterdir()):
                if not service_dir.is_dir():
                    continue
                if service_dir.name in ["mocks", "__pycache__"]:
                    continue
                    
                crud_file = service_dir / "crud.go"
                if not crud_file.exists():
                    continue
                
                service = ServiceInfo(
                    name=service_dir.name,
                    service_path=str(service_dir.relative_to(self.repo_root)),
                    api_type=api_type,
                    functions=[]
                )
                services.append(service)
        
        return services
    
    def run_tests(self) -> Dict[str, Dict[str, str]]:
        """Run tests and collect results."""
        results = {
            "unit": {},
            "acceptance": {}
        }
        
        print("  Running unit tests...")
        try:
            cmd = [
                "go", "test", "-v",
                "./jamfpro/services/..."
            ]
            result = subprocess.run(
                cmd,
                cwd=self.repo_root,
                capture_output=True,
                text=True,
                timeout=300,
                check=False
            )
            results["unit"] = self.parse_test_output_verbose(result.stdout)
            print(f"  ✅ Parsed {len(results['unit'])} test results")
        except (subprocess.TimeoutExpired, subprocess.SubprocessError, OSError) as e:
            print(f"  ⚠️  Unit tests failed: {e}")
        
        print("  ⚠️  Acceptance tests require live environment (marked as not_tested)")
        
        return results
    
    def parse_test_output_verbose(self, output: str) -> Dict[str, str]:
        """Parse Go test verbose output."""
        test_status = {}
        
        for line in output.splitlines():
            line = line.strip()
            
            if not (line.startswith("--- PASS:") or line.startswith("--- FAIL:")):
                continue
            
            parts = line.split()
            if len(parts) < 3:
                continue
            
            status = "pass" if "PASS" in parts[1] else "fail"
            test_name = parts[2]
            
            if not test_name.startswith("TestUnit_"):
                continue
            
            remaining = test_name[9:]
            parts = remaining.split("_")
            if len(parts) < 2:
                continue
            
            service = parts[0].lower()
            function = parts[1]
            
            key = f"{service}.{function}"
            if key not in test_status or status == "pass":
                test_status[key] = status
        
        return test_status
    
    def analyze_service(self, service: ServiceInfo, test_results: Dict):
        """Analyze a service to extract function information."""
        crud_file = self.services_base_path / service.api_type / service.name / "crud.go"
        
        with open(crud_file, 'r', encoding='utf-8') as f:
            content = f.read()
        
        functions = self.extract_functions(content)
        
        for func in functions:
            func.example_path = self.find_example(service.api_type, service.name, func.name)
            if func.example_path:
                service.examples_available += 1
            
            test_key = f"{service.name}.{func.name}"
            func.unit_test_status = test_results["unit"].get(test_key, "not_tested")
            func.acceptance_test_status = "not_tested"
            
            if func.unit_test_status == "pass":
                service.unit_tests_passing += 1
            
            if func.doc_url:
                service.docs_available += 1
        
        service.functions = functions
        service.total_functions = len(functions)
    
    def extract_functions(self, content: str) -> List[FunctionInfo]:
        """Extract function information from Go source code."""
        functions = []
        pattern = r'//\s*(.+?)\n(?://.*\n)*func\s+\([^)]+\)\s+(\w+)\s*\([^)]*\)'
        
        for match in re.finditer(pattern, content, re.MULTILINE):
            func_name = match.group(2)
            
            if not func_name[0].isupper():
                continue
            
            comment_block = self._extract_comment_block(content, match)
            
            functions.append(FunctionInfo(
                name=func_name,
                endpoint=self.extract_endpoint(comment_block, func_name),
                doc_url=self.extract_doc_url(comment_block),
                example_path=None,
                unit_test_status="not_tested",
                acceptance_test_status="not_tested",
                is_helper=self.is_helper_function(func_name, comment_block),
                description=self.extract_description(comment_block, func_name)
            ))
        
        return functions
    
    def _extract_comment_block(self, content: str, match: re.Match) -> str:
        """Extract comment block for a function match."""
        func_pos = match.start(2)
        comment_start = content.rfind('//', 0, func_pos - 50)
        if comment_start != -1:
            return content[comment_start:func_pos].strip()
        return match.group(1).strip()
    
    def extract_doc_url(self, comment: str) -> Optional[str]:
        """Extract Jamf Pro API documentation URL from comment."""
        url_pattern = r'https?://developer\.jamf\.com[^\s\)]*'
        match = re.search(url_pattern, comment)
        return match.group(0) if match else None
    
    def extract_endpoint(self, comment: str, func_name: str) -> str:
        """Extract API endpoint from comment or infer from function name."""
        url_match = re.search(r'URL:\s*(\w+\s+/[^\n]+)', comment)
        if url_match:
            return url_match.group(1).strip()
        
        endpoint_match = re.search(r'(GET|POST|PUT|PATCH|DELETE)\s+/api/[^\n\s]+', comment)
        if endpoint_match:
            return endpoint_match.group(0).strip()
        
        if any(x in func_name for x in ["ByName", "Helper"]):
            return "Helper"
        
        return "-"
    
    def is_helper_function(self, func_name: str, comment: str) -> bool:
        """Determine if function is a helper/convenience function."""
        helper_indicators = ["ByName", "convenience", "wrapper", "helper"]
        return any(indicator in func_name or indicator.lower() in comment.lower() 
                   for indicator in helper_indicators)
    
    def extract_description(self, comment: str, func_name: str) -> str:
        """Extract function description from comment."""
        lines = comment.split('\n')
        for line in lines:
            line = line.strip().lstrip('/').strip()
            if line.startswith(func_name):
                desc = line[len(func_name):].strip()
                for prefix in [' returns', ' gets', ' creates', ' updates', ' deletes', ' lists']:
                    if desc.lower().startswith(prefix):
                        desc = desc[len(prefix):].strip()
                        break
                return desc
        return ""
    
    def find_example(self, api_type: str, service_name: str, func_name: str) -> Optional[str]:
        """Find example file for a function."""
        service_examples = self.examples_base_path / api_type / service_name
        if not service_examples.exists():
            return None
        
        func_dir = self.function_name_to_dir(func_name)
        example_file = service_examples / func_dir / "main.go"
        if example_file.exists():
            return str(example_file.relative_to(self.repo_root))
        
        func_base = re.sub(r'V\d+$', '', func_name)
        func_words_list = re.sub('([A-Z][a-z]+)', r' \1', func_base).split()
        func_words = set(w.lower() for w in func_words_list if w)
        
        matches = []
        for subdir in service_examples.iterdir():
            if not subdir.is_dir():
                continue
            
            dir_name = subdir.name
            main_file = subdir / "main.go"
            
            if not main_file.exists():
                continue
            
            dir_words = set(dir_name.split('_'))
            overlap = len(dir_words & func_words)
            threshold = len(dir_words) * 0.7
            
            if dir_words and overlap >= threshold:
                score = overlap / len(dir_words)
                matches.append((score, overlap, dir_name, main_file))
        
        if matches:
            matches.sort(key=lambda x: (x[0], x[1]), reverse=True)
            return str(matches[0][3].relative_to(self.repo_root))
        
        return None
    
    def function_name_to_dir(self, func_name: str) -> str:
        """Convert function name to example directory name."""
        name = re.sub(r'V\d+$', '', func_name)
        
        patterns = [
            (r'^GetByID$', 'get'),
            (r'^GetByName$', 'get_by_name'),
            (r'^UpdateByID$', 'update'),
            (r'^UpdateByName$', 'update_by_name'),
            (r'^DeleteByID$', 'delete'),
            (r'^DeleteByName$', 'delete_by_name'),
            (r'^CreateByID$', 'create'),
            (r'^CreateByName$', 'create_by_name'),
            (r'^List$', 'list'),
            (r'^Create$', 'create'),
            (r'^Update$', 'update'),
            (r'^Delete$', 'delete'),
            (r'^Get$', 'get'),
        ]
        
        for pattern, replacement in patterns:
            if re.match(pattern, name):
                return replacement
        
        name = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', name)
        name = re.sub('([a-z0-9])([A-Z])', r'\1_\2', name)
        return name.lower()
    
    def generate_service_doc(self, service: ServiceInfo):
        """Generate markdown documentation for a service."""
        output_file = self.output_path / "services" / f"{service.name}.md"
        output_file.parent.mkdir(parents=True, exist_ok=True)
        
        api_functions = [f for f in service.functions if not f.is_helper]
        helper_functions = [f for f in service.functions if f.is_helper]
        
        unit_pct = (service.unit_tests_passing / service.total_functions * 100) if service.total_functions > 0 else 0
        acc_pct = (service.acceptance_tests_passing / service.total_functions * 100) if service.total_functions > 0 else 0
        example_pct = (service.examples_available / service.total_functions * 100) if service.total_functions > 0 else 0
        doc_pct = (service.docs_available / service.total_functions * 100) if service.total_functions > 0 else 0
        
        api_label = "Jamf Pro API" if service.api_type == "jamf_pro_api" else "Classic API"
        content = f"""# {service.name.replace('_', ' ').title()} Service

**API Type:** {api_label}  
**Service Path:** `{service.service_path}`  
**Last Updated:** {datetime.now().strftime('%Y-%m-%d')}  
**Tested Against:** Jamf Pro {self.jamf_version}

## 📊 Coverage Summary

| Metric | Count | Percentage |
|--------|-------|------------|
| Total Functions | {service.total_functions} | - |
| With Unit Tests | {service.unit_tests_passing} | {unit_pct:.1f}% |
| With Acceptance Tests | {service.acceptance_tests_passing} | {acc_pct:.1f}% |
| With Examples | {service.examples_available} | {example_pct:.1f}% |
| With Documentation | {service.docs_available} | {doc_pct:.1f}% |

"""
        
        if api_functions:
            content += """## 📚 API Functions

| Function | Endpoint | Docs | Example | Tests |
|----------|----------|------|---------|-------|
"""
            for func in api_functions:
                doc_link = f"[📖]({func.doc_url})" if func.doc_url else "-"
                example_link = f"[💻](../../../{func.example_path})" if func.example_path else "❌"
                
                unit_icon = "✅" if func.unit_test_status == "pass" else "❌" if func.unit_test_status == "fail" else "⚠️"
                acc_icon = "✅" if func.acceptance_test_status == "pass" else "❌" if func.acceptance_test_status == "fail" else "⚠️"
                tests = f"{unit_icon}{acc_icon}"
                
                content += f"| `{func.name}` | `{func.endpoint}` | {doc_link} | {example_link} | {tests} |\n"
        
        if helper_functions:
            content += """
## 🔧 Helper Functions

| Function | Description | Tests |
|----------|-------------|-------|
"""
            for func in helper_functions:
                description = func.description or "Helper function"
                unit_icon = "✅" if func.unit_test_status == "pass" else "❌" if func.unit_test_status == "fail" else "⚠️"
                acc_icon = "✅" if func.acceptance_test_status == "pass" else "❌" if func.acceptance_test_status == "fail" else "⚠️"
                tests = f"{unit_icon}{acc_icon}"
                
                content += f"| `{func.name}` | {description} | {tests} |\n"
        
        content += """
**Tests Legend:** First ✅/❌/⚠️ = Unit Test, Second ✅/❌/⚠️ = Acceptance Test
"""
        
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(content)
    
    def generate_index(self, services: List[ServiceInfo]):
        """Generate index README for all services."""
        self.output_path.mkdir(parents=True, exist_ok=True)
        output_file = self.output_path / "README.md"
        
        total_functions = sum(s.total_functions for s in services)
        total_unit_passing = sum(s.unit_tests_passing for s in services)
        total_examples = sum(s.examples_available for s in services)
        total_docs = sum(s.docs_available for s in services)
        
        content = f"""# Jamf Pro SDK Coverage Documentation

**Last Updated:** {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}  
**Tested Against:** Jamf Pro {self.jamf_version}  
**Total Services:** {len(services)}  
**Total Functions:** {total_functions}

## 📊 Overall Coverage

| Metric | Count | Percentage |
|--------|-------|------------|
| Functions with Unit Tests | {total_unit_passing} | {(total_unit_passing/total_functions*100):.1f}% |
| Functions with Examples | {total_examples} | {(total_examples/total_functions*100):.1f}% |
| Functions with Documentation | {total_docs} | {(total_docs/total_functions*100):.1f}% |

## 📚 Services

| Service | API Type | Functions | Unit Tests | Examples | Docs |
|---------|----------|-----------|------------|----------|------|
"""
        
        for service in sorted(services, key=lambda s: (s.api_type, s.name)):
            unit_pct = (service.unit_tests_passing / service.total_functions * 100) if service.total_functions > 0 else 0
            example_pct = (service.examples_available / service.total_functions * 100) if service.total_functions > 0 else 0
            doc_pct = (service.docs_available / service.total_functions * 100) if service.total_functions > 0 else 0
            
            api_label = "Pro" if service.api_type == "jamf_pro_api" else "Classic"
            service_link = f"[{service.name}](services/{service.name}.md)"
            content += f"| {service_link} | {api_label} | {service.total_functions} | {unit_pct:.0f}% | {example_pct:.0f}% | {doc_pct:.0f}% |\n"
        
        content += """
## Legend

- ✅ Pass / Available
- ❌ Fail / Not Available
- ⚠️ Not Tested
- 📖 Documentation link
- 💻 Example code

## How to Use

1. Navigate to a specific service page to see detailed coverage
2. Click on documentation links (📖) to view official Jamf Pro API docs
3. Click on example links (💻) to see working code examples
4. Check test status to understand what's been validated

## Contributing

To improve coverage:
1. Add missing examples in `examples/jamf_pro_api/<service>/`
2. Add missing unit tests in `jamfpro/services/jamf_pro_api/<service>/crud_test.go`
3. Add missing acceptance tests in `jamfpro/acceptance/jamf_pro_api/<service>_test.go`
"""
        
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(content)


def main():
    """Main entry point."""
    script_dir = Path(__file__).parent
    repo_root = script_dir.parent.parent.parent

    generator = CoverageGenerator(repo_root)
    generator.run()


if __name__ == "__main__":
    main()
