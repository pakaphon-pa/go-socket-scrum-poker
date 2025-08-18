package domain

type (
	CreateUserRequest struct {
		Name string `json:"name"`
	}

	CreateUserResponse struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}
)
