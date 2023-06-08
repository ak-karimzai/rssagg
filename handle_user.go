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

func (apicfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	fmt.Println(params)
	err := decoder.Decode(&params)
	if err != nil || params.Name == "" {
		respondWithError(w, 400, fmt.Sprint("error parsing json: ", err))
		return
	}

	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("couldn't create user: ", err))
		return
	}
	respondWithJSON(w, 201, models.DatabaseUserToUser(user))
}

func (apicfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (apicfg *apiConfig) handleGetPostsForUser(
	w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apicfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("coludn't get posts %v", err))
		return
	}
	respondWithJSON(w, 200, models.DatabasePostsToPosts(posts))
}
