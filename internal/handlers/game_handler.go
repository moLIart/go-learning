package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/moLIart/go-course/internal/dto"
	"github.com/moLIart/go-course/internal/model/game"
	"github.com/moLIart/go-course/internal/repository"
)

func CreateGameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	game := game.NewGame()
	repository.AddEntity(game)
	w.WriteHeader(http.StatusCreated)
}

func GetGamesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	games := repository.GetGames()
	gameDtos := make([]dto.GetGameDto, len(games))
	for i, game := range games {
		gameDtos[i] = dto.GetGameDto{ID: game.ID}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(gameDtos); err != nil {
		http.Error(w, "Failed to encode games", http.StatusInternalServerError)
	}
}

func GetGameByIDHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	game := repository.GetGameByID(id)
	if game == nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	gameDto := dto.GetGameDto{ID: game.ID}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(gameDto); err != nil {
		http.Error(w, "Failed to encode game", http.StatusInternalServerError)
	}
}

func DeleteGameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	if !repository.DeleteGameByID(id) {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
