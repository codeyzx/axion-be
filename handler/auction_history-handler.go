package handler

import (
	"axion/database"
	"axion/model/entity"
	"axion/model/request"
	"axion/model/response"
	"fmt"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/signintech/gopdf"
	"github.com/xuri/excelize/v2"
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

	var auction request.AuctionHistory
	err := database.DB.Table("auction_histories").Where("auction_histories.id = ?", ID).Preload("Auction").Preload("Auction.Product").Preload("Auction.User").Preload("User").First(&auction).Error

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

// @Summary Export Auction History to Excel
// @Description Export Auction History to Excel
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auction-histories-export-excel [get]
// @Security ApiKeyAuth
func AuctionHistoryExportToExcel(c *fiber.Ctx) error {
	temp := c.Locals("userId")

	var histories []response.AuctionHistory

	if temp != 0 {
		result := database.DB.Table("auction_histories").Select("auction_histories.id, auction_histories.auction_id, auction_histories.user_id, auction_histories.price, auction_histories.created_at, auction_histories.updated_at, auctions.name as auction_name, users.name as user_name").Joins("left join auctions on auctions.id = auction_histories.auction_id").Joins("left join products on products.id = auctions.product_id").Joins("left join users on users.id = auction_histories.user_id").Where("auction_histories.user_id = ?", temp).Find(&histories)
		if result.Error != nil {
			log.Println(result.Error)
		}
	} else {

		result := database.DB.Table("auction_histories").Select("auction_histories.id, auction_histories.auction_id, auction_histories.user_id, auction_histories.price, auction_histories.created_at, auction_histories.updated_at, auctions.name as auction_name, users.name as user_name").Joins("left join auctions on auctions.id = auction_histories.auction_id").Joins("left join products on products.id = auctions.product_id").Joins("left join users on users.id = auction_histories.user_id").Find(&histories)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}

	file := excelize.NewFile()
	const sheet = "Transaction"

	index, err := file.NewSheet(sheet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	file.SetCellValue(sheet, "A1", "ID")
	file.SetCellValue(sheet, "B1", "Auction")
	file.SetCellValue(sheet, "C1", "User")
	file.SetCellValue(sheet, "D1", "Price")
	file.SetCellValue(sheet, "E1", "Created At")
	file.SetCellValue(sheet, "F1", "Updated At")

	style, err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create style",
		})
	}

	file.SetCellStyle(sheet, "A1", "F1", style)

	for i, History := range histories {
		i = i + 2
		file.SetCellValue(sheet, "A"+strconv.Itoa(i), History.ID)
		file.SetCellValue(sheet, "B"+strconv.Itoa(i), History.AuctionName)
		file.SetCellValue(sheet, "C"+strconv.Itoa(i), History.UserName)
		file.SetCellValue(sheet, "D"+strconv.Itoa(i), History.Price)
		file.SetCellValue(sheet, "E"+strconv.Itoa(i), History.CreatedAt)
		file.SetCellValue(sheet, "F"+strconv.Itoa(i), History.UpdatedAt)
	}

	file.SetColWidth(sheet, "A", "A", 10)
	file.SetColWidth(sheet, "B", "B", 30)
	file.SetColWidth(sheet, "C", "C", 30)
	file.SetColWidth(sheet, "D", "D", 30)
	file.SetColWidth(sheet, "E", "E", 30)
	file.SetColWidth(sheet, "F", "F", 30)

	file.SetActiveSheet(index)

	c.Set("Content-Disposition", "attachment; filename=transaction-report.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	errWrite := file.Write(c.Response().BodyWriter())
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// @Summary Export Auction History to PDF
// @Description Export Auction History to PDF
// @Tags Auction History
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auction-histories-export-pdf [get]
// @Security ApiKeyAuth
func AuctionHistoryExportToPDF(c *fiber.Ctx) error {
	temp := c.Locals("userId")

	var histories []response.AuctionHistory

	if temp != 0 {
		result := database.DB.Table("auction_histories").Select("auction_histories.id, auction_histories.auction_id, auction_histories.user_id, auction_histories.price, auction_histories.created_at, auction_histories.updated_at, auctions.name as auction_name, users.name as user_name").Joins("left join auctions on auctions.id = auction_histories.auction_id").Joins("left join products on products.id = auctions.product_id").Joins("left join users on users.id = auction_histories.user_id").Where("auction_histories.user_id = ?", temp).Find(&histories)
		if result.Error != nil {
			log.Println(result.Error)
		}
	} else {

		result := database.DB.Table("auction_histories").Select("auction_histories.id, auction_histories.auction_id, auction_histories.user_id, auction_histories.price, auction_histories.created_at, auction_histories.updated_at, auctions.name as auction_name, users.name as user_name").Joins("left join auctions on auctions.id = auction_histories.auction_id").Joins("left join products on products.id = auctions.product_id").Joins("left join users on users.id = auction_histories.user_id").Find(&histories)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}

	var auctionNames []string
	for _, history := range histories {
		if history.AuctionName == "" {
			history.AuctionName = "-"
		}
		auctionNames = append(auctionNames, history.AuctionName)
	}

	file := excelize.NewFile()
	const sheet = "Transaction"

	index, err := file.NewSheet(sheet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	file.SetCellValue(sheet, "A1", "Price")
	file.SetCellValue(sheet, "B1", "Auction")
	file.SetCellValue(sheet, "C1", "User")

	style, err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create style",
		})
	}

	file.SetCellStyle(sheet, "A1", "C1", style)

	for i, History := range histories {
		j := i + 2
		file.SetCellValue(sheet, "A"+strconv.Itoa(j), History.Price)
		file.SetCellValue(sheet, "B"+strconv.Itoa(j), auctionNames[i])
		file.SetCellValue(sheet, "C"+strconv.Itoa(j), History.UserName)
	}

	file.SetColWidth(sheet, "A", "A", 10)
	file.SetColWidth(sheet, "B", "B", 30)
	file.SetColWidth(sheet, "C", "C", 30)

	file.SetActiveSheet(index)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	errFont := pdf.AddTTFFont("poppins", "/home/codeyzx/Data/programming/go/axion-be/assets/fonts/Poppins-Medium.ttf")
	if errFont != nil {
		log.Println("failed to add font")
	}
	errFont = pdf.SetFont("poppins", "", 12)
	if errFont != nil {
		log.Println("failed to set font")
	}

	pdf.AddPage()

	r, _ := file.GetRows(sheet)

	pdf.SetX(20)
	
	for row, rowCells := range r {
		for _, cell := range rowCells {
			err = pdf.Cell(nil, cell)
			if err != nil {
				log.Println(err)
			}

			if cell == "Price" || cell == "Auction" || cell == "User" {
				pdf.SetX(pdf.GetX() + 60)
			}

			if cell == "-" {
				pdf.SetX(pdf.GetX() + 150)
			} else {
				pdf.SetX(pdf.GetX() + 50)
			}
		}

		pdf.Br(40)
		pdf.SetX(20)

		if row%20 == 19 {
			pdf.AddPage()
			pdf.SetX(20)
		}

	}

	c.Set("Content-Disposition", "attachment; filename=history-report.pdf")
	c.Set("Content-Type", "application/pdf")
	errWrite := pdf.Write(c.Response().BodyWriter())
	if errWrite != nil {
		return errWrite
	}
	return nil

}
