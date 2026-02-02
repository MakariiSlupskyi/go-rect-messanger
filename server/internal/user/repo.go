package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	Create(ctx context.Context, u *User) error
}

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repo{db: db}
}

func (r *repo) GetAll(ctx context.Context) ([]User, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, username, displayName, passwordHash
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.DisplayName,
			&u.PasswordHash,
		); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, rows.Err()
}

func (r *repo) Create(ctx context.Context, u *User) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO users (username, displayName, passwordHash)
		VALUES ($1, $2, $3)
		RETURNING ID
	`,
		u.Username,
		u.DisplayName,
		u.PasswordHash,
	).Scan(&u.ID)
}
