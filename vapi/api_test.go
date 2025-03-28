package vapi_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestConst(t *testing.T) {
	vtest.Equal(t, "application/json", vapi.MimeJson)
	vtest.Equal(t, "application/x-www-form-urlencoded", vapi.MimePostForm)
	vtest.Equal(t, "multipart/form-data", vapi.MimeMultipartPostForm)
}
