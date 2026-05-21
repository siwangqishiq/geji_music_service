package component

import (
	"fmt"
	"geji/data"
	"geji/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func init() {
	util.Logi("component jwt init")
}

var secret = []byte("maolilan_gongtengxinyi")

// token存活时间  30天
var tokenDuring time.Duration = time.Hour * 24 * 30

type UserClaims struct {
	Uid int64 `json:"uid"`
}

func (u *UserClaims) IsValidated() bool {
	if u.Uid > 0 {
		return true
	} else {
		return false
	}
}

func GenerateToken(uid int64) (string, error) {
	claims := jwt.MapClaims{
		data.KEY_USER_ID: uid,
		data.KEY_EXPIRE:  time.Now().Add(tokenDuring).UnixMilli(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}

func ParseTokenToUserClaims(tokenString string) (UserClaims, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return UserClaims{
			Uid: -1,
		}, fmt.Errorf("parse token error")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		uid := int64(claims[data.KEY_USER_ID].(float64))

		return UserClaims{
			Uid: uid,
		}, fmt.Errorf("parse token error")
	} else {
		return UserClaims{
			Uid: -1,
		}, fmt.Errorf("parse token error")
	}
}
