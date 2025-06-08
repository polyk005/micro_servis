package queue

import (
	"context"
	"encoding/json"
	"fmt"
	
	"github.com/go-redis/redis/v8"
	"github.com/polyk005/micro_servis/pkg/models"
)

type RedisQueue struct {
	client *redis.Client
	stream string
}

func NewRedisQueue(addr, stream string) *RedisQueue {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	
	return &RedisQueue{
		client: client,
		stream: stream,
	}
}

func (q *RedisQueue) Publish(ctx context.Context, task models.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}
	
	return q.client.XAdd(ctx, &redis.XAddArgs{
		Stream: q.stream,
		Values: map[string]interface{}{"task": data},
	}).Err()
}

func (q *RedisQueue) Consume(ctx context.Context) (models.Task, error) {
	result, err := q.client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{q.stream, "0"},
		Count:   1,
		Block:   0,
	}).Result()

	if err != nil {
		return models.Task{}, err
	}

	if len(result) == 0 || len(result[0].Messages) == 0 {
		return models.Task{}, nil
	}

	taskData, ok := result[0].Messages[0].Values["task"].(string)
	if !ok {
		return models.Task{}, fmt.Errorf("invalid task data format")
	}

	var task models.Task
	if err := json.Unmarshal([]byte(taskData), &task); err != nil {
		return models.Task{}, err
	}

	return task, nil
}