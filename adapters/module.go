package adapters

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/ports/schedule"
	sch "tellmeac/extended-schedule/schedule"
)

// Module provides adapters like providers and repositories.
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
	fx.Provide(NewTargetProvider),

	fx.Provide(fx.Annotate(
		NewTeachersRepository,
		fx.As(new(schedule.TeacherProvider)),
	)),

	fx.Provide(fx.Annotate(
		NewStudyGroupRepository,
		fx.As(new(schedule.GroupProvider)),
	)),
)
