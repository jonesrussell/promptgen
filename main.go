package main

import (
	"fmt"
	"log"
)

func main() {
	generator := NewPromptGenerator()

	// Print supported AI models
	fmt.Println("Supported AI models:", generator.GetSupportedAIs())

	// Generate a test prompt
	prompt, err := generator.GeneratePrompt("Tell me a joke", "claude")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated prompt for %s:\n", prompt.AI.Name)
	fmt.Printf("System prompt: %s\n", prompt.AI.SystemPrompt)
	fmt.Printf("Content: %s\n", prompt.Content)
	fmt.Printf("Temperature: %.2f\n", prompt.Temperature)
	fmt.Printf("Max tokens: %d\n", prompt.MaxTokens)
}
