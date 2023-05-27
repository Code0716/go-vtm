// Package util is utility functions
package util

// import (
// 	"time"

// 	"github.com/Code0716/go-vtm/app/domain"
// 	"github.com/golang-jwt/jwt"
// )

// var iss = "vtm"

// // GetAdminNewToken is get new token for admins
// func GetAdminNewToken(user domain.AdminUser, signingKey string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["sub"] = user.AdminId
// 	claims["name"] = user.Name
// 	claims["iss"] = iss
// 	claims["iat"] = time.Now()
// 	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
// 	// 電子署名
// 	tokenString, err := token.SignedString([]byte(signingKey))

// 	if err != nil {
// 		return "", domain.WrapInternalError(err)
// 	}
// 	return tokenString, nil
// }

// // GetUserNewToken is get new token for users
// func GetUserNewToken(user domain.User, signingKey string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["sub"] = user.UserId
// 	claims["name"] = user.Name
// 	claims["iss"] = iss
// 	claims["phone_number"] = user.PhoneNumber
// 	claims["iat"] = time.Now()
// 	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

// 	tokenString, err := token.SignedString([]byte(signingKey))
// 	if err != nil {
// 		return "", domain.WrapInternalError(err)
// 	}
// 	return tokenString, nil
// }
