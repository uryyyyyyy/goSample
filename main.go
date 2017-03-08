package main

import (
	"fmt"
	"log"
	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// Sets your Google Cloud Platform project ID.
const projectID string = "gaesample-160801"

func CreateTopic(ctx context.Context ) {
	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new topic.
	topicName := "my-new-topic"

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)
}

func SendMessage(ctx context.Context ) {
	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	topicName := "my-new-topic"

	// Creates the new topic.
	topic := client.Topic(topicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	msgIDs, err := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("payload"),
	})
	fmt.Printf("Topic %v created.\n", msgIDs)
}

func CreateSubScription(ctx context.Context ) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	topicName := "my-new-topic"

	// Creates the new topic.
	topic := client.Topic(topicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}
	subName := "sub-name"

	sub, err := client.CreateSubscription(context.Background(), subName, topic, 0, nil)
	fmt.Printf("Topic %v created.\n", sub)
}

func Subscribe(ctx context.Context ) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	subName := "sub-name"

	sub := client.Subscription(subName)

	// Construct the iterator
	it, err := sub.Pull(context.Background())
	if err != nil {
		// handle err ...
	}
	defer it.Stop()

	N := 10

	// Consume N messages
	for i := 0; i < N; i++ {
		msg, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// handle err ...
			break
		}

		log.Print("got message: ", string(msg.Data))
		msg.Done(true)
	}
}

func main() {
	ctx := context.Background()
	Subscribe(ctx)
}