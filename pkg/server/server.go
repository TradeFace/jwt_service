package server

import (
	"errors"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo/v4"
	"github.com/tradeface/jwt_service/internal/conf"
	"github.com/tradeface/jwt_service/pkg/authorization"
	"github.com/tradeface/suggest_service/pkg/store"
)

type Server struct {
	cfg    *conf.Config
	stores *store.Stores
}

func NewServer(cfg *conf.Config, stores *store.Stores) (*Server, error) {

	return &Server{
		cfg:    cfg,
		stores: stores,
	}, nil
}

func (srv *Server) RegisterHandlers(e *echo.Echo) {

	e.GET("/user/login", srv.LoginUser)

	//http://localhost:8888/user/6262ce0dafd1acb9dfbc4f87
	e.GET("/user/:id", srv.GetUser)
}

func (srv *Server) GetAuthUser(c echo.Context) (*authorization.AuthUser, error) {

	user := c.Get("authUser")
	if user == nil {
		return nil, errors.New("no auth user available")
	}
	return user.(*authorization.AuthUser), nil
}

func (srv *Server) Output(c echo.Context, res interface{}, err error) error {

	if err != nil {
		srv.sendError(c, err)
		return nil
	}
	payload, err := jsonapi.Marshal(res)
	if err != nil {
		srv.sendError(c, err)
		return nil
	}
	return c.JSON(http.StatusOK, payload)
}

func (srv *Server) sendError(c echo.Context, err error) {

	c.Response().Header().Set(echo.HeaderContentType, jsonapi.MediaType)
	c.Response().WriteHeader(http.StatusBadRequest)
	jsonapi.MarshalErrors(c.Response().Writer, []*jsonapi.ErrorObject{{
		Title:  "Request Error",
		Detail: err.Error(),
		Status: "400",
	}})
}
