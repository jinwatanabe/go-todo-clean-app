package driver

import (
	"gorm.io/gorm"
)

type TodoDriver interface {
	GetAll() ([]Todo, error)
	GetById(id int) (Todo, error)
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

type Todo struct {
	Id 				int    `gorm:"primaryKey" json:"id"`
	Title 		string `gorm:"size:255" json:"title"j`
	Done 			bool  `gorm:"default:false" json:"done"`
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