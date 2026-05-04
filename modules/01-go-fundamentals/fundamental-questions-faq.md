# Fundamental Questions FAQ (Module 01)

## Why does Go prefer explicit error returns?
Because error handling is part of normal control flow. Explicit returns make failure paths visible and reviewable.

## Why avoid inheritance-heavy design?
Go favors composition. Small focused types are easier to test and evolve.

## Why are package boundaries strict?
They prevent hidden coupling. In this course:
- parser: input normalization
- audit: policy rules
- report: output formatting

## Why table-driven tests?
They scale naturally for policy engines and edge cases.

## Why not overuse interfaces early?
Define interfaces where multiple implementations are needed. Premature interfaces increase indirection without value.

## Why use context in production services?
To propagate deadlines/cancellation across API and worker boundaries.

## Why enforce formatting/lint/testing gates?
Consistency reduces review friction and production defects.
