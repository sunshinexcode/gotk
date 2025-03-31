package vauth

// ContextKey is a custom type for context keys used in authentication
// It provides type safety when storing and retrieving values from context
// This helps prevent key collisions and makes the code more maintainable
type ContextKey string

// String returns the string representation of the ContextKey
// This method implements the fmt.Stringer interface
// It's useful for logging and debugging purposes
//
// Example:
//
//	key := ContextKey("user_id")
//	fmt.Println(key.String()) // Output: "user_id"
func (c ContextKey) String() string {
	return string(c)
}
