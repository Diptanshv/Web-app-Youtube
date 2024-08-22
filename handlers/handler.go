package handlers

import (
	"net/http"
	"youtube-dashboard/models"
	"youtube-dashboard/utils"
)

func GetSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscriptions, err := models.FetchSubscriptions(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, subscriptions)
}

func GetLikedVideos(w http.ResponseWriter, r *http.Request) {
	likedVideos := models.FetchLikedVideos(r.Context())
	utils.RespondWithJSON(w, http.StatusOK, likedVideos)
}

func SearchVideosHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	videos := models.SearchVideos(query, r.Context())
	utils.RespondWithJSON(w, http.StatusOK, videos)
}

func GetChannelDetailsHandler(w http.ResponseWriter, r *http.Request) {
	channelID := r.URL.Query().Get("channelId")
	channelDetails := models.GetChannelDetails(channelID, r.Context())
	utils.RespondWithJSON(w, http.StatusOK, channelDetails)
}
