package redisstore

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewClient(addr string, password string) (*RedisClient, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	resp, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to ping redis client: %s", err.Error())
		return nil, err
	}
	if resp != "PONG" {
		fmt.Printf("Unexpected response to PING request: %s", resp)
		return nil, err
	}

	return &RedisClient{
		client: redisClient,
	}, nil
}

func (r *RedisClient) AddClient(ctx context.Context, username string) error {
	key := fmt.Sprintf("user:%s", username)
	id := uuid.New().String()
	if err := r.client.Set(ctx, key, id, 0).Err(); err != nil {
		fmt.Printf("Failed to save user to redis: %s", err.Error())
		return err
	}
	if err := r.client.RPush(ctx, "users", username).Err(); err != nil {
		fmt.Printf("Failed to save user to redis: %s", err.Error())
		return err
	}
	return nil
}

func (r *RedisClient) RemoveClient(ctx context.Context, username string) error {
	key := fmt.Sprintf("user:%s", username)
	if err := r.client.Del(ctx, key).Err(); err != nil {
		fmt.Printf("Failed to delete user from redis: %s", err.Error())
		return err
	}
	if err := r.client.LRem(ctx, "users", 1, username).Err(); err != nil {
		fmt.Printf("Failed to delete user from redis: %s", err.Error())
		return err
	}
	return nil
}

func (r *RedisClient) UsernameExist(ctx context.Context, username string) (bool, error) {
	key := fmt.Sprintf("user:%s", username)
	_, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		fmt.Printf("Failed to get user from redis: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (r *RedisClient) GetUsers(ctx context.Context) ([]string, error) {
	usernames, err := r.client.LRange(ctx, "users", 0, -1).Result()
	if err != nil {
		fmt.Printf("Failed to get users from redis: %s", err.Error())
		return nil, err
	}
	return usernames, err
}
