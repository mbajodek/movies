package main

import (
	"movies/db"
	"movies/db/character_repository"
	"movies/db/movie_repository"
	"movies/web"
	"movies/web/handlers/character_handler"
	"movies/web/handlers/movie_handler"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			db.New,
			movie_repository.New,
			character_repository.New,
			movie_handler.New,
			character_handler.New,
			web.NewHTTPServer,
			web.NewServeMux,
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
