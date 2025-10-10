package movie_handler

import (
	"encoding/json"
	"movies/db/movie_repository"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type MovieHandler struct {
	log *zap.Logger
	mr  *movie_repository.MovieRepository
}

func New(log *zap.Logger, mr *movie_repository.MovieRepository) *MovieHandler {
	return &MovieHandler{
		log: log,
		mr:  mr,
	}
}

func (mh *MovieHandler) Create(rwr http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	year := req.URL.Query().Get("year")

	if title == "" || year == "" {
		mh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	y, err := strconv.Atoi(year)
	if err != nil {
		mh.log.Error("Year must be number value")
		http.Error(rwr, "Internal server error: Year must be number value", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rwr).Encode(mh.mr.Create(title, y))
}

func (mh *MovieHandler) Get(rwr http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		mh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	movieUuid, error := uuid.Parse(id)
	if error != nil {
		mh.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rwr).Encode(mh.mr.Get(movieUuid))
}

func (mh *MovieHandler) GetAll(rwr http.ResponseWriter, req *http.Request) {
	json.NewEncoder(rwr).Encode(mh.mr.GetAll())
}

func (mh *MovieHandler) Update(rwr http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	title := req.URL.Query().Get("title")
	year := req.URL.Query().Get("year")

	if id == "" || title == "" || year == "" {
		mh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	y, err := strconv.Atoi(year)
	if err != nil {
		mh.log.Error("Year must be number value")
		http.Error(rwr, "Internal server error: Year must be number value", http.StatusInternalServerError)
		return
	}

	movieUuid, error := uuid.Parse(id)
	if error != nil {
		mh.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}
	movie, _ := mh.mr.Update(movieUuid, title, y)
	json.NewEncoder(rwr).Encode(movie)
}

func (mh *MovieHandler) Delete(rwr http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		mh.log.Error("Failed to read request")
		http.Error(rwr, "Internal server error: Failed to read request", http.StatusInternalServerError)
		return
	}

	movieUuid, error := uuid.Parse(id)
	if error != nil {
		mh.log.Error("Bad uuid value")
		http.Error(rwr, "Internal server error: Bad uuid valu", http.StatusInternalServerError)
		return
	}

	mh.mr.Delete(movieUuid)
}
