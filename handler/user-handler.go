package handler

import (
	"axion/database"
	"axion/model/entity"
	"axion/model/request"
	"axion/utils"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/signintech/gopdf"
	"github.com/xuri/excelize/v2"
)

// @Summary Get All User
// @Description Get All User
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Router /users [get]
// @Security ApiKeyAuth
func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)
}

// @Summary Get User By Id
// @Description Get User By Id
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /users/{id} [get]
func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// @Summary Create User
// @Description Create User
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body request.UserCreateRequest true "User Create Request"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /register [post]
// @Security ApiKeyAuth
func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	userId := ctx.Locals("userId")

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	if userId == 0 {
		if user.Role != "" && user.Role != entity.Users {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "Role must be Users",
			})
		}
	} else {
		if user.Role != "" && user.Role != entity.Admin && user.Role != entity.Operator && user.Role != entity.Users {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "Role must be Admin, Operator or Users",
			})
		}
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
		Role:    user.Role,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {

		errCreateUser := strings.Split(errCreateUser.Error(), ":")[0]
		log.Println(errCreateUser)
		if errCreateUser == "Error 1062 (23000)" {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "email already exist",
			})
		} else {
			return ctx.Status(500).JSON(fiber.Map{
				"message": "failed to store data",
			})
		}
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newUser,
	})
}

// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Param user body request.UserUpdateRequest true "User Update Request"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /users/{id} [put]
// @Security ApiKeyAuth
func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	temp := ctx.Locals("userId")
	authId := fmt.Sprintf("%v", temp)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")

	log.Println(userId + " " + authId)
	if temp != 0 {
		if userId != authId {
			return ctx.Status(403).JSON(fiber.Map{
				"message": "forbidden",
			})
		}
	}

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	if userRequest.Phone != "" {
		user.Phone = userRequest.Phone
	}

	if userRequest.Address != "" {
		user.Address = userRequest.Address
	}

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// @Summary Update User Email
// @Description Update User Email
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Param user body request.UserEmailRequest true "User Email Request"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /users/{id}/update-email [put]
// @Security ApiKeyAuth
func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)
	temp := ctx.Locals("userId")
	authId := fmt.Sprintf("%v", temp)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User
	var isEmailUserExist entity.User

	userId := ctx.Params("id")

	log.Println(userId + " " + authId)
	if temp != 0 {
		if userId != authId {
			return ctx.Status(403).JSON(fiber.Map{
				"message": "forbidden",
			})
		}
	}

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errCheckEmail := database.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error
	if errCheckEmail == nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "email already used.",
		})
	}

	user.Email = userRequest.Email

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// @Summary Update User Role
// @Description Update User Role
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Param user body request.UserRoleRequest true "User Role Request"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /users/{id}/update-role [put]
// @Security ApiKeyAuth
func UserHandlerUpdateRole(ctx *fiber.Ctx) error {
	userRequest := new(request.UserRoleRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	user.Role = userRequest.Role

	if user.Role == "" && user.Role != entity.Admin && user.Role != entity.Operator && user.Role != entity.Users {
		log.Println("masuk")
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Role must be Admin, Operator or Users",
		})
	}

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /users/{id} [delete]
// @Security ApiKeyAuth
func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User

	err := database.DB.Debug().First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}

// @Summary Export Users to Excel
// @Description Export Users to Excel
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /users-export-excel [get]
// @Security ApiKeyAuth
func UsersExportToExcel(c *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	f := excelize.NewFile()

	index, err := f.NewSheet("Sheet1")

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	f.SetCellValue("Sheet1", "A1", "No")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Email")
	f.SetCellValue("Sheet1", "D1", "Phone")
	f.SetCellValue("Sheet1", "E1", "Address")
	f.SetCellValue("Sheet1", "F1", "Role")

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	if err != nil {

		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create style",
		})
	}

	f.SetCellStyle("Sheet1", "A1", "F1", style)

	for i, user := range users {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), i+1)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), user.Name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), user.Email)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), user.Phone)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), user.Address)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(i+2), user.Role)
	}

	f.SetColWidth("Sheet1", "A", "A", 5)
	f.SetColWidth("Sheet1", "B", "B", 30)
	f.SetColWidth("Sheet1", "C", "C", 30)
	f.SetColWidth("Sheet1", "D", "D", 20)
	f.SetColWidth("Sheet1", "E", "E", 30)
	f.SetColWidth("Sheet1", "F", "F", 20)

	f.SetActiveSheet(index)

	c.Set("Content-Disposition", "attachment; filename=users-report.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	errWrite := f.Write(c.Response().BodyWriter())
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// @Summary Export Users to PDF
// @Description Export Users to PDF
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /users-export-pdf [get]
// @Security ApiKeyAuth
func UsersExportToPDF(c *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	f := excelize.NewFile()

	index, err := f.NewSheet("Sheet1")

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	f.SetCellValue("Sheet1", "A1", "No")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Email")

	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	if err != nil {

		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create style",
		})
	}

	f.SetCellStyle("Sheet1", "A1", "E1", style)

	for i, user := range users {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), user.ID)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), user.Name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), user.Email)
	}

	f.SetColWidth("Sheet1", "A", "A", 5)
	f.SetColWidth("Sheet1", "B", "B", 30)
	f.SetColWidth("Sheet1", "C", "C", 30)

	f.SetActiveSheet(index)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	errFont := pdf.AddTTFFont("poppins", "/home/codeyzx/Data/programming/go/axion-be/assets/fonts/Poppins-Medium.ttf")
	if errFont != nil {
		log.Println("failed to add font")
	}
	errFont = pdf.SetFont("poppins", "", 14)
	if errFont != nil {
		log.Println("failed to set font")
	}

	pdf.AddPage()

	r, err := f.GetRows("Sheet1")
	for row, rowCells := range r {
		for _, cell := range rowCells {

			err = pdf.Cell(nil, cell)
			if err != nil {
				log.Println(err)
			}

			pdf.SetX(pdf.GetX() + 100)
		}

		pdf.Br(30)
		pdf.SetX(20)

		if row%20 == 19 {
			pdf.AddPage()
			pdf.SetX(20)
		}

	}

	c.Set("Content-Disposition", "attachment; filename=users-report.pdf")
	c.Set("Content-Type", "application/pdf")
	errWrite := pdf.Write(c.Response().BodyWriter())
	if errWrite != nil {
		return errWrite
	}
	return nil

}
