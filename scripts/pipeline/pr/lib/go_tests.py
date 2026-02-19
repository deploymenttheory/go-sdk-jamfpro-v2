#!/usr/bin/env python3
"""Go test execution utilities.

Provides functions for running Go unit tests.
"""

import os
import subprocess
from pathlib import Path
from typing import List


def run_unit_tests(packages: List[str], output_dir: str = "coverage") -> Path:
    """Run Go unit tests with coverage profiling.

    Args:
        packages: List of package paths to test.
        output_dir: Directory for coverage output files.

    Returns:
        Path to merged coverage file.
    """
    print("\n" + "=" * 60)
    print("📊 Running Unit Tests with Coverage")
    print("=" * 60)

    coverage_dir = Path(output_dir)
    coverage_dir.mkdir(parents=True, exist_ok=True)

    merged_file = coverage_dir / "unit-coverage.txt"
    coverage_files = []

    for idx, package in enumerate(packages, 1):
        safe_name = package.replace("/", "_").replace(".", "_").strip("_")
        coverage_file = coverage_dir / f"{safe_name}.out"

        print(f"\n[{idx}/{len(packages)}] Testing: {package}")

        cmd = [
            "go",
            "test",
            "-v",
            f"-coverprofile={coverage_file}",
            "-covermode=atomic",
            f"./{package}",
        ]

        result = subprocess.run(cmd, env=os.environ.copy(), check=False)
        if result.returncode != 0:
            raise SystemExit(result.returncode)

        if coverage_file.exists():
            coverage_files.append(coverage_file)
            print(f"✅ Coverage generated for {package}")
        else:
            print(f"⚠️  No coverage file for {package}")

    if coverage_files:
        print(f"\n📊 Merging {len(coverage_files)} coverage file(s)...")
        _merge_coverage_files(coverage_files, merged_file)
        print(f"✅ Merged coverage file: {merged_file}")

    return merged_file


def _merge_coverage_files(
    coverage_files: List[Path], output_file: Path
) -> None:
    """Merge multiple Go coverage files into one."""
    with open(output_file, "w", encoding="utf-8") as out_f:
        out_f.write("mode: atomic\n")
        for cov_file in coverage_files:
            with open(cov_file, encoding="utf-8") as in_f:
                for line in in_f:
                    if not line.startswith("mode:"):
                        out_f.write(line)


def packages_missing_unit_tests(packages: List[str]) -> List[str]:
    """Return packages that have no *_test.go file in their directory."""
    missing = []
    for pkg in packages:
        pkg_path = Path(pkg)
        if not pkg_path.is_dir():
            continue
        if not any(pkg_path.glob("*_test.go")):
            missing.append(pkg)
    return missing
