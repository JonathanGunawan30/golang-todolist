package handler

import (
	"errors"
	"strconv"
	"todolist-v1/modules/activity/entities"
	"todolist-v1/modules/activity/models"
	"todolist-v1/modules/activity/repository"
	"todolist-v1/modules/activity/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type activityHandlerHttp struct {
	app      *fiber.App
	usecase  usecase.ActivityUsecase
	validate *validator.Validate
}

func NewActivityHttpHandler(app *fiber.App, usecase usecase.ActivityUsecase) ActivityHandler {
	return &activityHandlerHttp{
		app:      app,
		usecase:  usecase,
		validate: validator.New(),
	}
}

func (handler *activityHandlerHttp) GetAll(ctx *fiber.Ctx) error {
	activities, err := handler.usecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	var activityResponses []models.ActivityResponse
	for _, a := range activities {
		activityResponses = append(activityResponses, models.ActivityResponse{
			Id:           a.Id,
			Title:        a.Title,
			Category:     a.Category,
			Description:  a.Description,
			ActivityDate: a.ActivityDate,
			Status:       a.Status,
		})
	}

	return ctx.JSON(fiber.Map{
		"data":        activityResponses,
		"status_code": fiber.StatusOK,
		"message":     "Activities retrieved successfully",
	})
}

func (handler *activityHandlerHttp) Create(ctx *fiber.Ctx) error {
	var request models.ActivityCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusBadRequest,
			"message":     "Cannot parse JSON",
		})
	}

	if err := handler.validate.Struct(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
		})
	}

	activityEntity := entities.Activity{
		Title:        request.Title,
		Category:     request.Category,
		Description:  request.Description,
		ActivityDate: request.ActivityDate,
	}

	newActivity, err := handler.usecase.Create(activityEntity)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": models.ActivityResponse{
			Id:           newActivity.Id,
			Title:        newActivity.Title,
			Category:     newActivity.Category,
			Description:  newActivity.Description,
			ActivityDate: newActivity.ActivityDate,
			Status:       newActivity.Status,
		},
		"status_code": fiber.StatusCreated,
		"message":     "Activity created successfully",
	})
}

func (handler *activityHandlerHttp) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusBadRequest,
			"message":     "Invalid ID",
		})
	}

	var request models.ActivityUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusBadRequest,
			"message":     "Cannot parse JSON",
		})
	}

	if err := handler.validate.Struct(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
		})
	}

	activityEntity := entities.Activity{
		Title:        request.Title,
		Category:     request.Category,
		Description:  request.Description,
		ActivityDate: request.ActivityDate,
		Status:       request.Status,
	}

	updatedActivity, err := handler.usecase.Update(id, activityEntity)
	if err != nil {
		if errors.Is(err, repository.ErrActivityNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"data":        nil,
				"status_code": fiber.StatusNotFound,
				"message":     err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": models.ActivityResponse{
			Id:           updatedActivity.Id,
			Title:        updatedActivity.Title,
			Category:     updatedActivity.Category,
			Description:  updatedActivity.Description,
			ActivityDate: updatedActivity.ActivityDate,
			Status:       updatedActivity.Status,
		},
		"status_code": fiber.StatusOK,
		"message":     "Activity updated successfully",
	})
}

func (handler *activityHandlerHttp) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusBadRequest,
			"message":     "Invalid ID",
		})
	}

	if err := handler.usecase.Delete(id); err != nil {
		if errors.Is(err, repository.ErrActivityNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"data":        nil,
				"status_code": fiber.StatusNotFound,
				"message":     "Activity not found",
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":        nil,
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":        nil,
		"status_code": fiber.StatusOK,
		"message":     "Activity deleted successfully",
	})
}

func (handler *activityHandlerHttp) RegisterRoutes() {
	handler.app.Get("/api/activities", handler.GetAll)
	handler.app.Post("/api/activities", handler.Create)
	handler.app.Put("/api/activities/:id", handler.Update)
	handler.app.Delete("/api/activities/:id", handler.Delete)
}
