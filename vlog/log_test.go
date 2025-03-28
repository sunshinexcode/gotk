package vlog_test

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"go.uber.org/zap/zapcore"

	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vstr"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vtrace"
)

func TestConst(t *testing.T) {
	vtest.Equal(t, zapcore.Level(-1), vlog.DebugLevel)
	vtest.Equal(t, zapcore.Level(0), vlog.InfoLevel)
	vtest.Equal(t, zapcore.Level(1), vlog.WarnLevel)
	vtest.Equal(t, zapcore.Level(2), vlog.ErrorLevel)
	vtest.Equal(t, zapcore.Level(5), vlog.FatalLevel)

	vtest.Equal(t, "0", vstr.S("%d", vlog.OutputModeConsole))
	vtest.Equal(t, "0", vstr.S("%d", vlog.OutputFormatModeText))
}

func TestNew(t *testing.T) {
	log, err := vlog.New(vmap.M{"OutputMode": vlog.OutputModeFile, "OutputFormatMode": vlog.OutputFormatModeJson})

	vtest.Nil(t, err)
	vtest.Equal(t, vlog.OutputModeFile, log.Options.OutputMode)

	log, err = vlog.New(vmap.M{"OutputMode": vlog.OutputModeConsoleAndFile})

	vtest.Nil(t, err)
	vtest.Equal(t, vlog.OutputModeConsoleAndFile, log.Options.OutputMode)

	log, err = vlog.New(vmap.M{"OutputMode": vlog.OutputMode(100), "OutputFormatMode": vlog.OutputFormatMode(100)})

	vtest.Nil(t, err)
	vtest.Equal(t, vlog.OutputMode(100), log.Options.OutputMode)
}

func TestSetConfig(t *testing.T) {
	vtest.Equal(t, "/data/logs/app.log", vlog.GetLog().Options.File)
	vtest.Equal(t, 500, vlog.GetLog().Options.MaxSize)

	_, err := vlog.SetConfig(map[string]any{"File": "./logs/app.log", "MaxSize": 50})

	vtest.Nil(t, err)
	vtest.Equal(t, "./logs/app.log", vlog.GetLog().Options.File)
	vtest.Equal(t, 50, vlog.GetLog().Options.MaxSize)

	_, err = vlog.SetConfig(map[string]any{"File": "/tmp/logs/app.log", "MaxSize": 10})

	vtest.Nil(t, err)
	vtest.Equal(t, "/tmp/logs/app.log", vlog.GetLog().Options.File)
	vtest.Equal(t, 10, vlog.GetLog().Options.MaxSize)

	_, err = vlog.SetConfig(map[string]any{"Test": ""})

	vtest.NotNil(t, err)
	vtest.Equal(t, "no attr, attr:Test", err.Error())
}

func TestDebug(t *testing.T) {
	vlog.Debug("debug")
	vlog.Debug("debug", "key1", "val1")
	vlog.Debug("debug", "key1", "val1", "key2", "val2")
}

func TestDebugf(t *testing.T) {
	vlog.Debugf("debugf")
	vlog.Debugf("debugf-%s", "val1")
	vlog.Debugf("debugf-%s-%s", "val1", "val2")
}

func TestDebugc(t *testing.T) {
	ctx := vtrace.SetTraceId(context.TODO(), "")
	vlog.Debugc(ctx, "debugc")
	vlog.Debugc(ctx, "debugc", "key1", "val1")
	vlog.Debugc(ctx, "debugc", "key1", "val1", "key2", "val2")
}

func TestInfo(t *testing.T) {
	vlog.Info("info")
	vlog.Info("info", "key1", "val1")
	vlog.Info("info", "key1", "val1", "key2", "val2")
}

func TestInfof(t *testing.T) {
	vlog.Infof("infof")
	vlog.Infof("infof-%s", "val1")
	vlog.Infof("infof-%s-%s", "val1", "val2")
}

func TestInfoc(t *testing.T) {
	ctx := vtrace.SetTraceId(context.TODO(), "")
	vlog.Infoc(ctx, "infoc")
	vlog.Infoc(ctx, "infoc", "key1", "val1")
	vlog.Infoc(ctx, "infoc", "key1", "val1", "key2", "val2")
}

func TestWarn(t *testing.T) {
	vlog.Warn("warn")
	vlog.Warn("warn", "key1", "val1")
	vlog.Warn("warn", "key1", "val1", "key2", "val2")
}

func TestWarnf(t *testing.T) {
	vlog.Warnf("warnf")
	vlog.Warnf("warnf-%s", "val1")
	vlog.Warnf("warnf-%s-%s", "val1", "val2")
}

func TestWarnc(t *testing.T) {
	ctx := vtrace.SetTraceId(context.TODO(), "")
	vlog.Warnc(ctx, "warnc")
	vlog.Warnc(ctx, "warnc", "key1", "val1")
	vlog.Warnc(ctx, "warnc", "key1", "val1", "key2", "val2")
}

func TestError(t *testing.T) {
	vlog.Error("error")
	vlog.Error("error", "key1")
	vlog.Error("error", "key1", "val1")
	vlog.Error("error", "key1", "val1", "key2", "val2")
}

func TestErrorf(t *testing.T) {
	vlog.Errorf("errorf")
	vlog.Errorf("errorf-%s")
	vlog.Errorf("errorf-%s", "val1")
	vlog.Errorf("errorf-%s-%s", "val1", "val2")
}

func TestErrorc(t *testing.T) {
	ctx := vtrace.SetTraceId(context.TODO(), "")
	vlog.Errorc(ctx, "errorc")
	vlog.Errorc(ctx, "errorc", "key1", "val1")
	vlog.Errorc(ctx, "errorc", "key1", "val1", "key2", "val2")
}

func TestFatal(t *testing.T) {
	if os.Getenv("SUB_PROCESS") == "1" {
		vlog.Fatal("fatal")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "SUB_PROCESS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
}

func TestFatalf(t *testing.T) {
	if os.Getenv("SUB_PROCESS") == "1" {
		vlog.Fatalf("fatalf")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalf")
	cmd.Env = append(os.Environ(), "SUB_PROCESS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
}

func TestFatalc(t *testing.T) {
	ctx := vtrace.SetTraceId(context.TODO(), "")
	if os.Getenv("SUB_PROCESS") == "1" {
		vlog.Fatalc(ctx, "fatalc")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatalc")
	cmd.Env = append(os.Environ(), "SUB_PROCESS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
}
