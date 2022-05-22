package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"tellmeac/extended-schedule/adapters"
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/infrastructure"
	"tellmeac/extended-schedule/infrastructure/log"
	"tellmeac/extended-schedule/service/groups"
	"tellmeac/extended-schedule/service/schedule"
	"tellmeac/extended-schedule/service/userconfig"
)

var Module = fx.Options(
	fx.Invoke(log.ConfigureLogger),
	fx.Provide(config.MustLoad),
	fx.Provide(infrastructure.NewClient),
	adapters.Module,
	schedule.Module,
	userconfig.Module,
	groups.Module,
	fx.Invoke(bindRoutes),
	fx.Provide(infrastructure.NewServer),
	fx.Invoke(bootstrap),
)

// bootstrap function is an end point to all providing structs.
func bootstrap(lc fx.Lifecycle, cfg config.Config, server *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error
			go func() {
				log.Sugared.Infof("Starting server on: %s", cfg.ListenAddress)
				err = http.ListenAndServe(cfg.ListenAddress, server)
			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Sugared.Info("Stopping server")
			return nil
		},
	})
}
