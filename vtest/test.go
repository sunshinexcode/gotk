package vtest

import (
	"github.com/stretchr/testify/assert"
)

// Equal asserts that two objects are equal.
func Equal(t assert.TestingT, expected, actual any, msgAndArgs ...any) bool {
	return assert.Equal(t, expected, actual, msgAndArgs...)
}

// Nil asserts that the specified object is nil.
func Nil(t assert.TestingT, object any, msgAndArgs ...any) bool {
	return assert.Nil(t, object, msgAndArgs...)
}

// NotEqual asserts that the specified values are NOT equal.
func NotEqual(t assert.TestingT, expected, actual any, msgAndArgs ...any) bool {
	return assert.NotEqual(t, expected, actual, msgAndArgs...)
}

// NotNil asserts that the specified object is not nil.
func NotNil(t assert.TestingT, object any, msgAndArgs ...any) bool {
	return assert.NotNil(t, object, msgAndArgs...)
}
