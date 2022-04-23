package service

import (
	"errors"

	"github.com/tradeface/jwt_service/internal/conf"
)

type Service struct {
	Mongo *MongoService
}

func New(cfg *conf.Config) (service *Service, err error) {

	mongoConf := &MongoConfig{
		MongoURI: cfg.MongoURI,
		MongoDB:  cfg.MongoDB,
	}
	mongoService, err := NewMongoService(mongoConf)
	if err != nil {
		return service, errors.New("failed to connect to db")
	}

	return &Service{
		Mongo: mongoService,
	}, nil
}
