package model

type Location struct {
	BaseModel
	Nickname string  `json:"nickname"`
	Address  string  `json:"address"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}
