package public_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetResource(c echo.Context) error {
	resID := c.Param("resId")
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8002/int/v1/resources/%s", resID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	var resourceResp ResourceResponse
	if err := json.NewDecoder(resp.Body).Decode(&resourceResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(resp.StatusCode, resourceResp)
}

func GetAllResources(c echo.Context) error {
	log.Println("!")
	subjID := c.Param("subjId")
	log.Println(subjID)
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8002/int/v1/resources/%s/all", subjID))
	if err != nil {
		log.Println("!", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()
	log.Println(resp)
	var resourcesResp ResourcesResponse

	if err := json.NewDecoder(resp.Body).Decode(&resourcesResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Println(resourcesResp)
	return c.JSON(resp.StatusCode, resourcesResp)
}

func CreateResource(c echo.Context) error {
	subjID := c.Param("subjId")
	var createReq CreateResourceRequest
	if err := c.Bind(&createReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	reqBody, err := json.Marshal(createReq)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp, err := http.Post(fmt.Sprintf("http://127.0.0.1:8002/int/v1/resources/%s", subjID), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	var resourceResp ResourceResponse
	if err := json.NewDecoder(resp.Body).Decode(&resourceResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(resp.StatusCode, resourceResp)
}

func DeleteResource(c echo.Context) error {
	resID := c.Param("resId")
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://127.0.0.1:8002/int/v1/resources/%s", resID), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	var deleteResp DeleteResourceResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(resp.StatusCode, deleteResp)
}
