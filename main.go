package main

import (
	"image-api-go/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/download", controller.File)

	err := app.Listen(":9002")
	if err != nil {
		panic(err)
	}
}
