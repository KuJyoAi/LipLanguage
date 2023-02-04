package util

import (
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	ExpireTime = 3 * 24 * time.Hour
	Key        = "HASH256123456"
)

type Claim struct {
	jwt.StandardClaims
	ID       uint
	Nickname string
}

func GenerateToken(ID uint, Nickname string) (string, error) {
	claim := Claim{
		ID:       ID,
		Nickname: Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireTime).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(Key))
	if err != nil {
		logrus.Errorf("[util.GenerateJWT] %v", err)
		return "", err
	}
	return token, err
}

func ParseToken(token string) (*Claim, error) {
	TokenClaim, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
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
