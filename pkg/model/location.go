package model

type Location struct {
	BaseModel
	Nickname string
	Address  string
	Lat      float64
	Lon      float64
}
