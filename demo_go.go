/*
Dependency install:
go get github.com/sashabaranov/go-openai

Run (PowerShell):
$env:OPENAI_API_KEY="your_key"
go run demo/demo_go.go
*/

package main

import (
    "context"
    "errors"
    "fmt"
    "io"
    "log"
    "os"

    openai "github.com/sashabaranov/go-openai"
)

func newClient() *openai.Client {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        apiKey = "YOUR API KEY"
    }

    config := openai.DefaultConfig(apiKey)
    config.BaseURL = "https://api.chatanywhere.tech/v1"
    return openai.NewClientWithConfig(config)
}

// Non-stream response
func gpt35API(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage) error {
    req := openai.ChatCompletionRequest{
        Model:    openai.GPT3Dot5Turbo,
        Messages: messages,
    }

    resp, err := client.CreateChatCompletion(ctx, req)
    if err != nil {
        return err
    }

    if len(resp.Choices) > 0 {
        fmt.Println(resp.Choices[0].Message.Content)
    }
    return nil
}

// Stream response
func gpt35APIStream(ctx context.Context, client *openai.Client, messages []openai.ChatCompletionMessage) error {
    req := openai.ChatCompletionRequest{
        Model:    openai.GPT3Dot5Turbo,
        Messages: messages,
        Stream:   true,
    }

    stream, err := client.CreateChatCompletionStream(ctx, req)
    if err != nil {
        return err
    }
    defer stream.Close()

    for {
        response, err := stream.Recv()
        if errors.Is(err, io.EOF) {
            break
        }
        if err != nil {
            return err
        }

        if len(response.Choices) > 0 {
            fmt.Print(response.Choices[0].Delta.Content)
        }
    }

    fmt.Println()
    return nil
}

func main() {
    client := newClient()
    ctx := context.Background()

    messages := []openai.ChatCompletionMessage{
        {
            Role:    openai.ChatMessageRoleUser,
            Content: "What is the relationship between Lu Xun and Zhou Shuren?",
        },
    }

    // Non-stream call
    // if err := gpt35API(ctx, client, messages); err != nil {
    //     log.Fatal(err)
    // }

    // Stream call
    if err := gpt35APIStream(ctx, client, messages); err != nil {
        log.Fatal(err)
    }
}
