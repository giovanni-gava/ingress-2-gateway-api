# gateway-migrator 🚀

> Seamless, safe and automated migration from Kubernetes Ingress to Gateway API — built for SREs, DevOps, and Platform Engineers.

## 🧠 Why This Project?

Kubernetes is evolving — Gateway API is replacing Ingress as the future of traffic management. But the migration can be risky, breaking apps and affecting uptime.

This tool ensures:
- ✅ **Zero Downtime** migration
- 🔍 Deep **error tracing** and logging
- ⚠️ **Guard rails** to prevent bad deployments
- 🧠 Smart diff and validation engine
- ♻️ **Rollback-safe** transitions
- ⚡ Designed for **massive rollouts**

## 🧰 Key Features

- Clean Architecture & Hexagonal Design
- Go Concurrency + Channels + Mutexes
- Dependency Injection (Wire)
- CLI-powered (Cobra)
- Structured Logging
- Unit and E2E Test Suite
- Optional Slack Notifications
- gRPC / REST support
- Configurable rollouts (blue/green, canary)

## 🏗️ Project Modules (by domain)
- `core/domain` — domain entities & interfaces
- `core/usecase` — business logic & orchestrators
- `infra/k8s` — adapters to Kubernetes APIs
- `infra/http` — HTTP server & handlers
- `cli/cmd` — CLI entrypoint using Cobra
- `pkg/utils` — helpers, mappers, templates
- `test/` — mocks, fixtures, integration tests

## 🚀 Getting Started

```bash
git clone https://github.com/giovanni-gava/gateway-migrator.git
cd gateway-migrator
make run
