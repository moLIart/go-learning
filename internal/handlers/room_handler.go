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

func CreateRoomHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var roomDto dto.CreateRoomDto
	if err := json.NewDecoder(r.Body).Decode(&roomDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	room := room.NewRoom(roomDto.Code)
	repository.AddEntity(room)
	w.WriteHeader(http.StatusCreated)
}

func GetRoomsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rooms := repository.GetRooms()
	roomDtos := make([]dto.GetRoomDto, len(rooms))
	for i, room := range rooms {
		roomDtos[i] = dto.GetRoomDto{ID: room.ID, Code: room.Code}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(roomDtos); err != nil {
		http.Error(w, "Failed to encode rooms", http.StatusInternalServerError)
	}
}

func GetRoomByIDHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	room := repository.GetRoomByID(id)
	if room == nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	roomDto := dto.GetRoomDto{ID: room.ID, Code: room.Code}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(roomDto); err != nil {
		http.Error(w, "Failed to encode room", http.StatusInternalServerError)
	}
}

func DeleteRoomHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	if !repository.DeleteRoomByID(id) {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
