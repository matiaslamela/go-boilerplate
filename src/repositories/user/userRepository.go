package user

import (
	"context"
	"database/sql"

	"time"

	userModel "github.com/matiaslamela/go-boilerplate/src/database/user"
)

type Repository struct {
	DB *sql.DB
}

func (ur *Repository) GetAll(ctx context.Context) (*[]userModel.User, error) {
	q := `
    SELECT id, first_name, last_name, username, email, picture,
        created_at, updated_at
        FROM users;
    `
	rows, err := ur.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []userModel.User
	for rows.Next() {
		var u userModel.User
		rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username,
			&u.Email, &u.Picture, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, u)
	}
	if &users == nil {
		users = make([]userModel.User, 0)
	}
	return &users, nil
}

func (ur *Repository) GetOne(ctx context.Context, id uint) (*userModel.User, error) {
	q := `
    SELECT id, first_name, last_name, username, email, picture,
        created_at, updated_at
        FROM users WHERE id = $1;
    `

	row := ur.DB.QueryRowContext(ctx, q, id)

	var u userModel.User
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username, &u.Email,
		&u.Picture, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (ur *Repository) GetByUsername(ctx context.Context, username string) (*userModel.User, error) {
	q := `
    SELECT id, first_name, last_name, username, email, picture,
        password, created_at, updated_at
        FROM users WHERE username = $1;
    `

	row := ur.DB.QueryRowContext(ctx, q, username)

	var u userModel.User
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username,
		&u.Email, &u.Picture, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (ur *Repository) Create(ctx context.Context, u userModel.User) (*userModel.User, error) {
	q := `
    INSERT INTO users (first_name, last_name, username, email, picture, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id;
    `

	if u.Picture == "" {
		u.Picture = "https://placekitten.com/g/300/300"
	}

	if err := u.HashPassword(); err != nil {
		return nil, err
	}

	row := ur.DB.QueryRowContext(
		ctx, q, u.FirstName, u.LastName, u.Username, u.Email,
		u.Picture, u.PasswordHash, time.Now(), time.Now(),
	)

	var newUser userModel.User
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Username,
		&u.Email, &u.Picture, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (ur *Repository) Update(ctx context.Context, id uint, u userModel.User) (*userModel.User, error) {
	q := `
    UPDATE users set first_name=$1, last_name=$2, email=$3, picture=$4, updated_at=$5
        WHERE id=$6;
    `

	stmt, err := ur.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, u.FirstName, u.LastName, u.Email,
		u.Picture, time.Now(), id,
	)
	if err != nil {
		return nil, err
	}
	getQuery := `
    SELECT id, first_name, last_name, username, email, picture,
        password, created_at, updated_at
        FROM users WHERE username = $1;
    `
	row := stmt.QueryRowContext(ctx, getQuery, id)
	var updatedUser userModel.User
	err = row.Scan(&updatedUser.ID, &updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Username, &updatedUser.Email,
		&updatedUser.Picture, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (ur *Repository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM users WHERE id=$1;`

	stmt, err := ur.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
