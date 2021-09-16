package domain

import (
	"time"

	"github.com/Code0716/go-vtm/app/gen/api"
	jwt "github.com/dgrijalva/jwt-go"
)

// AuthenticationResponse  model
type AuthenticationResponse api.AuthenticationResponse

// JwtCustomClaims is model of Claims
type JwtCustomClaims struct {
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Iat         time.Time `json:"iat"` // time.Time型にしたいので上書き
	jwt.StandardClaims
}
