package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"net/http"
	"tellmeac/extended-schedule/pkg/config"
	"tellmeac/extended-schedule/pkg/ent"
	"tellmeac/extended-schedule/pkg/faculty"
	"tellmeac/extended-schedule/pkg/handlers"
	"tellmeac/extended-schedule/pkg/lesson"
	logger "tellmeac/extended-schedule/pkg/log"
	"tellmeac/extended-schedule/pkg/schedule"
	"tellmeac/extended-schedule/pkg/server"
	"tellmeac/extended-schedule/pkg/tsuschedule"
	"tellmeac/extended-schedule/pkg/userconfig"
)

// Module is a top tree module of application.
var Module = fx.Options(
	fx.Provide(config.MustLoad),
	logger.Module,
	ent.Module,
	tsuschedule.Module,
	userconfig.Module,
	faculty.Module,
	lesson.Module,
	schedule.Module,
	handlers.Module,
	fx.Provide(server.New),
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
