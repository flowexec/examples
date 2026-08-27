#!/usr/bin/env python3
# f:name=report f:verb=generate
# f:description=Summarize a JSON file, or the built-in sample when none is given.
# f:tags=python|fromfile
# f:args=pos:1:DATA_FILE
"""Imported by python/scripts.flow via `imports`.

The f: comments above are read at sync time, exactly as they are for .sh
scripts - Python uses the same # comment prefix, so the syntax is unchanged.
"""

import json
import os
import sys

SAMPLE = {"passed": 48, "failed": 0, "skipped": 2}


def main() -> int:
    path = os.environ.get("DATA_FILE", "")
    if path:
        with open(path, encoding="utf-8") as handle:
            data = json.load(handle)
    else:
        print("no DATA_FILE given, using the built-in sample")
        data = SAMPLE

    total = sum(data.values())
    print(f"total:  {total}")
    for key, value in sorted(data.items()):
        share = (value / total * 100) if total else 0
        print(f"  {key:<8} {value:>4}  ({share:.1f}%)")

    # A non-zero exit propagates: flow marks the run failed and shows the
    # traceback or message, the same as any other executable.
    return 1 if data.get("failed") else 0


if __name__ == "__main__":
    sys.exit(main())
