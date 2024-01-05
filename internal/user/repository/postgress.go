package repository

import (
	"context"
	"database/sql"

	"github.com/Hareshutit/ShopEase/internal/user/domain"
	"github.com/rs/zerolog"

	"github.com/google/uuid"
)

type UserPostgressRepository struct {
	users *sql.DB
	log   zerolog.Logger
}

func (t UserPostgressRepository) AuthGetId(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var (
		email        string
		phoneNumber  string
		login        string
		password     string
		name         string
		pathToAvatar string
	)

	row := t.users.QueryRow("SELECT email, phonenumber, login, password, name, pathtoavatar FROM users WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&email, &phoneNumber, &login, &password, &name, &pathToAvatar)

	return domain.User{
		Email:       email,
		PhoneNumber: phoneNumber,

		Login:    login,
		Password: password,

		Name: name,

		PathToAvatar: pathToAvatar,
	}, err
}

func (t UserPostgressRepository) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var (
		phoneNumber  string
		name         string
		pathToAvatar string
		email        string
		login        string
		password     string
	)

	row := t.users.QueryRow("SELECT  email, login, password, phonenumber, name, pathtoavatar FROM users WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&email, &login, &password, &phoneNumber, &name, &pathToAvatar)

	return domain.User{
		Email:    email,
		Login:    login,
		Password: password,

		PhoneNumber: phoneNumber,
		Name:        name,

		PathToAvatar: pathToAvatar,
	}, err
}

func (t UserPostgressRepository) FindById(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var (
		phoneNumber  string
		name         string
		pathToAvatar string
	)

	row := t.users.QueryRow("SELECT  phonenumber, name, pathtoavatar FROM users WHERE id = $1 LIMIT 1", id)
	err := row.Scan(&phoneNumber, &name, &pathToAvatar)

	return domain.User{
		PhoneNumber: phoneNumber,

		Name: name,

		PathToAvatar: pathToAvatar,
	}, err
}

func (t UserPostgressRepository) CheckUser(ctx context.Context, login string, password string) string {
	var checkUser string

	row := t.users.QueryRow("SELECT id FROM users WHERE login = $1 and password = $2", login, password)
	row.Scan(&checkUser)

	return checkUser
}

func (t *UserPostgressRepository) Create(ctx context.Context, user domain.User) error {
	_, err := t.users.Exec("insert into users (id, email,  phonenumber, login, password, name, pathtoavatar) values ($1, $2, $3, $4, $5, $6, $7)",
		user.Id, user.Email, user.PhoneNumber, user.Login, user.Password, user.Name,
		user.PathToAvatar)
	// Должен возвращаться или ссесия или jwt токен
	return err
}

func (t *UserPostgressRepository) Update(ctx context.Context, user domain.User) error {
	_, err := t.users.Exec("update users set email = $1,  phonenumber = $2, login = $3, password = $4, name = $5, pathtoavatar = $6 where id = $7",
		user.Email, user.PhoneNumber, user.Login, user.Password, user.Name,
		user.PathToAvatar, user.Id)
	return err
}

func (t *UserPostgressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := t.users.Exec("delete from users where id = $1", id)
	//logout
	return err
}
