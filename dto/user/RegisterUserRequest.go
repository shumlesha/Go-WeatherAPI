package user

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	City     string `json:"city" binding:"required"`
}
