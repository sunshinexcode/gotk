package vcache

import (
	"context"
	"time"

	"github.com/sunshinexcode/gotk/vvar"
)

// Ensure LocalCache implements ILocalCache interface
var _ ILocalCache = (*LocalCache)(nil)

// ILocalCache defines the interface for local in-memory cache operations
// It provides basic cache functionality for storing and retrieving values
type ILocalCache interface {
	// Get retrieves a value from the local cache
	// Returns a Var pointer containing the value and any error that occurred
	// If the key does not exist, returns nil and no error
	Get(ctx context.Context, key any) (*vvar.Var, error)

	// Set stores a value in the local cache with an expiration time
	// Returns an error if the operation fails
	Set(ctx context.Context, key any, value any, duration time.Duration) error
}

// LocalCache implements the ILocalCache interface
// It provides in-memory caching functionality using the global cache functions
type LocalCache struct {
}

// NewLocalCache creates a new LocalCache instance
// Returns an ILocalCache interface implementation
func NewLocalCache() ILocalCache {
	return &LocalCache{}
}

// Get retrieves a value from the local cache
// Delegates to the global Get function
func (r *LocalCache) Get(ctx context.Context, key any) (*vvar.Var, error) {
	return Get(ctx, key)
}

// Set stores a value in the local cache with an expiration time
// Delegates to the global Set function
func (r *LocalCache) Set(ctx context.Context, key any, value any, duration time.Duration) error {
	return Set(ctx, key, value, duration)
}
