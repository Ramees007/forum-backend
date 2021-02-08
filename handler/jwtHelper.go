package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JwtKey - JwtKey
var JwtKey = []byte("whyIstherainblackthatshines")

//Credentials - Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
}

//Claims -  Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func issueJwt(creds Credentials) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		UserId: creds.UserId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
