package redisdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient adalah instance global untuk koneksi Redis
var RedisClient *redis.Client

// InitRedis menginisialisasi koneksi Redis
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",               
		DB:       0,                
	})

	// Cek koneksi
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Redis connected!")
}

// SetKey menyimpan data ke Redis dengan expiry
func SetKey(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	err := RedisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetKey mengambil data dari Redis berdasarkan key
func GetKey(key string) (string, error) {
	ctx := context.Background()
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key not found")
		}
		return "", err
	}
	return val, nil
}
