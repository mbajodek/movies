package character_handler

import (
	"encoding/json"
	"movies/db/character_repository"
	"movies/db/movie_repository"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CharacterHandler struct {
	log *zap.Logger
	cr  *character_repository.CharacterRepository
	mr  *movie_repository.MovieRepository
}

func New(log *zap.Logger, cr *character_repository.CharacterRepository, mr *movie_repository.MovieRepository) *CharacterHandler {
	return &CharacterHandler{
		log: log,
		cr:  cr,
		mr:  mr,
	}
}

func (ch *CharacterHandler) Create(rwr http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	movieId := req.URL.Query().Get("movieId")

	if name == "" || movieId == "" {
		ch.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	movieUuid, error := uuid.Parse(movieId)
	if error != nil {
		ch.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid value", http.StatusInternalServerError)
		return
	}
	movie := ch.mr.Get(movieUuid)

	json.NewEncoder(rwr).Encode(ch.cr.Create(name, movie))
}

func (ch *CharacterHandler) Get(rwr http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		ch.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	characterUuid, error := uuid.Parse(id)
	if error != nil {
		ch.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rwr).Encode(ch.cr.Get(characterUuid))
}

func (ch *CharacterHandler) GetAll(rwr http.ResponseWriter, req *http.Request) {
	json.NewEncoder(rwr).Encode(ch.cr.GetAll())
}

func (ch *CharacterHandler) Update(rwr http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	name := req.URL.Query().Get("name")
	movieId := req.URL.Query().Get("movieId")

	if id == "" || name == "" || movieId == "" {
		ch.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	characterUuid, error := uuid.Parse(id)
	if error != nil {
		ch.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid value", http.StatusInternalServerError)
		return
	}

	movieUuid, err := uuid.Parse(movieId)
	if err != nil {
		ch.log.Error("Bad movie uuid value")
		http.Error(rwr, "Internal server error: Bad movie uuid value", http.StatusInternalServerError)
		return
	}
	movie := ch.mr.Get(movieUuid)

	character, _ := ch.cr.Update(characterUuid, name, movie)
	json.NewEncoder(rwr).Encode(character)
}

func (ch *CharacterHandler) Delete(rwr http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		ch.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	characterUuid, error := uuid.Parse(id)
	if error != nil {
		ch.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}

	ch.cr.Delete(characterUuid)
}
