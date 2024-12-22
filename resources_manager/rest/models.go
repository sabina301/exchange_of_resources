package rest

import "time"

type Resource struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Blob        []byte    `json:"blob"`
	UploadDate  time.Time `json:"upload_date"`
	AuthorName  string    `json:"author_name"`
	GroupNumber string    `json:"group_number"`
	SubjectID   int       `json:"subject_id"`
}

type GetResourceResponse struct {
	Resource Resource `json:"resource"`
	Error    string   `json:"error,omitempty"`
}

type GetAllResourcesResponse struct {
	Resources []Resource `json:"resources"`
	Error     string     `json:"error,omitempty"`
}

type CreateResourceRequest struct {
	Name        string `json:"name"`
	Blob        []byte `json:"blob"`
	AuthorName  string `json:"author_name"`
	GroupNumber string `json:"group_number"`
	SubjectID   int    `json:"subject_id"`
}

type CreateResourceResponse struct {
	Resource Resource `json:"resource"`
	Error    string   `json:"error,omitempty"`
}

type DeleteResourceResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
