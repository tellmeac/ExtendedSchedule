package adapters

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/adapters/builder"
	"tellmeac/extended-schedule/adapters/clients/tsuschedule"
	"tellmeac/extended-schedule/adapters/providers"
	"tellmeac/extended-schedule/adapters/repository"
)

var Module = fx.Options(
	fx.Provide(tsuschedule.MakeClient),
	fx.Provide(providers.NewBaseScheduleProvider),
	fx.Provide(repository.NewEntUserConfigRepository),
	fx.Provide(builder.NewUserScheduleBuilder),
)
