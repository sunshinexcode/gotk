package service_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/sunshinexcode/gotk/vcache"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vstruct"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vvar"

	"app/configs"
	"app/internal/req"
	"app/internal/resp"
	"app/internal/service"
	"app/internal/thirdparty"
)

func initCompanyService() (s *service.CompanyService) {
	s = &service.CompanyService{Config: configs.GetConfig(), ILocalCache: &vcache.LocalCache{}, IRedisCache: &vcache.RedisCache{}, IConsoleThirdParty: &thirdparty.ConsoleThirdParty{}}

	return
}

func TestNewCompanyService(t *testing.T) {
	s := service.NewCompanyService(service.CompanyServiceParam{Config: configs.GetConfig(), IConsoleThirdParty: &thirdparty.ConsoleThirdParty{}})

	vtest.Equal(t, "*service.CompanyService", reflect.TypeOf(s).String())
}

func TestCompanyServiceQueryMock(t *testing.T) {
	s := initCompanyService()

	patch := vmock.ApplyMethodReturn(s.IConsoleThirdParty, "GetBasicInfoByCid", &resp.ThirdPartyConsoleGetBasicInfoByCidResp{CompanyName: "test"}, nil)
	defer vmock.Reset(patch)

	response, err := s.Query(context.TODO(), &req.CompanyQueryReq{})

	vtest.Nil(t, err)
	vtest.Equal(t, "test", response.CompanyName)
}

func TestCompanyServiceQueryErrorGetBasicInfoByCidMock(t *testing.T) {
	s := initCompanyService()

	patch := vmock.ApplyMethodReturn(s.IConsoleThirdParty, "GetBasicInfoByCid", &resp.ThirdPartyConsoleGetBasicInfoByCidResp{}, verror.ErrDbOperation)
	defer vmock.Reset(patch)

	response, err := s.Query(context.TODO(), &req.CompanyQueryReq{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "10120|database operation error|<nil>", err.Error())
	vtest.Equal(t, "", response.CompanyName)
}

func TestCompanyServiceQueryErrorCopyMock(t *testing.T) {
	s := initCompanyService()

	patchGetBasicInfoByCid := vmock.ApplyMethodReturn(s.IConsoleThirdParty, "GetBasicInfoByCid", &resp.ThirdPartyConsoleGetBasicInfoByCidResp{}, nil)
	defer patchGetBasicInfoByCid.Reset()

	patchCopy := vmock.ApplyFuncReturn(vstruct.Copy, verror.ErrDataCopy)
	defer patchCopy.Reset()

	response, err := s.Query(context.TODO(), &req.CompanyQueryReq{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "10111|data copy error|<nil>", err.Error())
	vtest.Equal(t, "", response.CompanyName)
}

func TestCompanyServiceQueryCacheMock(t *testing.T) {
	s := initCompanyService()

	patch := vmock.ApplyMethodReturn(s.IRedisCache, "Get", `{"id": 1}`, nil)
	defer vmock.Reset(patch)

	response, err := s.QueryCache(context.TODO(), &req.CompanyQueryReq{})

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), response.Id)
}

func TestCompanyServiceQueryCacheDataEmptyMock(t *testing.T) {
	s := initCompanyService()

	patchGet := vmock.ApplyMethodReturn(s.IRedisCache, "Get", "", nil)
	defer patchGet.Reset()

	patchQuery := vmock.ApplyMethodReturn(s, "Query", &resp.CompanyQueryResp{Id: 1}, nil)
	defer patchQuery.Reset()

	patchSet := vmock.ApplyMethodReturn(s.IRedisCache, "Set", "", nil)
	defer patchSet.Reset()

	response, err := s.QueryCache(context.TODO(), &req.CompanyQueryReq{})

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), response.Id)
}

func TestCompanyServiceQueryCacheErrorQueryMock(t *testing.T) {
	s := initCompanyService()

	patchGet := vmock.ApplyMethodReturn(s.IRedisCache, "Get", "", nil)
	defer patchGet.Reset()

	patchQuery := vmock.ApplyMethodReturn(s, "Query", &resp.CompanyQueryResp{Id: 1}, verror.ErrDbOperation)
	defer patchQuery.Reset()

	response, err := s.QueryCache(context.TODO(), &req.CompanyQueryReq{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "10120|database operation error|<nil>", err.Error())
	vtest.Equal(t, int64(1), response.Id)
}

func TestCompanyServiceQueryCacheErrorJsonUnmarshalMock(t *testing.T) {
	s := initCompanyService()

	patchGet := vmock.ApplyMethodReturn(s.IRedisCache, "Get", "id:1", nil)
	defer patchGet.Reset()

	patchQuery := vmock.ApplyMethodReturn(s, "Query", &resp.CompanyQueryResp{Id: 1}, nil)
	defer patchQuery.Reset()

	response, err := s.QueryCache(context.TODO(), &req.CompanyQueryReq{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "10112|json decode error|id:1 &{CompanyName: Id:0}  \n-> invalid character 'i' looking for beginning of value", err.Error())
	vtest.Equal(t, int64(0), response.Id)
}

func TestCompanyServiceQueryLocalCacheMock(t *testing.T) {
	s := initCompanyService()

	patch := vmock.ApplyMethodReturn(s.ILocalCache, "Get", vvar.New(`{"id": 1}`), nil)
	defer vmock.Reset(patch)

	response, err := s.QueryLocalCache(context.TODO(), &req.CompanyQueryReq{})

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), response.Id)
}

func TestCompanyServiceQueryLocalCacheDataEmptyMock(t *testing.T) {
	s := initCompanyService()

	patchGet := vmock.ApplyMethodReturn(s.ILocalCache, "Get", vvar.New(""), nil)
	defer patchGet.Reset()

	patchQueryCache := vmock.ApplyMethodReturn(s, "QueryCache", &resp.CompanyQueryResp{Id: 1}, nil)
	defer patchQueryCache.Reset()

	patchSet := vmock.ApplyMethodReturn(s.ILocalCache, "Set", nil)
	defer patchSet.Reset()

	response, err := s.QueryLocalCache(context.TODO(), &req.CompanyQueryReq{})

	vtest.Nil(t, err)
	vtest.Equal(t, int64(1), response.Id)
}

func TestCompanyServiceQueryLocalCacheErrorQueryCacheMock(t *testing.T) {
	s := initCompanyService()

	patchGet := vmock.ApplyMethodReturn(s.ILocalCache, "Get", vvar.New(""), errors.New("get error"))
	defer patchGet.Reset()

	patchQueryCache := vmock.ApplyMethodReturn(s, "QueryCache", &resp.CompanyQueryResp{}, verror.ErrRedisOperation)
	defer patchQueryCache.Reset()

	response, err := s.QueryLocalCache(context.TODO(), &req.CompanyQueryReq{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "10140|redis operation error|<nil>", err.Error())
	vtest.Equal(t, int64(0), response.Id)
}

func TestCompanyServiceQueryLocalCacheErrorStructMock(t *testing.T) {
	s := initCompanyService()

	patchGet := vmock.ApplyMethodReturn(s.ILocalCache, "Get", vvar.New("id:1"), nil)
	defer patchGet.Reset()

	response, err := s.QueryLocalCache(context.TODO(), &req.CompanyQueryReq{})

	vtest.NotNil(t, err)
	vtest.Equal(t, "10114|struct decode error|id:1 &{CompanyName: Id:0}  \n-> convert params from \"\"id:1\"\" to \"map[string]interface{}\" failed", err.Error())
	vtest.Equal(t, int64(0), response.Id)
}

func TestCompanyServiceVerify(t *testing.T) {
	s := initCompanyService()

	vtest.Nil(t, s.Verify(context.TODO(), 1))
}
