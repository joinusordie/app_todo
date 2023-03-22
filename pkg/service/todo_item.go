package service

import (
	todo "github.com/joinusordie/app_todo"
	"github.com/joinusordie/app_todo/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(listId, userId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}