#!/usr/bin/env bash
set -euo pipefail
kind create cluster --name go-sec-tooling --config kind/kind-config.yaml
kubectl cluster-info
