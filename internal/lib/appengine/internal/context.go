package internal

import "context"

type contextKey string

const (
	namespaceKey = contextKey("namespace")
)

func GetAppEngineNamespace(ctx context.Context) string {
	value := ctx.Value(namespaceKey)

	if value == nil {
		return ""
	}

	return value.(string)
}

func WithAppEngineNamespace(ctx context.Context, namespace string) context.Context {
	return context.WithValue(ctx, namespaceKey, namespace)
}
