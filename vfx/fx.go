package vfx

import "go.uber.org/fx"

type (
	App       = fx.App
	Hook      = fx.Hook
	In        = fx.In
	Lifecycle = fx.Lifecycle
	Option    = fx.Option
)

var (
	NopLogger = fx.NopLogger
)

func New(opts ...Option) *App {
	return fx.New(opts...)
}

func Invoke(funcs ...interface{}) Option {
	return fx.Invoke(funcs...)
}

func Options(opts ...Option) Option {
	return fx.Options(opts...)
}

func Provide(constructors ...interface{}) Option {
	return fx.Provide(constructors...)
}

func Supply(values ...interface{}) Option {
	return fx.Supply(values...)
}
