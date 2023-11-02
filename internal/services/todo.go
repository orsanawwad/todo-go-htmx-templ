package services

import (
	// "fmt"
	"log"
	// "slices"

	"github.com/orsanawwad/htmxdemo/internal/database"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Text string
	Done bool
}

type Todos struct {
	// todos []Todo
	db *database.Database
}

func New(db *database.Database) *Todos {
	db.DB.AutoMigrate(&Todo{})
	return &Todos{db}
}

func (s *Todos) List() ([]Todo, error) {
	var todos []Todo
	query := s.db.DB.Find(&todos)
	if query.Error != nil {
		log.Fatal(query.Error)
	}
	return todos, nil
}

func (s *Todos) Create(val string) error {
	result := s.db.DB.Create(&Todo{Text: val})
	// s.todos = slices.Insert(s.todos, pos, Todo{Text: val, Done: false})
	return result.Error
}

func (s *Todos) Delete(id uint) error {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("panic occurred:", err)
	// 	}
	// }()
	// s.todos = slices.Delete(s.todos, pos, pos+1)
	result := s.db.DB.Delete(&Todo{}, id)
	return result.Error
}
