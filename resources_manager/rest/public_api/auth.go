package public_api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/auth/models"
	"github.com/sabina301/exchange_of_resources/auth/repo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret")

func Login(c echo.Context) error {
	var loginReq LoginRequest
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := repo.GetUserByUsername(loginReq.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, "invalid username or password")
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid username or password")
	}

	token, err := generateJWT(user.Username, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func generateJWT(username string, roles string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  username,
		"role": roles,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func Register(c echo.Context) error {
	var regReq RegisterRequest
	if err := c.Bind(&regReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user := &models.User{
		Username: regReq.Username,
		Password: string(hashedPassword),
		Role:     regReq.Role,
	}

	err = repo.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "User registered successfully")
}
