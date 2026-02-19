#!/usr/bin/env python3
"""Git operations for PR analysis.

Provides functions for querying Git history and identifying changed files.
"""

import subprocess
from typing import List, Optional, Set


def get_changed_files(base_ref: str, file_extension: str = ".go") -> List[str]:
    """Get list of changed files with specific extension.

    Args:
        base_ref: Base branch reference (e.g., 'origin/main').
        file_extension: File extension to filter (default: '.go').

    Returns:
        List of changed file paths.

    Raises:
        subprocess.CalledProcessError: If git command fails.
    """
    result = subprocess.run(
        ["git", "diff", "--name-only", f"{base_ref}...HEAD"],
        capture_output=True,
        text=True,
        check=True,
    )

    return [
        line.strip()
        for line in result.stdout.split("\n")
        if line.strip().endswith(file_extension)
    ]


def get_changed_packages(
    base_ref: str,
    exclude_prefixes: Optional[List[str]] = None,
) -> List[str]:
    """Get list of changed Go packages.

    Args:
        base_ref: Base branch reference.
        exclude_prefixes: Package paths starting with any of these are excluded
            (e.g. ['jamfpro/acceptance'] to skip acceptance tests).

    Returns:
        List of unique package directories (sorted).
    """
    go_files = get_changed_files(base_ref, ".go")

    if not go_files:
        return []

    exclude_prefixes = exclude_prefixes or []
    packages: Set[str] = set()
    for file_path in go_files:
        parts = file_path.split("/")
        if len(parts) > 1:
            pkg_path = "/".join(parts[:-1])
            if any(pkg_path.startswith(prefix) for prefix in exclude_prefixes):
                continue
            packages.add(pkg_path)

    return sorted(packages)
