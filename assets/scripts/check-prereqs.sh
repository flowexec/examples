#!/bin/sh
# Checks for common development tool prerequisites and prints a status line for each.

check() {
  tool="$1"
  if command -v "$tool" > /dev/null 2>&1; then
    echo "✅ $tool found at $(command -v "$tool")"
  else
    echo "❌ $tool not found"
  fi
}

echo "=== Prerequisite check ==="
check go
check git
check docker
check node
check kubectl
echo "=========================="
