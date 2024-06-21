package controller_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/controller/model/request"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserController_LoginUser(t *testing.T) {
	service, controller := SetupTest(t)

	t.Run("validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLoginRequest{
			Email:    "ERROR@_EMAIL",
			Password: "teste@",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("object_is_valid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLoginRequest{
			Email:    "teste@teste.com",
			Password: "teste@123!",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		model := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
		service.EXPECT().LoginUserServices(model).Return(nil, "", rest_err.NewInternalServerError("test error"))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("object_is_valid_and_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		userRequest := request.UserLoginRequest{
			Email:    "teste@teste.com",
			Password: "teste@123!",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		model := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
		service.EXPECT().LoginUserServices(model).Return(model, id, nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, recorder.Header().Values("Authorization")[0], id)
	})

}
