package web

import (
	"context"
	"fmt"
	"movies/web/handlers/character_handler"
	"movies/web/handlers/movie_handler"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				fmt.Println(err)
				return err
			}
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func NewServeMux(mh *movie_handler.MovieHandler, ch *character_handler.CharacterHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/movie/create", mh.Create)
	mux.HandleFunc("/movie/get", mh.Get)
	mux.HandleFunc("/movie/get-all", mh.GetAll)
	mux.HandleFunc("/movie/update", mh.Update)
	mux.HandleFunc("/movie/delete", mh.Delete)
	mux.HandleFunc("/character/create", ch.Create)
	mux.HandleFunc("/character/get", ch.Get)
	mux.HandleFunc("/character/get-all", ch.GetAll)
	mux.HandleFunc("/character/update", ch.Update)
	mux.HandleFunc("/character/delete", ch.Delete)

	return mux
}
