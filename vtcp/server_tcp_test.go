package vtcp_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtcp"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestNewServerErrorMock(t *testing.T) {
	patchNewServerKeyCrt := vmock.ApplyFuncReturn(vtcp.NewServerKeyCrt, nil, errors.New("new server error"))
	defer patchNewServerKeyCrt.Reset()

	patchFatal := vmock.ApplyFuncReturn(vlog.Fatal)
	defer patchFatal.Reset()

	s := &vtcp.Server{Address: ":6000", Handler: func(conn *vtcp.Conn) {}, StopHandler: func() {}}

	app := vfx.New(
		vfx.NopLogger,
		vfx.Options(
			vfx.Supply(s),
			vfx.Invoke(vtcp.NewServer),
		))
	ctx := context.TODO()

	vtest.Nil(t, app.Start(ctx))

	time.Sleep(1 * time.Second)

	vtest.Nil(t, app.Stop(ctx))
}
