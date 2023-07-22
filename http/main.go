package main

import _ "github.com/joho/godotenv/autoload"

import (
	"context"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"recipes-api/config"
	"recipes-api/domain/core"
	"recipes-api/http/middleware"
	"recipes-api/http/routes"
)

func main() {
	ctx := context.Background()
	r := render.New()
	router := chi.NewRouter()
	sqlc := config.NewDb()
	cH := core.NewContextHandler(ctx, sqlc, r)

	middleware.MakeMiddleware(router)
	routes.MakeRoutes(router, cH)

	addr := ":8090"
	log.Println("listen on", addr)
	http.ListenAndServe(addr, router)
}
