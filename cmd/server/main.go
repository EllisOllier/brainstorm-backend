package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/EllisOllier/brainstorm-backend/internal/ai"
	"github.com/EllisOllier/brainstorm-backend/internal/database"
	"github.com/EllisOllier/brainstorm-backend/internal/middleware"
	"github.com/EllisOllier/brainstorm-backend/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") // might not want log.fatal in production
	}
	log.Println("Successfully loaded .env file")
	port := os.Getenv("SERVER_PORT")

	db, err := database.Connect()
	if err != nil {
		log.Println("Failed to connect to database", err)
		panic(err) // probably not the best for production
	}
	log.Println("Successfully connected to database")

	// initialise repostiory and services
	aiService, err := ai.NewGeminiService(context.Background(), os.Getenv("GEMINI_API_KEY"))
	if err != nil {
		log.Fatal("Error starting AI client!", err)
	}
	log.Println("Successfully initialised ai client")

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)

	mux := http.NewServeMux()
	// API Routes below
	// ai retlated routes
	mux.Handle("POST /chat", middleware.Authenticate(http.HandlerFunc(aiService.ChatToProject)))

	// user related routes
	mux.HandleFunc("POST /user", userService.CreateAccount)
	mux.HandleFunc("POST /user/login", userService.Login)
	// API Routes above
	loggingMux := middleware.LoggingMiddleware(mux)
	http.ListenAndServe(port, loggingMux)

}
