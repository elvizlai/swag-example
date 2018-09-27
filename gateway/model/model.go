package model

type User struct {
	Name string
	Loc  Location
	Next *User
}

type Location struct {
	Addr string
}

type Coord struct {
	Lat float64
}
