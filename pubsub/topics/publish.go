package main

import (
	"context"
	"fmt"

	"log"

	"cloud.google.com/go/pubsub"
)

func main() {

	//msg := "Hello World"
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "ayopop-atul"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the id for the new topic.

	topicID := "transactionFailed"

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)
}
