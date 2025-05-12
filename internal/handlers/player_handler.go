package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/moLIart/go-course/internal/dto"
	"github.com/moLIart/go-course/internal/model/room"
	"github.com/moLIart/go-course/internal/repository"
)

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var playerDto dto.CreatePlayerDto
	if err := json.NewDecoder(r.Body).Decode(&playerDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	player := room.NewPlayer(playerDto.Name)
	repository.AddEntity(player)
	w.WriteHeader(http.StatusCreated)
}

func GetPlayersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	players := repository.GetPlayers()
	playerDtos := make([]dto.GetPlayerDto, len(players))
	for i, player := range players {
		playerDtos[i] = dto.GetPlayerDto{ID: player.ID, Name: player.Name}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(playerDtos); err != nil {
		http.Error(w, "Failed to encode players", http.StatusInternalServerError)
	}
}

func GetPlayerByIDHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	player := repository.GetPlayerByID(id)
	if player == nil {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	playerDto := dto.GetPlayerDto{ID: player.ID, Name: player.Name}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(playerDto); err != nil {
		http.Error(w, "Failed to encode player", http.StatusInternalServerError)
	}
}

func DeletePlayerByIDHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	if !repository.DeletePlayerByID(id) {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdatePlayerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var playerDto dto.UpdatePlayerDto
	if err := json.NewDecoder(r.Body).Decode(&playerDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !repository.UpdatePlayerByID(id, playerDto.Name) {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
