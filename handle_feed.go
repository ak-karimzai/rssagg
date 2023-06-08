package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ak-karimzai/rssagg/internal/database"
	"github.com/ak-karimzai/rssagg/models"
	"github.com/google/uuid"
)

func (apicfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	fmt.Println(params)
	err := decoder.Decode(&params)
	if err != nil || params.Name == "" {
		respondWithError(w, 400, fmt.Sprint("error parsing json: ", err))
		return
	}

	feed, err := apicfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("couldn't create feed: ", err))
		return
	}
	respondWithJSON(w, 201, models.DatabaseFeedToFeed(feed))
}

func (apicfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apicfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("coludn't fetch feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, models.DatabaseFeedsToFeeds(feeds))
}
