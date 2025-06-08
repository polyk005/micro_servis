package main

import (
	"context"
	"log"

	"github.com/spf13/viper"
)

func main() {
	q := queue.NewRedisQueue(
		viper.Getstring("redis.addr"),
		viper.Getstring("redis.stream"),
	)

	ctx := context.Background()
	for {
		task, err := q.Consume(ctx)
		if err != nil {
			log.Printf("Error consuming task: %v", err)
			continue
		}
		go processTask(task)
	}
}

func processTask(task models.Task) {
	switch task.Type {
	case "process_image":
		processImage(task.Data)
	case "process_video":
		processVideo(task.Data)
	}
}
