package models

type User struct {
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Nickname  string  `json:"nickname"`
	Birth     string  `json:"birth"`
	Point     int     `json:"point"`
	UserType  string  `json:"user_type"`
	Listening []Class `json:"listening"`
	Teaching  []Class `json:"teaching"`
}
