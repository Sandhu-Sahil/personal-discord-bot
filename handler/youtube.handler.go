package handler

import (
	"context"
	"log"
	"sandhu-sahil/bot/variables"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func StartYoutubeClient() {
	// client := &http.Client{
	// 	Transport: &transport.APIKey{Key: variables.YoutubeApiKey},
	// }

	var err error
	var ctx context.Context
	variables.YoutubeService, err = youtube.NewService(ctx, option.WithAPIKey(variables.YoutubeApiKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
}
