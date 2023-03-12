package route

import (
	"axion/handler"
	"axion/middleware"
	"axion/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RouteInit(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs")
	})

	r.Get("/docs/*", swagger.HandlerDefault)

	r.Static("/public", "./public")

	r.Post("/login", handler.LoginHandler)
	r.Post("/register", middleware.Auth, handler.UserHandlerCreate)
	r.Get("/check-jwt", handler.CheckJWT)

	r.Get("/users", middleware.Operator, handler.UserHandlerGetAll)
	r.Get("/users/:id", handler.UserHandlerGetById)
	r.Put("/users/:id", middleware.ByID, handler.UserHandlerUpdate)
	r.Put("/users/:id/update-email", middleware.ByID, handler.UserHandlerUpdateEmail)
	r.Put("/users/:id/update-role", middleware.Admin, handler.UserHandlerUpdateRole)
	r.Delete("/users/:id", middleware.Admin, handler.UserHandlerDelete)
	r.Get("/users-export-excel", middleware.Operator, handler.UsersExportToExcel)
	r.Get("/users-export-pdf", middleware.Operator, handler.UsersExportToPDF)

	r.Get("/auctions", handler.AuctionHandlerGetAll)
	r.Get("/auctions/:id", handler.AuctionHandlerGetById)
	r.Post("/auctions", middleware.Users, utils.HandleSingleFile, handler.AuctionHandlerCreate)
	r.Put("/auctions/:id", middleware.ByID, handler.AuctionHandlerUpdate)
	r.Delete("/auctions/:id", middleware.ByID, handler.AuctionHandlerDelete)
	r.Get("/auctions-export-excel", middleware.ByID, handler.AuctionExportToExcel)
	r.Get("/auctions-export-pdf", middleware.ByID, handler.AuctionExportToPDF)

	r.Get("/auction-histories", middleware.Admin, handler.AuctionHistoryHandlerGetAll)
	r.Get("/auction-histories/:id", handler.AuctionHistoryHandlerGetById)
	r.Get("/auction-histories/user/:id", middleware.ByID, handler.AuctionHistoryHandlerGetByUser)
	r.Post("/auction-histories", middleware.Users, handler.AuctionHistoryHandlerCreate)
	r.Put("/auction-histories/:id", middleware.Admin, handler.AuctionHistoryHandlerUpdate)
	r.Delete("/auction-histories/:id", middleware.Admin, handler.AuctionHistoryHandlerDelete)
	r.Get("/auction-histories-export-excel", middleware.ByID, handler.AuctionHistoryExportToExcel)
	r.Get("/auction-histories-export-pdf", middleware.ByID, handler.AuctionHistoryExportToPDF)

	r.Get("/products", handler.ProductHandlerGetAll)
	r.Get("/products/:id", handler.ProductHandlerGetById)
	r.Get("/products-by-users/:id", handler.ProductHandlerGetByUserId)
	r.Post("/products", middleware.Operator, utils.HandleSingleFile, handler.ProductHandlerCreate)
	r.Put("/products/:id", middleware.ByID, utils.HandleSingleFile, handler.ProductHandlerUpdate)
	r.Delete("/products/:id", middleware.ByID, handler.ProductHandlerDelete)
	r.Get("/products-export-excel", middleware.ByID, handler.ProductExportToExcel)
	r.Get("/products-export-pdf", middleware.ByID, handler.ProductExportToPDF)

	r.Get("/history", middleware.Admin, handler.HistoryHandlerGetAll)
	r.Get("/history/:id", middleware.Admin, handler.HistoryHandlerGetById)
	r.Post("/history", middleware.Admin, handler.HistoryHandlerCreate)
	r.Put("/history/:id", middleware.Admin, handler.HistoryHandlerUpdate)
	r.Delete("/history/:id", middleware.Admin, handler.HistoryHandlerDelete)
	r.Get("/history-export-excel", middleware.Admin, handler.HistoryExportToExcel)
	r.Get("/history-export-pdf", middleware.Admin, handler.HistoryExportToPDF)

}
