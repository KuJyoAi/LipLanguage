package dao

import (
	"LipLanguage/common"
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

type Claim struct {
	jwt.StandardClaims
	Phone    int64
	Nickname string
	UserID   int64
}

func GenerateTokenExpires(Phone int64, Nickname string, UserID int64, duration time.Duration) (string, error) {
	claim := Claim{
		Phone:    Phone,
		Nickname: Nickname,
		UserID:   UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(common.JwtExpireTime).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(common.JwtKey))
	if err != nil {
		logrus.Errorf("[util.GenerateJWT] %v", err)
		return "", err
	}
	//存储在Redis里
	key := fmt.Sprintf("%v_Token", Phone)

	SetRedisToken(key, token, duration)

	return token, err
}

func GenerateToken(Phone int64, Nickname string, UserID int64) (string, error) {
	claim := Claim{
		Phone:    Phone,
		Nickname: Nickname,
		UserID:   UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(common.JwtExpireTime).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(common.JwtKey))
	if err != nil {
		logrus.Errorf("[util.GenerateJWT] %v", err)
		return "", err
	}
	//存储在Redis里
	key := fmt.Sprintf("%v_Token", Phone)
	SetRedisToken(key, token, common.JwtExpireTime)
	return token, err
}

func ParseToken(token string) (*Claim, error) {
	TokenClaim, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.JwtKey), nil
	})

	if err != nil {
		logrus.Errorf("[util.ParseToken] %v", err)
		return nil, err
	}

	if TokenClaim != nil {
		if c, ok := TokenClaim.Claims.(*Claim); ok && TokenClaim.Valid {
			return c, nil
		}
	}

	logrus.Errorf("[util.ParseToken] %v,\n %v", err, TokenClaim)
	return nil, err
}

func Hash256(src string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(src)))
}
