package user

import (
	"app/pkg/database"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Repository interface {
	GetById(ctx context.Context, id int64) (User, error)
	GetByProvider(ctx context.Context, provider string, providerId string) (User, error)
	ListExternalAuths(ctx context.Context, userId int64) ([]ExternalAuth, error)
}

type repository struct {
	queries *database.Queries
}

func NewRepository(queries *database.Queries) Repository {
	return &repository{queries: queries}
}

func (r *repository) GetById(ctx context.Context, id int64) (User, error) {
	row, err := r.queries.GetUserById(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrUserNotFound
	}
	if err != nil {
		return User{}, err
	}
	return toUser(row), nil
}

func (r *repository) GetByProvider(ctx context.Context, provider string, providerId string) (User, error) {
	row, err := r.queries.GetUserByProvider(ctx, &database.GetUserByProviderParams{
		Provider:   provider,
		ProviderID: providerId,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrUserNotFound
	}
	if err != nil {
		return User{}, err
	}
	return toUser(row), nil
}

func (r *repository) ListExternalAuths(ctx context.Context, userId int64) ([]ExternalAuth, error) {
	rows, err := r.queries.ListExternalAuthsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	externalAuths := make([]ExternalAuth, len(rows))
	for i, row := range rows {
		externalAuths[i] = toExternalAuth(row)
	}
	return externalAuths, nil
}

func toUser(row *database.User) User {
	return User{
		ID:            row.ID,
		CreatedAt:     row.CreatedAt.Time,
		UpdatedAt:     row.UpdatedAt.Time,
		Name:          row.Name,
		Email:         row.Email,
		EmailVerified: row.EmailVerified,
	}
}

func toExternalAuth(row *database.ExternalAuth) ExternalAuth {
	return ExternalAuth{
		ID:         row.ID,
		CreatedAt:  row.CreatedAt.Time,
		UpdatedAt:  row.UpdatedAt.Time,
		UserID:     row.UserID,
		Provider:   row.Provider,
		ProviderID: row.ProviderID,
	}
}
