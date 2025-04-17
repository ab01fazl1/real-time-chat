package user

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func userExists(c *gin.Context, username string) (bool, error){
	_, err := get_one_user(c, username)
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

