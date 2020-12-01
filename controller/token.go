package controller

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/util"
)

func RefreshToken(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		util.SendJsonUnprocessableEntity(c, err.Error())
		return
	}
	
	refreshToken := mapToken["refresh_token"]
	token, err := util.MapToken(refreshToken)
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		util.SendJsonUnauthorized(c, err.Error())
		return
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		util.SendJsonUnauthorized(c, err.Error())
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshID, ok := claims["refresh_id"].(string)
		if !ok {
			util.SendJsonUnprocessableEntity(c, err.Error())
			return
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			util.SendJsonUnprocessableEntity(c, err.Error())
			return
		}

		deleted, delErr := util.DeleteAuth(refreshID)
		if delErr != nil || deleted == 0 { //if any goes wrong
			util.SendJsonUnauthorized(c, delErr.Error())
			return
		}

		ts, createErr := util.CreateToken(userID)
		if createErr != nil {
			util.SendJsonUnauthorized(c, createErr.Error())
			return
		}

		saveErr := util.CreateAuth(userID, ts)
		if saveErr != nil {
			util.SendJsonUnauthorized(c, saveErr.Error())
			return
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}

		util.SendJsonCreated(c, tokens)
	} else {
		util.SendJsonUnauthorized(c, errors.New("refresh expired"))
	}
}
