package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"youtube-dashboard/configs"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Video struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	ViewCount uint64 `json:"viewCount,omitempty"`
	LikeCount uint64 `json:"likeCount,omitempty"`
}

type Subscription struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
}

// getClient uses a Context and Config to retrieve a Token then generate a Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
func tokenCacheFile() (string, error) {
	tokenCacheDir := filepath.Join(".", ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir, url.QueryEscape("credd.json")), nil
}

// tokenFromFile retrieves a Token from a given file path.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	return t, err
}

// saveToken uses a file path to create a file and store the token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func FetchSubscriptions(ctx context.Context) ([]Subscription, error) {
	client := getClient(ctx, configs.Config)
	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("error creating YouTube client: %v", err)
	}

	var subscriptions []Subscription
	pageToken := ""
	for {
		call := service.Subscriptions.List([]string{"snippet"}).Mine(true)
		if pageToken != "" {
			call = call.PageToken(pageToken)
		}

		response, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("error making API call: %v", err)
		}

		for _, item := range response.Items {
			subscriptions = append(subscriptions, Subscription{
				Title:     item.Snippet.Title,
				Thumbnail: item.Snippet.Thumbnails.Default.Url,
				URL:       fmt.Sprintf("https://www.youtube.com/channel/%s", item.Snippet.ResourceId.ChannelId),
			})
			//fmt.Println(item.Snippet.ResourceId.ChannelId)
		}

		//log.Printf("Fetched %d subscriptions, total so far: %d", len(response.Items), len(subscriptions))

		if response.NextPageToken == "" {
			break
		}
		pageToken = response.NextPageToken
	}

	//log.Printf("Total subscriptions fetched: %d", len(subscriptions))
	return subscriptions, nil
}

func FetchLikedVideos(ctx context.Context) []Video {
	client := getClient(ctx, configs.Config)
	// if err != nil {
	// 	log.Fatalf("Unable to create client: %v", err)
	// }
	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	call := service.Videos.List([]string{"snippet", "contentDetails", "statistics"}).MyRating("like").MaxResults(500000)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Unable to retrieve liked videos: %v", err)
	}

	var videos []Video
	for _, item := range response.Items {
		videos = append(videos, Video{
			Title:     item.Snippet.Title,
			Thumbnail: item.Snippet.Thumbnails.Default.Url,
			ViewCount: item.Statistics.ViewCount,
			LikeCount: item.Statistics.LikeCount,
		})
	}
	fmt.Println(len(videos))

	return videos
}

// func SearchVideos(query string, ctx context.Context) []Video {
// 	client := getClient(ctx, configs.Config)
// 	// if err != nil {
// 	// 	log.Fatalf("Unable to create client: %v", err)
// 	// }
// 	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
// 	if err != nil {
// 		log.Fatalf("Unable to create YouTube service: %v", err)
// 	}

// 	call := service.Search.List([]string{"snippet"}).Q(query).Type("video").MaxResults(10)
// 	response, err := call.Do()
// 	if err != nil {
// 		log.Fatalf("Unable to search videos: %v", err)
// 	}

// 	var videos []Video
// 	for _, item := range response.Items {
// 		videos = append(videos, Video{
// 			Title:     item.Snippet.Title,
// 			Thumbnail: item.Snippet.Thumbnails.Default.Url,
// 		})
// 	}

// 	return videos
// }

func SearchVideos(query string, ctx context.Context) []*youtube.SearchResult {
	client := getClient(ctx, configs.Config)
	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	call := service.Search.List([]string{"id", "snippet"}).Q(query).MaxResults(25)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}

	return response.Items
}

func GetChannelDetails(channelID string, ctx context.Context) *youtube.Channel {
	client := getClient(ctx, configs.Config)
	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	call := service.Channels.List([]string{"snippet", "statistics"}).Id(channelID)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}

	if len(response.Items) == 0 {
		return nil
	}
	return response.Items[0]
}
