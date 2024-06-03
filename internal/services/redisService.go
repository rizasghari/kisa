package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"kisa/configs"
	"kisa/internal/models"
	"log"
	"time"
)

type RedisService struct {
	client *redis.Client
}

var (
	redisService  = &RedisService{}
	ctx           = context.Background()
	cacheDuration time.Duration
)

func InitializeRedisClient(config *configs.Config) *RedisService {
	host := config.Viper.GetString("redis.host")
	port := config.Viper.GetInt("redis.port")
	address := fmt.Sprintf("%s:%d", host, port)

	password := config.Viper.GetString("redis.password")

	cacheDuration = time.Duration(config.Viper.GetInt("redis.cache")) * time.Second

	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Println("Redis connection error: ", err)
		return nil
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	redisService.client = redisClient

	return redisService
}

func (rs *RedisService) Set(key string, url *models.URL) error {
	value, err := json.Marshal(url)
	if err != nil {
		return err
	}

	err = rs.client.Set(ctx, key, value, cacheDuration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rs *RedisService) Get(key string) (*models.URL, error) {
	result, err := rs.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var url *models.URL
	err = json.Unmarshal([]byte(result), &url)
	if err != nil {
		return nil, err
	}
	return url, err
}

func (rs *RedisService) Check(key string) bool {
	result, err := rs.client.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return result == 1
}

func (rs *RedisService) Delete(key string) error {
	err := rs.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rs *RedisService) Flush() error {
	err := rs.client.FlushAll(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
