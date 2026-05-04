# Setup: WSL + Go + Tooling

## Learning Objectives
- Install Go, Docker, kubectl, kind in WSL.
- Validate local developer loop with `go test`.

## Steps
1. Update apt and install prerequisites.
2. Install Go 1.22+ and add to PATH.
3. Install Docker Desktop + WSL integration.
4. Install kubectl and kind.

## Validation
```bash
go version
docker --version
kubectl version --client
kind --version
```
