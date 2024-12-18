package services

import (
	"context"

	"github.com/syntaxsdev/mercury/internal/repositories"
)

type RedisService struct {
	Repo *repositories.RedisClient
}

func (s *RedisService) Set(ctx context.Context, key string, value interface{}) error {
	return s.Repo.Client.Set(ctx, key, value, 0).Err()
}

func (s *RedisService) Get(ctx context.Context, key string) (string, error) {
	return s.Repo.Client.Get(ctx, key).Result()
}

func (s *RedisService) Add(ctx context.Context, key string, value interface{}) error {
	return s.Repo.Client.RPush(ctx, key, value, 0).Err()
}

func (s *RedisService) GetList(ctx context.Context, key string) ([]string, error) {
	return s.Repo.Client.LRange(ctx, key, 0, -1).Result()
}

func (s *RedisService) Pop(ctx context.Context, key string) (string, error) {
	return s.Repo.Client.RPop(ctx, key).Result()
}
