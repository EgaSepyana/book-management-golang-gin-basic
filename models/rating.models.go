package models

type Rating struct {
	ID         string `bson:"_id"`
	Rate       int    `json:"rate" validate:"required,numeric"`
	Comment    string `json:"comment"`
	Author     string `json:"author"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
}

type PostRating struct {
	Comment string `json:"comment" swaggertype:"string" format:"base64" example:"Comment....."`
	Rate    string `json:"rate" swaggertype:"integer" example:"1"`
}

type UpdateRating struct {
	Comment string `json:"comment" swaggertype:"string" format:"base64" example:"Comment....."`
}
