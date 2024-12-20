package main

// AI represents different AI models and their specific requirements
type AI struct {
	Name           string
	SystemPrompt   string
	MaxTokens      int
	TemperatureMin float64
	TemperatureMax float64
}

// Prompt represents a generated prompt with its configuration
type Prompt struct {
	AI          *AI
	Content     string
	Temperature float64
	MaxTokens   int
}

// PromptGenerator handles the generation of prompts for different AI models
type PromptGenerator interface {
	GeneratePrompt(content string, aiModel string) (*Prompt, error)
	GetSupportedAIs() []string
}
