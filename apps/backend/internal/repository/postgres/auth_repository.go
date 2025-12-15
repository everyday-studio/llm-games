package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	"github.com/everyday-studio/ollm/internal/domain"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) domain.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	//Add User Role if no default role
	if user.Role == "" {
		user.Role = domain.RoleUser
	}

	const query = `
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	if err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return nil, fmt.Errorf("email %s: %w", user.Email, domain.ErrAlreadyExists)
		}
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}
