package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	db "posts_ms/internal/database"
)

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error { return nil })
	app.Post("/", func(ctx *fiber.Ctx) error { return nil })
	app.Put("/", func(ctx *fiber.Ctx) error { return nil })
	app.Delete("/", func(ctx *fiber.Ctx) error { return nil })

	app.Listen("127.0.0.1:3000")
}
