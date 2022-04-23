package store

import (
	"github.com/tradeface/suggest_service/pkg/service"
)

type Stores struct {
	User *UserStore
	Auth *AuthStore
}

func New(service *service.Service) (*Stores, error) {
	return &Stores{
		User: NewUserStore(service.Mongo),
		Auth: NewAuthStore(),
	}, nil
}
