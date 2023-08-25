package web

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	GinContext GinContextKey = "GinContextKey"
)

type GinContextKey string

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContext, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContext)
	if ginContext == nil {
		return nil, errors.New("could not retrieve gin.Context")
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, errors.New("gin.Context has wrong type")
	}
	return gc, nil
}
