package verror_test

import (
	"errors"
	"testing"

	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestErrorNewError(t *testing.T) {
	vtest.Equal(t, "10000|unknown error|<nil>", verror.NewError(vcode.CodeErrUnknown).Error())
}

func TestErrorIs(t *testing.T) {
	e := verror.NewError(vcode.CodeErrUnknown)

	vtest.Equal(t, true, e.(*verror.Error).Is(verror.ErrUnknown))
	vtest.Equal(t, false, e.(*verror.Error).Is(verror.ErrParamBind))
	vtest.Equal(t, false, e.(*verror.Error).Is(errors.New("test error")))

	e1 := errors.New("test error")
	err := verror.Wrap(e1, vcode.CodeErrUnknown)
	err = verror.Wrap(err, vcode.CodeErrParamBind, "test data", vmap.M{"a": 1, "b": 2})

	vtest.Equal(t, true, errors.Is(err, e1))
	vtest.Equal(t, false, errors.Is(err, errors.New("test error")))
	vtest.Equal(t, true, errors.Is(err, verror.ErrUnknown))
	vtest.Equal(t, true, errors.Is(err, verror.ErrParamBind))
	vtest.Equal(t, false, errors.Is(err, verror.ErrHttpStatusNotOk))
}

func TestGetCode(t *testing.T) {
	code, message := verror.GetCode(nil)

	vtest.Equal(t, 0, code)
	vtest.Equal(t, "success", message)

	code, message = verror.GetCode(errors.New("test error"))

	vtest.Equal(t, 10000, code)
	vtest.Equal(t, "unknown error", message)

	code, message = verror.GetCode(verror.ErrParamBind)

	vtest.Equal(t, 10110, code)
	vtest.Equal(t, "param bind error", message)

	err := verror.Wrap(errors.New("test error"), vcode.CodeErrUnknown)
	err = verror.Wrap(err, vcode.CodeErrParamBind, "test data", vmap.M{"a": 1, "b": 2})
	code, message = verror.GetCode(err)

	vtest.Equal(t, 10110, code)
	vtest.Equal(t, "param bind error", message)
}

func TestGetCodeS(t *testing.T) {
	vtest.Equal(t, 0, verror.GetCodeS(nil).Code())
	vtest.Equal(t, "success", verror.GetCodeS(nil).Message())

	vtest.Equal(t, 10000, verror.GetCodeS(errors.New("test error")).Code())
	vtest.Equal(t, "unknown error", verror.GetCodeS(errors.New("test error")).Message())

	vtest.Equal(t, 10110, verror.GetCodeS(verror.ErrParamBind).Code())
	vtest.Equal(t, "param bind error", verror.GetCodeS(verror.ErrParamBind).Message())

	err := verror.Wrap(errors.New("test error"), vcode.CodeErrUnknown)
	err = verror.Wrap(err, vcode.CodeErrParamBind, "test data", vmap.M{"a": 1, "b": 2})

	vtest.Equal(t, 10110, verror.GetCodeS(err).Code())
	vtest.Equal(t, "param bind error", verror.GetCodeS(err).Message())
}

func TestRecoverException(t *testing.T) {
	defer verror.RecoverException("TestRecoverException", 0)
	num := 0

	vtest.NotNil(t, 1/num)
}

func TestWrap(t *testing.T) {
	err := verror.Wrap(nil, vcode.CodeErrUnknown)

	vtest.Nil(t, err)

	err = verror.Wrap(errors.New("test error"), vcode.CodeErrUnknown)
	err = verror.Wrap(err, vcode.CodeErrParamBind, "test data", vmap.M{"a": 1, "b": 2})

	vtest.Equal(t, "10110|param bind error|test data map[a:1 b:2]  \n-> 10000|unknown error|<nil> \n-> test error", err.Error())
}
