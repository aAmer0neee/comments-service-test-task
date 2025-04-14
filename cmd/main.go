package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aAmer0neee/comments-service-test-task/graph/resolver"
	"github.com/aAmer0neee/comments-service-test-task/graph/runtime"
	"github.com/aAmer0neee/comments-service-test-task/internal/config"
	"github.com/aAmer0neee/comments-service-test-task/internal/logger"
	"github.com/aAmer0neee/comments-service-test-task/internal/repository"
	"github.com/aAmer0neee/comments-service-test-task/internal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {

	cfg := config.LoadConfig()

	repository, err := repository.InitRepository(cfg)
	if err != nil {
		log.Fatalf("error init repository")
	}

	logger := logger.ConfigureLogger(cfg.Server.Env)

	service := service.InitService(repository, *logger)

	r := gin.Default()

	r.POST("/graphql", graphqlhandler(service))
	r.GET("/", func(ctx *gin.Context) {
		playground.Handler("GraphQL playground", "/graphql").ServeHTTP(ctx.Writer, ctx.Request)
	})

	go func() {

		if err := r.Run(":" + cfg.Server.Port); err != nil {
			log.Fatalf("ошибка запуска сервера %v", err.Error())
		}

	}()

	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: r,
	}

	shutdown(srv)
}

func graphqlhandler(service service.Service) gin.HandlerFunc {
	handler := handler.New(runtime.NewExecutableSchema(runtime.Config{
		Resolvers: &resolver.Resolver{Service: service},
	}))

	handler.AddTransport(transport.Options{})

	handler.AddTransport(transport.GET{})

	handler.AddTransport(transport.POST{})

	handler.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	handler.Use(extension.Introspection{})
	handler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func shutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Graceful Shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	srv.Shutdown(ctx)
}
