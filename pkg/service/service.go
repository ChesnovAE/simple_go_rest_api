package service

import (
	"github.com/ChesnovAE/simple_go_rest_api/pkg/repository"
)

// Authorization ...
type Authorization interface {
}

// TodoList ...
type TodoList interface {
}

// TodoItem ...
type TodoItem interface {
}

// Service ...
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService ...
// Сервисы обращаются к базе данных
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
