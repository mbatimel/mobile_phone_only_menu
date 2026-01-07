package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"

	"github.com/mbatimel/mobile_phone_only_menu/internal/config"
	menuAPIService "github.com/mbatimel/mobile_phone_only_menu/internal/service"
	"github.com/mbatimel/mobile_phone_only_menu/internal/storage/postgres"
	transportHttp "github.com/mbatimel/mobile_phone_only_menu/internal/transport/http"
	"github.com/mbatimel/mobile_phone_only_menu/internal/transport/jsonRPC/externalapi"
)

const serviceName = "menu"

func main() {
	log.Logger = config.Values().Logger().With().Str("serviceName", serviceName).Logger()
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	postgresStorage, err := postgres.New(config.Values().Postgres, log.Logger)
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("failed to connect to postgres")
	}

	svc := menuAPIService.NewMenuService(
		log.Logger,
		postgresStorage,

		config.Values().ServiceID,
	)

	services := []externalapi.Option{}

	services = append(services,
		externalapi.PublicApi(externalapi.NewPublicApi(svc)),
	)

	app := externalapi.New(log.Logger, services...).WithLog().WithMetrics()
	server := &fasthttp.Server{
		Handler:            app.Fiber().Handler(),
		MaxRequestBodySize: config.Values().MaxRequestBodySize,
		ReadBufferSize:     config.Values().MaxRequestHeaderSize,
		ReadTimeout:        time.Duration(config.Values().ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(config.Values().ExternalWriteTimeoutS) * time.Second,
	}

	healthServer := transportHttp.NewHealthServer()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		serveErr := server.ListenAndServe(config.Values().ServiceBind)
		if serveErr != nil {
			log.Fatal().Err(serveErr).Msg("failed to listen server")
		} else {
			log.Error().Msg("external api server stopped with no error")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		healthErr := healthServer.Start(config.Values().HealthBind)
		if healthErr != nil {
			log.Error().Err(healthErr).Msg("failed to start health server")
		} else {
			log.Error().Msg("health server stopped with no error")
		}
	}()

	<-shutdown

	err = healthServer.Stop()
	if err != nil {
		log.Error().Err(err).Msg("failed to stop health server")
	}

	err = server.Shutdown()
	if err != nil {
		log.Error().Err(err).Msg("failed to shutdown server")
	}

	wg.Wait()
}
