package handlers

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/pkg/handlers/faculty"
	"tellmeac/extended-schedule/pkg/handlers/lesson"
	"tellmeac/extended-schedule/pkg/handlers/schedule"
	"tellmeac/extended-schedule/pkg/handlers/userconfig"
)

var Module = fx.Options(
	fx.Provide(faculty.New),
	fx.Provide(lesson.New),
	fx.Provide(schedule.New),
	fx.Provide(userconfig.New),
)
