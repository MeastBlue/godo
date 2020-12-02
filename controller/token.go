package controller

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/meastblue/godo/util"
)

// RefreshToken controller function
func RefreshToken(c *gin.Context) {
	token, err := util.VerifyRefreshToken(c)
	if err != nil {
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
		if delErr != nil || deleted == 0 {
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
