package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tellmeac/ExtendedSchedule/userconfig/config"
	"github.com/tellmeac/ExtendedSchedule/userconfig/dao"
	inf "github.com/tellmeac/ExtendedSchedule/userconfig/infrastructure"
	"github.com/tellmeac/ExtendedSchedule/userconfig/services"
	"go.uber.org/fx"
	"net/http"
)

// Module is a top tree module of application.
var Module = fx.Options(
	fx.Provide(config.MustLoad),
	fx.Invoke(inf.InitLogger),

	fx.Provide(inf.NewEntClient),
	dao.Module,

	services.Module,
	fx.Provide(inf.NewServer),

	fx.Invoke(bind),
	fx.Invoke(bootstrap),
)

// bootstrap function is an endpoint to all provided structs.
func bootstrap(lc fx.Lifecycle, cfg config.Config, server *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info().Msg("Starting application")

			var err error
			go func() {
				err = http.ListenAndServe(cfg.ListenAddress, server)
			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Stop application")
			return nil
		},
	})
}
