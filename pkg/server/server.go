package server

import "github.com/gofiber/fiber/v2"

type Server interface {
	Start() error
	GetEngine() *fiber.App
}
