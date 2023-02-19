package handler

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/utils"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
// @Router /users/{id}/email [put]
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
// @Router /users/{id}/role [put]
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
