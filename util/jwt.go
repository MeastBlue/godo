package util

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/meastblue/godo/model"
)

func CreateToken(id string) (*model.Token, error) {
	td := model.Token{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	accessUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	td.AccessUUID = accessUUID.String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	td.RefreshUUID = refreshUUID.String()
	accessToken, err := createAccessToken(&td, id)
	if err != nil {
		return nil, err
	}

	td.AccessToken = accessToken
	refreshToken, err := createRefreshToken(&td, id)
	if err != nil {
		return nil, err
	}

	td.RefreshToken = refreshToken

	return &td, nil
}

func createAccessToken(token *model.Token, id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = token.AccessUUID
	atClaims["user_id"] = id
	atClaims["exp"] = token.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return at.SignedString([]byte(os.Getenv("jwt.Access")))
}

func createRefreshToken(token *model.Token, id string) (string, error) {
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = token.RefreshUUID
	rtClaims["user_id"] = id
	rtClaims["exp"] = token.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	return rt.SignedString([]byte(os.Getenv("jwt.Refresh")))
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := MapToken(tokenString)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := MapToken(tokenString)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func MapToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("jwt.Access")), nil
	})
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userId, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}

		return &model.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}

	return nil, err
}
