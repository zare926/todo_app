package models

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID         int
	UUID       string
	Name       string
	Email      string
	Password   string
	Created_At time.Time
	Todos      []Todo
}

type Session struct {
	ID         int
	UUID       string
	Email      string
	UserID     int
	Created_At time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values ($1,$2,$3,$4,$5)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password,created_at from users where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Created_At,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = $1, email = $2 where id = $3`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = $1`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	fmt.Println("todoiteru?")
	cmd := `select id, uuid, name, email, password, created_at from users where email = $1`

	err = Db.QueryRow(cmd, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.Created_At)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into session (
		uuid,
		email,
		user_id,
		created_at) values ($1,$2,$3,$4)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())

	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at
		from session where user_id = $1 and email = $2`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.Created_At)
	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
	from session where uuid = $1`

	err = Db.QueryRow(cmd, sess.UUID).Scan(&sess.ID, &sess.UUID, &sess.Email, &sess.UserID, &sess.Created_At)

	if err != nil {
		valid = false
		return
	}

	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from session where uuid = $1`

	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email,created_at from users where id = $1`

	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Created_At)

	return user, err
}
