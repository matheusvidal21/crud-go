package controller_test

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/controller"
	"github.com/matheusvidal21/crud-go/src/test/mocks"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(c *gin.Context, param gin.Params, u url.Values, method string, body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}

func SetupTest(t *testing.T) (*mocks.MockUserDomainService, controller.UserControllerInterface) {
	ctrl := gomock.NewController(t)

	t.Cleanup(func() {
		ctrl.Finish()
	})

	service := mocks.NewMockUserDomainService(ctrl)

	return service, controller.NewUserController(service)
}
