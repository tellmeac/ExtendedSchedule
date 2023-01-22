package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"tellmeac/extended-schedule/adapters"
	"tellmeac/extended-schedule/common/logger"
	"tellmeac/extended-schedule/common/tsu"
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/ports"
	"tellmeac/extended-schedule/schedule"
	"tellmeac/extended-schedule/server"
)

// Module is a root Module that aggregates dependencies for application.
var Module = fx.Options(
	fx.Invoke(logger.InitLogger),

	tsu.Module,
	adapters.Module,
	ports.Module,
	schedule.Module,

	fx.Provide(server.New),
	fx.Invoke(setupRouting),
	fx.Invoke(Bootstrap),
)

func Bootstrap(lc fx.Lifecycle, engine *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			go func() {
				err = http.ListenAndServe(config.Get().ListenAddress, engine)
			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
