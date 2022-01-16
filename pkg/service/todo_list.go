package service

import (
	todoServer "github.com/cortez1488/rest_todo"
	"github.com/cortez1488/rest_todo/pkg/repository"
)

type ListService struct {
	repo repository.TodoList
}

func NewListService(repo repository.TodoList) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) CreateList(userId int, list todoServer.TodoList) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *ListService) GetList(listId int, userId int) (todoServer.TodoList, error) {
	return s.repo.GetList(listId, userId)
}

func (s *ListService) GetAllLists(userId int) ([]todoServer.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

func (s *ListService) DeleteList(listId int, userId int) error {
	return s.repo.DeleteList(listId, userId)
}

func (s *ListService) UpdateList(listId, userId int, input todoServer.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err

	}
	return s.repo.UpdateList(listId, userId, input)
}
