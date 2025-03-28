package vcode

import "github.com/sunshinexcode/gotk/vconv"

var (
	CodeOk      = NewCode(0, "ok", nil)
	CodeSuccess = NewCode(0, "success", nil)

	CodeErrUnknown                = NewCode(10000, "unknown error", nil)
	CodeErrParamInvalid           = NewCode(10101, "param invalid", nil)
	CodeErrRequestExpired         = NewCode(10102, "request expired", nil)
	CodeErrSignInvalid            = NewCode(10103, "sign invalid", nil)
	CodeErrRateLimit              = NewCode(10104, "exceed request", nil)
	CodeErrParamBind              = NewCode(10110, "param bind error", nil)
	CodeErrDataCopy               = NewCode(10111, "data copy error", nil)
	CodeErrJsonDecode             = NewCode(10112, "json decode error", nil)
	CodeErrMapDecode              = NewCode(10113, "map decode error", nil)
	CodeErrStructDecode           = NewCode(10114, "struct decode error", nil)
	CodeErrDbOperation            = NewCode(10120, "database operation error", nil)
	CodeErrDbRecordNotFound       = NewCode(10121, "database record not found", nil)
	CodeErrHttpRequest            = NewCode(10130, "http request error", nil)
	CodeErrHttpStatusNotOk        = NewCode(10131, "http status not 200", nil)
	CodeErrRedisOperation         = NewCode(10140, "redis operation error", nil)
	CodeErrRedisAcquireLockFailed = NewCode(10141, "redis acquire lock failed", nil)
	CodeErrCronStartFailed        = NewCode(10150, "cron start failed", nil)
	CodeErrCloseConn              = NewCode(10160, "close connection error", nil)
	CodeErrInvalidData            = NewCode(10161, "invalid data", nil)
	CodeErrProtoUnmarshal         = NewCode(10162, "proto unmarshal error", nil)

	// Business code, 20000 - 29999
)

type Code struct {
	code    int
	message string
	data    any
}

func NewCode(code int, message string, data any) *Code {
	return &Code{code: code, message: message, data: data}
}

func (c *Code) Code() int {
	return c.code
}

func (c *Code) CodeStr() string {
	return vconv.String(c.code)
}

func (c *Code) Message() string {
	return c.message
}

func (c *Code) Data() any {
	return c.data
}

func (c *Code) SetData(data any) {
	c.data = data
}

func (c *Code) SetMessage(message string) {
	c.message = message
}
