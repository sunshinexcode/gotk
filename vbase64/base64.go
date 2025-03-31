package vbase64

import (
	"github.com/gogf/gf/v2/encoding/gbase64"
)

// Package vbase64 provides base64 encoding and decoding functionality.
// It re-exports all methods from github.com/gogf/gf/v2/encoding/gbase64 package.
//
// Example:
//   // Encode string to base64
//   encoded := vbase64.EncodeString("Hello World")
//   // Decode base64 string
//   decoded, err := vbase64.DecodeString("SGVsbG8gV29ybGQ=")
//   if err != nil {
//       // Handle error
//   }

// Export all methods from gbase64 package
var (
	Decode         = gbase64.Decode
	DecodeStr      = gbase64.DecodeString
	DecodeString   = gbase64.DecodeString
	DecodeToStr    = gbase64.DecodeToString
	DecodeToString = gbase64.DecodeToString
	Encode         = gbase64.Encode
	EncodeStr      = gbase64.EncodeString
	EncodeString   = gbase64.EncodeString
	EncodeToStr    = gbase64.EncodeToString
	EncodeToString = gbase64.EncodeToString
)
