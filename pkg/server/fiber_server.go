package server

import (
	"fmt"
	"todolist-v1/config"

	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app *fiber.App
	cfg *config.Config
}

func NewFiberServer(cfg *config.Config) Server {
	return &fiberServer{
		app: fiber.New(),
		cfg: cfg,
	}
}

func (s *fiberServer) Start() error {
	return s.app.Listen(fmt.Sprintf(":%s", s.cfg.Server.Port))
}

func (s *fiberServer) GetEngine() *fiber.App {
	return s.app
}
