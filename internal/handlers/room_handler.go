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

// CreateRoomHandler creates a new room.
// @Summary Create a new room
// @Description Creates a new room with the given code and adds it to the repository.
// @Tags rooms
// @Accept json
// @Produce json
// @Param room body dto.CreateRoomDto true "Room code"
// @Success 201 {object} dto.GetRoomDto
// @Failure 400 {string} string "Invalid request body"
// @Router /rooms [post]
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

// GetRoomsHandler retrieves all rooms.
// @Summary Get all rooms
// @Description Returns a list of all rooms.
// @Tags rooms
// @Produce json
// @Success 200 {array} dto.GetRoomDto
// @Failure 500 {string} string "Failed to encode rooms"
// @Router /rooms [get]
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

// GetRoomByIDHandler retrieves a room by its ID.
// @Summary Get room by ID
// @Description Returns a room by its ID.
// @Tags rooms
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} dto.GetRoomDto
// @Failure 400 {string} string "Invalid id parameter"
// @Failure 404 {string} string "Room not found"
// @Router /rooms/{id} [get]
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

// DeleteRoomHandler deletes a room by its ID.
// @Summary Delete room by ID
// @Description Deletes a room by its ID.
// @Tags rooms
// @Param id path int true "Room ID"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid id parameter"
// @Failure 404 {string} string "Room not found"
// @Router /rooms/{id} [delete]
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

// UpdateRoomHandler updates a room's code by its ID.
// @Summary Update room by ID
// @Description Updates the code of a room by its ID.
// @Tags rooms
// @Accept json
// @Param id path int true "Room ID"
// @Param room body dto.UpdateRoomDto true "Updated room code"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid id parameter or request body"
// @Failure 404 {string} string "Room not found"
// @Router /rooms/{id} [put]
func UpdateRoomHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	var roomDto dto.UpdateRoomDto
	if err := json.NewDecoder(r.Body).Decode(&roomDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !repository.UpdateRoomByID(id, roomDto.Code) {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
