package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/mahalahub/mahala/graph"
	"github.com/mahalahub/mahala/internal/web"
	"github.com/mahalahub/mahala/users"
	"net/http"
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

	api.POST("/users/login", users.GenerateLoginCode())

	web.ServeHttp(*httpAddr, "api", router)
}
