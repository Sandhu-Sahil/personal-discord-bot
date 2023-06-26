package handler

import (
	"log"
	"net/http"
	"sandhu-sahil/bot/variables"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func StartYoutubeClient() {
	client := &http.Client{
		Transport: &transport.APIKey{Key: variables.YoutubeApiKey},
	}

	var err error
	variables.YoutubeService, err = youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
}
