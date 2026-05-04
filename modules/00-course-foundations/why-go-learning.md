# Why This Course Teaches Go This Way

## Goal
This document explains *why* the course structure uses progressive labs, package boundaries, and production-like tooling from day one.

## Why modules are lab-first
- Reading syntax is not enough for engineering roles.
- Employers evaluate debugging, tradeoffs, and delivery speed.
- Labs force decision-making under constraints.

## Why package boundaries appear early
- Go codebases scale through clear package responsibilities.
- New developers often over-couple logic to `main`.
- We isolate parser/audit/report to teach maintainable design.

## Why tests are required in every major step
- Security tooling must be deterministic and reproducible.
- Tests document expected behavior better than prose alone.
- Tests make refactoring safe during later modules.

## Why we discuss tradeoffs, not only syntax
- Every Go style choice has costs (simplicity vs flexibility, explicitness vs abstraction).
- Kubernetes tooling adds reliability and operational constraints.
- Good engineers explain *why* a pattern fits a problem.

## If you're new and have foundational questions
You are expected to ask:
- Why use interfaces here?
- Why pass context?
- Why keep functions small?
- Why avoid clever abstractions early?

This repo now includes FAQ and design-rationale notes to answer those explicitly.
