package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mahalahub/mahala/graph"
	"github.com/mahalahub/mahala/integrations"
	"github.com/mahalahub/mahala/internal/mail"
	"github.com/mahalahub/mahala/internal/web"
	"github.com/mahalahub/mahala/user"
	"net/http"
	"os"
)

var (
	httpAddr = flag.String("http", ":8080", "Http address")
)

func main() {
	flag.Parse()

	router := web.NewRouter()
	router.Use(web.GinContextToContextMiddleware())

	router.POST("/graph", graph.GinHandler())

	api := router.Group("api")

	api.GET("healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	mailU := os.Getenv("MAIL_U")
	mailP := os.Getenv("MAIL_P")
	mailer := mail.NewServer(mail.Config{
		Host:          "smtp.gmail.com",
		Port:          587,
		Username:      mailU,
		Password:      mailP,
		UseEncryption: true,
	})

	userRepo := integrations.NewUserRepository()

	accMan := user.NewAccountManagement(userRepo, integrations.SendMailNotificationForLoginCode(
		mailU,
		mailer.Send,
	))

	api.POST("/users/login", user.GenerateLoginCodeHandler(accMan))

	web.ServeHttp(*httpAddr, "api", router)
}
