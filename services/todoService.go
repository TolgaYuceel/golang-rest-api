package services

import (
	"mongodb-api/dto"
	"mongodb-api/models"
	"mongodb-api/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -destination=../mocks/service/mockTodoService.go -package=services mongodb-api/services TodoService

type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	TodoInsert(todo models.Todo) (*dto.TodoDTO, error)
	TodoGetAll() ([]models.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
}

func (t DefaultTodoService) TodoInsert(todo models.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.Insert(todo)

	if err != nil || !result {
		res.Status = false
		return nil, err
	}

	res = dto.TodoDTO{Status: result}

	return &res, nil
}

func (t DefaultTodoService) TodoGetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultTodoService) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)
	if err != nil || !result {
		return false, err
	}
	return true, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
