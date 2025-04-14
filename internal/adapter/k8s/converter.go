package k8s

import (
	"context"
	"fmt"

	"github.com/giovanni-gava/gateway-migrator/internal/domain"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sIngressConverter struct {
	client *kubernetes.Clientset
}

func NewK8sIngressConverter() (*K8sIngressConverter, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		// fallback to local kubeconfig (para desenvolvimento)
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to load Kubernetes config: %w", err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	return &K8sIngressConverter{client: clientset}, nil
}

func (k *K8sIngressConverter) GenerateDiff(ctx context.Context, sourceNS, targetNS string) (*domain.DiffResult, error) {
	ingresses, err := k.client.NetworkingV1().Ingresses(sourceNS).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list Ingresses: %w", err)
	}

	summary := fmt.Sprintf("%d Ingress resources analyzed.", len(ingresses.Items))
	suggestions := "Add proper timeouts, retries, and ensure backendRefs match services."

	resources := make(map[string]string)
	for _, ing := range ingresses.Items {
		routeName := ing.Name + "-route"
		resources[routeName] = fmt.Sprintf(`
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: %s
spec:
  parentRefs:
  - name: my-gateway
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /
    backendRefs:
    - name: %s
      port: 80
`, routeName, ing.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.Service.Name)
	}

	return &domain.DiffResult{
		Summary:     summary,
		Suggestions: suggestions,
		Resources:   resources,
	}, nil
}
