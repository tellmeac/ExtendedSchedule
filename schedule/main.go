package main

import (
	"context"
	"github.com/tellmeac/ext-schedule/schedule/adapters"
	"github.com/tellmeac/ext-schedule/schedule/common/tsu"
	"github.com/tellmeac/ext-schedule/schedule/common/userconfig"
	"go.uber.org/fx"
)

// module is a root module that aggregates dependencies for application.
var module = fx.Options(
	tsu.Module,
	userconfig.Module,
	adapters.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lc fx.Lifecycle, p adapters.ScheduleProvider) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func main() {
	fx.New(module).Run()
}
