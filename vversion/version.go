package vversion

import (
	"runtime"

	"github.com/sunshinexcode/gotk/venv"
	"github.com/sunshinexcode/gotk/vlog"
	"github.com/sunshinexcode/gotk/vmetric"
)

const GotkVersion = "v0.0.1"

var (
	Branch    string
	BuildTime string
	Commit    string
	GoVersion = runtime.Version()
	Project   string
	Version   string
)

func Metric(metric *vmetric.Metric) {
	metric.BuildInfo.WithLabelValues(Branch, BuildTime, Commit, GoVersion, Version).Set(1)
}

func Print() {
	vlog.Info("version", "Branch", Branch, "BuildTime", BuildTime, "Commit", Commit, "Env", venv.GetEnv(venv.AppEnvKey, venv.AppEnvDev), "GoVersion", GoVersion, "Project", Project, "Version", Version)
}
