package public_api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/public_api_errors/errors"
	"net/http"
	"strings"
)

func HandlePublicApiError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		var errorCode int

		switch err.(type) {
		case *echo.HTTPError:
			return err
		case errors.AuthenticationFailed:
			errorCode = http.StatusUnauthorized
		case errors.MalformedBody:
			errorCode = http.StatusBadRequest
		default:
			err = errors.NewInternalError()
			errorCode = http.StatusInternalServerError
		}
		return c.JSON(errorCode, err)
	}
}

func AllowRoles(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing authorization header"})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("неизвестный метод подписи: %v", token.Header["alg"])
				}
				return []byte("secret"), nil
			})

			var userRole string

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userRole, ok = claims["role"].(string)
				if !ok {
					return fmt.Errorf("cant find role in token")
				}
			}

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}

			for _, role := range roles {
				if userRole == role {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, map[string]string{"error": "Access forbidden: insufficient permissions"})
		}
	}
}
