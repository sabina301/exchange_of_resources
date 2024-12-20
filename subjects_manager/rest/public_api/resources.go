package public_api

import (
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/models"
	"net/http"
	"strconv"
)

type SubjectController struct {
	subjectRepository repositories.SubjectRepository
}

func NewSubjectController(subjectRepository repositories.SubjectRepository) *SubjectController {
	return &SubjectController{subjectRepository: subjectRepository}
}

func (controller *SubjectController) GetSubject(c echo.Context) error {
	subjID, err := strconv.Atoi(c.Param("subjId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.GetSubjectResponse{Error: "Invalid subject ID"})
	}

	subject, err := controller.subjectRepository.GetSubjectByID(subjID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.GetSubjectResponse{Error: "Failed to get subject"})
	}

	return c.JSON(http.StatusOK, models.GetSubjectResponse{Subject: *subject})
}

func (controller *SubjectController) GetAllSubjects(c echo.Context) error {
	subjects, err := controller.subjectRepository.GetAllSubjects()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.GetAllSubjectsResponse{Error: "Failed to get subjects"})
	}

	return c.JSON(http.StatusOK, models.GetAllSubjectsResponse{Subjects: subjects})
}

func (controller *SubjectController) CreateSubject(c echo.Context) error {
	var createReq models.CreateSubjectRequest
	if err := c.Bind(&createReq); err != nil {
		return c.JSON(http.StatusBadRequest, models.CreateSubjectResponse{Error: "Invalid request"})
	}

	subject := &models.Subject{
		Name: createReq.Name,
	}

	err := controller.subjectRepository.CreateSubject(subject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.CreateSubjectResponse{Error: "Failed to create subject"})
	}

	return c.JSON(http.StatusOK, models.CreateSubjectResponse{Subject: *subject})
}

func (controller *SubjectController) DeleteSubject(c echo.Context) error {
	subjID, err := strconv.Atoi(c.Param("subjId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.DeleteSubjectResponse{Error: "Invalid subject ID"})
	}

	err = controller.subjectRepository.DeleteSubjectByID(subjID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.DeleteSubjectResponse{Error: "Failed to delete subject"})
	}

	return c.JSON(http.StatusOK, models.DeleteSubjectResponse{Message: "Subject deleted successfully"})
}
