package usecase

import (
	"context"
	"fmt"

	"github.com/giovanni-gava/gateway-migrator/internal/domain"
)

type GenerateDiffUseCase struct {
	Converter domain.IngressToGatewayConverter
}

func NewGenerateDiffUseCase(converter domain.IngressToGatewayConverter) *GenerateDiffUseCase {
	return &GenerateDiffUseCase{
		Converter: converter,
	}
}

func (uc *GenerateDiffUseCase) Execute(ctx context.Context, sourceNS, targetNS string) error {
	result, err := uc.Converter.GenerateDiff(ctx, sourceNS, targetNS)
	if err != nil {
		return fmt.Errorf("failed to generate diff: %w", err)
	}

	fmt.Println("📄 Diff Summary:", result.Summary)
	for _, warn := range result.Warnings {
		fmt.Println("⚠️ Warning:", warn)
	}
	fmt.Println("✨ Suggestions:", result.Suggestions)

	for name, yaml := range result.Resources {
		fmt.Printf("\n--- %s ---\n%s\n", name, yaml)
	}

	return nil
}
