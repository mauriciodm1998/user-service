package rest

type CreateUserRequest struct {
	Document string `json:"document"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

type UserResponse struct {
	Id            string `json:"id"`
	Login         string `json:"login"`
	AccessLevelID int    `json:"access_level_id"`
}

type Response struct {
	Message string `json:"message"`
}

type CustomerResponse struct {
	Id       string `json:"id"`
	UserID   string `json:"user_id"`
	Document string `json:"document"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
