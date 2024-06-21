package controller_test

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUserController_FindUserByEmail(t *testing.T) {
	service, controller := SetupTest(t)

	t.Run("email_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "TEST_ERROR",
			},
		}
		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}
		service.EXPECT().FindUserByEmailServices("test@test.com").Return(nil, rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}
		service.EXPECT().FindUserByEmailServices("test@test.com").
			Return(model.NewUserDomain("test@test.com", "test", "test", 20), nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestUserController_FindUserById(t *testing.T) {
	service, controller := SetupTest(t)

	t.Run("id_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "TEST_ERROR",
			},
		}
		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()
		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}
		service.EXPECT().FindUserByIDServices(id).Return(nil, rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()
		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}
		service.EXPECT().FindUserByIDServices(id).
			Return(model.NewUserDomain("test@test.com", "test", "test", 20), nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
