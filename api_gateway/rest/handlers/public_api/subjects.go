package public_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetSubject(c echo.Context) error {
	subjID := c.Param("subjId")
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:8001/int/v1/subjects/%s", subjID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	var subjectResp SubjectResponse
	if err := json.NewDecoder(resp.Body).Decode(&subjectResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(resp.StatusCode, subjectResp)
}

func GetAllSubjects(c echo.Context) error {
	resp, err := http.Get("http://127.0.0.1:8001/int/v1/subjects")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	var subjectsResp SubjectsResponse
	if err := json.NewDecoder(resp.Body).Decode(&subjectsResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(resp.StatusCode, subjectsResp)
}

func CreateSubject(c echo.Context) error {
	var createReq CreateSubjectRequest
	if err := c.Bind(&createReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	reqBody, err := json.Marshal(createReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Println("====createReq====", createReq)

	resp, err := http.Post("http://127.0.0.1:8001/int/v1/subjects", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	log.Println("====createReq====", resp.Body)

	var subjectResp SubjectResponse
	if err := json.NewDecoder(resp.Body).Decode(&subjectResp); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Println("====subjectResp====", subjectResp)

	return c.JSON(resp.StatusCode, subjectResp)
}

func DeleteSubject(c echo.Context) error {
	subjID := c.Param("subjId")
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://127.0.0.1:8001/int/v1/subjects/%s", subjID), nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	var deleteResp DeleteSubjectResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteResp); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(resp.StatusCode, deleteResp)
}
