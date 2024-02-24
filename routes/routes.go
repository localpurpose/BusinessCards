package routes

import (
	"github.com/buscard/controller"
	"github.com/gofiber/fiber/v2"
)

type (
	Host struct {
		Fiber *fiber.App
	}
)

func SetupDefaultRoutes(app *fiber.App) {

	//------------------------------------------------
	// Registration
	//------------------------------------------------
	app.Get("/create", controller.RenderRegister)
	app.Post("/create", controller.DoRegister)

	//------------------------------------------------
	// API
	//------------------------------------------------

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Redirect("/1") // Redirect to the creator page
	})

	app.Get("/api/user/:userid", controller.GetDetails)

	app.Post("/api/user/create", controller.CreateUser)
	app.Post("/api/user/edit/:userid", controller.EditUser)
	app.Post("/api/user/delete/:userid", controller.DeleteUser)

	//Upload Image Handler.
	app.Post("/api/uploadimage", controller.UploadImage)

	//------------------------------------------------
	// Main Application
	//------------------------------------------------
	app.Get("/", controller.RenderMain)
	app.Get("/user/:userid", controller.RenderUserProfile)

	//------------------------------------------------
	// 404 Middleware
	//------------------------------------------------
	app.Use(func(c *fiber.Ctx) error {
		return c.Render("home/404", nil) // => 404 "Not Found"
	})

}
