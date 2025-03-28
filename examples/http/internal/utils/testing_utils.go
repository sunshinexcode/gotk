package utils

import (
	"net/http"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vhttp"
	"github.com/sunshinexcode/gotk/vstr"

	"app/configs"
)

const (
	Authorization = "Basic "
)

func Request(url string, body any) (string, *vhttp.Response, error) {
	u := vstr.S("%s%s", configs.TestUrl, url)
	resp, err := vhttp.HttpClient.R().
		SetHeader("Content-Type", vapi.MimeJson).
		SetHeader("Authorization", Authorization).
		SetBody(body).
		Post(u)

	return u, resp, err
}

func SetHeader(req *http.Request) {
	req.Header.Set("Content-Type", vapi.MimeJson)
	req.Header.Set("Authorization", Authorization)
}
