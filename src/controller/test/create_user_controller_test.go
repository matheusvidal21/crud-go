package controller_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/controller/model/request"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserController_CreateUser(t *testing.T) {
	service, controller := SetupTest(t)

	t.Run("validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "ERROR@_EMAIL",
			Password: "teste@",
			Name:     "test",
			Age:      0,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("object_is_valid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "teste1!@",
			Name:     "Test User",
			Age:      20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		model := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
		service.EXPECT().CreateUserServices(model).Return(nil, rest_err.NewInternalServerError("test error"))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("object_is_valid_and_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "teste@teste.com",
			Password: "teste1!@",
			Name:     "Test User",
			Age:      20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		model := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
		service.EXPECT().CreateUserServices(model).Return(model, nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)

		assert.EqualValues(t, http.StatusCreated, recorder.Code)
	})

}
