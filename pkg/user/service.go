package user

import "context"

type Service interface {
	GetById(ctx context.Context, id int64) (User, error)
	GetByProvider(ctx context.Context, provider string, providerId string) (User, error)
	ListExternalAuths(ctx context.Context, userId int64) ([]ExternalAuth, error)
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
