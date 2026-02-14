package user

import "context"

type Service interface {
	GetById(ctx context.Context, id int64) (User, error)
	GetByProvider(ctx context.Context, provider string, providerId string) (User, error)
	ListExternalAuths(ctx context.Context, userId int64) ([]ExternalAuth, error)
	CreateUser(ctx context.Context, req CreateUserRequest) (User, error)
	CreateExternalAuth(ctx context.Context, userId int64, req CreateExternalAuthRequest) (ExternalAuth, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetById(ctx context.Context, id int64) (User, error) {
	return s.repository.GetById(ctx, id)
}

func (s *service) GetByProvider(ctx context.Context, provider string, providerId string) (User, error) {
	return s.repository.GetByProvider(ctx, provider, providerId)
}

func (s *service) ListExternalAuths(ctx context.Context, userId int64) ([]ExternalAuth, error) {
	return s.repository.ListExternalAuths(ctx, userId)
}

func (s *service) CreateUser(ctx context.Context, req CreateUserRequest) (User, error) {
	user := User{
		Name:          req.Name,
		Email:         req.Email,
		EmailVerified: req.EmailVerified,
	}
	return s.repository.Create(ctx, user)
}

func (s *service) CreateExternalAuth(ctx context.Context, userId int64, req CreateExternalAuthRequest) (ExternalAuth, error) {
	externalAuth := ExternalAuth{
		UserID:     userId,
		Provider:   req.Provider,
		ProviderID: req.ProviderID,
	}
	return s.repository.CreateExternalAuth(ctx, externalAuth)
}
