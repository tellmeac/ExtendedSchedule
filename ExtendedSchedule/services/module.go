package services

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/services/faculty"
	"tellmeac/extended-schedule/services/lesson"
	"tellmeac/extended-schedule/services/schedule"
	"tellmeac/extended-schedule/services/tsuschedule"
	"tellmeac/extended-schedule/services/userconfig"
)

var Module = fx.Options(
	fx.Provide(tsuschedule.NewBaseScheduleClient),

	fx.Provide(faculty.New),
	fx.Provide(faculty.NewEndpoints),

	fx.Provide(lesson.New),
	fx.Provide(lesson.NewEndpoints),

	fx.Provide(schedule.New),
	fx.Provide(schedule.NewEndpoints),

	fx.Provide(userconfig.New),
	fx.Provide(userconfig.NewEndpoints),
)
