package service

import (
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
	stmt, err := db.Preparex(`select * from users where id=?`)
	if err != nil {
		return nil, err
	}

	err = stmt.Get(&user, id)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func AddUser(user *model.User) (string, error) {
	id := ""
	db, err := database.GetDatabase()
	if err != nil {
		return "", err
	}

	defer db.Close()
	stmt, err := db.Preparex(`insert into users(nickname, email, password) values (?, ?, ?) returning id`)
	if err != nil {
		return "", err
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	row := stmt.QueryRowx(user.Nickname, user.Email, hashedPassword)
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateUser(user *model.User) error {
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	stmt, err := db.Preparex(`update users set nickname=?, email=?, password=? where id=?`)
	if err != nil {
		return err
	}

	_, err = stmt.Queryx(user.Nickname, user.Email, user.Password, user.ID)
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
	stmt, err := db.Preparex(`delete from users where id=?`)
	if err != nil {
		return err
	}

	_, err = stmt.Queryx(id)
	if err != nil {
		return err
	}

	return nil

}
