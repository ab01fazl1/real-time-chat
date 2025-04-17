package user

import (
	r "github.com/ab01fazl1/real-time-chat/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client = r.GetDb()

func get_one_user(c *gin.Context, username string) (string, error) {
	val, err := rdb.Get(c, username).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

func createUser(c *gin.Context, username string, password string) error {
	err := rdb.Set(c, username, password, 0).Err()
	return err
}