package utils

import (
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	JwtKey        = "HASH256123456"
	JwtExpireTime = 7 * 24 * time.Hour
)

type Claim struct {
	jwt.StandardClaims
	Phone    int64
	Nickname string
	UserID   uint
}

func GenerateToken(Phone int64, Nickname string, UserID uint) (string, error) {
	claim := Claim{
		Phone:    Phone,
		Nickname: Nickname,
		UserID:   UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JwtExpireTime).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(JwtKey))
	if err != nil {
		logrus.Errorf("[util.GenerateJWT] %v", err)
		return "", err
	}

	return token, err
}

func ParseToken(token string) (*Claim, error) {
	TokenClaim, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
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

	logrus.Debugf("[util.ParseToken] %v,\n %v", err, TokenClaim)
	return nil, err
}

func Hash256(src string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(src)))
}
