package main

import (
	"errors"
	"fmt"
)

type DefaultPromptGenerator struct {
	models map[string]*AI
}

func NewPromptGenerator() *DefaultPromptGenerator {
	return &DefaultPromptGenerator{
		models: map[string]*AI{
			"claude": {
				Name:           "claude",
				SystemPrompt:   "You are Claude, an AI assistant created by Anthropic.",
				MaxTokens:      100000,
				TemperatureMin: 0.0,
				TemperatureMax: 1.0,
			},
			"gpt4": {
				Name:           "gpt4",
				SystemPrompt:   "You are GPT-4, an AI assistant created by OpenAI.",
				MaxTokens:      8000,
				TemperatureMin: 0.0,
				TemperatureMax: 2.0,
			},
			// Add more AI models as needed
		},
	}
}

func (g *DefaultPromptGenerator) GeneratePrompt(content string, aiModel string) (*Prompt, error) {
	model, exists := g.models[aiModel]
	if !exists {
		return nil, errors.New(fmt.Sprintf("unsupported AI model: %s", aiModel))
	}

	// Create a new prompt with default settings
	return &Prompt{
		AI:          model,
		Content:     content,
		Temperature: model.TemperatureMin + (model.TemperatureMax-model.TemperatureMin)/2, // Default to middle temperature
		MaxTokens:   model.MaxTokens,
	}, nil
}

func (g *DefaultPromptGenerator) GetSupportedAIs() []string {
	models := make([]string, 0, len(g.models))
	for model := range g.models {
		models = append(models, model)
	}
	return models
}
