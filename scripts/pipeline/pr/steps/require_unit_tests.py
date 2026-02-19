#!/usr/bin/env python3
"""Ensure each changed package has at least one *_test.go file.

Exits with code 1 and prints errors for any package missing unit tests.
"""

import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent.parent))

from lib.go_tests import packages_missing_unit_tests  # noqa: E402


def main():
    """Check packages for unit tests and exit 1 if any are missing."""
    import argparse

    parser = argparse.ArgumentParser(
        description="Require unit tests for changed packages"
    )
    parser.add_argument(
        "--packages",
        required=True,
        help="Space-separated list of package paths",
    )
    args = parser.parse_args()

    packages = [p.strip() for p in args.packages.split() if p.strip()]
    if not packages:
        return 0

    missing = packages_missing_unit_tests(packages)
    if not missing:
        print("✅ All changed packages have unit tests.")
        return 0

    print("::error::The following changed packages have no unit tests (*_test.go):", file=sys.stderr)
    for pkg in missing:
        print(f"::error::  - {pkg}", file=sys.stderr)
    print("::error::Add at least one *_test.go file in each changed package.", file=sys.stderr)
    return 1


if __name__ == "__main__":
    sys.exit(main())
