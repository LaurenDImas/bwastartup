package user

type RegisterUserInput struct {
	Name       string `json:"Name" binding:"required"`
	Occupation string `json:"Occupation" binding:"required"`
	Email      string `json:"Email" binding:"required,email"`
	Password   string `json:"Password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
