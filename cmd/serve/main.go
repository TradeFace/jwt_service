package main

import (
	"io/ioutil"

	"github.com/labstack/echo/v4"

	echolog "github.com/labstack/gommon/log"
	"github.com/rs/zerolog/log"
	"github.com/tradeface/jwt_service/internal/conf"
	"github.com/tradeface/jwt_service/pkg/middleware"
	"github.com/tradeface/jwt_service/pkg/server"
	"github.com/tradeface/jwt_service/pkg/store"
	"github.com/tradeface/suggest_service/pkg/service"
)

//TODO: config cli/dockersecrets

//  https://pkg.go.dev/github.com/golang-jwt/jwt/v4

const (
	// APPNAME contains the name of the program
	APPNAME = "jwt_service"
	// APPVERSION contains the version of the program
	APPVERSION = "0.0.2"
)

func main() {

	cfg, err := conf.NewDefaultConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	srvConf := &service.Config{
		MongoURI: cfg.MongoURI,
		MongoDB:  cfg.MongoDB,
	}

	services, err := service.New(srvConf)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to db")
	}

	stores, err := store.New(services)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to db")
	}

	srv, err := server.NewServer(cfg, stores)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to bind api")
	}

	e := echo.New()
	// shut up
	e.Logger.SetOutput(ioutil.Discard)
	e.Logger.SetLevel(echolog.OFF)

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.Use(middleware.JWTWithConfig(&middleware.JWTConfig{}, stores.Auth))

	srv.RegisterHandlers(e)

	log.Info().Str("addr", cfg.Addr).Msg("starting http listener")
	err = e.Start(cfg.Addr)
	log.Fatal().Err(err).Msg("Server failed")
}
