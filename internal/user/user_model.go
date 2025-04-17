package user

// import "github.com/google/uuid"

type User struct {
	// Id *uuid.UUID `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string	`json:"password" binding:"required"`
}
