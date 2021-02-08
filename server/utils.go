package server

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/rameesThattarath/qaForum/handler"
)

func isAuthorized(endPoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		tknStr := r.Header.Get("token")
		if tknStr == "" {
			respondError(w, http.StatusUnauthorized, "Not Authorized")
			return
		}

		claims := &handler.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return handler.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				respondError(w, http.StatusUnauthorized, "Not Authorized")
				return
			}
			respondError(w, http.StatusUnauthorized, "Not Authorized")
			return
		}
		if !tkn.Valid {
			respondError(w, http.StatusUnauthorized, "Not Authorized")
			return
		}
		endPoint(w, r)
	}

}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
