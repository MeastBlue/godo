package util

import (
	"log"
	"time"

	"github.com/meastblue/godo/database"
	"github.com/meastblue/godo/model"
)

func CreateAuth(id string, token *model.Token) error {
	at := time.Unix(token.AtExpires, 0)
	rt := time.Unix(token.RtExpires, 0)
	now := time.Now()

	client := database.IniStorage()

	errAccess := client.Set(token.AccessUUID, id, at.Sub(now)).Err()
	if errAccess != nil {
		log.Fatalf("NICKNAME: %s\n", errAccess.Error())
		return errAccess
	}

	errRefresh := client.Set(token.RefreshUUID, id, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

func FetchAuth(authD *model.AccessDetails) (string, error) {
	client := database.IniStorage()
	userid, err := client.Get(authD.AccessUuid).Result()
	if err != nil {
		return "", err
	}

	return userid, nil
}

func DeleteAuth(id string) (int64, error) {
	client := database.IniStorage()
	deleted, err := client.Del(id).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
