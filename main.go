package main

import (
	"context"
	"encoding/json"
	"fmt"
	"movies/db"
	"movies/db/character_repository"
	"movies/db/movie_repository"
	"net"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var DB *db.MemoryDb
var Mr *movie_repository.MovieRepository
var Cr *character_repository.CharacterRepository

type CreateMovieHandler struct{
	log *zap.Logger
}

type GetMovieHandler struct {
   	log *zap.Logger
}

type GetAllMovieHandler struct {
   	log *zap.Logger
}

type UpdateMovieHandler struct {
   	log *zap.Logger
}

type DeleteMovieHandler struct {
   	log *zap.Logger
}

type CreateCharacterHandler struct{
	log *zap.Logger
}

type GetCharacterHandler struct {
   	log *zap.Logger
}

type GetAllCharacterHandler struct {
   	log *zap.Logger
}

type UpdateCharacterHandler struct {
   	log *zap.Logger
}

type DeleteCharacterHandler struct {
   	log *zap.Logger
}

type Route interface {
   http.Handler

   // Pattern reports the path at which this is registered.
   Pattern() string
}

func NewCreateMovieHandler(log *zap.Logger) *CreateMovieHandler {
   	return &CreateMovieHandler{log: log}
}

func NewGetMovieHandler(log *zap.Logger) *GetMovieHandler {
   return &GetMovieHandler{log: log}
}

func NewGetAllMovieHandler(log *zap.Logger) *GetAllMovieHandler {
   return &GetAllMovieHandler{log: log}
}

func NewUpdateMovieHandler(log *zap.Logger) *UpdateMovieHandler {
   return &UpdateMovieHandler{log: log}
}

func NewDeleteMovieHandler(log *zap.Logger) *DeleteMovieHandler {
   return &DeleteMovieHandler{log: log}
}

func NewCreateCharacterHandler(log *zap.Logger) *CreateCharacterHandler {
   	return &CreateCharacterHandler{log: log}
}

func NewGetCharacterHandler(log *zap.Logger) *GetCharacterHandler {
   return &GetCharacterHandler{log: log}
}

func NewGetAllCharacterHandler(log *zap.Logger) *GetAllCharacterHandler {
   return &GetAllCharacterHandler{log: log}
}

func NewUpdateCharacterHandler(log *zap.Logger) *UpdateCharacterHandler {
   return &UpdateCharacterHandler{log: log}
}

func NewDeleteCharacterHandler(log *zap.Logger) *DeleteCharacterHandler {
   return &DeleteCharacterHandler{log: log}
}

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(
				NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
			AsRoute(NewCreateMovieHandler),
			AsRoute(NewGetMovieHandler),
			AsRoute(NewGetAllMovieHandler),
			AsRoute(NewUpdateMovieHandler),
			AsRoute(NewDeleteMovieHandler),
			AsRoute(NewCreateCharacterHandler),
			AsRoute(NewGetCharacterHandler),
			AsRoute(NewGetAllCharacterHandler),
			AsRoute(NewUpdateCharacterHandler),
			AsRoute(NewDeleteCharacterHandler),
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				DB = db.New()
				Mr = movie_repository.New(DB)
				Cr = character_repository.New(DB)
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

func (cmh *CreateMovieHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	year := req.URL.Query().Get("year")

	if title == "" || year == "" {
		cmh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	y, err := strconv.Atoi(year)
    if err != nil {
		cmh.log.Error("Year must be number value")
       	http.Error(rwr, "Internal server error: Year must be number value", http.StatusInternalServerError)
		return
    }

	json.NewEncoder(rwr).Encode(Mr.Create(title, y))
}

func (gmh *GetMovieHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
   	id := req.URL.Query().Get("id")
	if id == "" {
		gmh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	movieUuid, error := uuid.Parse(id)
	if error != nil {
		gmh.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rwr).Encode(Mr.Get(movieUuid))
}

func (gmh *GetAllMovieHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
	json.NewEncoder(rwr).Encode(Mr.GetAll())
}

func (umh *UpdateMovieHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	title := req.URL.Query().Get("title")
	year := req.URL.Query().Get("year")

	if id =="" || title == "" || year == "" {
		umh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	y, err := strconv.Atoi(year)
    if err != nil {
		umh.log.Error("Year must be number value")
       	http.Error(rwr, "Internal server error: Year must be number value", http.StatusInternalServerError)
		return
    }

	movieUuid, error := uuid.Parse(id)
	if error != nil {
		umh.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}
	movie, _ := Mr.Update(movieUuid, title, y)
	json.NewEncoder(rwr).Encode(movie)
}

func (dmh *DeleteMovieHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
   	id := req.URL.Query().Get("id")
	if id == "" {
		dmh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	movieUuid, error := uuid.Parse(id)
	if error != nil {
		dmh.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}

	Mr.Delete(movieUuid)
}

func (cch *CreateCharacterHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	movieId := req.URL.Query().Get("movieId")

	if name == "" || movieId == "" {
		cch.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	movieUuid, error := uuid.Parse(movieId)
	if error != nil {
		cch.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid value", http.StatusInternalServerError)
		return
	}
	movie := Mr.Get(movieUuid)

	json.NewEncoder(rwr).Encode(Cr.Create(name, movie))
}

func (gch *GetCharacterHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
   	id := req.URL.Query().Get("id")
	if id == "" {
		gch.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	characterUuid, error := uuid.Parse(id)
	if error != nil {
		gch.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rwr).Encode(Cr.Get(characterUuid))
}

func (gmh *GetAllCharacterHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
	json.NewEncoder(rwr).Encode(Cr.GetAll())
}

func (uch *UpdateCharacterHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	movieId := req.URL.Query().Get("movieId")

	if id =="" || name == "" || movieId == "" {
		uch.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	characterUuid, error := uuid.Parse(id)
	if error != nil {
		uch.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid value", http.StatusInternalServerError)
		return
	}

	movieUuid, err := uuid.Parse(movieId)
	if err != nil {
		uch.log.Error("Bad movie uuid value")
		http.Error(rwr, "Internal server error: Bad movie uuid value", http.StatusInternalServerError)
		return
	}
	movie := Mr.Get(movieUuid)

	character, _ := Cr.Update(characterUuid, name, movie)
	json.NewEncoder(rwr).Encode(character)
}

func (dmh *DeleteCharacterHandler) ServeHTTP(rwr http.ResponseWriter, req *http.Request) {
   	id := req.URL.Query().Get("id")
	if id == "" {
		dmh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	characterUuid, error := uuid.Parse(id)
	if error != nil {
		dmh.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}

	Cr.Delete(characterUuid)
}

func (*CreateMovieHandler) Pattern() string {
   return "/movie/create"
}

func (*GetMovieHandler) Pattern() string {
   return "/movie/get"
}

func (*GetAllMovieHandler) Pattern() string {
   return "/movie/get-all"
}

func (*UpdateMovieHandler) Pattern() string {
   return "/movie/update"
}

func (*DeleteMovieHandler) Pattern() string {
   return "/movie/delete"
}

func (*CreateCharacterHandler) Pattern() string {
   return "/character/create"
}

func (*GetCharacterHandler) Pattern() string {
   return "/character/get"
}

func (*GetAllCharacterHandler) Pattern() string {
   return "/character/get-all"
}

func (*UpdateCharacterHandler) Pattern() string {
   return "/character/update"
}

func (*DeleteCharacterHandler) Pattern() string {
   return "/character/delete"
}

func AsRoute(h any) any {
	return fx.Annotate(
		h,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return mux
}

