package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/replicate/replicate-go"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	r8, err := replicate.NewClient(replicate.WithTokenFromEnv())
	if err != nil {
		log.Fatal(err)
	}
	version := "527d2a6296facb8e47ba1eaf17f142c240c19a30894f437feee9b91cc29d8e4f"
	input := replicate.PredictionInput{
		"prompt": "An astronaut riding a rainbow unicorn",
	}

	webhook := replicate.Webhook{
		URL:    "https://webhook.site/9579c641-42cf-40d9-b91d-ceaa5901856e",
		Events: []replicate.WebhookEventType{"start", "completed"},
	}
	output, err := r8.CreatePrediction(ctx, version, input, &webhook, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Output: ", output)

	// prediction, err := r8.CreatePrediction(ctx, version, input, &webhook, false)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
