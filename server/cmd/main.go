package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/Hritikpandey-ops/blog-site/config"
	"github.com/Hritikpandey-ops/blog-site/db"
	"github.com/Hritikpandey-ops/blog-site/handlers"
	"github.com/go-chi/chi/v5"
)

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func setupRoutes(r chi.Router) {
	db := db.GetDB()

	r.Post("/register", handlers.Register(db))
	r.Post("/login", handlers.Login(db))
}

func main() {

	// Connect to the database
	handleError(db.Connect(), "Failed to connect to the database")

	// Recover from panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	config.RunMigrations()

	r := chi.NewRouter()

	// Middleware for logging requests
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Serving: %s %s", r.Method, r.URL.Path)
			log.Printf("Headers: %v", r.Header)
			next.ServeHTTP(w, r)
		})
	})

	// Serve static files
	absStatic, _ := filepath.Abs("static")
	fs := http.FileServer(http.Dir(absStatic))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	setupRoutes(r)

	// Start the server with graceful shutdown
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		fmt.Println("Server is starting on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	// Listen for OS shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Gracefully shutdown the server
	fmt.Println("Shutting down server...")
	if err := server.Close(); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}
}

// func RecoveryMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				log.Printf("PANIC: %v", err)
// 				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 			}
// 		}()
// 		next.ServeHTTP(w, r)
// 	})
// }

// func StaticDebugMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Serving static file: %s", r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Request: %s %s", r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }
