package cmd

import (
	"errors"
	"sandhu-sahil/bot/framework"

	"google.golang.org/api/youtube/v3"
)

func YoutubeCommandPreIntractions(client *youtube.Service, query string) (string, error) {
	list := []string{"id", "snippet"}

	// Make the API call to YouTube.
	call := client.Search.List(list).Order("relevance").Q(query).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		return "", err
	}

	var id string
	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			id = item.Id.VideoId
		}
	}

	if id == "" {
		// return not found anything
		return "", errors.New("sorry didn't found anything on youtube")
	}

	return id, nil
}

func YoutubeCommandIntractions(ctx *framework.Context, id string) string {
	return "https://www.youtube.com/watch?v=" + id
}
