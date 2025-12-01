package model

type Location struct {
	BaseModel
	Name    string
	Address string
	Lat     float64
	Lon     float64
}
