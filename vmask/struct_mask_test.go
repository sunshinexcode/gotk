package vmask_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vmask"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestMaskStruct(t *testing.T) {
	// User represents a user with sensitive information
	type User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Password string `json:"password" mask:"secret"`              // Will keep 8 chars by default
		Phone    string `json:"phone" mask:"secret:5"`               // Will keep 5 chars
		Email    string `json:"email" mask:"secret:12"`              // Will keep 12 chars
		Token    string `json:"token" mask:"secret" maskKey:"token"` // Will be masked in maps
	}

	user := &User{
		ID:       "1",
		Username: "test",
		Password: "mySecurePassword123",
		Phone:    "13812345678",
		Email:    "test@example.com",
		Token:    "secret-token-12345",
	}

	// Test struct masking
	masked := vmask.MaskStruct(user)

	maskedUser := masked.(*User)

	vtest.Equal(t, "mySecure*****", maskedUser.Password)
	vtest.Equal(t, "13812*****", maskedUser.Phone)
	vtest.Equal(t, "test@example*****", maskedUser.Email)
}
