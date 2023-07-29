package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/unrolled/render"
	"recipes-api/config"
	"recipes-api/domain/core"
	"recipes-api/http/middleware"
	"recipes-api/http/routes"
)

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	r := render.New()
	router := chi.NewRouter()
	sqlc := config.NewDb()
	cH := core.NewContextHandler(ctx, sqlc, r)

	middleware.MakeMiddleware(router)

	routes.AuthRoute(router, cH)
	routes.MakePrivateRoutes(router, cH)

	addr := ":8089"
	log.Println("listen on", addr)

	http.ListenAndServe(addr, router)
}
