package redis_repo

import (
	"context"
	"github.com/redis/go-redis/v9"
	"ideadeck/domain/model"
)

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository(client *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{
		client: client,
	}
}

func (r *RedisUserRepository) StartSession(email *model.Email) error {
	token := model.NewUUID().ID()

	if err := r.client.Set(context.Background(), email.Email(), token, 3600).Err(); err != nil {
		return err
	}

	return nil
}

func (r *RedisUserRepository) GetSession(email *model.Email) (string, error) {
	token, err := r.client.Get(context.Background(), email.Email()).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *RedisUserRepository) DeleteSession(email *model.Email) error {
	if err := r.client.Del(context.Background(), email.Email()).Err(); err != nil {
		return err
	}

	return nil
}
