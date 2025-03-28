package vmiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/voutput"
	"github.com/sunshinexcode/gotk/vtime"
)

type (
	Expired struct {
		ExpiredSecond  int64
		ExpiredKey     string
		GetExpiredFunc GetExpiredFunc
	}

	GetExpiredFunc func(c *vapi.Context, s *Expired, expiredReq int64) int64
)

var DefaultExpired = NewExpired()

// NewExpired create expired object
func NewExpired() (e *Expired) {
	e = &Expired{
		ExpiredSecond:  10 * 60,
		ExpiredKey:     "ts",
		GetExpiredFunc: GetExpired,
	}
	return
}

// SetExpiredKey set secret
func (e *Expired) SetExpiredKey(expiredKey string) *Expired {
	e.ExpiredKey = expiredKey
	return e
}

// SetGetExpiredFunc set GetExpiredFunc
func (e *Expired) SetGetExpiredFunc(getExpiredFunc GetExpiredFunc) *Expired {
	e.GetExpiredFunc = getExpiredFunc
	return e
}

// GetExpired get expired
func GetExpired(c *vapi.Context, s *Expired, expiredReq int64) (expired int64) {
	var existed bool
	var expiredStr string

	if expiredStr, existed = c.GetQuery(s.ExpiredKey); existed {
		return vconv.Int64(expiredStr)
	}
	if expiredStr, existed = c.GetPostForm(s.ExpiredKey); existed {
		return vconv.Int64(expiredStr)
	}

	return expiredReq
}

// ExpiredMiddleware check expired
func ExpiredMiddleware[T any](expired *Expired) vapi.HandlerFunc {
	return func(c *vapi.Context) {
		var r T
		var expiredReq int64

		if c.GetHeader("Content-Type") == vapi.MimeJson {
			_ = c.ShouldBindBodyWith(&r, binding.JSON)
		} else {
			_ = c.Bind(&r)
		}

		m := vconv.Map(r)
		if c.GetHeader("Content-Type") == vapi.MimeJson {
			if expiredM, ok := m[expired.ExpiredKey].(int64); ok {
				expiredReq = expiredM
			}
		}

		now := vtime.Timestamp()
		ts := vconv.Int64(expired.GetExpiredFunc(c, expired, expiredReq))
		if now-ts > expired.ExpiredSecond || ts > now {
			voutput.E(c, verror.ErrRequestExpired, http.StatusRequestTimeout).Abort()
			return
		}

		c.Next()
	}
}
