package domain

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
