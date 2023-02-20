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
	"github.com/google/uuid"
)

// @Summary Get All Auction
// @Description Get All Auction
// @Tags Auction
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Router /auctions [get]
func AuctionHandlerGetAll(ctx *fiber.Ctx) error {
	var auctions []response.Auction

	result := database.DB.Preload("Product").Preload("User").Find(&auctions)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(auctions)
}

// @Summary Get Auction By Id
// @Description Get Auction By Id
// @Tags Auction
// @Accept  json
// @Produce  json
// @Param id path string true "Auction Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /auctions/{id} [get]
func AuctionHandlerGetById(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	var auction response.Auction

	err := database.DB.Where("auctions.id = ?", ID).Preload("Product").Preload("User").Preload("AuctionHistory.User").First(&auction).Error

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

// @Summary Create Auction
// @Description Create Auction
// @Tags Auction
// @Accept  json
// @Produce  json
// @Param auction body request.AuctionCreateRequest true "Auction"
// @Success 200
// @Failure 400
// @Failure 401
// @Router /auctions [post]
// @Security ApiKeyAuth
func AuctionHandlerCreate(ctx *fiber.Ctx) error {

	productId := uuid.New()
	auction := new(request.AuctionCreateRequest)

	if err := ctx.BodyParser(auction); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(auction)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var pathFileString string
	pathFile := ctx.Locals("pathFile")
	if pathFile == nil {
		pathFileString = ""

	} else {
		pathFileString = fmt.Sprintf("%v", pathFile)
	}

	newProduct := entity.Product{

		ID:          productId,
		Name:        auction.Product.Name,
		Description: auction.Product.Description,
		Price:       auction.Product.Price,

		Image: pathFileString,
	}

	newAuction := entity.Auction{

		ProductID: productId,

		LastPrice: auction.Product.Price,
		UserId:    auction.UserId,

		Status:       entity.Open,
		BiddersCount: auction.BiddersCount,
		EndAt:        auction.EndAt,
	}

	errCreateProduct := database.DB.Create(&newProduct).Error
	errCreateAuction := database.DB.Create(&newAuction).Error

	if errCreateAuction != nil || errCreateProduct != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newAuction,
	})
}

// @Summary Update Auction
// @Description Update Auction
// @Tags Auction
// @Accept  json
// @Produce  json
// @Param id path string true "Auction Id"
// @Param auction body request.AuctionUpdateRequest true "Auction"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /auctions/{id} [put]
// @Security ApiKeyAuth
func AuctionHandlerUpdate(ctx *fiber.Ctx) error {
	auctionRequest := new(request.AuctionUpdateRequest)
	if err := ctx.BodyParser(auctionRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var auction entity.Auction

	ID := ctx.Params("id")

	err := database.DB.First(&auction, "id = ?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "auction not found",
		})
	}

	if auctionRequest.Status != "" {
		if auctionRequest.Status != entity.Open && auctionRequest.Status != entity.Closed {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "status must be Open or Closed",
			})
		}
		auction.Status = auctionRequest.Status
	}

	if auctionRequest.LastPrice != 0 {
		auction.LastPrice = auctionRequest.LastPrice
	}

	if auctionRequest.UserId != 0 {
		auction.UserId = auctionRequest.UserId
	}

	if auctionRequest.EndAt != "" {
		auction.EndAt = auctionRequest.EndAt
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

// @Summary Delete Auction
// @Description Delete Auction
// @Tags Auction
// @Accept  json
// @Produce  json
// @Param id path string true "Auction Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /auctions/{id} [delete]
// @Security ApiKeyAuth
func AuctionHandlerDelete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	var auction entity.Auction

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
