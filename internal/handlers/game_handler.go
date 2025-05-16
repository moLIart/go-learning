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

// CreateGameHandler creates a new game.
//
//	@Summary		Create a new game (Requires authorization)
//	@Description	Creates a new game and adds it to the repository.
//	@Tags			games
//	@Param			Authorization	header		string	true	"Authorization"
//	@Success		201				{object}	dto.GetGameDto
//	@Security		BearerAuth
//	@Router			/games [post]
func CreateGameHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	game := game.NewGame()
	repository.AddEntity(game)
	w.WriteHeader(http.StatusCreated)
}

// GetGamesHandler retrieves all games.
//
//	@Summary		Get all games
//	@Description	Returns a list of all games.
//	@Tags			games
//	@Produce		json
//	@Success		200	{array}		dto.GetGameDto
//	@Failure		500	{string}	string	"Failed to encode games"
//	@Router			/games [get]
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

// GetGameByIDHandler retrieves a game by its ID.
//
//	@Summary		Get game by ID
//	@Description	Returns a game by its ID.
//	@Tags			games
//	@Produce		json
//	@Param			id	query		int	true	"Game ID"
//	@Success		200	{object}	dto.GetGameDto
//	@Failure		400	{string}	string	"Invalid id parameter"
//	@Failure		404	{string}	string	"Game not found"
//	@Router			/games/{id} [get]
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

// DeleteGameHandler deletes a game by its ID.
//
//	@Summary		Delete game by ID (Requires authorization)
//	@Description	Deletes a game by its ID.
//	@Tags			games
//	@Param			Authorization	header		string	true	"Authorization"
//	@Param			id				query		int		true	"Game ID"
//	@Success		200				{string}	string	"OK"
//	@Failure		400				{string}	string	"Invalid id parameter"
//	@Failure		404				{string}	string	"Game not found"
//	@Security		BearerAuth
//	@Router			/games/{id} [delete]
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
