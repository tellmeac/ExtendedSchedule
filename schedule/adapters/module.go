package adapters

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/ports/schedule"
	sch "tellmeac/extended-schedule/schedule"
)

// Module provides adapters like providers.
var Module = fx.Options(
	fx.Provide(NewEntClient),
	fx.Provide(fx.Annotate(
		NewUserConfigRepository,
		fx.As(new(sch.ConfigProvider)),
	)),
	fx.Provide(fx.Annotate(
		NewScheduleProvider,
		fx.As(new(sch.Provider)),
		fx.As(new(schedule.Provider)),
	)),
	fx.Provide(fx.Annotate(
		NewFacultyProvider,
		fx.As(new(schedule.FacultyProvider)),
	)),
)
