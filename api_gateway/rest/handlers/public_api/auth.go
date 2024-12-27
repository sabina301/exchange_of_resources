package public_api

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	var loginReq LoginRequest
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	reqBody, err := json.Marshal(loginReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp, err := http.Post("http://127.0.0.1:8000/int/v1/auth/login", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(resp.StatusCode, loginResp)
}

func Register(c echo.Context) error {
	var registerReq RegisterRequest
	if err := c.Bind(&registerReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	reqBody, err := json.Marshal(registerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp, err := http.Post("http://127.0.0.1:8000/int/v1/auth/register", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	var registerResp RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&registerResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(resp.StatusCode, registerResp)
}
