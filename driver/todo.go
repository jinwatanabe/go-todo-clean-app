package driver

import (
	"gorm.io/gorm"
)

type TodoDriver interface {
	GetAll() ([]Todo, error)
	GetById(id int) (Todo, error)
	Create(todo CreateTodo) (error)
	Update(id int, todo UpdateTodo) (error)
}

type TodoDriverImpl struct {
	conn *gorm.DB
}

func (t TodoDriverImpl) GetAll() ([]Todo, error) {
	todos := []Todo{}
	t.conn.Find(&todos)

	return todos, nil
}

func (t TodoDriverImpl) GetById(id int) (Todo, error) {
	todo := Todo{}
	t.conn.First(&todo, id)
	return todo, nil
}

func (t TodoDriverImpl) Create(todo CreateTodo) (error) {
	err := t.conn.Create(&todo)

	if err != nil {
		return err.Error
	}

	return nil
}

func (t TodoDriverImpl) Update(id int, todo UpdateTodo) (error) {
	err := t.conn.Model(&todo).Where("id = ?", id).Updates(todo)

	if err != nil {
		return err.Error
	}
	
	return nil
}

type Todo struct {
	Id 				int    `gorm:"primaryKey" json:"id"`
	Title 		string `gorm:"size:255" json:"title"`
	Done 			bool  `gorm:"default:false" json:"done"`
}

type CreateTodo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (CreateTodo) TableName() string {
	return "todos"
}

type UpdateTodo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (UpdateTodo) TableName() string {
	return "todos"
}

func NewTodo(id int, title string, done bool) Todo {
	return Todo{
		Id: id,
		Title: title,
		Done: done,
	}
}

func ProvideTodoDriver(conn *gorm.DB) TodoDriver {
	return TodoDriverImpl{conn: conn}
}