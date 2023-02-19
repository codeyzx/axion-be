package handler

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/model/response"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get All Auction History
// @Description Get All Auction History
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Router /auction-histories [get]
// @Security ApiKeyAuth
func AuctionHistoryHandlerGetAll(ctx *fiber.Ctx) error {
	var histories []entity.AuctionHistory

	result := database.DB.Debug().Find(&histories)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(histories)
}

// @Summary Get Auction History By Id
// @Description Get Auction History By Id
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Param id path string true "Auction History Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /auction-histories/{id} [get]
func AuctionHistoryHandlerGetById(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	var auction response.AuctionHistory

	err := database.DB.Table("auction_histories").Where("auction_histories.id = ?", ID).Preload("Auction.Product").Preload("Auction.User").Preload("User").First(&auction).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "auction not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    auction,
	})
}

// @Summary Get Auction History By User Id
// @Description Get Auction History By User Id
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auction-histories/user/{id} [get]
// @Security ApiKeyAuth
func AuctionHistoryHandlerGetByUser(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	temp := ctx.Locals("userId")

	authId := fmt.Sprintf("%v", temp)

	var auction []response.AuctionHistory

	log.Println(ID + " " + authId)
	if temp != 0 {
		if ID != authId {
			return ctx.Status(403).JSON(fiber.Map{
				"message": "forbidden",
			})
		}
	}

	err := database.DB.Table("auction_histories").Where("auction_histories.user_id = ?", ID).Preload("Auction.Product").Preload("Auction.User").Preload("User").Find(&auction).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "auction not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    auction,
	})
}

// @Summary Get Auction History By Auction Id
// @Description Get Auction History By Auction Id
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Param id path string true "Auction Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auction-histories/auction/{id} [get]
// @Security ApiKeyAuth
func AuctionHistoryHandlerCreate(ctx *fiber.Ctx) error {
	auction := new(request.AuctionHistoryCreateRequest)
	if err := ctx.BodyParser(auction); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(auction)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newAuctionHistory := entity.AuctionHistory{
		AuctionID: auction.AuctionID,
		UserId:    auction.UserId,
		Price:     auction.Price,
	}

	errCreateAuctionHistory := database.DB.Create(&newAuctionHistory).Error
	if errCreateAuctionHistory != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newAuctionHistory,
	})
}

// @Summary Update Auction History
// @Description Update Auction History
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Param id path string true "Auction History Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auction-histories/{id} [put]
// @Security ApiKeyAuth
func AuctionHistoryHandlerUpdate(ctx *fiber.Ctx) error {
	auctionRequest := new(request.AuctionHistoryCreateRequest)
	if err := ctx.BodyParser(auctionRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var auction entity.AuctionHistory

	ID := ctx.Params("id")

	err := database.DB.First(&auction, "id = ?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "auction not found",
		})
	}

	if auctionRequest.AuctionID != 0 {
		auction.AuctionID = auctionRequest.AuctionID
	}

	if auctionRequest.UserId != 0 {
		auction.UserId = auctionRequest.UserId
	}
	if auctionRequest.Price != 0 {
		auction.Price = auctionRequest.Price
	}

	errUpdate := database.DB.Save(&auction).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    auction,
	})
}

// @Summary Delete Auction History
// @Description Delete Auction History
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Param id path string true "Auction History Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auction-histories/{id} [delete]
// @Security ApiKeyAuth
func AuctionHistoryHandlerDelete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	var auction entity.AuctionHistory

	err := database.DB.Debug().First(&auction, "id=?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "auction not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&auction).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "auction was deleted",
	})
}
