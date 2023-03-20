package middleware

import (
	"axion/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Admin(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "Admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	return ctx.Next()
}

func Operator(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "Operator" && role != "Admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	return ctx.Next()
}

func Users(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "Users" && role != "Admin" && role != "Operator" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	return ctx.Next()
}

func ByID(ctx *fiber.Ctx) error {
	var userId float64
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)

	if role != "Users" && role != "Admin" && role != "Operator" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	if role == "Users" {

		id := claims["id"].(float64)
		userId = id

	} else {
		userId = 0
	}
	if userId != 0 {

		ctx.Locals("userId", userId)
	} else {
		ctx.Locals("userId", 0)
	}

	return ctx.Next()
}

func Auth(ctx *fiber.Ctx) error {
	var userId float64
	token := ctx.Get("Authorization")
	if token == "" {
		ctx.Locals("userId", 0)
		return ctx.Next()
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)

	if role != "Admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	if role == "Admin" {
		id := claims["id"].(float64)
		userId = id
	} else {
		userId = 0
	}

	ctx.Locals("userId", userId)

	return ctx.Next()
}
