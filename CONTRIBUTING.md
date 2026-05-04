# Contributing Standards

## Engineering Expectations
- Prefer small, testable packages with clear responsibilities.
- Use table-driven tests for rule engines and parsers.
- Treat security checks as deterministic and reproducible.
- Return explicit exit codes for CI usage.
- Run `go fmt ./...` and `go test ./...` before opening a PR.

## Project 1 Design Notes
- `cmd/` contains CLI glue only.
- `internal/parser` handles manifest discovery + extraction.
- `internal/audit` handles pure policy logic.
- `internal/report` handles output rendering.
