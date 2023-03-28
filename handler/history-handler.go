package handler

import (
	"axion/database"
	"axion/model/entity"
	"axion/model/request"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/signintech/gopdf"
	"github.com/xuri/excelize/v2"
)

// @Summary Get All History
// @Description Get All History
// @Tags History
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Router /history [get]
// @Security ApiKeyAuth
func HistoryHandlerGetAll(ctx *fiber.Ctx) error {
	var Histories []entity.History

	result := database.DB.Debug().Find(&Histories)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(Histories)
}

// @Summary Get History By Id
// @Description Get History By Id
// @Tags History
// @Accept  json
// @Produce  json
// @Param id path string true "History Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /history/{id} [get]
// @Security ApiKeyAuth
func HistoryHandlerGetById(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	var history entity.History
	err := database.DB.First(&history, "id = ?", ID).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "History not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    history,
	})
}

// @Summary Create History
// @Description Create History
// @Tags History
// @Accept  json
// @Produce  json
// @Param History body request.HistoryRequest true "History"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /history [post]
// @Security ApiKeyAuth
func HistoryHandlerCreate(ctx *fiber.Ctx) error {
	History := new(request.HistoryRequest)
	if err := ctx.BodyParser(History); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(History)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newHistory := entity.History{
		Log: History.Log,
	}

	errCreateHistory := database.DB.Create(&newHistory).Error
	if errCreateHistory != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newHistory,
	})
}

// @Summary Update History
// @Description Update History
// @Tags History
// @Accept  json
// @Produce  json
// @Param id path string true "History Id"
// @Param History body request.HistoryRequest true "History"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /history/{id} [put]
// @Security ApiKeyAuth
func HistoryHandlerUpdate(ctx *fiber.Ctx) error {
	HistoryRequest := new(request.HistoryRequest)
	if err := ctx.BodyParser(HistoryRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var History entity.History

	ID := ctx.Params("id")

	err := database.DB.First(&History, "id = ?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "History not found",
		})
	}

	if HistoryRequest.Log != "" {
		History.Log = HistoryRequest.Log
	}

	errUpdate := database.DB.Save(&History).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    History,
	})
}

// @Summary Delete History
// @Description Delete History
// @Tags History
// @Accept  json
// @Produce  json
// @Param id path string true "History Id"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /history/{id} [delete]
// @Security ApiKeyAuth
func HistoryHandlerDelete(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	var History entity.History

	err := database.DB.Debug().First(&History, "id=?", ID).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "History not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&History).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "History was deleted",
	})
}

// @Summary Export History to Excel
// @Description Export History to Excel
// @Tags History
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /history-export-excel [get]
// @Security ApiKeyAuth
func HistoryExportToExcel(c *fiber.Ctx) error {
	var Histories []entity.History

	result := database.DB.Debug().Find(&Histories)
	if result.Error != nil {
		log.Println(result.Error)
	}

	file := excelize.NewFile()
	const sheet = "History"

	index, err := file.NewSheet(sheet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	file.SetCellValue(sheet, "A1", "ID")
	file.SetCellValue(sheet, "B1", "Log")
	file.SetCellValue(sheet, "C1", "Created At")
	file.SetCellValue(sheet, "D1", "Updated At")

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

	for i, History := range Histories {
		i = i + 2
		file.SetCellValue(sheet, "A"+strconv.Itoa(i), History.ID)
		file.SetCellValue(sheet, "B"+strconv.Itoa(i), History.Log)
		file.SetCellValue(sheet, "C"+strconv.Itoa(i), History.CreatedAt)
		file.SetCellValue(sheet, "D"+strconv.Itoa(i), History.UpdatedAt)
	}

	file.SetColWidth(sheet, "A", "A", 10)
	file.SetColWidth(sheet, "B", "B", 30)
	file.SetColWidth(sheet, "C", "C", 30)
	file.SetColWidth(sheet, "D", "D", 30)

	file.SetActiveSheet(index)

	c.Set("Content-Disposition", "attachment; filename=history-report.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	errWrite := file.Write(c.Response().BodyWriter())
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// @Summary Export History to PDF
// @Description Export History to PDF
// @Tags History
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Router /history-export-pdf [get]
// @Security ApiKeyAuth
func HistoryExportToPDF(c *fiber.Ctx) error {
	var Histories []entity.History

	result := database.DB.Debug().Find(&Histories)
	if result.Error != nil {
		log.Println(result.Error)
	}

	file := excelize.NewFile()
	const sheet = "History"

	index, err := file.NewSheet(sheet)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create sheet",
		})
	}

	file.SetCellValue(sheet, "A1", "ID")
	file.SetCellValue(sheet, "B1", "Log")

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

	file.SetCellStyle(sheet, "A1", "B1", style)

	for i, History := range Histories {
		i = i + 2
		file.SetCellValue(sheet, "A"+strconv.Itoa(i), History.ID)
		file.SetCellValue(sheet, "B"+strconv.Itoa(i), History.Log)
	}

	file.SetColWidth(sheet, "A", "A", 10)
	file.SetColWidth(sheet, "B", "B", 30)

	file.SetActiveSheet(index)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	errFont := pdf.AddTTFFont("poppins", "assets/fonts/Poppins-Medium.ttf")
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

			pdf.SetX(pdf.GetX() + 50)
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
