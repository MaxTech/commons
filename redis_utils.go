package commons

import (
    "fmt"
    "github.com/go-redis/redis"
    "os"
)

func InitRedisClient(_address, _password string, _dbNum int) *redis.Client {
    redisClient := redis.NewClient(&redis.Options{
        Addr:     _address,
        Password: _password, // no password set
        DB:       _dbNum,    // use default DB
    })
    return redisClient
}

func TestRedisClient(_redisClient *redis.Client) bool {
    _, err := _redisClient.Ping().Result()
    if err != nil {
        _, _ = fmt.Fprintln(os.Stderr, "redis connect error: ", err)
        return false
    }
    return true
}
