package service_test

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vcache"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vmock"
	"github.com/sunshinexcode/gotk/vtest"

	"app/configs"
	"app/internal/entity"
	"app/internal/service"
)

func initCompanyCronService() (s *service.CompanyCronService, patches []*vmock.Patches) {
	s = &service.CompanyCronService{Config: configs.GetConfig(), Metric: &vmetric.Metric{}, IRedisCache: &vcache.RedisCache{}, ICompanyService: &service.CompanyService{}}
	_, patches = vmetric.Mock()

	return
}

func TestNewCompanyCronService(t *testing.T) {
	s := service.NewCompanyCronService(service.CompanyCronServiceParam{})

	vtest.Equal(t, "*service.CompanyCronService", reflect.TypeOf(s).String())
}

func TestCompanyCronServiceCreateTaskForVerifyMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	list1 := make([]*entity.Company, 0)
	for i := 1; i <= service.CreateTaskForVerifyPageSize+1; i++ {
		list1 = append(list1, &entity.Company{Id: int64(i)})
	}

	list2 := make([]*entity.Company, 0)
	list2 = append(list2, &entity.Company{Id: 100})

	patchGetIdsByPage := vmock.ApplyMethodSeq(reflect.TypeOf(s.CompanyModel), "GetIdsByPage", []vmock.OutputCell{{Values: vmock.Params{list1, nil}}, {Values: vmock.Params{list2, nil}}})
	defer patchGetIdsByPage.Reset()

	patchSAdd := vmock.ApplyMethodReturn(s.IRedisCache, "SAdd", int64(1), nil)
	defer patchSAdd.Reset()

	vtest.Nil(t, s.CreateTaskForVerify(context.TODO()))
}

func TestCompanyCronServiceCreateTaskForVerifyErrorMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchGetIdsByPage := vmock.ApplyMethodReturn(s.CompanyModel, "GetIdsByPage", nil, verror.ErrDbOperation)
	defer patchGetIdsByPage.Reset()

	err := s.CreateTaskForVerify(context.TODO())

	vtest.NotNil(t, err)
	vtest.Equal(t, "10120|database operation error|<nil>", err.Error())
}

func TestCompanyCronServiceCronCreateTaskForVerifyMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchCronPatternCreateTaskForVerify := vmock.ApplyGlobalVar(&s.Config.AppCustom.CronPatternCreateTaskForVerify, "*/1 * * * * *")
	defer patchCronPatternCreateTaskForVerify.Reset()

	patchUnlock := vmock.ApplyMethodReturn(s.IRedisCache, "Unlock", int64(0), nil)
	defer patchUnlock.Reset()

	patchLock := vmock.ApplyMethodReturn(s.IRedisCache, "Lock", true, nil)
	defer patchLock.Reset()

	patchCreateTaskForVerify := vmock.ApplyMethodReturn(s, "CreateTaskForVerify", nil)
	defer patchCreateTaskForVerify.Reset()

	s.CronCreateTaskForVerify()
	time.Sleep(2 * time.Second)
}

func TestCompanyCronServiceCronCreateTaskForVerifyErrorRedisLockMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchCronPatternCreateTaskForVerify := vmock.ApplyGlobalVar(&s.Config.AppCustom.CronPatternCreateTaskForVerify, "*/1 * * * * *")
	defer patchCronPatternCreateTaskForVerify.Reset()

	patchUnlock := vmock.ApplyMethodReturn(s.IRedisCache, "Unlock", int64(0), nil)
	defer patchUnlock.Reset()

	patchLock := vmock.ApplyMethodReturn(s.IRedisCache, "Lock", false, verror.ErrRedisAcquireLockFailed)
	defer patchLock.Reset()

	patchCreateTaskForVerify := vmock.ApplyMethodReturn(s, "CreateTaskForVerify", nil)
	defer patchCreateTaskForVerify.Reset()

	s.CronCreateTaskForVerify()
	time.Sleep(2 * time.Second)
}

func TestCompanyCronServiceCronProcessTaskForVerifyMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchCronPatternProcessTaskForVerify := vmock.ApplyGlobalVar(&s.Config.AppCustom.CronPatternProcessTaskForVerify, "*/1 * * * * *")
	defer patchCronPatternProcessTaskForVerify.Reset()

	patchProcessTaskForVerify := vmock.ApplyMethodReturn(s, "ProcessTaskForVerify", nil)
	defer patchProcessTaskForVerify.Reset()

	s.CronProcessTaskForVerify()
	time.Sleep(2 * time.Second)
}

func TestCompanyCronServiceCronProcessTaskForVerifyErrorMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchCronPatternProcessTaskForVerify := vmock.ApplyGlobalVar(&s.Config.AppCustom.CronPatternProcessTaskForVerify, "*/1 * * * * *")
	defer patchCronPatternProcessTaskForVerify.Reset()

	patchProcessTaskForVerify := vmock.ApplyMethodReturn(s, "ProcessTaskForVerify", verror.ErrRedisOperation)
	defer patchProcessTaskForVerify.Reset()

	s.CronProcessTaskForVerify()
	time.Sleep(2 * time.Second)
}

func TestCompanyCronServiceProcessTaskForVerifyMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchSPop := vmock.ApplyMethodSeq(reflect.TypeOf(s.IRedisCache), "SPop", []vmock.OutputCell{{Values: vmock.Params{"1", nil}}, {Values: vmock.Params{"2", verror.ErrRedisOperation}}, {Values: vmock.Params{"", nil}}})
	defer patchSPop.Reset()

	patchProcessTaskForVerifySingle := vmock.ApplyMethodReturn(s, "ProcessTaskForVerifySingle", nil)
	defer patchProcessTaskForVerifySingle.Reset()

	vtest.Nil(t, s.ProcessTaskForVerify(context.TODO()))
}

func TestCompanyCronServiceProcessTaskForVerifyErrorMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchSPop := vmock.ApplyMethodSeq(reflect.TypeOf(s.IRedisCache), "SPop", []vmock.OutputCell{{Values: vmock.Params{"1", nil}}, {Values: vmock.Params{"", nil}}})
	defer patchSPop.Reset()

	patchProcessTaskForVerifySingle := vmock.ApplyMethodReturn(s, "ProcessTaskForVerifySingle", verror.ErrDbOperation)
	defer patchProcessTaskForVerifySingle.Reset()

	vtest.Nil(t, s.ProcessTaskForVerify(context.TODO()))
}

func TestCompanyCronServiceProcessTaskForVerifySingleMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchVerify := vmock.ApplyMethodReturn(s.ICompanyService, "Verify", nil)
	defer patchVerify.Reset()

	vtest.Nil(t, s.ProcessTaskForVerifySingle(context.TODO(), 1))
}

func TestCompanyCronServiceProcessTaskForVerifySingleErrorMock(t *testing.T) {
	s, patches := initCompanyCronService()
	defer vmock.ResetMock(patches)

	patchVerify := vmock.ApplyMethodReturn(s.ICompanyService, "Verify", errors.New("verify error"))
	defer patchVerify.Reset()

	err := s.ProcessTaskForVerifySingle(context.TODO(), 1)

	vtest.NotNil(t, err)
	vtest.Equal(t, "verify error", err.Error())
}
