package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
)

func main() {
	// Load Go html template engine
	engine := html.New("./views", ".html")

	// Creating the Fiber App
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure the App
	app.Static("/", "./public")

	// Middleware
	m := func(c *fiber.Ctx) error {
		fmt.Println("I'm in the middleware")
		return c.Next()
	}

	// Add Routes
	pages := app.Group("/pages", m)

	pages.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(struct {
			Name string
			Age  int
		}{
			Name: "Alf",
			Age:  34,
		})
	})

	pages.Get("/names/:name/age/:age", func(c *fiber.Ctx) error {
		name := c.Params("name")
		age := c.Params("age")

		return c.Render("index", fiber.Map{
			"Name": name,
			"Age":  age,
		})
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		var body struct {
			Message string
		}
		if err := ctx.BodyParser(&body); err != nil {
			return err
		}

		return ctx.Render("index", fiber.Map{
			"Name":    "Alf!",
			"Message": body.Message,
		})
	})
	// Start App
	log.Fatal(app.Listen(":3000"))
}
