package repo

import (
	"context"
	"time"

	"github.com/amjadjibon/microservices/auth/model"
	"github.com/amjadjibon/microservices/pkg/db"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type RoleRepo interface {
}

type UserPermissionRepo interface {
}

type AuthRepo interface {
	UserRepo
	RoleRepo
	UserPermissionRepo
}

type authRepo struct {
	db *db.Postgres
}

func NewAuthRepo(db *db.Postgres) AuthRepo {
	return &authRepo{db}
}

func (r *authRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// SQL query to insert a new user into the User table
	query := `
		INSERT INTO User (username, name, email, is_verified, gender, password, role_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	var userID int
	err := r.db.Pool.QueryRow(ctx, query,
		user.Username,
		user.Name,
		user.Email,
		user.IsVerified,
		user.Gender,
		user.Password,
		user.RoleID,
		time.Now(),
	).Scan(&userID)

	if err != nil {
		return nil, err
	}

	user.ID = userID

	return user, nil
}
