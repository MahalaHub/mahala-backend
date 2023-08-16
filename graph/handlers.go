package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/mahalahub/mahala/graph/generated"
)

func GinHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
