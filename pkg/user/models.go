package user

type User struct {
	ID string
}

type UserResponse struct {
	ID string `json:"id"`
}

func (u User) ToResponse() UserResponse {
	return UserResponse{
		ID: u.ID,
	}
}
