package service

import (
	"context"
	"time"

	"github.com/sunshinexcode/gotk/vcache"
	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/vcron"
	"github.com/sunshinexcode/gotk/verror"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmetric"
	"github.com/sunshinexcode/gotk/vtime"
	"github.com/sunshinexcode/gotk/vtrace"

	"app/configs"
	"app/internal/entity"
	"app/internal/model"
)

var _ ICompanyCronService = (*CompanyCronService)(nil)

type ICompanyCronService interface {
	CreateTaskForVerify(ctx context.Context) (err error)
	CronCreateTaskForVerify()
	CronProcessTaskForVerify()
	ProcessTaskForVerify(ctx context.Context) (err error)
	ProcessTaskForVerifySingle(ctx context.Context, cid int64) (err error)
}

type CompanyCronServiceParam struct {
	vfx.In

	Config      *configs.Config
	Metric      *vmetric.Metric
	IRedisCache vcache.IRedisCache

	CompanyModel *model.CompanyModel

	ICompanyService ICompanyService
}

type CompanyCronService struct {
	Config      *configs.Config
	Metric      *vmetric.Metric
	IRedisCache vcache.IRedisCache

	CompanyModel *model.CompanyModel

	ICompanyService ICompanyService
}

const (
	CreateTaskForVerifyPageSize      = 100
	CreateTaskForVerifyRedisCacheKey = "CreateTaskForVerifyRedisCacheKey"

	CronCreateTaskForVerifyRedisLockName     = "lock-CronCreateTaskForVerify"
	CronCreateTaskForVerifyRedisLockDuration = 5 * time.Minute
)

func NewCompanyCronService(p CompanyCronServiceParam) ICompanyCronService {
	return &CompanyCronService{Config: p.Config, Metric: p.Metric, IRedisCache: p.IRedisCache, CompanyModel: p.CompanyModel, ICompanyService: p.ICompanyService}
}

func (service *CompanyCronService) CreateTaskForVerify(ctx context.Context) (err error) {
	vlog.Infoc(ctx, "CreateTaskForVerify start", "pageSize", CreateTaskForVerifyPageSize)

	var list []*entity.Company
	lastId := int64(0)
	count := 0

	for {
		list, err = service.CompanyModel.GetIdsByPage(ctx, CreateTaskForVerifyPageSize, &entity.Company{Id: lastId})
		count = len(list)

		vlog.Infoc(ctx, "CreateTaskForVerify GetIdsByPage", "err", err, "lastId", lastId, "count", count)

		if err != nil {
			return
		}

		for _, item := range list {
			num, err := service.IRedisCache.SAdd(ctx, CreateTaskForVerifyRedisCacheKey, item.Id)
			vlog.Infoc(ctx, "CreateTaskForVerify SAdd", "err", err, "id", item.Id, "num", num)
			vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "CreateTaskForVerify-SAdd", verror.GetCodeS(err).CodeStr())
		}

		if count < CreateTaskForVerifyPageSize {
			vlog.Infoc(ctx, "CreateTaskForVerify no more data break", "lastId", lastId, "count", count)
			break
		}

		lastId = list[count-1].Id
	}

	vlog.Infoc(ctx, "CreateTaskForVerify end", "err", err, "lastId", lastId)
	return
}

func (service *CompanyCronService) CronCreateTaskForVerify() {
	var err error

	pattern := service.Config.AppCustom.CronPatternCreateTaskForVerify
	vlog.Info("CronCreateTaskForVerify start", "pattern", pattern)

	defer func() {
		if err != nil {
			vlog.Error("CronCreateTaskForVerify", "err", err, "pattern", pattern)
			vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "CronCreateTaskForVerify-start", vcode.CodeErrCronStartFailed.CodeStr())
		}
	}()

	_, err = vcron.Add(context.TODO(), pattern, func(ctx context.Context) {
		var err error

		timeStart := vtime.GetNow()
		ctx = vtrace.SetTraceId(ctx, configs.CronTraceIdPrefix)

		vlog.Infoc(ctx, "CronCreateTaskForVerify run start", "pattern", pattern)

		defer func() {
			codeS := verror.GetCodeS(err)

			if err != nil {
				vlog.Errorc(ctx, "CronCreateTaskForVerify", "err", err, "code", codeS.CodeStr())
			}

			// Release lock
			_, _ = service.IRedisCache.Unlock(ctx, CronCreateTaskForVerifyRedisLockName)

			vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "CronCreateTaskForVerify-run", codeS.CodeStr())
			vmetric.MetricHttpRequestDurationTypeCron(service.Metric, timeStart, "CronCreateTaskForVerify-run", codeS.CodeStr())

			vlog.Infoc(ctx, "CronCreateTaskForVerify run end", "pattern", pattern)
		}()

		// Acquire lock
		if _, err = service.IRedisCache.Lock(ctx, CronCreateTaskForVerifyRedisLockName, CronCreateTaskForVerifyRedisLockDuration); err != nil {
			return
		}

		err = service.CreateTaskForVerify(ctx)
	}, "CronCreateTaskForVerify")

	if err != nil {
		return
	}

	vlog.Info("CronCreateTaskForVerify end", "pattern", pattern)
	vcron.Start()
}

func (service *CompanyCronService) CronProcessTaskForVerify() {
	var err error

	pattern := service.Config.AppCustom.CronPatternProcessTaskForVerify
	vlog.Info("CronProcessTaskForVerify start", "pattern", pattern)

	defer func() {
		if err != nil {
			vlog.Error("CronProcessTaskForVerify", "err", err, "pattern", pattern)
			vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "CronProcessTaskForVerify-start", vcode.CodeErrCronStartFailed.CodeStr())
		}
	}()

	_, err = vcron.Add(context.TODO(), pattern, func(ctx context.Context) {
		var err error

		timeStart := vtime.GetNow()
		ctx = vtrace.SetTraceId(ctx, configs.CronTraceIdPrefix)

		vlog.Infoc(ctx, "CronProcessTaskForVerify run start", "pattern", pattern)

		defer func() {
			codeS := verror.GetCodeS(err)

			if err != nil {
				vlog.Errorc(ctx, "CronProcessTaskForVerify", "err", err, "code", codeS.CodeStr())
			}

			vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "CronProcessTaskForVerify-run", codeS.CodeStr())
			vmetric.MetricHttpRequestDurationTypeCron(service.Metric, timeStart, "CronProcessTaskForVerify-run", codeS.CodeStr())

			vlog.Infoc(ctx, "CronProcessTaskForVerify run end", "pattern", pattern)
		}()

		err = service.ProcessTaskForVerify(ctx)
	}, "CronProcessTaskForVerify")

	if err != nil {
		return
	}

	vlog.Info("CronProcessTaskForVerify end", "pattern", pattern)
	vcron.Start()
}

func (service *CompanyCronService) ProcessTaskForVerify(ctx context.Context) (err error) {
	vlog.Infoc(ctx, "ProcessTaskForVerify start")

	for {
		cidStr, err := service.IRedisCache.SPop(ctx, CreateTaskForVerifyRedisCacheKey)
		vlog.Infoc(ctx, "ProcessTaskForVerify SPop", "err", err, "cidStr", cidStr)

		if err != nil {
			vlog.Errorc(ctx, "ProcessTaskForVerify SPop", "err", err, "cidStr", cidStr)
			vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "ProcessTaskForVerify-SPop", vcode.CodeErrRedisOperation.CodeStr())
			continue
		}

		if vcache.CheckDataEmpty(cidStr) {
			vlog.Infoc(ctx, "ProcessTaskForVerify data empty break")
			vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "ProcessTaskForVerify-dataEmptyBreak", vcode.CodeOk.CodeStr())
			break
		}

		cid := vconv.Int64(cidStr)
		if err = service.ProcessTaskForVerifySingle(ctx, cid); err != nil {
			vlog.Errorc(ctx, "ProcessTaskForVerify ProcessTaskForVerifySingle", "err", err, "cid", cid)
			vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "ProcessTaskForVerify-ProcessTaskForVerifySingle", verror.GetCodeS(err).CodeStr())
		}
	}

	vlog.Infoc(ctx, "ProcessTaskForVerify end")
	return
}

func (service *CompanyCronService) ProcessTaskForVerifySingle(ctx context.Context, cid int64) (err error) {
	vlog.Infoc(ctx, "ProcessTaskForVerifySingle start", "cid", cid)
	timeStart := vtime.GetNow()

	if err = service.ICompanyService.Verify(ctx, cid); err != nil {
		return
	}

	vmetric.MetricHttpRequestTotalTypeCron(service.Metric, "ProcessTaskForVerifySingle", vcode.CodeOk.CodeStr())
	vmetric.MetricHttpRequestDurationTypeCron(service.Metric, timeStart, "ProcessTaskForVerifySingle", vcode.CodeOk.CodeStr())

	vlog.Infoc(ctx, "ProcessTaskForVerifySingle end", "err", err, "cid", cid)
	return
}
