package vmiddleware

import (
	"context"

	"github.com/gin-gonic/gin/binding"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vvalid"
)

// ValidMiddleware valid params
func ValidMiddleware[T any](errCode ...*vcode.Code) vapi.HandlerFunc {
	return func(c *vapi.Context) {
		var r T

		if c.GetHeader("Content-Type") == vapi.MimeJson {
			_ = c.ShouldBindBodyWith(&r, binding.JSON)
		} else {
			_ = c.Bind(&r)
		}

		err := vvalid.New().Data(r).Run(context.TODO())
		if err != nil {
			if len(errCode) == 0 {
				errCode = append(errCode, vcode.CodeErrParamInvalid)
			}

			errCode[0].SetMessage(err.FirstError().Error())
			voutput.E(c, verror.NewError(errCode[0])).Abort()
			return
		}

		c.Next()
	}
}
