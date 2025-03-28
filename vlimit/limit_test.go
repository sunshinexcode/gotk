package vlimit_test

import (
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vlimit"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestEvery(t *testing.T) {
	limiter := vlimit.New(vlimit.Every(time.Millisecond*31), 2)

	vtest.Equal(t, 2, limiter.Burst())

	for i := 0; i < 10; i++ {
		ok := limiter.Allow()
		if ok {
			vtest.Equal(t, true, ok)
		} else {
			vtest.Equal(t, false, ok)
		}
		time.Sleep(time.Millisecond * 20)
	}
}

func TestNew(t *testing.T) {
	limiter := vlimit.New(1, 1)

	vtest.Equal(t, 1, limiter.Burst())

	for i := 0; i < 10; i++ {
		ok := limiter.Allow()
		if ok {
			vtest.Equal(t, true, ok)
		} else {
			vtest.Equal(t, false, ok)
		}
	}
}
