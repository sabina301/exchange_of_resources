package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/sabina301/exchange_of_resources/resources_manager/repo"
	"log"
	"net/http"
	"strconv"
)

type ResourceController struct {
	repo repo.ResourceRepository
}

func NewResourceController(repo repo.ResourceRepository) *ResourceController {
	return &ResourceController{repo: repo}
}

func (rc *ResourceController) GetResource(c echo.Context) error {
	resID, err := strconv.Atoi(c.Param("resId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, GetResourceResponse{Error: "Invalid resource ID"})
	}

	resource, err := rc.repo.GetResourceByID(resID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetResourceResponse{Error: "Failed to get resource"})
	}

	return c.JSON(http.StatusOK, GetResourceResponse{Resource: Resource(*resource)})
}

func (rc *ResourceController) GetAllResources(c echo.Context) error {
	subjId := c.Param("subjId")
	intSubjId, err := strconv.Atoi(subjId)
	if err != nil {
		log.Println("Failed to parse subjId")
		return c.JSON(http.StatusBadRequest, GetAllResourcesResponse{Error: "Invalid subject ID"})
	}
	resources, err := rc.repo.GetAllResourcesBySubjectID(intSubjId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetAllResourcesResponse{Error: "Failed to get resources"})
	}

	log.Println(resources)
	resp := make([]Resource, len(resources))
	for i, val := range resources {
		resp[i] = Resource(val)
	}
	log.Println(resp)
	log.Println(GetAllResourcesResponse{Resources: resp})
	return c.JSON(http.StatusOK, GetAllResourcesResponse{Resources: resp})
}

func (rc *ResourceController) CreateResource(c echo.Context) error {
	var createReq CreateResourceRequest
	if err := c.Bind(&createReq); err != nil {
		return c.JSON(http.StatusBadRequest, CreateResourceResponse{Error: "Invalid request"})
	}

	subjId, err := strconv.Atoi(c.Param("subjId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, CreateResourceResponse{Error: "Invalid subject ID"})
	}

	resource := &Resource{
		Name:        createReq.Name,
		Blob:        createReq.Blob,
		AuthorName:  createReq.AuthorName,
		GroupNumber: createReq.GroupNumber,
		SubjectID:   subjId,
	}
	log.Println(resource)

	err = rc.repo.CreateResource((*repo.Resource)(resource))
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, CreateResourceResponse{Error: "Failed to create resource"})
	}

	return c.JSON(http.StatusOK, CreateResourceResponse{Resource: *resource})
}

func (controller *ResourceController) DeleteResource(c echo.Context) error {
	resID, err := strconv.Atoi(c.Param("resId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, DeleteResourceResponse{Error: "Invalid resource ID"})
	}

	err = controller.repo.DeleteResourceByID(resID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, DeleteResourceResponse{Error: "Failed to delete resource"})
	}

	return c.JSON(http.StatusOK, DeleteResourceResponse{Message: "Resource deleted successfully"})
}
