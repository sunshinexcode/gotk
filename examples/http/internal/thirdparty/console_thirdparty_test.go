package thirdparty_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vhttp"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
	"app/internal/thirdparty"
)

func initConsoleThirdParty() (tp thirdparty.IConsoleThirdParty, c *vapi.Context, patches []*vmock.Patches) {
	metric, patches := vmetric.Mock()

	tp = thirdparty.NewConsoleThirdParty(thirdparty.ConsoleThirdPartyParam{Config: configs.GetConfig(), Metric: metric})
	c, _ = vapi.CreateTestContext(httptest.NewRecorder())

	return
}

func TestConsoleThirdPartyGetBasicInfoByCidMock(t *testing.T) {
	tp, c, patches := initConsoleThirdParty()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(&vhttp.Request{}, "Get", &vhttp.Response{RawResponse: &http.Response{StatusCode: 200}}, nil)
	defer vmock.Reset(patch)

	response, err := tp.GetBasicInfoByCid(c, 1)

	vtest.Nil(t, err)
	vtest.Equal(t, "", response.CompanyName)
}

func TestConsoleThirdPartyGetBasicInfoByCidErrorHttpRequestMock(t *testing.T) {
	tp, c, patches := initConsoleThirdParty()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(&vhttp.Request{}, "Get", nil, errors.New("http test error"))
	defer vmock.Reset(patch)

	response, err := tp.GetBasicInfoByCid(c, 1)

	vtest.NotNil(t, err)
	vtest.Equal(t, "", response.CompanyName)
	vtest.Equal(t, true, errors.Is(err, verror.ErrHttpRequest))
}

func TestConsoleThirdPartyGetBasicInfoByCidErrorHttpStatusNotOkMock(t *testing.T) {
	tp, c, patches := initConsoleThirdParty()
	defer vmock.ResetMock(patches)

	patch := vmock.ApplyMethodReturn(&vhttp.Request{}, "Get", &vhttp.Response{RawResponse: &http.Response{StatusCode: 500}}, nil)
	defer vmock.Reset(patch)

	response, err := tp.GetBasicInfoByCid(c, 1)

	vtest.NotNil(t, err)
	vtest.Equal(t, "", response.CompanyName)
	vtest.Equal(t, true, errors.Is(err, verror.ErrHttpRequest))
	vtest.Equal(t, true, errors.Is(err, verror.ErrHttpStatusNotOk))
}
