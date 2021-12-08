//go:generate stringer -type=ErrorType -linecomment -output error_string.gen.go

package domain

import (
	"fmt"
)

// TODO:errorのメッセージを適切に返すようにする。
// app/infrastructure/db/sqlhandler.go

// ErrorType is error type
type ErrorType int

const (
	// ErrorTypeValidationFailed バリデーションエラー
	ErrorTypeValidationFailed ErrorType = iota + 1 // validation_failed
	// ErrorTypeRegistAdminValidationFailed バリデーションエラー
	ErrorTypeRegistAdminValidationFailed // regist admin validation_failed
	// ErrorTypeAdminEmailValidationFailed validate admin email failed
	ErrorTypeAdminEmailValidationFailed
	// ErrorTypeAdminLoginValidationFailed admin login validatein_failed
	ErrorTypeAdminLoginValidationFailed
	// ErrorTypeRegistItemAlreadyRegistered regest error
	ErrorTypeRegistItemAlreadyRegistered // regist item validation_failed
	// ErrorTypePasswordOrEmailValidationFailed admin login validatein_failed
	ErrorTypePasswordOrEmailValidationFailed
	// ErrorTypeRegistMemberValidationFailed regist member validete failed
	ErrorTypeRegistMemberValidationFailed
	// ErrorTypeUUIDValidationFailed invalid uuid
	ErrorTypeUUIDValidationFailed // invalid_uuid
	// ErrorTypeMemberAlreadyDeleted requested member already deleted
	ErrorTypeMemberAlreadyDeleted // requested_member_already_deleted
	// ErrorTypeAuthenticationFailed 認証エラー
	ErrorTypeAuthenticationFailed // authentication_faild
	// ErrorTypeNotFound not found
	ErrorTypeNotFound // not_found
	// ErrorTypeContentNotFound content not found
	ErrorTypeContentNotFound // content_not_found
	// ErrorTypeInternalError 内部エラー
	ErrorTypeInternalError // internal_error

)

// ErrorMessageMap map error　message
var ErrorMessageMap = map[ErrorType]string{
	// admin
	ErrorTypeRegistAdminValidationFailed: "name,password,mail_addressは必須です",
	ErrorTypeRegistItemAlreadyRegistered: "既に登録済みです",
	ErrorTypeAdminEmailValidationFailed:  "メールアドレスの形式が不正です",
	ErrorTypeAdminLoginValidationFailed:  "name,passwordは必須です",
	// member
	ErrorTypeRegistMemberValidationFailed: "name,phone_numberは必須です",
	ErrorTypeMemberAlreadyDeleted:         "リクエストされたmemberは既に削除されています。",
	// other
	ErrorTypeNotFound:                        "request uri found",
	ErrorTypeValidationFailed:                "validate failed",
	ErrorTypeAuthenticationFailed:            "認証エラー",
	ErrorTypeInternalError:                   "internal server error",
	ErrorTypePasswordOrEmailValidationFailed: "invalid passwrod or email",
	ErrorTypeUUIDValidationFailed:            "invalid uuid",
	ErrorTypeContentNotFound:                 "content not found",
}

// Error is Error struct
type Error struct {
	Type       ErrorType
	Status     int
	Message    string
	InnerError error
}

// NewError create error
func NewError(errType ErrorType) Error {
	return Error{
		Type: errType,
	}
}

// NewErrorf create error
func NewErrorf(errType ErrorType, msg string, args ...interface{}) Error {
	return Error{
		Type:    errType,
		Message: fmt.Sprintf(msg, args...),
	}
}

// WrapError wrap error
func WrapError(errType ErrorType, innerError error) Error {
	return Error{
		Type:       errType,
		InnerError: innerError,
	}
}

// WrapInternalError wrap error
func WrapInternalError(innerError error) Error {
	return Error{
		Type:       ErrorTypeInternalError,
		InnerError: innerError,
	}
}

// Unwrap unrap InnerError
func (e Error) Unwrap() error {
	return e.InnerError
}

func (e Error) Error() string {
	return e.Message
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// エラー
	Error `json:"error"`
}
