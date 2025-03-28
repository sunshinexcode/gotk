package run

import (
	"net/http"
	"sync"
	"time"

	"github.com/sunshinexcode/gotk/vcmd"
	"github.com/sunshinexcode/gotk/vfx"
	"github.com/sunshinexcode/gotk/vhttp"
	"github.com/sunshinexcode/gotk/vlog"

	"app/configs"
	"app/internal/bootstrap"
)

type AbTest struct {
}

var (
	concurrent int
	interval   int
	url        string

	abCmd = &vcmd.Command{
		Use:   "ab",
		Short: "Start abtest",
		// make run-ab param="--concurrent=10 --interval=100 --url=http://localhost:8080"
		Run: initAbTest,
	}
)

func init() {
	abCmd.Flags().IntVarP(&concurrent, "concurrent", "c", 10, "concurrent")
	abCmd.Flags().IntVarP(&interval, "interval", "i", 100, "interval millisecond")
	abCmd.Flags().StringVarP(&url, "url", "u", "http://localhost:8080", "url")
}

func initAbTest(cmd *vcmd.Command, args []string) {
	vfx.New(
		vfx.NopLogger,
		vfx.Options(
			vfx.Provide(
				configs.NewConfig,
			),
			vfx.Invoke(
				bootstrap.NewLog,
				abtest,
			),
		)).Run()
}

func abtest() {
	var wg sync.WaitGroup

	vlog.Info("abtest start", "concurrent", concurrent, "interval", interval, "url", url)

	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			ab := &AbTest{}
			ab.do()
		}()
	}

	wg.Wait()
	vlog.Info("abtest end")
}

func (ab *AbTest) do() {
	for {
		resp, err := vhttp.C.R().EnableTrace().Get(url)

		if err != nil {
			vlog.Error("Get", "err", err)
			return
		}
		if resp.StatusCode() != http.StatusOK {
			vlog.Error("Get", "StatusCode", resp.StatusCode())
			return
		}

		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
