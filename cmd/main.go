package main

import (
	"fmt"
	"os"

	"github.com/giovanni-gava/gateway-migrator/internal/adapter/cli"
)

func main() {
	if err := cli.NewRootCommand().Execute(); err != nil {
		fmt.Println("❌ Erro ao executar comando:", err)
		os.Exit(1)
	}
}
