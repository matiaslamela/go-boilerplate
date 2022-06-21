package postRepository

import (
	"context"
	"database/sql"
	"time"

	postModel "github.com/matiaslamela/go-boilerplate/src/database/post"
)

type PostRepository struct {
	DB *sql.DB
}

func (pr *PostRepository) GetAll(ctx context.Context) ([]postModel.Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts;
    `

	rows, err := pr.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []postModel.Post
	for rows.Next() {
		var p postModel.Post
		rows.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) GetOne(ctx context.Context, id uint) (postModel.Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts WHERE id = $1;
    `

	row := pr.DB.QueryRowContext(ctx, q, id)

	var p postModel.Post
	err := row.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return postModel.Post{}, err
	}

	return p, nil
}

func (pr *PostRepository) GetByUser(ctx context.Context, userID uint) ([]postModel.Post, error) {
	q := `
    SELECT id, body, user_id, created_at, updated_at
        FROM posts
        WHERE user_id = $1;
    `

	rows, err := pr.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []postModel.Post
	for rows.Next() {
		var p postModel.Post
		rows.Scan(&p.ID, &p.Body, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) Create(ctx context.Context, p *postModel.Post) error {
	q := `
    INSERT INTO posts (body, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `

	stmt, err := pr.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Body, p.UserID, time.Now(), time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) Update(ctx context.Context, id uint, p postModel.Post) error {
	q := `
    UPDATE posts set body=$1, updated_at=$2
        WHERE id=$3;
    `

	stmt, err := pr.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Body, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM posts WHERE id=$1;`

	stmt, err := pr.DB.PrepareContext(ctx, q)
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
