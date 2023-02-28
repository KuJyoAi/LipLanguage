package dao

import (
	"fmt"
	"time"
)

func DeleteRedisToken(Phone int64) error {
	key := fmt.Sprintf("%v_Token", Phone)
	return RDB.Del(key).Err()
}

func SetRedisToken(key string, token string, duration time.Duration) {
	RDB.Set(key, token, duration)
}
