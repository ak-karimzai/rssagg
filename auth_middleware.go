package main

import (
	"fmt"
	"net/http"

	"github.com/ak-karimzai/rssagg/internal/auth"
	"github.com/ak-karimzai/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apicfg.DB.GetUserByApikey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprint("coludn't found user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
