package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sunshinexcode/gotk/vcache"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vjson"
	"github.com/sunshinexcode/gotk/vstruct"

	"app/configs"
	"app/internal/req"
	"app/internal/resp"
	"app/internal/thirdparty"
)

var _ ICompanyService = (*CompanyService)(nil)

type ICompanyService interface {
	Query(ctx context.Context, req *req.CompanyQueryReq) (response *resp.CompanyQueryResp, err error)
	QueryCache(ctx context.Context, req *req.CompanyQueryReq) (response *resp.CompanyQueryResp, err error)
	QueryLocalCache(ctx context.Context, req *req.CompanyQueryReq) (response *resp.CompanyQueryResp, err error)
	Verify(ctx context.Context, cid int64) (err error)
}

type CompanyServiceParam struct {
	vfx.In

	Config *configs.Config

	ILocalCache vcache.ILocalCache
	IRedisCache vcache.IRedisCache

	IConsoleThirdParty thirdparty.IConsoleThirdParty
}

type CompanyService struct {
	Config *configs.Config

	ILocalCache vcache.ILocalCache
	IRedisCache vcache.IRedisCache

	IConsoleThirdParty thirdparty.IConsoleThirdParty
}

const (
	CompanyServiceQueryLocalCacheKey      = "CompanyServiceQuery"
	CompanyServiceQueryLocalCacheDuration = 1 * time.Hour

	CompanyServiceQueryRedisCacheKey      = "CompanyServiceQuery"
	CompanyServiceQueryRedisCacheDuration = 24 * time.Hour
)

func NewCompanyService(p CompanyServiceParam) ICompanyService {
	return &CompanyService{Config: p.Config, ILocalCache: p.ILocalCache, IRedisCache: p.IRedisCache, IConsoleThirdParty: p.IConsoleThirdParty}
}

func (service *CompanyService) Query(ctx context.Context, req *req.CompanyQueryReq) (response *resp.CompanyQueryResp, err error) {
	response = &resp.CompanyQueryResp{}

	companyInfo, err := service.IConsoleThirdParty.GetBasicInfoByCid(ctx, req.Id)
	if err != nil {
		return
	}

	if err = vstruct.Copy(response, companyInfo); err != nil {
		return
	}

	return
}

func (service *CompanyService) QueryCache(ctx context.Context, req *req.CompanyQueryReq) (response *resp.CompanyQueryResp, err error) {
	response = &resp.CompanyQueryResp{}

	data, err := service.IRedisCache.Get(ctx, CompanyServiceQueryRedisCacheKey)

	// Data empty
	if err != nil || vcache.CheckDataEmpty(data) {
		// Get data
		if response, err = service.Query(ctx, req); err != nil {
			return
		}

		// Set cache
		if data, err = vjson.Encode(response); err == nil {
			_, _ = service.IRedisCache.Set(ctx, CompanyServiceQueryRedisCacheKey, data, CompanyServiceQueryLocalCacheDuration)
		}

		return response, nil
	}

	// Decode data
	if err = json.Unmarshal([]byte(data), response); err != nil {
		return response, verror.Wrap(err, vcode.CodeErrJsonDecode, data, response)
	}

	return
}

func (service *CompanyService) QueryLocalCache(ctx context.Context, req *req.CompanyQueryReq) (response *resp.CompanyQueryResp, err error) {
	response = &resp.CompanyQueryResp{}

	data, err := service.ILocalCache.Get(ctx, CompanyServiceQueryLocalCacheKey)

	// Data empty
	if err != nil || data.IsEmpty() {
		// Get data
		if response, err = service.QueryCache(ctx, req); err != nil {
			return
		}

		// Set cache
		_ = service.ILocalCache.Set(ctx, CompanyServiceQueryLocalCacheKey, response, CompanyServiceQueryLocalCacheDuration)

		return response, nil
	}

	// Decode data
	if err = data.Struct(response); err != nil {
		return response, verror.Wrap(err, vcode.CodeErrStructDecode, data, response)
	}

	return
}

func (service *CompanyService) Verify(ctx context.Context, cid int64) (err error) {
	return
}
