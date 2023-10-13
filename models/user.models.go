package models

type User struct {
	ID         string `bson:"_id"`
	First_Name string `json:"first_name" validate:"required"`
	Last_Name  string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required,min=6"`
	Gender     string `json:"gender" validate:"required,oneof=LAKI-LAKI PEREMPUAN"`
	Address    string `json:"address" validate:"required"`
	Role       string `json:"role"`
	Disable    bool   `json:"disable"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
}

type SignIn struct {
	First_Name string `json:"first_name" swaggertype:"string" format:"base64" example:"Foo"`
	Last_Name  string `json:"last_name" swaggertype:"string" format:"base64" example:"Bar"`
	Email      string `json:"email" swaggertype:"string" format:"base64" example:"Example@gmail.com"`
	Password   string `json:"password" swaggertype:"string" format:"base64" example:"FooBar"`
	Gender     string `json:"gender" swaggertype:"string" format:"base64" example:"LAKI-LAKI"`
	Address    string `json:"Address" swaggertype:"string" format:"base64" example:"Cimahi"`
}

type Profile struct {
	First_Name string `json:"first_name" swaggertype:"string" format:"base64" example:"Foo"`
	Last_Name  string `json:"last_name" swaggertype:"string" format:"base64" example:"Bar"`
	Gender     string `json:"gender" swaggertype:"string" format:"base64" example:"LAKI-LAKI"`
	Address    string `json:"Address" swaggertype:"string" format:"base64" example:"Cimahi"`
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type Login struct {
	Email    string `json:"email" swaggertype:"string" format:"base64" example:"Example@gmail.com"`
	Password string `json:"password" swaggertype:"string" format:"base64" example:"FooBar"`
}

type ChangePassword struct {
	OldPassword string `json:"oldPassword" swaggertype:"string" format:"base64" example:"oldpass"`
	NewPassword string `json:"newPassword" swaggertype:"string" format:"base64" example:"newpass"`
}
