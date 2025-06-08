package queue

import (
	"github.com/go-redis/redis"
)

type RedisQueue struct {
	client *redis.Client
	stream string
}

func NewRedisQueue(addr, stream string) *RedisQueue {
	return &RedisQueue{
		client: redis.NewClient(&redis.Options{Addr: addr}),
		stream: stream,
}
}

func (q *RedisQueue) Publish(ctx context.Context, task models.Task) error {
	data, _ := json.Marshal(task)
	return q.client.XAdd(ctx, &redis.XAddArgs{
		Stream: q.stream,
		Values: map[string]interface{}{"task": data},
	}).Err()
}

func (q *RedisQueue) Consume(ctx context.Context) (models.Task, error) {
	result, err := q.client.Xread(ctx, &redis.XReadArgs{
		Streams: []string{q.stream, "0"},
		Block: 0,
	}).Result()

	if err != nil {
		return models.Task{}, err
	}
}