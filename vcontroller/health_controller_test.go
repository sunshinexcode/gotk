package vcontroller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcontroller"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestHealthControllerCheck(t *testing.T) {
	g := vapi.Default()
	_ = vcontroller.NewHealthController(vcontroller.HealthControllerParam{Engine: g})

	req := httptest.NewRequest(http.MethodGet, "/health/check", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	vtest.Equal(t, http.StatusOK, w.Code)
	vtest.Equal(t, int64(0), vjson.Parse(w.Body.String()).Get("code").Int())
}
