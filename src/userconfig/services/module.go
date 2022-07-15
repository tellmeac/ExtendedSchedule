package services

import (
	"github.com/tellmeac/ExtendedSchedule/userconfig/services/faculty"
	"github.com/tellmeac/ExtendedSchedule/userconfig/services/lesson"
	"github.com/tellmeac/ExtendedSchedule/userconfig/services/schedule"
	"github.com/tellmeac/ExtendedSchedule/userconfig/services/tsuschedule"
	"github.com/tellmeac/ExtendedSchedule/userconfig/services/userconfig"
	"go.uber.org/fx"
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
