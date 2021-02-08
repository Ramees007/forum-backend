package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/rameesThattarath/gooo/app/handler"
)

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

func getLimitAndOffset(v url.Values) (int, int, error) {
	limit := v.Get("limit")
	offset := v.Get("offset")
	var l, o int

	if limit != "" {
		lim, err := strconv.ParseUint(limit, 10, 64)
		if err != nil {
			return l, o, err
		}
		l = int(lim)
	}

	if offset != "" {
		off, err := strconv.ParseUint(offset, 10, 64)
		if err != nil {
			return l, o, err
		}
		o = int(off)
	}

	if l == 0 {
		l = 100
	}
	return l, o, nil
}

func isAuthorized(endPoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		tknStr := r.Header.Get("token")
		if tknStr == "" {
			respondError(w, http.StatusUnauthorized, "Not Authorized")
			return
		}

		claims := &Claims{}

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

func getIDOfAuthorizedUser(r *http.Request) uint {
	tknStr := r.Header.Get("token")

	claims := &Claims{}

	jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return handler.JwtKey, nil
	})

	return claims.UserId

}
