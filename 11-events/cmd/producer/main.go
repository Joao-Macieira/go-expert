package main

import (
	"events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World from application!")
}
