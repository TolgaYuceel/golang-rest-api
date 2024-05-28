package services

import (
	"mongodb-api/mocks/repository"
	"mongodb-api/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{primitive.NewObjectID(), "Title 1", "Content 1"},
	{primitive.NewObjectID(), "Title 2", "Content 2"},
	{primitive.NewObjectID(), "Title 3", "Content 3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)

	return func() {
		service = nil
		defer ct.Finish()
	}
}

// Add function test.
func TestDefaultTodoService_TodoInsert(t *testing.T) {
	td := setup(t)
	defer td()

	todo := models.Todo{
		Id:      primitive.ObjectID{},
		Title:   "Title Insert Test",
		Content: "Content for Insert test.",
	}
	mockRepo.EXPECT().Insert(todo).Return(true, nil)

	result, err := service.TodoInsert(todo)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}

// GetAll function test.
func TestDefaultTodoService_TodoGetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.TodoGetAll()

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}

// Delete function test.
func TestDefaultTodoService_TodoDelete(t *testing.T) {
	td := setup(t)
	defer td()

	fakeId := FakeData[0].Id
	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	mockRepo.EXPECT().Delete(fakeId).Return(true, nil)

	result, err := service.TodoDelete(fakeId)
	assert.Equal(t, result, true)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}
