package services

import (
	"fmt"
	"slices"
)

type Todo struct {
	Text string
	Done bool
}

type Todos struct {
	todos []Todo
}

func New() *Todos {
	return &Todos{
		[]Todo{
			{"Wash car", false},
			{"Continue working on linux updates", false},
		},
	}
}

func (s *Todos) List() ([]Todo, error) {
	return s.todos, nil
}

func (s *Todos) Create(pos int, val string) error {
	s.todos = slices.Insert(s.todos, pos, Todo{val, false})
	return nil
}

func (s *Todos) Delete(pos int) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()
	s.todos = slices.Delete(s.todos, pos, pos+1)
	return nil
}
