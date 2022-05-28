package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if strings.TrimSpace(config.Port) == "" {
		return nil, errors.New("Port is required")
	}
	if strings.TrimSpace(config.JWTSecret) == "" {
		return nil, errors.New("JWTSecret is required")
	}
	if strings.TrimSpace(config.DatabaseUrl) == "" {
		return nil, errors.New("DatabaseUrl is required")
	}
	return &Broker{
		config: config,
		router: mux.NewRouter(),
	}, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Println("Starting server on port", b.config.Port)
	if err := http.ListenAndServe(":"+b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
