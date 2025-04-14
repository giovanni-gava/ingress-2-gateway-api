package domain

import (
	"context"
)

// DiffResult representa o resultado da análise de conversão
type DiffResult struct {
	Summary     string
	Warnings    []string
	Suggestions string
	Resources   map[string]string // nome -> YAML sugerido
}

// IngressToGatewayConverter define o contrato de um conversor
type IngressToGatewayConverter interface {
	GenerateDiff(ctx context.Context, sourceNS, targetNS string) (*DiffResult, error)
}
