package usecases

import (
	"github.com/sinestrea/todo-list/entities"
)

type TodoRepository interface {
	Create(todo entities.Todo) error
	Update(todo entities.Todo, id uint) error
	FindById(id uint) (entities.Todo, error)
	FindTodoList() ([]entities.Todo, error)
	Delete(id uint) error
}
