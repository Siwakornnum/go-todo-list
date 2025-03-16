package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/sinestrea/todo-list/adapters"
	"github.com/sinestrea/todo-list/usecases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host = "localhost"
	port = 5432
	databaseName = "mydatabase"
	username = "myuser"
	password = "mypassword"
)

func SetupDatabase() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, username, password, databaseName)

		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Enable color
			},
		)

  db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: newLogger,
	})

  if err != nil {
		panic("failed to connect database")
  }

  return db
}

func main() {
	app := fiber.New()

	db := SetupDatabase()

	// db.AutoMigrate(&entities.Todo{})

	todoRepo := adapters.NewGormTodoRepository(db)
	todoService := usecases.NewTodoService(todoRepo)
	todoHandler := adapters.NewHttpTodoHandler(todoService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/todo/", todoHandler.FindTodoList)

	app.Get("/todo/:id", todoHandler.FindTodoById)

	app.Post("/todo", todoHandler.CreateTodoList)

	app.Put("/todo/:id", todoHandler.EditTodoList)

	app.Delete("/todo/:id", todoHandler.DeleteTodo)

	app.Listen(":8080")
}