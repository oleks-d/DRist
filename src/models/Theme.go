package models

type Theme struct {
	Id          string
	Title       string
	Description string
	Side1       string
	Side2       string
}

func NewTheme(id, title, description, side1, side2 string) *Theme {
	return &Theme{id, title, description, side1, side2}
}
