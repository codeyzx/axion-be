package handler

import (
	"axion/database"
	"axion/model/entity"
	"axion/model/request"
	"axion/model/response"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/signintech/gopdf"
	"github.com/xuri/excelize/v2"
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

	result := database.DB.Preload("Product").Preload("User").Preload("Bidder").Find(&auctions)
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

	err := database.DB.Where("auctions.id = ?", ID).Preload("Product").Preload("User").Preload("Bidder").Preload("AuctionHistory").First(&auction).Error

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
			"message": "failed to parse request body",
			"error":   err.Error(),
		})
	}

	log.Println("auction : ", auction)
	validate := validator.New()
	errValidate := validate.Struct(auction)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed to validate request body",
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

	var bidId *uint
	if auction.BidderId == 0 {
		bidId = nil
	} else {
		bidId = &auction.BidderId
	}

	var userId *uint
	if auction.UserId == 0 {
		userId = nil
	} else {
		userId = &auction.UserId
	}

	newProduct := entity.Product{
		ID:          productId,
		Name:        auction.ProductName,
		Description: auction.Description,
		Price:       auction.Price,
		Image:       pathFileString,
	}

	newAuction := entity.Auction{
		Name:         auction.Name,
		ProductID:    &productId,
		LastPrice:    auction.Price,
		UserId:       userId,
		BidderId:     bidId,
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

	if auctionRequest.Name != "" {
		auction.Name = auctionRequest.Name
	}

	if auctionRequest.LastPrice != 0 {
		auction.LastPrice = auctionRequest.LastPrice
	}

	if auctionRequest.EndAt != "" {
		auction.EndAt = auctionRequest.EndAt
	}

	if auctionRequest.BiddersCount != 0 {
		auction.BiddersCount = auctionRequest.BiddersCount
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

// @Summary Export Auction to Excel
// @Description Export Auction to Excel
// @Tags Auction
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auctions-export-excel [get]
// @Security ApiKeyAuth
func AuctionExportToExcel(c *fiber.Ctx) error {
	temp := c.Locals("userId")

	var auctions []response.Auction

	if temp != 0 {
		result := database.DB.Preload("Product").Preload("User").Preload("Bidder").Where("user_id = ?", temp).Find(&auctions)
		if result.Error != nil {
			log.Println(result.Error)
		}
	} else {
		result := database.DB.Preload("Product").Preload("User").Preload("Bidder").Find(&auctions)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}

	file := excelize.NewFile()
	const sheet = "Auctions"

	index, err := file.NewSheet(sheet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	file.SetCellValue(sheet, "A1", "ID")
	file.SetCellValue(sheet, "B1", "Name")
	file.SetCellValue(sheet, "C1", "Last Price")
	file.SetCellValue(sheet, "D1", "Status")
	file.SetCellValue(sheet, "E1", "BidCount")
	file.SetCellValue(sheet, "F1", "Bidder")
	file.SetCellValue(sheet, "G1", "User")
	file.SetCellValue(sheet, "H1", "Product")
	file.SetCellValue(sheet, "I1", "Image")
	file.SetCellValue(sheet, "J1", "End At")

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

	file.SetCellStyle(sheet, "A1", "J1", style)

	for i, auction := range auctions {
		i = i + 2
		file.SetCellValue(sheet, "A"+strconv.Itoa(i), auction.ID)
		file.SetCellValue(sheet, "B"+strconv.Itoa(i), auction.Name)
		file.SetCellValue(sheet, "C"+strconv.Itoa(i), auction.LastPrice)
		file.SetCellValue(sheet, "D"+strconv.Itoa(i), auction.Status)
		file.SetCellValue(sheet, "E"+strconv.Itoa(i), auction.BiddersCount)
		file.SetCellValue(sheet, "F"+strconv.Itoa(i), auction.Bidder.Name)
		file.SetCellValue(sheet, "G"+strconv.Itoa(i), auction.User.Name)
		file.SetCellValue(sheet, "H"+strconv.Itoa(i), auction.Product.Name)
		file.SetCellValue(sheet, "J"+strconv.Itoa(i), auction.EndAt)

		graphicOptions := excelize.GraphicOptions{
			AutoFit: true,
		}

		log.Println(auction.Product)
		if auction.Product.Image != "" {
			var imagePath = strings.Replace(auction.Product.Image, "./public/covers/", "", -1)
			errImage := file.AddPicture(sheet, "I"+strconv.Itoa(i), "/home/codeyzx/Data/programming/go/axion-be/public/covers/"+imagePath, &graphicOptions)

			if errImage != nil {
				fmt.Println("err:::", errImage)
			}
		}
	}

	file.SetColWidth(sheet, "A", "A", 10)
	file.SetColWidth(sheet, "B", "B", 20)
	file.SetColWidth(sheet, "C", "C", 20)
	file.SetColWidth(sheet, "D", "D", 20)
	file.SetColWidth(sheet, "E", "E", 20)
	file.SetColWidth(sheet, "F", "F", 20)
	file.SetColWidth(sheet, "G", "G", 20)
	file.SetColWidth(sheet, "H", "H", 20)
	file.SetColWidth(sheet, "I", "I", 20)
	file.SetColWidth(sheet, "J", "J", 20)

	file.SetActiveSheet(index)

	c.Set("Content-Disposition", "attachment; filename=auction-report.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	errWrite := file.Write(c.Response().BodyWriter())
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// @Summary Export Auction to PDF
// @Description Export Auction to PDF
// @Tags Auction
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /auctions-export-pdf [get]
// @Security ApiKeyAuth
func AuctionExportToPDF(c *fiber.Ctx) error {
	temp := c.Locals("userId")

	var auctions []response.Auction

	if temp != 0 {
		result := database.DB.Preload("Product").Preload("User").Where("user_id = ?", temp).Find(&auctions)
		if result.Error != nil {
			log.Println(result.Error)
		}
	} else {
		result := database.DB.Preload("Product").Preload("User").Find(&auctions)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}

	var auctionNames []string
	for _, auction := range auctions {
		if auction.Name == "" {
			auction.Name = "-"
		}
		auctionNames = append(auctionNames, auction.Name)
	}

	file := excelize.NewFile()
	const sheet = "Auctions"

	index, err := file.NewSheet(sheet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	file.SetCellValue(sheet, "A1", "Name")
	file.SetCellValue(sheet, "B1", "Last Price")
	file.SetCellValue(sheet, "C1", "Status")
	file.SetCellValue(sheet, "D1", "Product")

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

	file.SetCellStyle(sheet, "A1", "D1", style)

	for i, auction := range auctions {

		j := i + 2
		file.SetCellValue(sheet, "A"+strconv.Itoa(j), auctionNames[i])
		file.SetCellValue(sheet, "B"+strconv.Itoa(j), auction.LastPrice)
		file.SetCellValue(sheet, "C"+strconv.Itoa(j), auction.Status)
		file.SetCellValue(sheet, "D"+strconv.Itoa(j), auction.Product.Name)
	}

	file.SetColWidth(sheet, "A", "A", 10)
	file.SetColWidth(sheet, "B", "B", 20)
	file.SetColWidth(sheet, "C", "C", 20)
	file.SetColWidth(sheet, "D", "D", 20)

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
	for row, rowCells := range r {
		for _, cell := range rowCells {

			err = pdf.Cell(nil, cell)
			if err != nil {
				log.Println(err)
			}

			if cell == "-" {
				pdf.SetX(pdf.GetX() + 100)
			} else {
				pdf.SetX(pdf.GetX() + 50)
			}
		}

		pdf.Br(30)
		pdf.SetX(20)

		if row%20 == 19 {
			pdf.AddPage()
			pdf.SetX(20)
		}

	}

	c.Set("Content-Disposition", "attachment; filename=auction-report.pdf")
	c.Set("Content-Type", "application/pdf")
	errWrite := pdf.Write(c.Response().BodyWriter())
	if errWrite != nil {
		return errWrite
	}
	return nil
}
