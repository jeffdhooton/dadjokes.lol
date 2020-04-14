package models

type Joke struct {
	ID			uint		`json:"id"`
	Title		string	`json:"joke"`
}

type CreateJokeInput struct {
	Title		string	`json:"joke" binding:"required"`
}

type UpdateJokeInput struct {
	Title 	string `json:"joke"`
}
