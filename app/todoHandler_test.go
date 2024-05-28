package app

import (
	services "mongodb-api/mocks/service"
	"mongodb-api/models"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var td TodoHandler
var mockService *services.MockTodoService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = services.NewMockTodoService(ctrl)
	td = TodoHandler{mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_GetAllTodo(t *testing.T) {
	trd := setup(t)
	defer trd()

	router := fiber.New()
	router.Get("/api/getAllTodos", td.GetAllTodo)

	var FakeData = []models.Todo{
		{primitive.NewObjectID(), "Title 1", "Content 1"},
		{primitive.NewObjectID(), "Title 2", "Content 2"},
		{primitive.NewObjectID(), "Title 3", "Content 3"},
	}

	mockService.EXPECT().TodoGetAll().Return(FakeData, nil)

	req := httptest.NewRequest("GET", "/api/getAllTodos", nil)
	resp, _ := router.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}
