package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/covers/"

func HandleSingleFile(ctx *fiber.Ctx) error {

	file, errFile := ctx.FormFile("image")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}

	var filename *string
	var pathFile *string

	if file != nil {
		errCheckContentType := checkContentType(file, "image/jpg", "image/jpeg", "image/webp", "image/jfif", "image/png", "image/gif")
		if errCheckContentType != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errCheckContentType.Error(),
			})
		}

		filename = &file.Filename

		newFilename := fmt.Sprintf("%d-%s", time.Now().Unix(), *filename)

		path := fmt.Sprintf("./public/covers/%s", newFilename)
		errSaveFile := ctx.SaveFile(file, path)

		pathFile = &path

		if errSaveFile != nil {
			log.Println("Fail to store file into directory.")
		}
	} else {
		log.Println("Nothing file to uploading.")
	}

	if pathFile != nil {
		ctx.Locals("pathFile", *pathFile)
	} else {
		ctx.Locals("pathFile", nil)
	}

	return ctx.Next()
}

func HandleRemoveFile(filename string, pathFile ...string) error {
	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	}

	return nil
}

func checkContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}

		return errors.New("not allowed file type")
	} else {
		return errors.New("not found content type to be checking")
	}
}
