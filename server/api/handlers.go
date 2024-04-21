package api

import (
	"encoding/json"
	"net/http"
	"tournament_api/server/types"

	"github.com/go-playground/validator/v10"
)

func (s *Server) handleGetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("siema")
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	//find if user exsits
	//create a new user in db

	var user types.User
	errDecode := json.NewDecoder(r.Body).Decode(&user)
	if errDecode != nil {
		http.Error(w, "Error while decoding: ", http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	errValidation := validate.Struct(user)
	if errValidation != nil {
		http.Error(w, "Validation failed: "+errValidation.Error(), http.StatusBadRequest)
		return
	}

	accessToken := newAccessToken()
	refreshToken := newRefreshToken()

	accessTokenString, errAccess := accessToken.generateTokenString(user.FirstName, user.LastName)
	refreshTokenString, errRefresh := refreshToken.generateTokenString(user.FirstName, user.LastName)

	if errAccess != nil || errRefresh != nil {
		http.Error(w, "Error generating token string", http.StatusInternalServerError)
		return
	}

	accessToken.saveTokenAsCookie(w, accessTokenString)
	refreshToken.saveTokenAsCookie(w, refreshTokenString)

	json.NewEncoder(w).Encode(map[string]string{"message": "login successful"})
}

func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {}