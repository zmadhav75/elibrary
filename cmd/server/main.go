package main

import (
	"elibrary/internal/config"
	"elibrary/internal/handlers"
	"elibrary/internal/repository"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.LoadConfig()

	db, err := repository.NewDB(cfg.DBConnection)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	bookRepo := repository.NewBookRepository(db)
	bookHandler := handlers.NewBookHandler(bookRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Serve static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Serve HTML template
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		books, _ := bookRepo.GetAllBooks()
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
		tmpl.Execute(w, map[string]interface{}{"Books": books})
	})

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/books/{id}", bookHandler.GetBook)
		r.Post("/borrow", bookHandler.BorrowBook)
	})

	log.Printf("Server running on port %s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, r)
}
