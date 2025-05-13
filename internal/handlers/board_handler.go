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
