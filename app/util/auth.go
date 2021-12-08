package util

import (
	"time"

	"github.com/Code0716/go-vtm/app/domain"
	jwt "github.com/dgrijalva/jwt-go"
)

var iss = "vtm"

// GetAdminNewToken is get new token for admins
func GetAdminNewToken(member domain.AdminUser, signingKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = member.AdminId
	claims["name"] = member.Name
	claims["iss"] = iss
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	// 電子署名
	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return "", domain.WrapInternalError(err)
	}
	return tokenString, nil
}

// GetMemberNewToken is get new token for members
func GetMemberNewToken(member domain.Member, signingKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = member.MemberId
	claims["name"] = member.Name
	claims["iss"] = iss
	claims["phone_number"] = member.PhoneNumber
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", domain.WrapInternalError(err)
	}
	return tokenString, nil
}
