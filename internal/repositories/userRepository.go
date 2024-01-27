package repository

import (
	"context"
	"user-service/internal/canonical"
)

type UserRepository interface {
	CreateUser(context.Context, canonical.User) error
	GetUserById(context.Context, string) (*canonical.User, error)
	GetUserByLogin(context.Context, string) (*canonical.User, error)
	GetAllUsers(ctx context.Context) ([]canonical.User, error)
}

type userRepository struct {
	db PgxIface
}

func NewUserRepo(db PgxIface) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(ctx context.Context, user canonical.User) error {
	sqlStatement := "INSERT INTO \"User\" (Id, AccessLevelID, Login, Password, Createdat) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.db.Exec(ctx, sqlStatement, user.Id, user.AccessLevelID, user.Login, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByLogin(ctx context.Context, login string) (*canonical.User, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"User\" WHERE LOGIN = $1",
		login,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user canonical.User
	if rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.AccessLevelID,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, ErrorNotFound
}

func (r *userRepository) GetUserById(ctx context.Context, id string) (*canonical.User, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"User\" WHERE ID = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user canonical.User
	if rows.Next() {
		if err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.AccessLevelID,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, ErrorNotFound
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]canonical.User, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM \"User\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []canonical.User

	for rows.Next() {
		var user canonical.User

		if err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.AccessLevelID,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
