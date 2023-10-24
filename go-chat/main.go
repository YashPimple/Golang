package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("Missing API-KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apikey)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"The first things you should know about golang "},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)
}
