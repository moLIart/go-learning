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

// CreateBoardHandler creates a new board with the specified size.
//
//	@Summary		Create a new board (Requires authorization)
//	@Description	Creates a new board with the given size and adds it to the repository.
//	@Tags			boards
//	@Accept			json
//	@Produce		json
//	@Param			board			body		dto.CreateBoardDto	true	"Board size"
//	@Success		201				{object}	dto.GetBoardDto
//	@Failure		400				{string}	string	"Invalid request body"
//	@Security		BearerAuth
//	@Router			/boards [post]
func CreateBoardHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var boardDto dto.CreateBoardDto
	if err := json.NewDecoder(r.Body).Decode(&boardDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	board := game.NewBoard(boardDto.Size)
	repository.AddEntity(board)
	w.WriteHeader(http.StatusCreated)
}

// GetBoardsHandler retrieves all boards.
//
//	@Summary		Get all boards
//	@Description	Returns a list of all boards.
//	@Tags			boards
//	@Produce		json
//	@Success		200	{array}		dto.GetBoardDto
//	@Failure		500	{string}	string	"Failed to encode boards"
//	@Router			/boards [get]
func GetBoardsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	boards := repository.GetBoards()
	boardDtos := make([]dto.GetBoardDto, len(boards))
	for i, board := range boards {
		boardDtos[i] = dto.GetBoardDto{ID: board.ID, Size: board.Size}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(boardDtos); err != nil {
		http.Error(w, "Failed to encode boards", http.StatusInternalServerError)
	}
}

// GetBoardByIDHandler retrieves a board by its ID.
//
//	@Summary		Get board by ID
//	@Description	Returns a board by its ID.
//	@Tags			boards
//	@Produce		json
//	@Param			id	path		int	true	"Board ID"
//	@Success		200	{object}	dto.GetBoardDto
//	@Failure		400	{string}	string	"Invalid id parameter"
//	@Failure		404	{string}	string	"Board not found"
//	@Router			/boards/{id} [get]
func GetBoardByIDHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	board := repository.GetBoardByID(id)
	if board == nil {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	boardDto := dto.GetBoardDto{ID: board.ID, Size: board.Size}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(boardDto); err != nil {
		http.Error(w, "Failed to encode board", http.StatusInternalServerError)
	}
}

// DeleteBoardHandler deletes a board by its ID.
//
//	@Summary		Delete board by ID (Requires authorization)
//	@Description	Deletes a board by its ID.
//	@Tags			boards
//	@Param			id				path		int		true	"Board ID"
//	@Success		200				{string}	string	"OK"
//	@Failure		400				{string}	string	"Invalid id parameter"
//	@Failure		404				{string}	string	"Board not found"
//	@Security		BearerAuth
//	@Router			/boards/{id} [delete]
func DeleteBoardHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	if !repository.DeleteBoardByID(id) {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateBoardHandler updates a board's size by its ID.
//
//	@Summary		Update board by ID (Requires authorization)
//	@Description	Updates the size of a board by its ID.
//	@Tags			boards
//	@Accept			json
//	@Param			id				path		int					true	"Board ID"
//	@Param			board			body		dto.UpdateBoardDto	true	"Updated board size"
//	@Success		200				{string}	string				"OK"
//	@Failure		400				{string}	string				"Invalid id parameter or request body"
//	@Failure		404				{string}	string				"Board not found"
//	@Security		BearerAuth
//	@Router			/boards/{id} [put]
func UpdateBoardHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var boardDto dto.UpdateBoardDto
	if err := json.NewDecoder(r.Body).Decode(&boardDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !repository.UpdateBoardByID(id, boardDto.Size) {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
