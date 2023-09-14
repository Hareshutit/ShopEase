package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/Hareshutit/ShopEase/internal/post/domain"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PostPostgressRepository struct {
	post *sql.DB
}

func (t PostPostgressRepository) GetById(ctx context.Context,
	id uuid.UUID) (*domain.Post, int, error) {

	var result domain.Post
	result.New()

	err := t.post.QueryRow(`SELECT userid, title, description,
	views, price, status, tags, images, time FROM post WHERE id = $1
	LIMIT 1`, id).Scan(result.UserID, result.Title, result.Description,
		result.Views, result.Price, result.Status, result.Category,
		pq.Array(result.PathImages), result.Time)

	if err == sql.ErrNoRows {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	// Удалить от сюда
	//*result.Views = *result.Views + 1
	//_, err = t.post.Exec("update post set views = $1 where id = $2",
	//	*result.Views, id)
	//До сюда
	return &result, http.StatusOK, nil
}

func (t PostPostgressRepository) GetMiniObject(ctx context.Context,
	par domain.Parameters) ([]domain.Post, int, error) {

	if *par.Offset < 0 {
		return nil, http.StatusBadRequest, errors.New("The page number cannot be less than zero")
	}

	var rows *sql.Rows
	var err error

	if par.Sort == nil {
		rows, err = t.post.Query(`SELECT id, userid, title, description, price, images, time 
		FROM post
		WHERE (tags = COALESCE($1, tags) OR $1 IS NULL) AND
			  (status = COALESCE($2, status) OR $2 IS NULL) AND
			  (userid = COALESCE($3, userid) OR $3 IS NULL)
		ORDER BY time LIMIT $4 OFFSET $5`, par.Category, par.Status, par.UserId,
			par.Limit, par.Offset)
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var post []domain.Post

	for rows.Next() {
		var p domain.Post
		p.New()
		err = rows.Scan(p.Id, p.UserID, p.Title, p.Description,
			p.Price, pq.Array(p.PathImages), p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		post = append(post, p)
	}

	if err = rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = rows.Close(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return post, http.StatusOK, nil
}

func (t *PostPostgressRepository) Create(ctx context.Context,
	post domain.Post) (int, error) {

	_, err := t.post.Exec(`insert into post (id, userid, title,
		description, price, status, tags, images, time)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		post.Id, post.UserID, post.Title, post.Description, post.Price,
		true, post.Category, pq.Array(post.PathImages), post.Time)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) Update(ctx context.Context,
	post domain.Post) (int, error) {

	_, err := t.post.Exec(`update post set
	title = CASE WHEN $1::text IS NOT NULL THEN $1::text ELSE title END,
	description = CASE WHEN $2::text IS NOT NULL THEN $2::text ELSE description END,
	price =  CASE WHEN $3::int IS NOT NULL THEN $3::int ELSE price END,
	tags =  CASE WHEN $4::text IS NOT NULL THEN $4::text ELSE tags END,
	status =  CASE WHEN $5::boolean IS NOT NULL THEN $5::boolean ELSE status END,
	images = CASE WHEN $6::text[] IS NOT NULL THEN $6::text[] ELSE images END
	where id = $7 and userid = $8`,
		post.Title, post.Description, post.Price,
		post.Description, post.Status, pq.Array(post.PathImages), post.Id, post.UserID)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) IncrementViews(ctx context.Context,
	PostId uuid.UUID, UserId uuid.UUID) (int, error) {

	_, err := t.post.Exec(`update views set count=count+1 where PostId = $1`, PostId)

	if err != nil {
		//Добавить логирование
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) Delete(ctx context.Context,
	id uuid.UUID) (int, error) {

	_, err := t.post.Exec(`delete from post where id = $1`, id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
