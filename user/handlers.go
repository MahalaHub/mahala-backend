package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func GenerateLoginCodeHandler(accMan AccountManagement) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		loginCode, err := accMan.GenerateLoginCode(loginRequest.Username, loginRequest.Email)
		if err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, LoginDetailsResponse{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, LoginDetailsResponse{Code: loginCode})
	}
}
