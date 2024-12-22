package rest

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/auth/models"
	"github.com/sabina301/exchange_of_resources/auth/repo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret")

type UserController struct {
	repo repo.UserRepository
}

func NewUserController(repo repo.UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (uc *UserController) Login(c echo.Context) error {
	var loginReq LoginRequest
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := uc.repo.GetUserByUsername(loginReq.Username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return c.JSON(http.StatusUnauthorized, LoginResponse{Token: "", Error: "Cant find user"})
		}
		return c.JSON(http.StatusInternalServerError, LoginResponse{Token: "", Error: err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, LoginResponse{Token: "", Error: "Wrong password"})
	}

	token, err := generateJWT(user.Username, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, LoginResponse{Token: "", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func (c *UserController) Register(ctx echo.Context) error {
	var regReq RegisterRequest
	if err := ctx.Bind(&regReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	user := &models.User{
		Username: regReq.Username,
		Password: string(hashedPassword),
		Role:     regReq.Role,
	}

	err = c.repo.CreateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, RegisterResponse{"User created", ""})
}

func generateJWT(username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  username,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
