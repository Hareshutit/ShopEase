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
	posts *sql.DB
}

func (t PostPostgressRepository) GetIdPost(ctx context.Context,
	id uuid.UUID) (*domain.Post, int, error) {

	var result domain.Post
	result.New()

	err := t.posts.QueryRow(`SELECT userid, title, description,
	views, price, close, tags, images, time FROM posts WHERE id = $1
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
	//_, err = t.posts.Exec("update posts set views = $1 where id = $2",
	//	*result.Views, id)
	//До сюда
	return &result, http.StatusOK, nil
}

func (t PostPostgressRepository) GetMiniPostSortNew(ctx context.Context,
	par domain.Parameters) ([]domain.Post, int, error) {

	if *par.Offset < 0 {
		return nil, http.StatusBadRequest, errors.New("The page number cannot be less than zero")
	}

	var rows *sql.Rows
	var err error

	if par.Sort == nil {
		rows, err = t.posts.Query(`SELECT id, userid, title, description, price, images, time 
		FROM posts
		WHERE (tags = COALESCE($1, tags) OR $1 IS NULL) AND
			  (close = COALESCE($2, close) OR $2 IS NULL) AND
			  (userid = COALESCE($3, userid) OR $3 IS NULL)
		ORDER BY time LIMIT $4 OFFSET $5`, par.Category, par.Status, par.UserId,
			par.Limit, par.Offset)
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var posts []domain.Post

	for rows.Next() {
		var p domain.Post
		p.New()
		err = rows.Scan(p.Id, p.UserID, p.Title, p.Description,
			p.Price, pq.Array(p.PathImages), p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = rows.Close(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return posts, http.StatusOK, nil
}

func (t PostPostgressRepository) GetFavorite(ctx context.Context,
	userId uuid.UUID) ([]uuid.UUID, int, error) {

	rows, err := t.posts.Query(`SELECT idpost FROM favorite
	WHERE userid = $1`, userId)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var postId []uuid.UUID

	for rows.Next() {
		p := uuid.UUID{}
		err = rows.Scan(pq.Array(&postId))
		if err != nil {
			fmt.Println(err)
			continue
		}
		q := uuid.UUID{}
		if p != q {
			postId = append(postId, p)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err = rows.Close(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return postId, http.StatusOK, nil
}

func (t PostPostgressRepository) GetCart(ctx context.Context,
	postId []uuid.UUID) ([]domain.Post, int, error) {
	var posts []domain.Post

	for _, idp := range postId {
		fmt.Println(idp)
		row := t.posts.QueryRow(`SELECT id, userid, title, description,
		price,  tags, images, time FROM posts WHERE close = false and id = $1`, idp)
		fmt.Println(row)
		p := domain.Post{}
		err := row.Scan(&p.Id, &p.UserID, &p.Title, &p.Description,
			&p.Price, &p.Status, pq.Array(&p.PathImages), &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(p)
		posts = append(posts, p)
	}
	return posts, http.StatusOK, nil
}

func (t PostPostgressRepository) SearchPost(ctx context.Context,
	search string) ([]uuid.UUID, int, error) {

	_, err := t.posts.Exec(`UPDATE posts SET fts = setweight(to_tsvector(title), 'A')
	|| setweight(to_tsvector(description), 'B')`)

	rows, err := t.posts.Query(`Select id FROM posts WHERE fts @@ to_tsquery($1)`, search)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var posts []uuid.UUID

	for rows.Next() {
		p := uuid.UUID{}
		err = rows.Scan(&p)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}

	// Подумать нд именованием ошибки
	if err = rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	// Подумать нд именованием ошибки
	if err = rows.Close(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return posts, http.StatusOK, nil
}

func (t *PostPostgressRepository) Create(ctx context.Context,
	post domain.Post) (int, error) {

	_, err := t.posts.Exec(`insert into posts (id, userid, title,
		description, price, close, tags, images, time, views)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		post.Id, post.UserID, post.Title, post.Description, post.Price,
		true, post.Category, pq.Array(post.PathImages), post.Time, 1)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) Update(ctx context.Context,
	post domain.Post) (int, error) {

	_, err := t.posts.Exec(`update posts set
	title = CASE WHEN $1::text IS NOT NULL THEN $1::text ELSE title END,
	description = CASE WHEN $2::text IS NOT NULL THEN $2::text ELSE description END,
	price =  CASE WHEN $3::text IS NOT NULL THEN $3::text ELSE price END,
	tags =  CASE WHEN $4::text IS NOT NULL THEN $4::text ELSE tags END,
	close =  CASE WHEN $5::boolean IS NOT NULL THEN $5::boolean ELSE close END,
	images = CASE WHEN $6::text[] IS NOT NULL THEN $6::text[] ELSE images END
	where id = $7 and userid = $8`,
		post.Title, post.Description, post.Price,
		post.Description, post.Status, pq.Array(post.PathImages), post.Id, post.UserID)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) Delete(ctx context.Context,
	id uuid.UUID) (int, error) {

	_, err := t.posts.Exec(`delete from posts where id = $1`, id)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) AddFavorite(ctx context.Context, userId uuid.UUID,
	postId uuid.UUID) (int, error) {

	pid, status, err := t.GetFavorite(ctx, userId)

	if status != http.StatusOK {
		return status, err
	}

	pid = append(pid, postId)

	_, err = t.posts.Exec("insert into favorite (idpost, userid) values ($1, $2)",
		pq.Array(pid), userId)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (t *PostPostgressRepository) RemoveFavorite(ctx context.Context, userId uuid.UUID,
	postId uuid.UUID) (int, error) {

	pid, status, err := t.GetFavorite(ctx, userId)

	if status != http.StatusOK {
		return status, err
	}

	var pidClear []uuid.UUID
	for _, fpid := range pid {
		if fpid != postId {
			pidClear = append(pidClear, fpid)
		}
	}

	_, err = t.posts.Exec("update favorite set idpost = $1 where UserId = $2",
		pq.Array(pidClear), userId)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
