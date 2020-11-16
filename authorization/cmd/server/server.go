package main

import (
	"authorization/internal/middleware"
	"authorization/pkg/db"
	"authorization/pkg/graph"
	"authorization/pkg/graph/generated"
	"authorization/pkg/graph/models"
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)
// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

const defaultPort = "4002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	gormDb := db.InitDB()
	gormDb.AutoMigrate(&models.User{})
	db, err := gormDb.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()
	router.Use(middleware.GinContextToContextMiddleware())
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(router.Run(":" + port))
}
