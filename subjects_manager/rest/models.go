package rest

import "github.com/sabina301/exchange_of_resources/models"

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetSubjectResponse struct {
	Subject models.Subject `json:"subject"`
	Error   string         `json:"error,omitempty"`
}

type GetAllSubjectsResponse struct {
	Subjects []models.Subject `json:"subjects"`
	Error    string           `json:"error,omitempty"`
}

type CreateSubjectRequest struct {
	Name string `json:"name"`
}

type CreateSubjectResponse struct {
	Subject models.Subject `json:"subject"`
	Error   string         `json:"error,omitempty"`
}

type DeleteSubjectResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
