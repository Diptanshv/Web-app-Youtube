package routes

import (
	"net/http"
	"youtube-dashboard/handlers"
)

func SubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	handlers.GetSubscriptions(w, r)
}

func LikedVideosHandler(w http.ResponseWriter, r *http.Request) {
	handlers.GetLikedVideos(w, r)
}

func SearchVideosHandler(w http.ResponseWriter, r *http.Request) {
	handlers.SearchVideosHandler(w, r)
}

func GetChannelDetailsHandler(w http.ResponseWriter, r *http.Request) {
	handlers.GetChannelDetailsHandler(w, r)
}
