package routes

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/ruyjfs/example-golang/graphql"
	"github.com/ruyjfs/example-golang/graphql/generated"
)

func Graphql(router *gin.Engine) *gin.Engine {

	router.POST("/graphql", graphqlHandler())
	router.GET("/graphiql", playgroundHandler())

	return router
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
