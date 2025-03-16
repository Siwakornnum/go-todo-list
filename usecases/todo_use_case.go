package usecases

import (
	"github.com/sinestrea/todo-list/entities"
)

type TodoUseCase interface {
	CreateTodoList(todo entities.Todo) error
	EditTodoList(todo entities.Todo, id uint) error
	FindTodoById(id uint) (entities.Todo, error)
	FindTodoList() ([]entities.Todo, error)
	DeleteTodo(id uint) error
}

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoUseCase {
	return &TodoService{
		repo: repo,
	}
}

func (t *TodoService) FindTodoById(id uint) (entities.Todo, error) {
	return t.repo.FindById(id)
}

func (t *TodoService) CreateTodoList(todo entities.Todo) error {
	return t.repo.Create(todo)
}

func (t *TodoService) EditTodoList(todo entities.Todo, id uint) error {
	return t.repo.Update(todo, id)
}

func (t *TodoService) FindTodoList() ([]entities.Todo, error) {
	return t.repo.FindTodoList()
}

func (t *TodoService) DeleteTodo(id uint) error {
	return t.repo.Delete(id)
}
