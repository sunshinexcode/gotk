package vrand

import (
	"github.com/gogf/gf/v2/util/grand"
	"github.com/matoous/go-nanoid/v2"
)

// GetNum returns a random int between min and max: [min, max].
// The `min` and `max` also support negative numbers.
func GetNum(min int, max int) int {
	return grand.N(min, max)
}

// GetStr generates secure URL-friendly unique ID.
// Accepts optional parameter - length of the ID to be generated (21 by default).
func GetStr(num ...int) (string, error) {
	return gonanoid.New(num...)
}
