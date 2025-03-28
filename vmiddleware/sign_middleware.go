package vmiddleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin/binding"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vhmac"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vmd5"
	"github.com/sunshinexcode/gotk/voutput"
)

type (
	SignAlgorithm int

	Sign struct {
		Algorithm   SignAlgorithm
		CalSignFunc CalSignFunc
		GetSignFunc GetSignFunc
		Secret      string
		SignKey     string
	}

	CalSignFunc func(m vmap.M, secret string, signAlgorithm SignAlgorithm) (sign string, err error)
	GetSignFunc func(c *vapi.Context, s *Sign, signReq string) string
)

const (
	SignAlgorithmMd5  = 1
	SignAlgorithmHmac = 2
)

var DefaultSign = NewSign()

// NewSign create sign object
func NewSign() (s *Sign) {
	s = &Sign{
		Algorithm:   SignAlgorithmMd5,
		CalSignFunc: CalSign,
		GetSignFunc: GetSign,
		Secret:      "",
		SignKey:     "sign",
	}
	return
}

// SetAlgorithm set CalSignFunc
func (s *Sign) SetAlgorithm(algorithm SignAlgorithm) *Sign {
	s.Algorithm = algorithm
	return s
}

// SetCalSignFunc set CalSignFunc
func (s *Sign) SetCalSignFunc(calSignFunc CalSignFunc) *Sign {
	s.CalSignFunc = calSignFunc
	return s
}

// SetGetSignFunc set GetSignFunc
func (s *Sign) SetGetSignFunc(getSignFunc GetSignFunc) *Sign {
	s.GetSignFunc = getSignFunc
	return s
}

// SetSecret set secret
func (s *Sign) SetSecret(secret string) *Sign {
	s.Secret = secret
	return s
}

// SetSignKey set secret
func (s *Sign) SetSignKey(signKey string) *Sign {
	s.SignKey = signKey
	return s
}

// GetSign get sign
func GetSign(c *vapi.Context, s *Sign, signReq string) (sign string) {
	var existed bool

	if sign, existed = c.GetQuery(s.SignKey); existed {
		return
	}
	if sign, existed = c.GetPostForm(s.SignKey); existed {
		return
	}

	return signReq
}

// GetSignByHeader get sign by header
func GetSignByHeader(c *vapi.Context, s *Sign, _ string) string {
	return c.GetHeader(s.SignKey)
}

// CalSign calculate sign
func CalSign(m vmap.M, secret string, signAlgorithm SignAlgorithm) (sign string, err error) {
	var buf strings.Builder

	keys := vmap.SortKey(m)

	for _, k := range keys {
		buf.WriteString(fmt.Sprintf("%v", fmt.Sprintf("%v", m[k])))
	}

	switch signAlgorithm {
	case SignAlgorithmMd5:
		buf.WriteString(fmt.Sprintf("%v", secret))
		sign, err = vmd5.Get(buf.String())
	case SignAlgorithmHmac:
		sign = vhmac.Sha256(buf.String(), secret)
	}

	return
}

// SignMiddleware sign
func SignMiddleware[T any](sign *Sign) vapi.HandlerFunc {
	return func(c *vapi.Context) {
		var r T
		var signReq string

		// Get sign
		if c.GetHeader("Content-Type") == vapi.MimeJson {
			_ = c.ShouldBindBodyWith(&r, binding.JSON)
		} else {
			_ = c.Bind(&r)
		}

		m := vconv.Map(r)
		if c.GetHeader("Content-Type") == vapi.MimeJson {
			if signM, ok := m[sign.SignKey].(string); ok {
				signReq = signM
				delete(m, sign.SignKey)
			}
		}

		// Verify sign
		signCal, _ := sign.CalSignFunc(m, sign.Secret, sign.Algorithm)
		if sign.GetSignFunc(c, sign, signReq) != signCal {
			voutput.E(c, verror.ErrSignInvalid, http.StatusUnauthorized).Abort()
			return
		}

		c.Next()
	}
}
