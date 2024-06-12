package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
)

func FindUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("User not found")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {

}
