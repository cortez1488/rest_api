package service

import (
	"errors"
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/cortez1488/rest_todo/pkg/repository"
)

type ItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewItemService(repo repository.TodoItem, listRepo repository.TodoList) *ItemService {
	return &ItemService{repo: repo, listRepo: listRepo}
}

func (s *ItemService) CreateItem(userId, listId int, input todoServer.TodoItem) (int, error) {
	if _, err := s.listRepo.GetList(listId, userId); err != nil {
		return 0, err
	}
	return s.repo.CreateItem(listId, input)
}

func (s *ItemService) GetItems(userId, listId int) ([]todoServer.TodoItem, error) {
	_, err := s.listRepo.GetList(listId, userId)
	if err != nil {
		return nil, errors.New("you are not owner")
	}
	return s.repo.GetItems(listId)
}