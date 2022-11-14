package users

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"time"

	"howToUseMod/modules/db"

	"github.com/google/uuid"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

var Db *sql.DB = db.Db

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
		from users 
		where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) Create() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	Result, err := Db.Exec(
		cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(Result)
	return err
}

func (u *User) Update() (err error) {
	cmd := `update users set name = ?, email = ?
		where id = ?`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) Delete() (err error) {
	cmd := `delete from users
		where id = ?`

	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// UUIDを作成
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// 暗号化
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
