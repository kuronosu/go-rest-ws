package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kuronosu/go-rest-ws/database"
	"github.com/kuronosu/go-rest-ws/handlers"
	"github.com/kuronosu/go-rest-ws/repository"
	"github.com/kuronosu/go-rest-ws/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("SERVER_PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URL := os.Getenv("DB_URL")
	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DB_URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	InitRepository(DB_URL)
	s.Start(BindRoutes)
	// user := models.User{"", "", models.BaseModel{}}
	// fmt.Println(user)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
}

func InitRepository(DB_URL string) {
	pgrepo, err := database.NewPostgresRepository(DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetUserRepository(pgrepo.User)
}
