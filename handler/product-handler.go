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
		pathFileString = ""

	} else {
		pathFileString = fmt.Sprintf("%v", pathFile)
	}

	Product.Image = pathFileString

	if ProductRequest.Price != 0 {
		Product.Price = ProductRequest.Price
	}

	if ProductRequest.Description != "" {
		Product.Description = ProductRequest.Description
	}

	if ProductRequest.Name != "" {
		Product.Name = ProductRequest.Name
	}

	log.Println("userId", userId)
	if userId != 0 {
		log.Println("MASUK UPDATE")
		result := database.DB.Debug().Exec("CALL update_product(?, ?, ?, ?, ?, ?)", Product.ID, userId, Product.Name, Product.Description, Product.Price, Product.Image)

		log.Println(result)
		log.Println(result.RowsAffected)

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

// func ProductHandlerDelete(ctx *fiber.Ctx) error {
// 	ID := ctx.Params("id")
// 	var Product entity.Product

// 	userId := ctx.Locals("userId")

// 	err := database.DB.Debug().First(&Product, "id=?", ID).Error
// 	if err != nil {
// 		return ctx.Status(404).JSON(fiber.Map{
// 			"message": "Product not found",
// 		})
// 	}

// 	log.Println("userId", userId)
// 	if userId != 0 {
// 		result := database.DB.Debug().Exec("CALL delete_product(?, ?)", Product.ID, userId)

// 		if result.Error != nil {
// 			return ctx.Status(500).JSON(fiber.Map{
// 				"message": "internal server error",
// 			})
// 		}

// 		if result.RowsAffected == 0 {
// 			return ctx.Status(500).JSON(fiber.Map{
// 				"message": "cannot delete product",
// 			})
// 		}

// 	} else {
// 		errDelete := database.DB.Debug().Delete(&Product).Error
// 		if errDelete != nil {
// 			return ctx.Status(500).JSON(fiber.Map{
// 				"message": "internal server error",
// 			})
// 		}
// 	}

// 	return ctx.JSON(fiber.Map{
// 		"message": "Product was deleted",
// 	})
// }
