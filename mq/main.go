package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

//RECIEVER

//left arrow means channel can only be read from
//excluding arrow means channel read from and written to
func listen(msgs <-chan amqp.Delivery) {
	log.Println("listening for new messages...")
	for msg := range msgs {
		log.Println(string(msg.Body))
	}
}

func main() {
	mqAddr := os.Getenv("MQADDR")
	if len(mqAddr) == 0 {
		mqAddr = "192.168.99.100:5672"
	}

	//connect to RabbitMQ
	mqURL := fmt.Sprintf("amqp://%s", mqAddr)
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		log.Fatalf("error connecting to RabbitMQ: %v", err)
	}

	//connect to a channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("error creating channel: %v", err)
	}

	//declare (initialize) a queue
	q, err := channel.QueueDeclare("testQ", false, false, false, false, nil)

	//listen for messages
	msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)
	go listen(msgs)

	//ensures main doesn't exit until explicitly stopped
	neverEnd := make(chan bool)
	<-neverEnd
}
