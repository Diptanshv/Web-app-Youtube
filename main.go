package main

import (
	"fmt"
	"net/http"
	"youtube-dashboard/configs"
	"youtube-dashboard/routes"
)

func main() {
	configs.LoadConfig()

	http.HandleFunc("/api/subscriptions", routes.SubscriptionsHandler)
	http.HandleFunc("/api/likedvideos", routes.LikedVideosHandler)
	http.HandleFunc("/api/searchvideos", routes.SearchVideosHandler)
	http.HandleFunc("/api/channeldetails", routes.GetChannelDetailsHandler)
	// http.HandleFunc("/api/search", routes.SearchVideosHandler)

	fs := http.FileServer(http.Dir("./frontend/build"))
	http.Handle("/", fs)
	fmt.Println("Listening to port 8080:")
	http.ListenAndServe(":8080", nil)
}
