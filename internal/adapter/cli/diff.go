package cli

import (
	"context"
	"fmt"

	"github.com/giovanni-gava/gateway-migrator/internal/adapter/k8s"
	"github.com/giovanni-gava/gateway-migrator/internal/domain"
	"github.com/giovanni-gava/gateway-migrator/internal/usecase"
	"github.com/spf13/cobra"
)

// mock temporário até termos o adapter real
type mockConverter struct{}

func (m *mockConverter) GenerateDiff(ctx context.Context, sourceNS, targetNS string) (*domain.DiffResult, error) {
	return &domain.DiffResult{
		Summary:     "2 Ingress resources analyzed.",
		Warnings:    []string{"Ingress 'api-v1' missing backend."},
		Suggestions: "Consider enabling retries and setting proper timeouts.",
		Resources: map[string]string{
			"http-app-route": `
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: http-app
spec:
  parentRefs:
  - name: my-gateway
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /
    backendRefs:
    - name: my-service
      port: 80
`,
		},
	}, nil
}

func NewDiffCommand() *cobra.Command {
	var (
		sourceNamespace string
		targetNamespace string
		outputFile      string
	)

	cmd := &cobra.Command{
		Use:   "diff",
		Short: "Generate a diff between Ingress and Gateway API resources",
		Run: func(cmd *cobra.Command, args []string) {
			converter, err := k8s.NewK8sIngressConverter()
			if err != nil {
				fmt.Println("❌ Erro ao inicializar converter:", err)
				return
			}
			useCase := usecase.NewGenerateDiffUseCase(converter)

			ctx := context.Background()
			if err := useCase.Execute(ctx, sourceNamespace, targetNamespace); err != nil {
				fmt.Println("❌ Erro:", err)
			}
		},
	}

	cmd.Flags().StringVarP(&sourceNamespace, "source", "s", "default", "Namespace de origem")
	cmd.Flags().StringVarP(&targetNamespace, "target", "t", "default", "Namespace de destino")
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Arquivo de saída (não implementado ainda)")

	return cmd
}
