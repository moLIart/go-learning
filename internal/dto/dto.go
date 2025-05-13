package dto

type CreatePlayerDto struct {
	Name string `json:"name"`
}

type CreateRoomDto struct {
	Code string `json:"code"`
}

type CreateBoardDto struct {
	Size int `json:"size"`
}

type GetPlayerDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetRoomDto struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

type GetBoardDto struct {
	ID   int `json:"id"`
	Size int `json:"size"`
}

type GetGameDto struct {
	ID int `json:"id"`
}
