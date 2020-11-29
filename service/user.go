package service

import (
	"time"

	"github.com/meastblue/godo/database"
	"github.com/meastblue/godo/model"
	"github.com/meastblue/godo/util"
)

func GetUsers() (*model.Users, error) {
	users := model.Users{}
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	stmt, err := db.Preparex(`select * from users`)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func GetUser(id string) (*model.User, error) {
	user := model.User{}
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	stmt, err := db.Preparex(`select * from users where id=$1`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&user, id)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func AddUser(user *model.User) error {

	db, err := database.GetDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	stmt, err := db.Preparex(`insert into users(nickname, email, password) values ($1, $2 ,$3) returning *`)
	if err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Queryx(user.Nickname, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user *model.User) error {
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}

	user.UpdatedAt = time.Now()

	defer db.Close()
	stmt, err := db.Preparex(`update users set nickname=$1, email=$2, password=$3, updated_at=$4 where id=$5`)
	if err != nil {
		return err
	}

	_, err = stmt.Queryx(user.Nickname, user.Email, user.Password, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	stmt, err := db.Preparex(`delete from users where id=$1`)
	if err != nil {
		return err
	}

	_, err = stmt.Queryx(id)
	if err != nil {
		return err
	}

	return nil

}
