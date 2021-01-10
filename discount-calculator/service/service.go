package service

import (
	"github.com/iurybranco/backend-challenge/discount-calculator/service/controller"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/server"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	grpcServer *server.Server
	dbClient   database.Client
}

func New(c *Config) (*Service, error) {
	var srvice Service
	var err error
	srvice.dbClient, err = database.New(c.Database)
	if err != nil {
		return nil, err
	}
	cntroller := controller.New(srvice.dbClient)
	srvice.grpcServer, _ = server.New(c.ServerPort, cntroller)
	return &srvice, nil
}

func (s *Service) Run() <-chan error {
	return s.grpcServer.Run()
}

func (s *Service) Shutdown() {
	if err := s.dbClient.Close(); err != nil {
		log.Errorln(errors.Wrap(err, "failed to close database client"))
	}
	s.grpcServer.Close()
}
