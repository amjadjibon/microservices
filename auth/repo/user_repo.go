package repo

import (
	"context"
	"time"

	"github.com/amjadjibon/microservices/auth/model"
	"github.com/amjadjibon/microservices/pkg/db"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserById(ctx context.Context, userID int) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetAllUser(ctx context.Context) ([]*model.User, error)
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
	sql, args, err := r.db.Builder.Insert("auth_user").Columns(
		"username",
		"name",
		"email",
		"is_verified",
		"gender",
		"password",
		"role",
		"created_at",
		"updated_at",
	).Values(
		user.Username,
		user.Name,
		user.Email,
		user.IsVerified,
		user.Gender,
		user.Password,
		user.Role,
		time.Now(),
		time.Now(),
	).Suffix("RETURNING id").ToSql()
	if err != nil {
		return nil, err
	}

	var userID int
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(&userID)
	if err != nil {
		return nil, err
	}

	user.ID = userID
	return user, nil
}

func (r *authRepo) GetUserById(ctx context.Context, userID int) (*model.User, error) {
	query, args, err := r.db.Builder.
		Select("*").
		From("auth_user").
		Where("id = ?", userID).
		Where("deleted_at IS NULL").
		ToSql()

	if err != nil {
		return nil, err
	}

	return r.getUser(ctx, query, args)
}

func (r *authRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query, args, err := r.db.Builder.
		Select(
			"id",
			"username",
			"name",
			"email",
			"is_verified",
			"gender",
			"password",
			"role",
			"created_at",
			"updated_at",
		).
		From("auth_user").
		Where("username = ?", username).
		Where("deleted_at IS NULL").
		ToSql()

	if err != nil {
		return nil, err
	}

	return r.getUser(ctx, query, args)
}

func (r *authRepo) GetAllUser(ctx context.Context) ([]*model.User, error) {
	query, _, err := r.db.Builder.
		Select(
			"id",
			"username",
			"name",
			"email",
			"is_verified",
			"gender",
			"password",
			"role",
			"created_at",
			"updated_at",
		).
		From("auth_user").
		Where("deleted_at IS NULL").
		ToSql()

	if err != nil {
		return nil, err
	}

	return r.getUsers(ctx, query)
}

func (r *authRepo) getUser(ctx context.Context, query string, args []interface{}) (*model.User, error) {
	var user model.User
	if err := r.db.Pool.
		QueryRow(ctx, query, args...).
		Scan(
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Email,
			&user.IsVerified,
			&user.Gender,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepo) getUsers(ctx context.Context, query string) ([]*model.User, error) {
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Name,
			&user.Email,
			&user.IsVerified,
			&user.Gender,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
