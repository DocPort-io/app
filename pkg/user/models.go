package user

import (
	"time"
)

type User struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Email         string
	EmailVerified bool
}

type ExternalAuth struct {
	ID         int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     int64
	Provider   string
	ProviderID string
}

type CreateUserRequest struct {
	Name          string `json:"name" validate:"required" example:"John Doe"`
	Email         string `json:"email" validate:"required,email" example:"john.doe@example.org"`
	EmailVerified bool   `json:"emailVerified" example:"false" default:"false"`
}

type CreateExternalAuthRequest struct {
	Provider   string `json:"provider" validate:"required" example:"https://keycloak.docport.io/realms/docport-dev"`
	ProviderID string `json:"providerId" validate:"required" example:"4191a0e2-c347-46d4-97bf-7d274ad201d7"`
}

type UserResponse struct {
	ID            int64  `json:"id" example:"1"`
	CreatedAt     string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt     string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	Name          string `json:"name" example:"John Doe"`
	Email         string `json:"email" example:"john.doe@example.org"`
	EmailVerified bool   `json:"emailVerified" example:"true"`
}

type ExternalAuthResponse struct {
	ID         int64  `json:"id" example:"1"`
	CreatedAt  string `json:"createdAt" example:"2026-01-01T00:00:00.000Z"`
	UpdatedAt  string `json:"updatedAt" example:"2026-01-01T00:00:00.000Z"`
	UserID     int64  `json:"userId" example:"1"`
	Provider   string `json:"provider" example:"https://github.com"`
	ProviderID string `json:"providerId" example:"4191a0e2-c347-46d4-97bf-7d274ad201d7"`
}

type ListExternalAuthResponse struct {
	ExternalAuths []ExternalAuthResponse `json:"externalAuths"`
}

func (u User) ToResponse() UserResponse {
	return UserResponse{
		ID:            u.ID,
		CreatedAt:     u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     u.UpdatedAt.Format(time.RFC3339),
		Name:          u.Name,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
	}
}

func (e ExternalAuth) ToResponse() ExternalAuthResponse {
	return ExternalAuthResponse{
		ID:         e.ID,
		CreatedAt:  e.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  e.UpdatedAt.Format(time.RFC3339),
		UserID:     e.UserID,
		Provider:   e.Provider,
		ProviderID: e.ProviderID,
	}
}

func ToListExternalAuthResponse(externalAuths []ExternalAuth) ListExternalAuthResponse {
	response := make([]ExternalAuthResponse, len(externalAuths))
	for i, externalAuth := range externalAuths {
		response[i] = externalAuth.ToResponse()
	}
	return ListExternalAuthResponse{
		ExternalAuths: response,
	}
}
