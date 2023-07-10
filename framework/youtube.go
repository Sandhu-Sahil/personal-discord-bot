package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"strings"

	"google.golang.org/api/youtube/v3"
)

const (
	ERROR_TYPE    = -1
	VIDEO_TYPE    = 0
	PLAYLIST_TYPE = 1
)

type (
	videoResponse struct {
		Formats []struct {
			Url string `json:"url"`
		} `json:"formats"`
		Title string `json:"title"`
	}

	VideoResult struct {
		Media string
		Title string
	}

	PlaylistVideo struct {
		Id string `json:"id"`
	}

	YTSearchContent struct {
		Id           string `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		ChannelTitle string `json:"channel_title"`
		Duration     string `json:"duration"`
		Thumbnail    *youtube.Thumbnail
	}

	Youtube struct {
		ApiKey string
		Search *YTSearchContent
	}
)

func NewYoutube(apiKey string) *Youtube {
	return &Youtube{
		ApiKey: apiKey,
		Search: &YTSearchContent{},
	}
}

func (y *Youtube) SearchYoutube(client *youtube.Service, query string) error {
	list := []string{"id", "snippet"}

	// Make the API call to YouTube.
	call := client.Search.List(list).Order("relevance").Q(query).MaxResults(1)
	response, err := call.Do()
	if err != nil {
		return err
	}

	var item *youtube.SearchResult
	// Iterate through each item and add it to the correct list.
	for _, search := range response.Items {
		switch search.Id.Kind {
		case "youtube#video":
			if search.Snippet.LiveBroadcastContent == "live" {
				continue
			}
			item = search
		}
	}

	if item == nil {
		// return not found anything
		return errors.New("sorry didn't found anything on youtube")
	}

	y.Search.Id = item.Id.VideoId
	y.Search.Title = item.Snippet.Title
	y.Search.ChannelTitle = item.Snippet.ChannelTitle
	y.Search.Thumbnail = item.Snippet.Thumbnails.High
	y.Search.Description = item.Snippet.Description

	return nil
}

func (y *Youtube) getType(input string) int {
	if strings.Contains(input, "upload_date") {
		return VIDEO_TYPE
	}
	if strings.Contains(input, "_type") {
		return PLAYLIST_TYPE
	}
	return ERROR_TYPE
}

func (y *Youtube) GetFromYT() (int, *string, error) {
	cmd := exec.Command("yt-dlp", "--skip-download", "--print-json", "--flat-playlist", y.Search.Id)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return ERROR_TYPE, nil, err
	}
	str := out.String()
	return y.getType(str), &str, nil
}

func (y *Youtube) Video(input string) (*VideoResult, error) {
	var resp videoResponse
	err := json.Unmarshal([]byte(input), &resp)
	if err != nil {
		return nil, err
	}
	var url string
	// extract the which ha itag 599 and mime type audio/mp4
	// itag 599 is the highest quality audio and audio/mp4 is the highest quality audio format
	// it is the recommended format for ffmpeg
	for _, format := range resp.Formats {
		if strings.Contains(format.Url, "itag=599") && strings.Contains(format.Url, "mime=audio%2Fmp4") {
			url = format.Url
			break
		}
	}
	return &VideoResult{url, resp.Title}, nil
}

func (y Youtube) ExtractPlaylistId(input string) (string, error) {
	// in input find list= and extract the id
	// example input: https://www.youtube.com/watch?v=5qap5aO4i9A&list=PLx0sYbCqOb8TBPRdmBHs5Iftvv9TPboYG
	// output: PLx0sYbCqOb8TBPRdmBHs5Iftvv9TPboYG
	id := strings.Split(input, "list=")
	if len(id) < 2 {
		return "", errors.New("invalid youtube playlist url")
	}
	final_id := strings.Split(id[1], "&")
	return final_id[0], nil
}

func (y Youtube) GetFromYTPlaylist(input string) (string, error) {
	cmd := exec.Command("yt-dlp", "--skip-download", "--print-json", "--flat-playlist", input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	str := out.String()

	return str, nil
}

// func (youtube Youtube) Playlist(input string) (*[]PlaylistVideo, error) {
// 	lines := strings.Split(input, "\n")
// 	videos := make([]PlaylistVideo, 0)
// 	for _, line := range lines {
// 		if len(line) == 0 {
// 			continue
// 		}
// 		var video PlaylistVideo
// 		fmt.Println("line,", line)
// 		err := json.Unmarshal([]byte(line), &video)
// 		if err != nil {
// 			return nil, err
// 		}
// 		videos = append(videos, video)
// 	}
// 	return &videos, nil
// }

/*func (youtube Youtube) OldGet(id string) (*VideoResult, error) {
	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "--flat-playlist", id)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error getting youtube info,", err)
		return nil, err
	}
    fmt.Println(string(out.Bytes()))
	var resp response
	json.Unmarshal(out.Bytes(), &resp)
	u := resp.Formats[0].Url
	return &VideoResult{u, resp.Title}, nil
}*/
