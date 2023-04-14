package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	db "posts_ms/internal/database"
	"posts_ms/internal/posts"
)

type test struct {
	Desc string `json:"desc"`
}

func main() {
	// Connect to database
	err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		err := posts.CreatePost(ctx)
		if err != nil {
			return err
		}
		return ctx.SendString("Document added succesfully!")
	})
	app.Post("/", func(ctx *fiber.Ctx) error { return nil })
	app.Put("/", func(ctx *fiber.Ctx) error { return nil })
	app.Delete("/", func(ctx *fiber.Ctx) error { return nil })

	app.Listen("127.0.0.1:3000")
}
