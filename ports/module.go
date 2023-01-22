package ports

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/ports/schedule"
)

var Module = fx.Options(
	fx.Provide(schedule.NewServerHandler),
)
