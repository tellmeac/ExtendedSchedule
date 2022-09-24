package main

import (
	"context"
	"go.uber.org/fx"
	"tellmeac/extended-schedule/adapters"
	"tellmeac/extended-schedule/common/tsu"
)

// module is a root module that aggregates dependencies for application.
var module = fx.Options(
	tsu.Module,
	adapters.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lc fx.Lifecycle) {
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
