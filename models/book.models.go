package models

type Books struct {
	ID          string   `bson:"_id"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Content     string   `json:"content" validate:"required"`
	Author      string   `json:"author"`
	Price       int      `json:"price" validate:"required"`
	Create_at   int64    `json:"Create_at"`
	Update_at   int64    `json:"Update_at"`
	Rating_id   []string `json:"rating_id"`
}

type PostBook struct {
	Title       string `json:"title" swaggertype:"string" format:"base64" example:"Foo"`
	Description string `json:"description" swaggertype:"string" format:"base64" example:"Some Description"`
	Content     string `json:"content" swaggertype:"string" format:"base64" example:"Content....."`
	Price       string `json:"price" swaggertype:"integer" example:"100"`
}

type UpdatedBook struct {
	Title       string `json:"title" swaggertype:"string" format:"base64" example:""`
	Description string `json:"description" swaggertype:"string" format:"base64" example:""`
	Content     string `json:"content" swaggertype:"string" format:"base64" example:""`
	Price       string `json:"price" swaggertype:"integer" example:"0"`
}
