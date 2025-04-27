package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"

	"INF368/internal/models"
	"INF368/internal/controller"
	"INF368/internal/repository"
	"INF368/internal/service"
	"INF368/internal/server"
)

// @title           Book API
// @version         1.0
// @description     This is a sample server for managing books.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)
	h := controller.NewBookHandler(svc)

	repo.AddBook(models.Book{ID: 1, Title: "The Alchemist"})
	repo.AddBook(models.Book{ID: 2, Title: "1984"})
	repo.AddBook(models.Book{ID: 3, Title: "Moby Dick"})

	router := h.InitRoutes()

	srv := server.NewServer(router, "8080")

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server failed to shut down: %v", err)
	}

	log.Println("Server stopped")
}
