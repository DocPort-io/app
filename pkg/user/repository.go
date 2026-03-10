package user

import (
	"app/pkg/database"
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type Repository interface {
	GetById(ctx context.Context, id int64) (User, error)
	GetByProvider(ctx context.Context, provider string, providerId string) (User, error)
	Create(ctx context.Context, user User) (User, error)
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

func isPgUniqueViolation(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "23505") || strings.Contains(err.Error(), "unique"))
}
