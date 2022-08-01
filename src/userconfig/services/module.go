package services

import (
	"github.com/tellmeac/extended-schedule/userconfig/services/faculty"
	"github.com/tellmeac/extended-schedule/userconfig/services/lesson"
	"github.com/tellmeac/extended-schedule/userconfig/services/schedule"
	"github.com/tellmeac/extended-schedule/userconfig/services/tsuschedule"
	"github.com/tellmeac/extended-schedule/userconfig/services/userconfig"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(tsuschedule.NewBaseScheduleClient),

	fx.Provide(fx.Annotate(
		faculty.New,
		fx.ResultTags(`name:"facultyService"`),
	)),
	fx.Provide(fx.Annotate(
		faculty.NewCacheService,
		fx.ParamTags(`name:"facultyService"`),
	)),
	fx.Provide(faculty.NewEndpoints),

	fx.Provide(lesson.New),
	fx.Provide(lesson.NewEndpoints),

	fx.Provide(schedule.New),
	fx.Provide(schedule.NewEndpoints),

	fx.Provide(userconfig.New),
	fx.Provide(userconfig.NewEndpoints),
)
