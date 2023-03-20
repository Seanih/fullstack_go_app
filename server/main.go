package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("Hey pussies")

	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		if len(todos) > 0 {
			return c.JSON(todos)
		}

		return c.SendString("no items have been added yet")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// initialize new Todo struct
		newTodo := &Todo{}

		// handle parsing error if needed
		if err := c.BodyParser(newTodo); err != nil {
			return err
		}

		// create new IDs for new Todo struct
		newTodo.ID = len(todos) + 1

		// add new Todo struct to slice
		todos = append(todos, *newTodo)

		// return the updated slice as JSON
		return c.JSON(todos)
	})

	app.Get("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		for _, value := range todos {
			if value.ID == id {
				return c.JSON(value)
			}
		}

		return c.SendString("no item matching that ID")
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}

		index := 0
		for i, val := range todos {
			if val.ID == id {
				todos[i].Done = true
				index = i
				break
			}
		}

		return c.JSON(todos[index])
	})

	log.Fatal(app.Listen(":4000"))
}
