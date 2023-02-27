package handler

import (
	"axion/database"
	"axion/model/entity"
	"axion/model/request"
	"axion/model/response"
	"fmt"
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
	var histories []response.AuctionHistory

	// result := database.DB.Debug().Find(&histories)
	// result := database.DB.Table("auction_histories").Preload("Auction.Product").Preload("Auction.User").Find(&histories)
	// get auction histories with product name, user name, and auction user name
	result := database.DB.Table("auction_histories").Select("auction_histories.id, auction_histories.auction_id, auction_histories.user_id, auction_histories.price, auction_histories.created_at, auction_histories.updated_at, auctions.name as auction_name, users.name as user_name").Joins("left join auctions on auctions.id = auction_histories.auction_id").Joins("left join products on products.id = auctions.product_id").Joins("left join users on users.id = auction_histories.user_id").Find(&histories)
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

	var auction entity.AuctionHistory

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

	var auction []entity.AuctionHistory

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

// @Summary Create Auction History
// @Description Create Auction History
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Param auction body request.AuctionHistoryCreateRequest true "Auction History"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auction-histories [post]
// @Security ApiKeyAuth
func AuctionHistoryHandlerCreate(ctx *fiber.Ctx) error {
	auction := new(request.AuctionHistoryCreateRequest)
	if err := ctx.BodyParser(auction); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}
	log.Println("ini auction : ", auction)

	validate := validator.New()
	errValidate := validate.Struct(auction)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	log.Println("auction: ", auction)
	log.Println("auction: ", auction.AuctionID)
	log.Println("auction: ", auction.UserId)

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
// @Param auction body request.AuctionHistoryUpdateRequest true "Auction History"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auction-histories/{id} [put]
// @Security ApiKeyAuth
func AuctionHistoryHandlerUpdate(ctx *fiber.Ctx) error {
	auctionRequest := new(request.AuctionHistoryUpdateRequest)
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

	auction.Price = auctionRequest.Price

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
			"message": "transaction not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&auction).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "transaction was deleted",
	})
}
