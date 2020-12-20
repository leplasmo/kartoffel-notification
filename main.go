// kartoffel-notification/main.go

package main

import (
	"encoding/json"
	"log"

	pb "github.com/leplasmo/kartoffel-user/proto/user"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
)

const topic = "user.created"

func main() {

	svc := micro.NewService(
		micro.Name("kartoffel.notification"),
		micro.Version("v0.1.0"),
	)

	svc.Init()

	// Get the broker instance using our environment variables
	pubsub := svc.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe to messages on the broker
	_, err := pubsub.Subscribe(topic, func(p broker.Event) error {
		var user *pb.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			return err
		}
		log.Println(user)
		go sendEmail(user)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	// Run the server
	if err := svc.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error {
	// stub function - only prints to console for now
	log.Println("Sending email to:", user.Name)
	return nil
}
