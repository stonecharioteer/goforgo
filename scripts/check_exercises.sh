#!/usr/bin/env bash
set -euo pipefail

# GoForGo Exercise Completeness Checker
# Delegates to the Go implementation so the check is portable and testable.

go run ./cmd/check-exercises
