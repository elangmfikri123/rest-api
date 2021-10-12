package user

type RequestUser struct {
	Name     string `json:"name" validator:"required"`
	Email    string `json:"email" validator:"required,email"`
	Password string `json:"password" validator:"required"`
}

type RequestUserLogin struct {
	Email    string `json:"email" validator:"required,email"`
	Password string `json:"password" validator:"required"`
}
