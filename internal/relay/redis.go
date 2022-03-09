package relay

import (
	"context"
	"fmt"
	"time"

	"github.com/DaivikDave/walletconnect-server/internal/relay/jsonrpc"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
	prefix string
}

type RedisConfig struct {
	URL    string
	Prefix string
}

func NewRedis(config RedisConfig) Redis {
	redis := Redis{
		Client: redis.NewClient(&redis.Options{
			Addr: config.URL,
		}),
	}
	redis.setPrefix(config.Prefix)
	return redis
}

func (r *Redis) SetMessage(params jsonrpc.PublishParams) error {
	key := fmt.Sprintf("message:%s", params.Topic)
	hash := Sha256(params.Message)
	val := fmt.Sprintf("%s:%s", hash, params.Message)
	ctx := context.Background()

	if err := r.Client.SAdd(ctx, r.addPrefix(key), val).Err(); err != nil {
		return err
	}

	ttl := time.Duration(int64(params.Ttl))
	if err := r.Client.Expire(ctx, r.addPrefix(key), ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Redis) GetMessage(topic string, hash string) error {

}

func (r *Redis) GetMessages(topic string) ([]string, error) {

}

func (r *Redis) DeleteMessage(topic string, hash string) error {

}

func (r *Redis) SetLegacyCached(message LegacySocketMessage) error {

}

func (r *Redis) GetLegacyCached(topic string) ([]LegacySocketMessage, error) {

}

func (r *Redis) SetNotification(notification Notification) error {

}

func (r *Redis) GetNotification(topic string) ([]Notification, error) {

}

func (r *Redis) SetPendingRequest(topic string, id int, message string) error {

}

func (r *Redis) GetPendingRequest(id int) (string, error) {

}

func (r *Redis) DeletePendingRequest(id int) error {

}

func (r *Redis) addPrefix(key string) string {
	return r.prefix + key
}

func (r *Redis) setPrefix(prefix string) {
	r.prefix = prefix + ":"
}
