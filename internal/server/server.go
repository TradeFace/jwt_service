package server

import (
	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
	"github.com/tradeface/jwt_service/internal/conf"
	"github.com/tradeface/jwt_service/internal/controller"
	"github.com/tradeface/jwt_service/internal/provider"
)

type Server struct {
	controller *controller.Provider
}

func NewServer(cfg *conf.Config, providers *provider.Provider) (*Server, error) {

	controllerProvider, err := controller.NewProvider(cfg, providers.Store)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to db")
	}
	return &Server{
		controller: controllerProvider,
	}, nil
}

func (srv *Server) RegisterHandlers(e *echo.Echo) {

	e.GET("/user/login", srv.controller.User.Login)
	e.GET("/user/renew", srv.controller.User.Renew)

	//http://localhost:8888/user/6262ce0dafd1acb9dfbc4f87
	e.GET("/user/:id", srv.controller.User.Get)
}
