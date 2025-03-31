package vqueue

import (
	"github.com/gogf/gf/v2/container/gqueue"
)

// New returns an empty queue object.
// Optional parameter `limit` is used to limit the size of the queue, which is unlimited in default.
// When `limit` is given, the queue will be static and high performance which is comparable with stdlib channel.
func New(limit ...int) *Queue {
	return gqueue.New(limit...)
}
