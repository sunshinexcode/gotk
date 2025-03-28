package vcron_test

import (
	"context"
	"testing"
	"time"

	"github.com/sunshinexcode/gotk/vcron"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestAdd(t *testing.T) {
	_, err := vcron.Add(context.TODO(), "* * * * * *", func(ctx context.Context) {
		vlog.Debug("cron doing 1")
	}, "CronJob")

	vtest.Nil(t, err)

	_, err = vcron.Add(context.TODO(), "*/2 * * * * *", func(ctx context.Context) {
		vlog.Debug("cron doing 2")
	})

	vtest.Nil(t, err)

	_, err = vcron.Add(context.TODO(), "* * * * * *", func(ctx context.Context) {
		vlog.Debug("cron doing 3")
	}, "CronJob")

	vtest.NotNil(t, err)
	vtest.Equal(t, `duplicated cron job name "CronJob", already exists`, err.Error())

	vcron.Start()
	time.Sleep(3 * time.Second)
}

func TestAddSingleton(t *testing.T) {
	_, err := vcron.AddSingleton(context.TODO(), "* * * * * *", func(ctx context.Context) {
		vlog.Debug("cron doing 1")
	}, "CronJobSingleton")

	vtest.Nil(t, err)

	_, err = vcron.AddSingleton(context.TODO(), "@every 2s", func(ctx context.Context) {
		vlog.Debug("cron doing 2")
	})

	vtest.Nil(t, err)
	_, err = vcron.AddSingleton(context.TODO(), "@every 2s", func(ctx context.Context) {
		vlog.Debug("cron doing 3")
	}, "CronJobSingleton")

	vtest.NotNil(t, err)
	vtest.Equal(t, `duplicated cron job name "CronJobSingleton", already exists`, err.Error())

	vcron.Start()
	time.Sleep(3 * time.Second)
}

func TestStop(t *testing.T) {
	_, err := vcron.Add(context.TODO(), "* * * * * *", func(ctx context.Context) {
		vlog.Debug("cron doing 1")
	}, "CronJob2")

	vtest.Nil(t, err)

	vcron.Start()
	time.Sleep(2 * time.Second)
	vcron.Stop()
}
