package main

import (
	"log"
	"os"

	"github.com/sunzhqr/book_library/internal/adapter"
	booksrepository "github.com/sunzhqr/book_library/internal/books/repository"
	booksservice "github.com/sunzhqr/book_library/internal/books/service"
	"github.com/sunzhqr/book_library/internal/config"
	router "github.com/sunzhqr/book_library/internal/transport/http"
	"github.com/sunzhqr/book_library/pkg/postgres"
)

var logger = log.New(os.Stdout, "book_library", log.LstdFlags)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	logger = log.New(file, "BOOKSERVICE: ", log.Ldate|log.Ltime|log.Lshortfile)
	cfg := config.NewConfig()
	db, err := postgres.NewPostgres(cfg.DB)
	if err != nil {
		logger.Fatal("Error creating table:", err)
	}

	repo := booksrepository.NewRepository(db)

	wheatherAdapter := adapter.NewAdapter()

	service := booksservice.NewService(repo, wheatherAdapter)

	handler := router.NewHandler(service)

	r := router.NewRouter(cfg.RouterConfig, handler)
	r.Run()
}
