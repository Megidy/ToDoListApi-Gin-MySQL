package models

import (
	"database/sql"
	"log"

	"github.com/Megidy/To-Do-List-Api/pkj/config"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int    `json:"id"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var db *sql.DB

func init() {
	config.Connect()
	db = config.GetDb()
}

func CreateUser(user *User) (*User, error) {
	_, err := db.Exec("insert into users (nickname,email,password) values(?,?,?)",
		user.NickName, user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, nil

}
func IsSignedUp(NewUser User) (bool, string) {
	var user User
	row := db.QueryRow("select email, nickname from users where email = ? or nickname = ?",
		NewUser.Email, NewUser.NickName)
	err := row.Scan(&user.Email, &user.NickName)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows found, meaning the user is not signed up
			return false, ""
		}

	}
	if NewUser.Email == user.Email || NewUser.NickName == user.NickName {
		return true, "user with this nickname or email is already signed up"
	}

	return false, ""
}
func FindUserByEmail(user User) (User, error) {
	var SignedUser User
	row := db.QueryRow("select * from users where email =?",
		user.Email)
	err := row.Scan(&SignedUser.Id, &SignedUser.NickName, &SignedUser.Email, &SignedUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, err
		}

	}
	return SignedUser, nil

}

func FindUserById(id float64) (User, error) {
	var user User
	row := db.QueryRow("select * from users where id= ?",
		id)
	err := row.Scan(&user.Id, &user.NickName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, err
		}
	}
	return user, nil
}
