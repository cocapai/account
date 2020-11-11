package common

import (
	"github.com/dgrijalva/jwt-go"
	"scutrobot.buff/go_demo/model"
	"time"
)

// 私钥是不能泄露的
var jwtKey = []byte("a_secret_key")

type Claim struct {
	UserId uint
	jwt.StandardClaims
}

//发放jwt
func ReleaseToken(user model.User) (string, error) {
	// 一个星期的过期时间
	expirationTime := time.Now().Add(7*24*time.Hour)
	claims := &Claim{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			// 签发token的时间的
			IssuedAt: time.Now().Unix(),
			// 签发token的人
			Issuer: "scutrobot.buff",
			// 加密为用户token名
			Subject: "user token",
		},
	}
	// 256加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "",err
	}

	return tokenString, nil
}
//jwt解包
func ParseToken(tokenString string) (*jwt.Token, *Claim, error){
	claims := &Claim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}