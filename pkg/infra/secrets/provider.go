package secrets
import (
	"context"
	"fmt"
	"os"
)
type SecretProvider interface {
	GetSecret(ctx context.Context, key string) (string, error)
}
type VaultProvider struct {
	Address string
	Token   string
}
func (v *VaultProvider) GetSecret(ctx context.Context, key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("secret %s not found in secure store", key)
	}
	return val, nil
}
