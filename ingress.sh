#!/bin/bash

set -e

echo "📁 Criando estrutura de diretórios..."

mkdir -p cmd
mkdir -p internal/{domain,usecase,adapter/{k8s,http,cli},infra/{k8sclient,logger,config},di}
mkdir -p api
mkdir -p pkg/diff
mkdir -p test
mkdir -p deployments
mkdir -p scripts
mkdir -p docs

echo "✅ Estrutura de diretórios criada."

echo "📦 Inicializando módulo Go..."
go mod init github.com/giovanni-gava/gateway-migrator
go mod tidy

echo "🐍 Instalando Cobra CLI..."
go install github.com/spf13/cobra-cli@latest
cobra-cli init --pkg-name github.com/giovanni-gava/gateway-migrator --viper --author "Giovanni Gava" --license mit

echo "🚀 Criando entrypoint principal..."

cat > cmd/main.go <<EOF
package main

import "github.com/giovanni-gava/gateway-migrator/cmd"

func main() {
	cmd.Execute()
}
EOF

echo "📦 Instalando Wire para DI..."
go install github.com/google/wire/cmd/wire@latest

echo "⚙️ Criando Makefile..."

cat > Makefile <<'EOF'
run:
	go run cmd/main.go

build:
	go build -o bin/gateway-migrator cmd/main.go

test:
	go test ./...

wire:
	cd internal/di && wire

clean:
	rm -rf bin
EOF

echo "📄 Criando README.md base..."

cat > README.md <<EOF
# gateway-migrator 🚀

> Seamless migration from Ingress to Gateway API — with zero downtime, observability, and clean architecture.

## 📦 Features

- Clean Architecture (Hexagonal)
- CLI with Cobra
- Dependency Injection with Wire
- Kubernetes API interaction
- Smart diff & safe rollout

## 🚀 Usage

```bash
make run
