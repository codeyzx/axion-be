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

// @Summary Get All Product
// @Description Get All Product
// @Tags Product
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Router /products [get]
func ProductHandlerGetAll(ctx *fiber.Ctx) error {
	var Products []entity.Product

	result := database.DB.Debug().Find(&Products)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(Products)
}

// @Summary Get Product By Id
// @Description Get Product By Id
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /products/{id} [get]
func ProductHandlerGetById(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	var product response.Product
	err := database.DB.First(&product, "id = ?", ID).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    product,
	})
}

// @Summary Get Product Users By Id
// @Description Get Product Users By Id
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /products-by-users/{id} [get]
func ProductHandlerGetByUserId(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var Products []entity.Product

	result := database.DB.Debug().Raw("CALL get_product(?)", userId).Scan(&Products)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	return ctx.JSON(Products)
}

// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param productRequest body request.ProductRequest true "Product Request"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /products [post]
// @Security ApiKeyAuth
func ProductHandlerCreate(ctx *fiber.Ctx) error {

	productId := uuid.New()
	Product := new(request.ProductRequest)
	if err := ctx.BodyParser(Product); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(Product)
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
		Name:        Product.Name,
		Description: Product.Description,
		Price:       Product.Price,

		Image: pathFileString,
	}

	errCreateProduct := database.DB.Create(&newProduct).Error
	if errCreateProduct != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newProduct,
	})
}

// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product Id"
// @Param productRequest body request.ProductRequest true "Product Request"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func ProductHandlerUpdate(ctx *fiber.Ctx) error {
	ProductRequest := new(request.ProductRequest)
	if err := ctx.BodyParser(ProductRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var Product entity.Product

	ID := ctx.Params("id")

	err := database.DB.First(&Product, "id = ?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	userId := ctx.Locals("userId")
	var pathFileString string
	pathFile := ctx.Locals("pathFile")

	if pathFile == nil {
		pathFileString = Product.Image

	} else {
		pathFileString = fmt.Sprintf("%v", pathFile)
	}

	Product.Image = pathFileString
	Product.Price = ProductRequest.Price
	Product.Description = ProductRequest.Description
	Product.Name = ProductRequest.Name

	if userId != 0 {
		result := database.DB.Debug().Exec("CALL update_product(?, ?, ?, ?, ?, ?)", Product.ID, userId, Product.Name, Product.Description, Product.Price, Product.Image)

		if result.Error != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "internal server error",
			})
		}

		if result.RowsAffected == 0 {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "cannot update product",
			})
		}

	} else {
		errUpdate := database.DB.Save(&Product).Error
		if errUpdate != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "internal server error",
			})
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    Product,
	})
}

// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func ProductHandlerDelete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	var Product entity.Product

	userId := ctx.Locals("userId")

	err := database.DB.Debug().First(&Product, "id=?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	if userId != 0 {
		var Auction entity.Auction
		errAuction := database.DB.Debug().First(&Auction, "product_id=? AND user_id=?", ID, userId).Error

		if errAuction != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		updateAuction := database.DB.Model(&Auction).Update("product_id", nil)
		if updateAuction.Error != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "failed to update auction",
			})
		}

		deleteProduct := database.DB.Delete(&Product)
		if deleteProduct.Error != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "failed to delete product",
			})
		}

	} else {
		var Auction entity.Auction
		errAuction := database.DB.Debug().First(&Auction, "product_id=?", ID).Error

		if errAuction == nil {
			errDeleteAuction := database.DB.Debug().Model(&Product).Association("Auctions").Delete(&Auction)

			if errDeleteAuction != nil {
				return ctx.Status(500).JSON(fiber.Map{
					"message": "internal server error",
				})
			}
		}

		errDeleteProduct := database.DB.Debug().Delete(&Product).Error
		if errDeleteProduct != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "internal server error",
			})
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "Product was deleted",
	})
}

// @Summary Export Product to Excel
// @Description Export Product to Excel
// @Tags Product
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /products-export-excel [get]
// @Security ApiKeyAuth
func ProductExportToExcel(c *fiber.Ctx) error {
	temp := c.Locals("userId")

	var Products []entity.Product

	if temp != 0 {
		result := database.DB.Debug().Raw("CALL get_product(?)", temp).Scan(&Products)
		if result.Error != nil {
			log.Println(result.Error)
		}
	} else {
		result := database.DB.Debug().Find(&Products)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}

	file := excelize.NewFile()
	const sheet = "Products"

	index, err := file.NewSheet(sheet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	file.SetCellValue(sheet, "A1", "ID")
	file.SetCellValue(sheet, "B1", "Name")
	file.SetCellValue(sheet, "C1", "Description")
	file.SetCellValue(sheet, "D1", "Price")
	file.SetCellValue(sheet, "E1", "Image")

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

	for i, Product := range Products {
		i = i + 2
		file.SetCellValue(sheet, "A"+strconv.Itoa(i), Product.ID)
		file.SetCellValue(sheet, "B"+strconv.Itoa(i), Product.Name)
		file.SetCellValue(sheet, "C"+strconv.Itoa(i), Product.Description)
		file.SetCellValue(sheet, "D"+strconv.Itoa(i), Product.Price)

		graphicOptions := excelize.GraphicOptions{
			AutoFit: true,
		}
		if Product.Image != "" {
			var imagePath = strings.Replace(Product.Image, "./public/covers/", "", -1)
			errImage := file.AddPicture(sheet, "E"+strconv.Itoa(i), "/home/codeyzx/Data/programming/go/axion-be/public/covers/"+imagePath, &graphicOptions)

			if errImage != nil {
				fmt.Println(errImage)
			}
		}
	}

	file.SetColWidth(sheet, "A", "A", 20)
	file.SetColWidth(sheet, "B", "B", 30)
	file.SetColWidth(sheet, "C", "C", 30)
	file.SetColWidth(sheet, "D", "D", 30)
	file.SetColWidth(sheet, "E", "E", 30)

	file.SetActiveSheet(index)

	c.Set("Content-Disposition", "attachment; filename=transaction-report.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	errWrite := file.Write(c.Response().BodyWriter())
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// @Summary Export Product to PDF
// @Description Export Product to PDF
// @Tags Product
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /products-export-pdf [get]
// @Security ApiKeyAuth
func ProductExportToPDF(c *fiber.Ctx) error {
	temp := c.Locals("userId")

	var Products []entity.Product

	if temp != 0 {
		result := database.DB.Debug().Raw("CALL get_product(?)", temp).Scan(&Products)
		if result.Error != nil {
			log.Println(result.Error)
		}
	} else {
		result := database.DB.Debug().Find(&Products)
		if result.Error != nil {
			log.Println(result.Error)
		}
	}

	file := excelize.NewFile()
	const sheet = "Products"

	index, err := file.NewSheet(sheet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	file.SetCellValue(sheet, "A1", "Image")
	file.SetCellValue(sheet, "B1", "Name")
	file.SetCellValue(sheet, "C1", "Price")

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

	for i, Product := range Products {
		i = i + 2

		file.SetCellValue(sheet, "A"+strconv.Itoa(i), Product.Image)
		file.SetCellValue(sheet, "B"+strconv.Itoa(i), Product.Name)
		file.SetCellValue(sheet, "C"+strconv.Itoa(i), Product.Price)
		graphicOptions := excelize.GraphicOptions{
			AutoFit: true,
		}
		if Product.Image != "" {
			var imagePath = strings.Replace(Product.Image, "./public/covers/", "", -1)
			errImage := file.AddPicture(sheet, "C"+strconv.Itoa(i), "/home/codeyzx/Data/programming/go/axion-be/public/covers/"+imagePath, &graphicOptions)

			if errImage != nil {
				fmt.Println(errImage)
			}
		}
	}

	file.SetColWidth(sheet, "A", "A", 20)
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
	for row, rowCells := range r {
		for _, cell := range rowCells {

			if row > 0 && cell != "" && rowCells[0] != "" && strings.Contains(cell, "./public/covers/") {

				errImage := pdf.Image(cell, pdf.GetX(), pdf.GetY(), &gopdf.Rect{W: 50, H: 50})
				if errImage != nil {
					log.Println(errImage)
				}
				pdf.SetX(pdf.GetX() + 50)
				pdf.SetY(pdf.GetY() + 10)
				continue
			}

			err = pdf.Cell(nil, cell)
			if err != nil {
				log.Println(err)
			}

			pdf.SetX(pdf.GetX() + 50)
			pdf.SetY(pdf.GetY() + 10)
		}

		pdf.Br(30)
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
