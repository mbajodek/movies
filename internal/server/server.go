package server

import (
	"context"
	"movies/db/character_repository"
	"movies/db/movie_repository"
	"movies/internal/api"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var port string = ":8080"

type Server struct {
	Mr *movie_repository.MovieRepository
	Cr *character_repository.CharacterRepository
	StarWarsValidator *StarWarsValidator
}

func New(mr *movie_repository.MovieRepository, cr *character_repository.CharacterRepository, starWarsValidator *StarWarsValidator) *Server {
	return &Server{
		Mr: mr,
		Cr: cr,
		StarWarsValidator: starWarsValidator,
	}
}

func NewEchoServer(lc fx.Lifecycle, log *zap.Logger, mr *movie_repository.MovieRepository, cr *character_repository.CharacterRepository, starWarsValidator *StarWarsValidator) *echo.Echo {
	e := echo.New()
	apiServer := New(mr, cr, starWarsValidator)
	strictHandler := api.NewStrictHandler(apiServer, []api.StrictMiddlewareFunc{})
	api.RegisterHandlersWithBaseURL(e, strictHandler, "")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting echo HTTP server", zap.String("addr", e.Server.Addr))
			go e.Start(port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
	return e
}
