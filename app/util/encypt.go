package util

// import (
// 	"github.com/Code0716/go-vtm/app/domain"
// 	"golang.org/x/crypto/bcrypt"
// )

// // GetHush is get encrypt
// func GetHush(password string) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
// 	if err != nil {
// 		return "", domain.WrapInternalError(err)
// 	}
// 	return string(hash), nil
// }

// // CheckHush is check password
// func CheckHush(savedPass, checkPass string) bool {
// 	hash := []byte(savedPass)
// 	err := bcrypt.CompareHashAndPassword(hash, []byte(checkPass))
// 	return err == nil
// }
