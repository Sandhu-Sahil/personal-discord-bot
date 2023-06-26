package cmd

import (
	"sandhu-sahil/bot/framework"

	"google.golang.org/api/youtube/v3"
)

func YoutubeCommandPreIntractions(client *youtube.Service, query string) map[string]string {
	list := []string{"id", "snippet"}
	videos := make(map[string]string)

	// Make the API call to YouTube.
	call := client.Search.List(list).Q(query).MaxResults(2)
	response, err := call.Do()
	if err != nil {
		panic(err)
	}

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			if item.Snippet.LiveBroadcastContent == "live" {
				continue
			}
			// check length of title
			if stringLength := len(item.Snippet.Title); stringLength > 100 {
				item.Snippet.Title = item.Snippet.Title[:95]
			}
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}
	return videos
}

func YoutubeCommandIntractions(ctx *framework.Context, id string) string {
	return "https://www.youtube.com/watch?v=" + id
}
