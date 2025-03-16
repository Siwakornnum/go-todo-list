package adapters

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sinestrea/todo-list/entities"
	"github.com/sinestrea/todo-list/usecases"
)

type HttpTodoHandler struct {
	usecase usecases.TodoUseCase
}

func NewHttpTodoHandler(usecase usecases.TodoUseCase) *HttpTodoHandler {
	return &HttpTodoHandler{
		usecase: usecase,
	}
}

func (h *HttpTodoHandler) CreateTodoList(c *fiber.Ctx) error {
	todo := entities.Todo{}
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	if err := h.usecase.CreateTodoList(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "success",
		"data":    todo,
	})
}

func (h *HttpTodoHandler) EditTodoList(c *fiber.Ctx) error {
	todoId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	todo := entities.Todo{}
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	if err := h.usecase.EditTodoList(todo, uint(todoId)); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "success",
		"data":    todo,
	})
}

func (h *HttpTodoHandler) FindTodoById(c *fiber.Ctx) error {
	todoId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	todo, err := h.usecase.FindTodoById(uint(todoId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "success",
		"data":    todo,
	})
}

func (h *HttpTodoHandler) FindTodoList(c *fiber.Ctx) error {
	todos, err := h.usecase.FindTodoList()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "success",
		"data":    todos,
	})
}

func (h *HttpTodoHandler) DeleteTodo(c *fiber.Ctx) error {
	todoId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.usecase.DeleteTodo(uint(todoId)); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "success",
	})
}