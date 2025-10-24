package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	// Create a new Fiber app instance
	app := fiber.New()

	todos := []Todo{}

	// Define a GET route for the root path ("/")
	app.Get("/", func(c *fiber.Ctx) error {
		// Respond with a JSON object {"msg": "hello world"} and HTTP status 200
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	// get todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Define a POST route for creating a new todo item
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// Create a new empty Todo struct to hold the request data
		todo := &Todo{}

		// Parse the JSON request body into the todo struct
		if err := c.BodyParser(todo); err != nil {
			// If parsing fails, return a 400 Bad Request with error details
			return c.Status(400).JSON(fiber.Map{"error": "failed to parse request body", "details": err.Error()})
		}

		// Validate that the 'body' field is not empty
		if todo.Body == "" {
			// If 'body' is missing, return a 400 Bad Request
			return c.Status(400).JSON(fiber.Map{"error": "body field is required"})
		}

		// Assign a unique ID to the new todo (based on current length of todos slice)
		todo.ID = len(todos) + 1
		// Add the new todo to the todos slice (in-memory storage)
		todos = append(todos, *todo)
		// Return the created todo as JSON with a 201 Created status
		return c.Status(201).JSON(todo)
	})

	// update todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		// If no todo is found with the given ID, return a 404 Not Found
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})

	// delete todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				// Remove the todo from the slice
				todos = append(todos[:i], todos[i+1:]...)
				return c.SendStatus(204) // No Content
			}
		}
		// If no todo is found with the given ID, return a 404 Not Found
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})

	// Start the Fiber app and listen on port 4000
	// If the server fails to start, log the error and exit
	log.Fatal(app.Listen(":4000"))
}
