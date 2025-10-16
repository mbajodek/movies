package main

//go:generate go tool oapi-codegen -config ./api/codegen_server.yaml ./api/openapi.yaml

import (
	"movies/internal/db"
	"movies/internal/repository/character_repository"
	"movies/internal/repository/movie_repository"
	"movies/internal/server"
	"movies/internal/server/validator"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			db.New,
			movie_repository.New,
			character_repository.New,
			zap.NewExample,
			resty.New,
			validator.NewStarWarsValidator,
		),
		fx.Invoke(
			server.NewEchoServer,
		),
	).Run()
}
