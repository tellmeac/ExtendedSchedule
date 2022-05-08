package adapters

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/adapters/builder"
	"tellmeac/extended-schedule/adapters/client/tsuschedule"
	"tellmeac/extended-schedule/adapters/provider"
	"tellmeac/extended-schedule/adapters/repository"
)

var Module = fx.Options(
	fx.Provide(tsuschedule.MakeClient),
	fx.Provide(provider.NewBaseScheduleProvider),
	fx.Provide(repository.NewEntUserConfigRepository),
	fx.Provide(builder.NewUserScheduleBuilder),
)
