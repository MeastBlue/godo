package service

import (
	"github.com/meastblue/godo/database"
	"github.com/meastblue/godo/model"
	"github.com/meastblue/godo/util"
)

func Login(auth *model.Auth) error {
	user := model.User{}
	db, err := database.GetDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	stmt, err := db.Preparex(`select * from users where nickname=? or email=?`)
	if err != nil {
		return err
	}

	err = stmt.Get(&user, auth.Username)
	if err != nil {
		return err
	}

	err = util.VerifyPassword(user.Password, auth.Password)
	if err != nil {
		return err
	}

	return nil
}
