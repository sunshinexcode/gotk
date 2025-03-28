package vqueue_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vqueue"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNew(t *testing.T) {
	vtest.Equal(t, int64(0), vqueue.New().Size())

	q := vqueue.New()
	q.Push("test")

	vtest.Equal(t, int64(1), q.Len())
	vtest.Equal(t, int64(1), q.Size())
	vtest.Equal(t, "test", q.Pop())
}
