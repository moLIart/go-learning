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

// CreatePlayerHandler creates a new player.
//
//	@Summary		Create a new player (Requires authorization)
//	@Description	Creates a new player with the given name and adds it to the repository.
//	@Tags			players
//	@Accept			json
//	@Produce		json
//	@Param			player	body		dto.CreatePlayerDto	true	"Player name"
//	@Success		201		{object}	dto.GetPlayerDto
//	@Failure		400		{string}	string	"Invalid request body"
//	@Security		BearerAuth
//	@Router			/players [post]
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

// GetPlayersHandler retrieves all players.
//
//	@Summary		Get all players
//	@Description	Returns a list of all players.
//	@Tags			players
//	@Produce		json
//	@Success		200	{array}		dto.GetPlayerDto
//	@Failure		500	{string}	string	"Failed to encode players"
//	@Router			/players [get]
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

// GetPlayerByIDHandler retrieves a player by their ID.
//
//	@Summary		Get player by ID
//	@Description	Returns a player by their ID.
//	@Tags			players
//	@Produce		json
//	@Param			id	path		int	true	"Player ID"
//	@Success		200	{object}	dto.GetPlayerDto
//	@Failure		400	{string}	string	"Invalid id parameter"
//	@Failure		404	{string}	string	"Player not found"
//	@Router			/players/{id} [get]
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

// DeletePlayerByIDHandler deletes a player by their ID.
//
//	@Summary		Delete player by ID (Requires authorization)
//	@Description	Deletes a player by their ID.
//	@Tags			players
//	@Param			id	path		int		true	"Player ID"
//	@Success		200	{string}	string	"OK"
//	@Failure		400	{string}	string	"Invalid id parameter"
//	@Failure		404	{string}	string	"Player not found"
//	@Security		BearerAuth
//	@Router			/players/{id} [delete]
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

// UpdatePlayerHandler updates a player's name by their ID.
//
//	@Summary		Update player by ID (Requires authorization)
//	@Description	Updates the name of a player by their ID.
//	@Tags			players
//	@Accept			json
//	@Param			id		path		int					true	"Player ID"
//	@Param			player	body		dto.UpdatePlayerDto	true	"Updated player name"
//	@Success		200		{string}	string				"OK"
//	@Failure		400		{string}	string				"Invalid id parameter or request body"
//	@Failure		404		{string}	string				"Player not found"
//	@Security		BearerAuth
//	@Router			/players/{id} [put]
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
