package vsafe_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vsafe"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestMaskPassword(t *testing.T) {
	vtest.Equal(t, `"Password":"***"`, vsafe.MaskPassword(`"Password":"123"`))
	vtest.Equal(t, `"Secret":"***"`, vsafe.MaskPassword(`"Secret":"123"`))
}

func TestMaskUrl(t *testing.T) {
	vtest.Equal(t, "http://test:***@localhost:9200", vsafe.MaskUrl("http://test:test@localhost:9200"))
	vtest.Equal(t, "https://test:***@localhost:9200", vsafe.MaskUrl("https://test:test@localhost:9200"))
	vtest.Equal(t, "mongodb://test:***@localhost:9200", vsafe.MaskUrl("mongodb://test:test@localhost:9200"))
}
