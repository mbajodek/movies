package server

import (
	"context"
	"movies/internal/api"
	"movies/internal/cert"
	"movies/internal/repository/character_repository"
	"movies/internal/repository/movie_repository"
	"movies/internal/server/validator"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var port string = ":8080"

type Server struct {
	Mr *movie_repository.MovieRepository
	Cr *character_repository.CharacterRepository
	StarWarsValidator *validator.StarWarsValidator
	CertGenerator *cert.CertGenerator
}

func New(mr *movie_repository.MovieRepository, cr *character_repository.CharacterRepository, starWarsValidator *validator.StarWarsValidator, certGenerator *cert.CertGenerator) *Server {
	return &Server{
		Mr: mr,
		Cr: cr,
		StarWarsValidator: starWarsValidator,
		CertGenerator: certGenerator,
	}
}

func NewEchoServer(lc fx.Lifecycle, log *zap.Logger, mr *movie_repository.MovieRepository, cr *character_repository.CharacterRepository, starWarsValidator *validator.StarWarsValidator, certGenerator *cert.CertGenerator) *echo.Echo {
	e := echo.New()
	apiServer := New(mr, cr, starWarsValidator, certGenerator)
	strictHandler := api.NewStrictHandler(apiServer, []api.StrictMiddlewareFunc{})
	api.RegisterHandlers(e, strictHandler)

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
