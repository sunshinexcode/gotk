package verror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/vdebug"
	"github.com/sunshinexcode/gotk/vlog"
)

var (
	ErrUnknown                = NewError(vcode.CodeErrUnknown)
	ErrParamInvalid           = NewError(vcode.CodeErrParamInvalid)
	ErrRequestExpired         = NewError(vcode.CodeErrRequestExpired)
	ErrSignInvalid            = NewError(vcode.CodeErrSignInvalid)
	ErrRateLimit              = NewError(vcode.CodeErrRateLimit)
	ErrParamBind              = NewError(vcode.CodeErrParamBind)
	ErrDataCopy               = NewError(vcode.CodeErrDataCopy)
	ErrJsonDecode             = NewError(vcode.CodeErrJsonDecode)
	ErrMapDecode              = NewError(vcode.CodeErrMapDecode)
	ErrStructDecode           = NewError(vcode.CodeErrStructDecode)
	ErrDbOperation            = NewError(vcode.CodeErrDbOperation)
	ErrDbRecordNotFound       = NewError(vcode.CodeErrDbRecordNotFound)
	ErrHttpRequest            = NewError(vcode.CodeErrHttpRequest)
	ErrHttpStatusNotOk        = NewError(vcode.CodeErrHttpStatusNotOk)
	ErrRedisOperation         = NewError(vcode.CodeErrRedisOperation)
	ErrRedisAcquireLockFailed = NewError(vcode.CodeErrRedisAcquireLockFailed)
	ErrCronStartFailed        = NewError(vcode.CodeErrCronStartFailed)
	ErrCloseConn              = NewError(vcode.CodeErrCloseConn)
	ErrInvalidData            = NewError(vcode.CodeErrInvalidData)
	ErrProtoUnmarshal         = NewError(vcode.CodeErrProtoUnmarshal)
)

type Error struct {
	code *vcode.Code
}

func NewError(code *vcode.Code) error {
	return &Error{code: code}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d|%s|%+v", e.code.Code(), e.code.Message(), e.code.Data())
}

func (e *Error) Is(target error) bool {
	if t, ok := target.(*Error); ok {
		return e.code.Code() == t.code.Code()
	}
	return false
}

func GetCode(err error) (code int, message string) {
	codeS := GetCodeS(err)

	return codeS.Code(), codeS.Message()
}

func GetCodeS(err error) *vcode.Code {
	if err == nil {
		return vcode.CodeSuccess
	}

	var e *Error
	if errors.As(err, &e) {
		return e.code
	}

	return vcode.CodeErrUnknown
}

// RecoverException recover exception
func RecoverException(tag string, params ...interface{}) {
	if err := recover(); err != nil {
		vlog.Errorf("RecoverException, tag:%s, error:%s, params:%v, stack:%s", tag, err, params, vdebug.Stack())
	}
}

func Wrap(err error, code *vcode.Code, data ...any) error {
	if err == nil {
		return nil
	}

	if len(data) > 0 {
		code.SetData(fmt.Sprintf(strings.Repeat("%+v ", len(data)), data...))
	}

	return fmt.Errorf("%w \n-> %w", NewError(code), err)
}
