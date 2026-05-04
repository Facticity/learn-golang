# Module 01 — Go Fundamentals for Security Tooling

## Learning Objectives
- Build a multi-package Go CLI with flags and clean architecture.
- Parse files, unmarshal YAML-like documents, and evaluate rules.
- Write table-driven tests for security checks.

## Estimated Time
4 hours

## Explanation
You will implement the first production-style project end-to-end: a CLI auditor for Kubernetes manifests.

## Hands-on Lab
See `labs/lab-01-build-auditor-checks.md`.

## Expected Output
A CLI that recursively scans YAML manifests and reports:
- containers running privileged
- missing CPU/memory limits
- containers using image tag `latest`

## Stretch Task
Add JSON output mode and rule severity filtering.

## Solution File
`solutions/lab-01-build-auditor-checks-solution.md`
