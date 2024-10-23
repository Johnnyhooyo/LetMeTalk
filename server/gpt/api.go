package gpt

import (
	"context"
	"github.com.johnyooho.lmt/server/global"
	"github.com/sashabaranov/go-openai"
)

func Chat(ctx context.Context) {
	client := openai.NewClient(global.ProjectConfig.ChatAIKey)

	openai.ImageRequest{
		Prompt:         "",
		Model:          openai.CreateImageModelDallE3,
		N:              0,
		Quality:        "",
		Size:           "",
		Style:          "",
		ResponseFormat: "",
		User:           "",
	}
}
