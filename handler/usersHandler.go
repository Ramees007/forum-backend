package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rameesThattarath/qaForum/entity"
	"github.com/rameesThattarath/qaForum/usecase"
)

//UserHandler - UserHandler
type UserHandler struct {
	RegisterUC usecase.RegisterUser
	LoginUC    usecase.LoginUser
	ProfileUC  usecase.GetProfile
}

//HandleSignup - HandleSignup
func (h *UserHandler) HandleSignup() http.HandlerFunc {

	type res struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		user := entity.User{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}

		validity := user.IsValidForReg()
		if validity != "" {
			respondError(w, http.StatusBadRequest, validity)
			return
		}

		userID, err := h.RegisterUC.RegisterUser(user)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		token, err := issueJwt(Credentials{
			UserId:   userID,
			Username: user.Email,
		})

		response := res{
			token,
		}

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusCreated, response)
	}

}

//HandleLogin - HandleLogin
func (h *UserHandler) HandleLogin() http.HandlerFunc {
	type res struct {
		Token string `json:"token"`
		Email string `user:"email"`
		Name  string `user:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		user := entity.User{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}

		validity := user.IsValidForLogin()
		if validity != "" {
			respondError(w, http.StatusBadRequest, validity)
			return
		}

		id, usrOut, err := h.LoginUC.LoginUser(user)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		token, err := issueJwt(Credentials{
			UserId:   id,
			Username: user.Email,
		})

		response := res{
			token,
			usrOut.Email,
			usrOut.Name,
		}

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondJSON(w, http.StatusCreated, response)
	}
}

//HandleProfile - HandleProfile
func (h *UserHandler) HandleProfile() http.HandlerFunc {
	type res struct {
		Email string `user:"email"`
		Name  string `user:"name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := getIDOfAuthorizedUser(r)

		u, err := h.ProfileUC.GetProfile(id)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusCreated, res{
			u.Email, u.Name,
		})

	}

}
