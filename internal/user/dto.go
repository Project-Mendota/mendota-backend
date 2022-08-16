package user

type LoginUserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RemoveUserRequest struct {
	UserID uint `json:"user_id"`
}

type GetProfileRequest struct {
	UserId uint `json:"user_id"`
}
