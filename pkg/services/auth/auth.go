package auth

import (
	"time"

	"github.com/SassoStorTo/studenti-italici/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

type ValidInt interface { // Todo: rename interface
	Redirect(location string, status ...int) error
}

func IsValidToken(cookieValue string, isRefresh bool, c ValidInt) (*models.User, error) {
	user, err := ParseToken(cookieValue)
	if err != nil || user.Exp < time.Now().Unix() || user.Refresh != isRefresh {
		return nil, c.Redirect("/refresh-access-token") //Todo: check the route
	}
	if !user.IsEditor {
		return nil, c.Redirect("/wait-accept")
	}

	savedUser, err := models.GetUserById(user.Id)
	if err != nil || !IsCookieUpToDate(savedUser, user) {
		return nil, err
		return nil, c.Redirect("/refresh-access-token") //Todo: check the route
	}

	return savedUser, nil
}

func IsCookieUpToDate(s *models.User, u *UserClaims) bool {
	return (s.IsAdmin == u.IsAdmin) && (s.IsEditor == u.IsEditor)
}

func GetAccessToken(usr *models.User) (string, time.Time, error) {
	exp := time.Now().Add(time.Hour * 24) // 24 hours
	token, err := getToken(usr, false, exp)
	return token, exp, err
}

func GetRefreshToken(usr *models.User) (string, time.Time, error) {
	exp := time.Now().Add(time.Hour * 24 * 30 * 6) // 6 months
	token, err := getToken(usr, true, exp)
	return token, exp, err
}

func getToken(usr *models.User, isRefresh bool, exp time.Time) (string, error) {
	token, err := NewToken(usr.Id, usr.IsAdmin, usr.IsEditor, isRefresh, exp)
	return token, err
}

type UserClaims struct {
	Id       int   `json:"id"`
	IsAdmin  bool  `json:"is_admin"`
	IsEditor bool  `json:"is_editor"`
	Refresh  bool  `json:"refresh"`
	Exp      int64 `json:"exp"`
	jwt.RegisteredClaims
}

func NewToken(id int, isAdmin bool, isEditor bool, refresh bool, exp_time time.Time) (string, error) {
	claims := UserClaims{id, isAdmin, isEditor, refresh, exp_time.Unix(), jwt.RegisteredClaims{}}
	unsigend_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token, err := unsigend_token.SignedString([]byte(config.Secret)) //todo: set a proper secret
	return unsigend_token.SignedString([]byte("segreto"))
}

func ParseToken(accessToken string) (*UserClaims, error) {
	claims := &UserClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("segreto"), nil // Todo: change the way of getting the secret
	})
	return claims, err
}
