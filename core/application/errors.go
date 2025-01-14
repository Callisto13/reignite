package application

import (
	"errors"
	"fmt"

	"github.com/weaveworks/flintlock/core/models"
)

var (
	errIDRequired        = errors.New("microvm id is required")
	errNamespaceRequired = errors.New("namespace is required")
)

type specAlreadyExistsError struct {
	name      string
	namespace string
}

// Error returns the error message.
func (e specAlreadyExistsError) Error() string {
	return fmt.Sprintf("microvm spec %s/%s already exists", e.namespace, e.name)
}

type specNotFoundError struct {
	name      string
	namespace string
}

// Error returns the error message.
func (e specNotFoundError) Error() string {
	return fmt.Sprintf("microvm spec %s/%s not found", e.namespace, e.name)
}

type reachedMaximumRetryError struct {
	vmid    models.VMID
	retries int
}

func (e reachedMaximumRetryError) Error() string {
	return fmt.Sprintf(
		"microvm reconciliation for %s failed %d times",
		e.vmid.String(),
		e.retries,
	)
}
