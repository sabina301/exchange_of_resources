package errors

import (
	"encoding/json"
	"fmt"
	"github.com/sabina301/exchange_of_resources/public_api_errors/templates"
	"reflect"
)

const (
	TypeInternalError        = "InternalError"
	TypeAuthenticationFailed = "AuthenticationFailed"
	TypeMalformedBody        = "MalformedBody"
)

var errorInstances = map[string]any{
	TypeInternalError:        InternalError{},
	TypeAuthenticationFailed: AuthenticationFailed{},
	TypeMalformedBody:        MalformedBody{},
}

type BaseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func newBase(t, message string) BaseError {
	return BaseError{t, message}
}

type InternalError struct {
	BaseError
}

func NewInternalError() InternalError {
	msg := templates.CreateMessage(TypeInternalError, nil)
	return InternalError{
		BaseError: newBase(TypeInternalError, msg),
	}
}

type AuthenticationFailed struct {
	BaseError
	Username string `json:"username"`
}

func NewAuthenticationFailed(username string) AuthenticationFailed {

	msg := templates.CreateMessage(TypeAuthenticationFailed, map[string]any{
		"Username": username,
	})
	return AuthenticationFailed{
		BaseError: newBase(TypeAuthenticationFailed, msg),
		Username:  username,
	}
}

type MalformedBody struct {
	BaseError
}

func NewMalformedBody() MalformedBody {
	msg := templates.CreateMessage(TypeMalformedBody, nil)
	return MalformedBody{
		BaseError: newBase(TypeMalformedBody, msg),
	}
}

func Unmarshal(data []byte) (error, bool) {
	var baseError BaseError
	if err := json.Unmarshal(data, &baseError); err != nil {
		return err, false
	}

	errorInstance, exists := errorInstances[baseError.Type]
	if !exists {
		return fmt.Errorf("cant find this error: %q", baseError.Type), false
	}

	errorPointer := reflect.New(reflect.TypeOf(errorInstance))
	if err := json.Unmarshal(data, errorPointer.Interface()); err != nil {
		return err, false
	}

	parsedError, ok := errorPointer.Elem().Interface().(error)
	if !ok {
		return fmt.Errorf("cant unmarshal it: %T", errorPointer.Elem().Interface()), false
	}
	return parsedError, true
}
