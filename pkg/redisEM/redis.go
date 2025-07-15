package redisEM

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"testAPI/internal/entity"
	"time"
)

type NotificaitonCache struct {
	client *redis.Client
}

func NewNotificaitonCache(addr string) *NotificaitonCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
	return &NotificaitonCache{client: rdb}
}

func (nc *NotificaitonCache) Set(ctx context.Context, uuid string, notif *entity.NotificationEntity, ttl time.Duration) error {
	data, _ := json.Marshal(notif)
	return nc.client.Set(ctx, uuid, data, ttl).Err()
}

func (nc *NotificaitonCache) Get(ctx context.Context, uuid string) (*entity.NotificationEntity, error) {
	val, err := nc.client.Get(ctx, uuid).Result()
	if err != nil {
		return nil, err
	}
	var notif entity.NotificationEntity
	json.Unmarshal([]byte(val), &notif)
	return &notif, nil
}

func (nc *NotificaitonCache) Delete(ctx context.Context, uuid string) error {
	return nc.client.Del(ctx, uuid).Err()
}
