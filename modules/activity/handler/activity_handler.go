package handler

import "github.com/gofiber/fiber/v2"

type ActivityHandler interface {
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	RegisterRoutes()
}
