package adapters

import (
	"github.com/sinestrea/todo-list/entities"
	"github.com/sinestrea/todo-list/usecases"
	"gorm.io/gorm"
)

type GormTodoRepository struct {
	db *gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) usecases.TodoRepository {
	return &GormTodoRepository{
		db: db,
	}
}

func (r *GormTodoRepository) Create(todo entities.Todo) error {
	return r.db.Create(&todo).Error
}

func (r *GormTodoRepository) Update(todo entities.Todo, id uint) error {
	return r.db.Where("id = ?", id).Updates(&todo).Error
}

func (r *GormTodoRepository) FindById(id uint) (entities.Todo, error) {
	var todo entities.Todo
	return todo, r.db.First(&todo, id).Error
}

func (r *GormTodoRepository) FindTodoList() ([]entities.Todo, error) {
	var todos []entities.Todo
	return todos, r.db.Find(&todos).Error
}

func (r *GormTodoRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Todo{}, id).Error
}