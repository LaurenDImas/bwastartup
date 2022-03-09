package user

type RegisterUserInput struct {
	Name       string `json:"Name" binding:"required"`
	Occupation string `json:"Occupation" binding:"required"`
	Email      string `json:"Email" binding:"required,email"`
	Password   string `json:"Password" binding:"required"`
}
