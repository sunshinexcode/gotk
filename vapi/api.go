package vapi

// Common MIME type constants used in HTTP requests and responses
const (
	// MimeJson represents the MIME type for JSON data
	// Used when sending or receiving JSON-formatted data
	// Example: Content-Type: application/json
	MimeJson = "application/json"

	// MimeMultipartPostForm represents the MIME type for multipart form data
	// Used when sending form data that includes file uploads
	// Example: Content-Type: multipart/form-data
	MimeMultipartPostForm = "multipart/form-data"

	// MimePostForm represents the MIME type for URL-encoded form data
	// Used when sending traditional form data
	// Example: Content-Type: application/x-www-form-urlencoded
	MimePostForm = "application/x-www-form-urlencoded"
)
