package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mahalahub/mahala/internal/random"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username           string `json:"username"`
	EmailOrPhoneNumber string `json:"emailOrPhoneNumber"`
}

type LoginDetailsResponse struct {
	Code string `json:"code"`
}

func GenerateLoginCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			logrus.Error(err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		logrus.Infof("login: %s : %s", loginRequest.Username, loginRequest.EmailOrPhoneNumber)

		c.JSON(http.StatusOK, LoginDetailsResponse{Code: random.Str(6)})
	}
}
