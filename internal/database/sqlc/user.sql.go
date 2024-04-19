// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateUserParams struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const createUserAddress = `-- name: CreateUserAddress :exec
INSERT INTO address (id, user_id, cep, ibge, uf, city, complement, street, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`

type CreateUserAddressParams struct {
	ID         string
	UserID     string
	Cep        string
	Ibge       string
	Uf         string
	City       string
	Complement sql.NullString
	Street     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (q *Queries) CreateUserAddress(ctx context.Context, arg CreateUserAddressParams) error {
	_, err := q.db.ExecContext(ctx, createUserAddress,
		arg.ID,
		arg.UserID,
		arg.Cep,
		arg.Ibge,
		arg.Uf,
		arg.City,
		arg.Complement,
		arg.Street,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const findManyUsers = `-- name: FindManyUsers :many
SELECT u.id, u.name, u.email, u.created_At, u.updated_at, a.cep, a.uf, a.city, a.complement, a.street
FROM users u
         JOIN address a ON a.user_id = u.id
ORDER BY u.created_at DESC
`

type FindManyUsersRow struct {
	ID         string
	Name       string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Cep        string
	Uf         string
	City       string
	Complement sql.NullString
	Street     string
}

func (q *Queries) FindManyUsers(ctx context.Context) ([]FindManyUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, findManyUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindManyUsersRow
	for rows.Next() {
		var i FindManyUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Cep,
			&i.Uf,
			&i.City,
			&i.Complement,
			&i.Street,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT u.id, u.name, u.email FROM users u WHERE u.email = $1
`

type FindUserByEmailRow struct {
	ID    string
	Name  string
	Email string
}

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i FindUserByEmailRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT u.id, u.name, u.email, u.created_At, u.updated_at, a.cep, a.uf, a.city, a.complement, a.street
FROM users u
         JOIN address a ON a.user_id = u.id
WHERE u.id = $1
`

type FindUserByIDRow struct {
	ID         string
	Name       string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Cep        string
	Uf         string
	City       string
	Complement sql.NullString
	Street     string
}

func (q *Queries) FindUserByID(ctx context.Context, id string) (FindUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, findUserByID, id)
	var i FindUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Cep,
		&i.Uf,
		&i.City,
		&i.Complement,
		&i.Street,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, email, password, created_at, updated_at from users u where u.id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserPassword = `-- name: GetUserPassword :one
SELECT u.password FROM users u WHERE u.id = $1
`

func (q *Queries) GetUserPassword(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserPassword, id)
	var password string
	err := row.Scan(&password)
	return password, err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users SET password = $2, updated_at = $3 WHERE id = $1
`

type UpdatePasswordParams struct {
	ID        string
	Password  string
	UpdatedAt time.Time
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.ID, arg.Password, arg.UpdatedAt)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET
             name = COALESCE($3, name),
             email = COALESCE($4, email),
             updated_at = $2
WHERE id = $1
`

type UpdateUserParams struct {
	ID        string
	UpdatedAt time.Time
	Name      sql.NullString
	Email     sql.NullString
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.ID,
		arg.UpdatedAt,
		arg.Name,
		arg.Email,
	)
	return err
}

const updateUserAddress = `-- name: UpdateUserAddress :exec
UPDATE address SET
               cep = COALESCE($3, cep),
               ibge = COALESCE($4, ibge),
               uf = COALESCE($5, uf),
               city = COALESCE($6, city),
               complement = COALESCE($7, complement),
               street = COALESCE($8, street),
               updated_at = $2
WHERE user_id = $1
`

type UpdateUserAddressParams struct {
	UserID     string
	UpdatedAt  time.Time
	Cep        sql.NullString
	Ibge       sql.NullString
	Uf         sql.NullString
	City       sql.NullString
	Complement sql.NullString
	Street     sql.NullString
}

func (q *Queries) UpdateUserAddress(ctx context.Context, arg UpdateUserAddressParams) error {
	_, err := q.db.ExecContext(ctx, updateUserAddress,
		arg.UserID,
		arg.UpdatedAt,
		arg.Cep,
		arg.Ibge,
		arg.Uf,
		arg.City,
		arg.Complement,
		arg.Street,
	)
	return err
}
