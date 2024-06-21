package controller_test

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUserController_DeleteUser(t *testing.T) {
	service, controller := SetupTest(t)

	t.Run("userId_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "test",
			},
		}

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("userId_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()
		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}
		service.EXPECT().DeleteUserServices(id).Return(rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("userId_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()
		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}
		service.EXPECT().DeleteUserServices(id).Return(nil)

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
