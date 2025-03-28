package thirdparty

import (
	"context"
	"net/http"

	"github.com/sunshinexcode/gotk/vapi"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vhttp"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vtime"

	"app/configs"
	"app/internal/resp"
)

var _ IConsoleThirdParty = (*ConsoleThirdParty)(nil)

type IConsoleThirdParty interface {
	GetBasicInfoByCid(ctx context.Context, cid int64) (response *resp.ThirdPartyConsoleGetBasicInfoByCidResp, err error)
}

type ConsoleThirdPartyParam struct {
	vfx.In

	Config *configs.Config
	Metric *vmetric.Metric
}

type ConsoleThirdParty struct {
	Config *configs.Config
	Metric *vmetric.Metric
}

func NewConsoleThirdParty(p ConsoleThirdPartyParam) IConsoleThirdParty {
	return &ConsoleThirdParty{Config: p.Config, Metric: p.Metric}
}

func (t *ConsoleThirdParty) GetBasicInfoByCid(ctx context.Context, cid int64) (response *resp.ThirdPartyConsoleGetBasicInfoByCidResp, err error) {
	var res *vhttp.Response

	response = &resp.ThirdPartyConsoleGetBasicInfoByCidResp{}
	code := vcode.CodeSuccess
	timeStart := vtime.GetNow()
	url := vstr.S("%s/api/for-package/v1/company/%d/info", t.Config.AppCustom.ConsoleThirdPartyHost, cid)

	defer func() {
		if err != nil {
			vlog.Errorc(ctx, "GetBasicInfoByCid", "err", err, "cid", cid, "url", url, "res", res)
		}

		vmetric.MetricHttpRequestTotalTypeThirdPartyApi(t.Metric, "ConsoleThirdParty-GetBasicInfoByCid", code.CodeStr())
		vmetric.MetricHttpRequestDurationTypeThirdPartyApi(t.Metric, timeStart, "ConsoleThirdParty-GetBasicInfoByCid", code.CodeStr())
	}()

	res, err = vhttp.HttpClient.R().
		SetHeader("Content-Type", vapi.MimeJson).
		SetHeader("Authorization", t.Config.AppCustom.ConsoleThirdPartyAuthorization).
		SetResult(response).
		Get(url)

	if err != nil {
		code = vcode.CodeErrHttpRequest
		return response, verror.Wrap(err, code, url, res)
	}

	if res.StatusCode() != http.StatusOK {
		code = vcode.CodeErrHttpStatusNotOk
		return response, verror.Wrap(verror.ErrHttpRequest, code, url, res, res.StatusCode())
	}

	vlog.Infoc(ctx, "GetBasicInfoByCid", "cid", cid, "url", url)
	return
}
