package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vtest"

	"app/internal/utils"
)

func TestConst(t *testing.T) {
	vtest.Equal(t, "Basic ", utils.Authorization)
}

func TestRequest(t *testing.T) {
	url, resp, err := utils.Request("/test", nil)

	vtest.Nil(t, err)
	vtest.Equal(t, "http://localhost:8080/test", url)
	vtest.Equal(t, http.StatusNotFound, resp.StatusCode())
}

func TestSetHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	utils.SetHeader(req)

	vtest.Equal(t, utils.Authorization, req.Header.Get("Authorization"))
}
