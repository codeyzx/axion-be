package handler

import (
	"axion/database"
	"axion/model/entity"
	"axion/model/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get All History
// @Description Get All History
// @Tags History
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Router /history [get]
// @Security ApiKeyAuth
func HistoryHandlerGetAll(ctx *fiber.Ctx) error {
	var Historys []entity.History

	result := database.DB.Debug().Find(&Historys)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(Historys)
}

// @Summary Get History By Id
// @Description Get History By Id
// @Tags History
// @Accept  json
// @Produce  json
// @Param id path string true "History Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /history/{id} [get]
// @Security ApiKeyAuth
func HistoryHandlerGetById(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	var history entity.History
	err := database.DB.First(&history, "id = ?", ID).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "History not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    history,
	})
}

// @Summary Create History
// @Description Create History
// @Tags History
// @Accept  json
// @Produce  json
// @Param History body request.HistoryRequest true "History"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /history [post]
// @Security ApiKeyAuth
func HistoryHandlerCreate(ctx *fiber.Ctx) error {
	History := new(request.HistoryRequest)
	if err := ctx.BodyParser(History); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(History)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newHistory := entity.History{
		Log: History.Log,
	}

	errCreateHistory := database.DB.Create(&newHistory).Error
	if errCreateHistory != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newHistory,
	})
}

// @Summary Update History
// @Description Update History
// @Tags History
// @Accept  json
// @Produce  json
// @Param id path string true "History Id"
// @Param History body request.HistoryRequest true "History"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /history/{id} [put]
// @Security ApiKeyAuth
func HistoryHandlerUpdate(ctx *fiber.Ctx) error {
	HistoryRequest := new(request.HistoryRequest)
	if err := ctx.BodyParser(HistoryRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var History entity.History

	ID := ctx.Params("id")

	err := database.DB.First(&History, "id = ?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "History not found",
		})
	}

	if HistoryRequest.Log != "" {
		History.Log = HistoryRequest.Log
	}

	errUpdate := database.DB.Save(&History).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    History,
	})
}

// @Summary Delete History
// @Description Delete History
// @Tags History
// @Accept  json
// @Produce  json
// @Param id path string true "History Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /history/{id} [delete]
// @Security ApiKeyAuth
func HistoryHandlerDelete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	var History entity.History

	err := database.DB.Debug().First(&History, "id=?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "History not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&History).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "History was deleted",
	})
}
