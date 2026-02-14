package user

import (
	"app/pkg/database"
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

var (
	ErrUserNotFound              = errors.New("user not found")
	ErrUserAlreadyExists         = errors.New("user already exists")
	ErrExternalAuthAlreadyExists = errors.New("external auth already exists")
)

type Repository interface {
	GetById(ctx context.Context, id int64) (User, error)
	GetByProvider(ctx context.Context, provider string, providerId string) (User, error)
	ListExternalAuths(ctx context.Context, userId int64) ([]ExternalAuth, error)
	Create(ctx context.Context, user User) (User, error)
	CreateExternalAuth(ctx context.Context, externalAuth ExternalAuth) (ExternalAuth, error)
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

func (r *repository) Create(ctx context.Context, user User) (User, error) {
	row, err := r.queries.CreateUser(ctx, &database.CreateUserParams{
		Name:          user.Name,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
	})
	if err != nil {
		if isPgUniqueViolation(err) {
			return User{}, ErrUserAlreadyExists
		}
		return User{}, err
	}
	return toUser(row), nil
}

func (r *repository) CreateExternalAuth(ctx context.Context, externalAuth ExternalAuth) (ExternalAuth, error) {
	row, err := r.queries.CreateExternalAuth(ctx, &database.CreateExternalAuthParams{
		UserID:     externalAuth.UserID,
		Provider:   externalAuth.Provider,
		ProviderID: externalAuth.ProviderID,
	})
	if err != nil {
		if isPgUniqueViolation(err) {
			return ExternalAuth{}, ErrExternalAuthAlreadyExists
		}
		return ExternalAuth{}, err
	}
	return toExternalAuth(row), nil
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

func isPgUniqueViolation(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "23505") || strings.Contains(err.Error(), "unique"))
}
