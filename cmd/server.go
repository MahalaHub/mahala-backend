package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mahalahub/mahala/internal/web"
	"net/http"
)

var (
	httpAddr = flag.String("http", ":8000", "Http address")
)

func main() {
	flag.Parse()

	router := web.NewRouter()

	api := router.Group("api")

	api.GET("healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	web.ServeHttp(*httpAddr, "api", router)
}
