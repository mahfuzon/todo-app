package todo_test

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/to-do-list-app/controller"
	"github.com/mahfuzon/to-do-list-app/library"
	"github.com/mahfuzon/to-do-list-app/model"
	"github.com/mahfuzon/to-do-list-app/repository"
	"github.com/mahfuzon/to-do-list-app/response"
	"github.com/mahfuzon/to-do-list-app/service"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateTodoSuccess(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableTodos(db)

	// create dummy data
	todo := model.Todo{
		Title: "membuat kue",
	}
	err := db.Create(&todo).Error
	assert.NoError(t, err)

	// setup controller
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	// make example request
	requestJsonString := `{
	"title" :"membuat sayur"
}`

	// setup router
	router := libraries.SetRouter()
	router.PUT("api/v1/todo/:id", todoController.Update)

	// make request test
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/todo/1", strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 201, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestUpdateTodoValidationError(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableTodos(db)

	// create dummy data
	todo := model.Todo{
		Title: "membuat kue",
	}
	err := db.Create(&todo).Error
	assert.NoError(t, err)

	// setup controller
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	// make example request
	requestJsonString := `{
}`

	// setup router
	router := libraries.SetRouter()
	router.PUT("api/v1/todo/:id", todoController.Update)

	// make request test
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/todo/4", strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 422, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestUpdateTodoNotFound(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableTodos(db)

	// create dummy data
	todo := model.Todo{
		Title: "membuat kue",
	}
	err := db.Create(&todo).Error
	assert.NoError(t, err)

	// setup controller
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	// make example request
	requestJsonString := `{
"title" : "memasak nasi"
}`

	// setup router
	router := libraries.SetRouter()
	router.PUT("api/v1/todo/:id", todoController.Update)

	// make request test
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/todo/4", strings.NewReader(requestJsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// get result recorder
	result := rec.Result()
	assert.Equal(t, 400, result.StatusCode)

	// get result body
	body := result.Body

	// get response body
	responseBody, _ := io.ReadAll(body)
	var apiResponse response.ApiResponse
	err = json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}
