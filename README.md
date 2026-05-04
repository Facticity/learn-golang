# Go for Kubernetes Security Tooling

A 40-hour, markdown-first, lab-driven course for developers who are new to Go and want to build real Kubernetes security tooling from scratch in WSL.

## What you'll build

By the end of this course you will have built and shipped four serious projects:

1. **CLI Kubernetes YAML security auditor** (complete in this repo)
2. **REST API scan service**
3. **Concurrent Kubernetes event detector**
4. **Production-style Kubernetes security controller**

## Course shape (40 hours)

- **Module 00–01**: complete and fully implemented now
- **Module 02–12**: full file scaffolding with objectives, labs, and solution templates
- **Projects 02–04**: scoped with architecture, milestones, and implementation plan

| Module | Theme | Time |
|---|---|---:|
| 00 | Foundations + WSL + workflow | 3h |
| 01 | Go fundamentals for security tooling | 4h |
| 02 | Packages, testing, and quality gates | 3h |
| 03 | Concurrency patterns | 4h |
| 04 | Kubernetes APIs with client-go | 3h |
| 05 | YAML, policy checks, and rule engines | 3h |
| 06 | Container security context | 3h |
| 07 | REST scan service | 4h |
| 08 | Event-driven detectors | 3h |
| 09 | Controller-runtime operator patterns | 4h |
| 10 | Observability + hardening | 2h |
| 11 | Production delivery + CI/CD | 2h |
| 12 | Capstone readiness + portfolio | 2h |

## Quick start

```bash
make setup
make test
make project1-run
```

## Repository map

- `setup/` — WSL, Go, Docker, and Kind setup guides
- `modules/` — 13 modules with labs + solution tracks
- `projects/` — four real projects
- `kind/` — local Kubernetes cluster bootstrap
- `docker/` — Dockerfiles for projects/services
- `examples/` — reusable snippets and sample manifests


## Pedagogy and "Why"

To support foundational learners and avoid hidden assumptions, review:
- `docs-learning-philosophy.md`
- `modules/00-course-foundations/why-go-learning.md`
- `modules/01-go-fundamentals/fundamental-questions-faq.md`

## Success criteria

You should finish this course able to:

- write clean idiomatic Go with tests and concurrency
- parse and evaluate Kubernetes YAML at scale
- design REST scanning services and background workers
- build controllers with reconciliation loops
- ship production-grade tooling with observability and CI
