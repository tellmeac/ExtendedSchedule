package adapters

import "go.uber.org/fx"

// Module provides adapters like providers.
var Module = fx.Options(
	fx.Provide(NewScheduleProvider),
)
