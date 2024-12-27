package public_api

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type Resource struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Blob        []byte    `json:"blob"`
	UploadDate  time.Time `json:"upload_date"`
	AuthorName  string    `json:"author_name"`
	GroupNumber string    `json:"group_number"`
	SubjectID   int       `json:"subject_id"`
}

type ResourceResponse struct {
	Resources Resource `json:"resources"`
	Error     string   `json:"error,omitempty"`
}

type ResourcesResponse struct {
	Resources []Resource `json:"resources"`
	Error     string     `json:"error,omitempty"`
}

type CreateResourceRequest struct {
	Name        string    `json:"name"`
	Blob        []byte    `json:"blob"`
	UploadDate  time.Time `json:"upload_date"`
	AuthorName  string    `json:"author_name"`
	GroupNumber string    `json:"group_number"`
}

type DeleteResourceResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SubjectResponse struct {
	Subject Subject `json:"subject"`
	Error   string  `json:"error,omitempty"`
}

type SubjectsResponse struct {
	Subjects []Subject `json:"subjects"`
	Error    string    `json:"error,omitempty"`
}

type CreateSubjectRequest struct {
	Name string `json:"name"`
}

type DeleteSubjectResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
