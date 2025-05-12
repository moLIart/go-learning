package dto

type UpdatePlayerDto struct {
	Name string `json:"name"`
}

type UpdateRoomDto struct {
	Code string `json:"code"`
}

type UpdateBoardDto struct {
	Size int `json:"size"`
}
