package main

import (
	"mongodb-api/app"
	"mongodb-api/configs"
	"mongodb-api/repository"
	"mongodb-api/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDb()
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDb := repository.NewTodoRepositoryDb(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/addTodo", td.CreateTodo)
	appRoute.Get("/api/getAllTodos", td.GetAllTodo)
	appRoute.Delete("/api/todo/delete/:id", td.DeleteTodo)
	appRoute.Listen(":8080")
}