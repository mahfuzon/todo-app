package todo_test

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahfuzon/to-do-list-app/controller"
	"github.com/mahfuzon/to-do-list-app/library"
	"github.com/mahfuzon/to-do-list-app/repository"
	"github.com/mahfuzon/to-do-list-app/response"
	"github.com/mahfuzon/to-do-list-app/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TruncateTableTodos(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE todos")
}

func TestAddTodoSuccess(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableTodos(db)

	// setup controller
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	// make example request
	requestJsonString := `{
	"title" :"membuat kue"
}`

	// setup router
	router := libraries.SetRouter()
	router.POST("api/v1/todo", todoController.Create)

	// make request test
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/todo", strings.NewReader(requestJsonString))
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
	err := json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}

func TestAddTodoValidationError(t *testing.T) {
	// instansiasi db
	db := libraries.SetDbTest()

	//truncate table contacts
	TruncateTableTodos(db)

	// setup controller
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoController := controller.NewTodoController(todoService)

	// make example request
	requestJsonString := `{
}`

	// setup router
	router := libraries.SetRouter()
	router.POST("api/v1/todo", todoController.Create)

	// make request test
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/todo", strings.NewReader(requestJsonString))
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
	err := json.Unmarshal(responseBody, &apiResponse)
	assert.NoError(t, err)

	fmt.Println(apiResponse.Data)
}
