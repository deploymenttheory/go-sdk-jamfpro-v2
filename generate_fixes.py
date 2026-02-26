#!/usr/bin/env python3
"""
Generate a list of files and test patterns that need fixing.
"""

import subprocess
import re
from pathlib import Path

def get_service_name_from_path(file_path):
    """Extract service name from file path and convert to PascalCase."""
    parts = Path(file_path).parts
    service_dir = parts[-2]
    
    # Convert snake_case to PascalCase
    words = service_dir.split('_')
    service_name = ''.join(word.capitalize() for word in words)
    
    return service_name

def find_non_conforming_files():
    """Find all files with non-conforming test names."""
    result = subprocess.run(
        ['find', 'jamfpro/services', '-name', 'crud_test.go', '-type', 'f'],
        capture_output=True,
        text=True,
        cwd='/Users/dafyddwatkins/GitHub/deploymenttheory/go-sdk-jamfpro-v2'
    )
    
    files = result.stdout.strip().split('\n')
    non_conforming = []
    
    for file_path in files:
        full_path = f'/Users/dafyddwatkins/GitHub/deploymenttheory/go-sdk-jamfpro-v2/{file_path}'
        
        # Check if file has non-conforming test names
        grep_result = subprocess.run(
            ['grep', '-E', '^func TestUnit[^_]', full_path],
            capture_output=True,
            text=True
        )
        
        if grep_result.returncode == 0:
            service_name = get_service_name_from_path(file_path)
            
            # Get unique test prefixes
            test_names = grep_result.stdout.strip().split('\n')
            prefixes = set()
            for test in test_names:
                match = re.match(r'func (TestUnit\w+?)(_|Success|Error|NotFound|EmptyID|ZeroID|NegativeID|NilRequest|Conflict|EmptyName)', test)
                if match:
                    prefixes.add(match.group(1))
            
            non_conforming.append({
                'file': full_path,
                'service': service_name,
                'prefixes': sorted(prefixes)
            })
    
    return non_conforming

if __name__ == '__main__':
    files = find_non_conforming_files()
    
    print(f"Found {len(files)} files with non-conforming test names:\n")
    
    for item in files:
        print(f"File: {item['file']}")
        print(f"Service: {item['service']}")
        print(f"Prefixes to fix: {len(item['prefixes'])}")
        for prefix in item['prefixes'][:3]:
            print(f"  - {prefix}")
        if len(item['prefixes']) > 3:
            print(f"  ... and {len(item['prefixes']) - 3} more")
        print()
