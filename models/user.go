package models

type User struct {
	Email     string  `json:"email" validate:"email"`
	Password  string  `json:"password" validate:"min=8"`
	Nickname  string  `json:"nickname" validate:"required"`
	Birth     string  `json:"birth" validate:"required"`
	Point     int     `json:"point"`
	UserType  string  `json:"user_type" validate:"required"`
	Listening []Class `json:"listening"`
	Teaching  []Class `json:"teaching"`
}
