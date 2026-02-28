#!/usr/bin/env python3
"""
Generate coverage documentation for Jamf Pro SDK services.

This script analyzes the codebase to generate comprehensive documentation
showing test coverage, examples, and API documentation for each service.
"""

import os
import re
import json
import subprocess
from pathlib import Path
from datetime import datetime
from typing import Dict, List, Optional, Tuple
from dataclasses import dataclass, asdict


@dataclass
class FunctionInfo:
    """Information about a service function."""
    name: str
    endpoint: str
    doc_url: Optional[str]
    example_path: Optional[str]
    unit_test_status: str  # "pass", "fail", "not_tested"
    acceptance_test_status: str  # "pass", "fail", "not_tested"
    is_helper: bool = False
    description: str = ""


@dataclass
class ServiceInfo:
    """Information about a service."""
    name: str
    service_path: str
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
        self.services_path = repo_root / "jamfpro" / "services" / "jamf_pro_api"
        self.examples_path = repo_root / "examples" / "jamf_pro_api"
        self.output_path = repo_root / "docs" / "coverage"
    
    def get_jamf_version(self) -> str:
        """Get Jamf Pro version from live instance or use default."""
        # Try to get version from environment variable first
        env_version = os.environ.get('JAMF_PRO_VERSION')
        if env_version:
            return env_version
        
        # Try to get version from live instance using SDK
        print("  Attempting to fetch Jamf Pro version from live instance...")
        try:
            script_dir = Path(__file__).parent
            version_helper = script_dir / "get_jamf_version.go"
            
            result = subprocess.run(
                ["go", "run", str(version_helper)],
                cwd=self.repo_root,
                capture_output=True,
                text=True,
                timeout=15
            )
            
            if result.returncode == 0:
                version = result.stdout.strip()
                if version and version != "unknown":
                    print(f"  ✅ Detected Jamf Pro version: {version}")
                    return version
                else:
                    print(f"  ℹ️  Using default version (no live instance available)")
            else:
                print(f"  ℹ️  Using default version (error: {result.stderr.strip()[:100]})")
        except Exception as e:
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
        
        for service_dir in sorted(self.services_path.iterdir()):
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
        
        # Run unit tests
        print("  Running unit tests...")
        try:
            cmd = [
                "go", "test", "-v",
                "./jamfpro/services/jamf_pro_api/..."
            ]
            result = subprocess.run(
                cmd,
                cwd=self.repo_root,
                capture_output=True,
                text=True,
                timeout=300
            )
            results["unit"] = self.parse_test_output_verbose(result.stdout)
            print(f"  ✅ Parsed {len(results['unit'])} test results")
        except Exception as e:
            print(f"  ⚠️  Unit tests failed: {e}")
        
        # Note: Acceptance tests require environment variables
        # We'll mark them as "not_tested" for now
        print("  ⚠️  Acceptance tests require live environment (marked as not_tested)")
        
        return results
    
    def parse_test_output_verbose(self, output: str) -> Dict[str, str]:
        """Parse Go test verbose output."""
        test_status = {}
        
        for line in output.splitlines():
            line = line.strip()
            
            # Look for PASS/FAIL lines with test names
            if line.startswith("--- PASS:") or line.startswith("--- FAIL:"):
                parts = line.split()
                if len(parts) >= 3:
                    status = "pass" if "PASS" in parts[1] else "fail"
                    test_name = parts[2]
                    
                    # Extract service and function from TestUnit_ServiceName_FunctionName_TestCase
                    # Example: TestUnit_Accounts_ListV1_Success
                    if test_name.startswith("TestUnit_"):
                        # Remove TestUnit_ prefix
                        remaining = test_name[9:]  # len("TestUnit_") = 9
                        
                        # Split on underscore: [ServiceName, FunctionName, TestCase, ...]
                        parts = remaining.split("_")
                        if len(parts) >= 2:
                            service = parts[0].lower()
                            function = parts[1]  # The function name is always the second part
                            
                            key = f"{service}.{function}"
                            # Mark as pass if any test for this function passes
                            if key not in test_status or status == "pass":
                                test_status[key] = status
        
        return test_status
    
    def analyze_service(self, service: ServiceInfo, test_results: Dict):
        """Analyze a service to extract function information."""
        crud_file = self.services_path / service.name / "crud.go"
        
        with open(crud_file, 'r') as f:
            content = f.read()
        
        # Extract functions with their documentation
        functions = self.extract_functions(content, service.name)
        
        # Check for examples
        for func in functions:
            func.example_path = self.find_example(service.name, func.name)
            if func.example_path:
                service.examples_available += 1
            
            # Check test status
            test_key = f"{service.name}.{func.name}"
            func.unit_test_status = test_results["unit"].get(test_key, "not_tested")
            func.acceptance_test_status = "not_tested"  # Would need env vars
            
            if func.unit_test_status == "pass":
                service.unit_tests_passing += 1
            
            if func.doc_url:
                service.docs_available += 1
        
        service.functions = functions
        service.total_functions = len(functions)
    
    def extract_functions(self, content: str, service_name: str) -> List[FunctionInfo]:
        """Extract function information from Go source code."""
        functions = []
        
        # Pattern to match function declarations with comments
        pattern = r'//\s*(.+?)\n(?://.*\n)*func\s+\([^)]+\)\s+(\w+)\s*\([^)]*\)'
        
        for match in re.finditer(pattern, content, re.MULTILINE):
            comment_first_line = match.group(1).strip()
            func_name = match.group(2)
            
            # Skip unexported functions
            if not func_name[0].isupper():
                continue
            
            # Extract full comment block
            func_pos = match.start(2)
            comment_start = content.rfind('//', 0, func_pos - 50)
            if comment_start != -1:
                comment_block = content[comment_start:func_pos].strip()
            else:
                comment_block = comment_first_line
            
            # Extract doc URL
            doc_url = self.extract_doc_url(comment_block)
            
            # Extract endpoint
            endpoint = self.extract_endpoint(comment_block, func_name)
            
            # Determine if helper function
            is_helper = self.is_helper_function(func_name, comment_block)
            
            # Extract description
            description = self.extract_description(comment_block, func_name)
            
            func_info = FunctionInfo(
                name=func_name,
                endpoint=endpoint,
                doc_url=doc_url,
                example_path=None,
                unit_test_status="not_tested",
                acceptance_test_status="not_tested",
                is_helper=is_helper,
                description=description
            )
            functions.append(func_info)
        
        return functions
    
    def extract_doc_url(self, comment: str) -> Optional[str]:
        """Extract Jamf Pro API documentation URL from comment."""
        url_pattern = r'https?://developer\.jamf\.com[^\s\)]*'
        match = re.search(url_pattern, comment)
        return match.group(0) if match else None
    
    def extract_endpoint(self, comment: str, func_name: str) -> str:
        """Extract API endpoint from comment or infer from function name."""
        # Look for URL: pattern
        url_match = re.search(r'URL:\s*(\w+\s+/[^\n]+)', comment)
        if url_match:
            return url_match.group(1).strip()
        
        # Look for endpoint pattern
        endpoint_match = re.search(r'(GET|POST|PUT|PATCH|DELETE)\s+/api/[^\n\s]+', comment)
        if endpoint_match:
            return endpoint_match.group(0).strip()
        
        # Check if it's a helper function
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
                # Remove function name and return the rest
                desc = line[len(func_name):].strip()
                # Remove leading verb connectors
                for prefix in [' returns', ' gets', ' creates', ' updates', ' deletes', ' lists']:
                    if desc.lower().startswith(prefix):
                        desc = desc[len(prefix):].strip()
                        break
                return desc
        return ""
    
    def find_example(self, service_name: str, func_name: str) -> Optional[str]:
        """Find example file for a function."""
        service_examples = self.examples_path / service_name
        if not service_examples.exists():
            return None
        
        # Try exact match first (converted to snake_case)
        func_dir = self.function_name_to_dir(func_name)
        example_file = service_examples / func_dir / "main.go"
        if example_file.exists():
            return str(example_file.relative_to(self.repo_root))
        
        # Try fuzzy matching - look for directories that might match
        # Remove version suffix and split camelCase into words
        func_base = re.sub(r'V\d+$', '', func_name)
        # Split camelCase: "DeleteBuildingsByID" -> ["Delete", "Buildings", "By", "ID"]
        func_words_list = re.sub('([A-Z][a-z]+)', r' \1', func_base).split()
        func_words = set(w.lower() for w in func_words_list if w)
        
        # Find all potential matches and score them
        matches = []
        for subdir in service_examples.iterdir():
            if not subdir.is_dir():
                continue
            
            dir_name = subdir.name
            main_file = subdir / "main.go"
            
            if not main_file.exists():
                continue
            
            # Check if the function name contains key words from the directory name
            # e.g., "Delete Buildings By ID" contains "delete" and matches "delete_multiple"
            # or "Get Building History" contains "get" and "history" and matches "get_history"
            dir_words = set(dir_name.split('_'))
            overlap = len(dir_words & func_words)
            threshold = len(dir_words) * 0.7
            
            # If most directory words are in the function name, it's a potential match
            if dir_words and overlap >= threshold:
                # Score: prefer matches with more overlapping words and fewer total dir words
                # This prefers "get_history" (2/2 match) over "get" (1/1 match) for "GetBuildingHistory"
                score = overlap / len(dir_words)
                matches.append((score, overlap, dir_name, main_file))
        
        # Return the best match (highest score, then most overlapping words)
        if matches:
            matches.sort(key=lambda x: (x[0], x[1]), reverse=True)
            return str(matches[0][3].relative_to(self.repo_root))
        
        return None
    
    def function_name_to_dir(self, func_name: str) -> str:
        """Convert function name to example directory name."""
        # Remove version suffix
        name = re.sub(r'V\d+$', '', func_name)
        
        # Handle common patterns first
        # GetByID -> get, GetByName -> get_by_name, etc.
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
        
        # Convert camelCase to snake_case for other cases
        name = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', name)
        name = re.sub('([a-z0-9])([A-Z])', r'\1_\2', name)
        result = name.lower()
        
        # Remove redundant service name prefixes (e.g., delete_buildings_by_id -> delete_multiple)
        # This is a heuristic: if the result contains the pattern "service_name" in it, simplify
        # For now, just return the snake_case conversion
        return result
    
    def generate_service_doc(self, service: ServiceInfo):
        """Generate markdown documentation for a service."""
        output_file = self.output_path / "services" / f"{service.name}.md"
        
        # Separate API functions from helpers
        api_functions = [f for f in service.functions if not f.is_helper]
        helper_functions = [f for f in service.functions if f.is_helper]
        
        # Calculate percentages
        unit_pct = (service.unit_tests_passing / service.total_functions * 100) if service.total_functions > 0 else 0
        acc_pct = (service.acceptance_tests_passing / service.total_functions * 100) if service.total_functions > 0 else 0
        example_pct = (service.examples_available / service.total_functions * 100) if service.total_functions > 0 else 0
        doc_pct = (service.docs_available / service.total_functions * 100) if service.total_functions > 0 else 0
        
        content = f"""# {service.name.replace('_', ' ').title()} Service

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
        
        with open(output_file, 'w') as f:
            f.write(content)
    
    def generate_index(self, services: List[ServiceInfo]):
        """Generate index README for all services."""
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

| Service | Functions | Unit Tests | Examples | Docs |
|---------|-----------|------------|----------|------|
"""
        
        for service in sorted(services, key=lambda s: s.name):
            unit_pct = (service.unit_tests_passing / service.total_functions * 100) if service.total_functions > 0 else 0
            example_pct = (service.examples_available / service.total_functions * 100) if service.total_functions > 0 else 0
            doc_pct = (service.docs_available / service.total_functions * 100) if service.total_functions > 0 else 0
            
            service_link = f"[{service.name}](services/{service.name}.md)"
            content += f"| {service_link} | {service.total_functions} | {unit_pct:.0f}% | {example_pct:.0f}% | {doc_pct:.0f}% |\n"
        
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
        
        with open(output_file, 'w') as f:
            f.write(content)


def main():
    """Main entry point."""
    import sys
    
    # Get repo root (script is in scripts/pipeline/coverage)
    script_dir = Path(__file__).parent
    repo_root = script_dir.parent.parent.parent
    
    # Jamf Pro version will be auto-detected by CoverageGenerator
    generator = CoverageGenerator(repo_root)
    generator.run()


if __name__ == "__main__":
    main()
