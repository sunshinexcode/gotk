package vhttp

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	C = NewClient().SetTimeout(5 * time.Second)

	HttpClient = NewClient().
			SetRetryCount(0).
			SetTimeout(5 * time.Second).
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	HttpClientRetry = NewClient().
			SetRetryCount(3).
			SetTimeout(5 * time.Second).
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
)

// NewClient method creates a new Resty client.
func NewClient() *resty.Client {
	return resty.New()
}
