package vuuid

import (
	"github.com/google/uuid"
)

// Get return a random (version 4) UUID
func Get() string {
	return uuid.New().String()
}
