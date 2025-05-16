package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// LoginHandler godoc
//
//	@Summary		User login
//	@Description	Authenticates user with login and password from environment variables. Returns JWT token on success.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			login	body		LoginRequest	true	"Login credentials"
//	@Success		200		{object}	LoginResponse
//	@Failure		400		{string}	string	"Invalid request"
//	@Failure		401		{string}	string	"Unauthorized"
//	@Failure		500		{string}	string	"Could not generate token"
//	@Router			/login [get]
func LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	envLogin := os.Getenv("LOGIN")
	envPassword := os.Getenv("PASSWORD")
	jwtSecret := os.Getenv("JWT_SECRET")

	if req.Login != envLogin || req.Password != envPassword {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": req.Login,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	resp := LoginResponse{Token: tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
