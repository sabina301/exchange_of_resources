package public_api

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetSubjectResponse struct {
	Subject Subject `json:"subject"`
	Error   string  `json:"error,omitempty"`
}

type GetAllSubjectsResponse struct {
	Subjects []Subject `json:"subjects"`
	Error    string    `json:"error,omitempty"`
}

type CreateSubjectRequest struct {
	Name string `json:"name"`
}

type CreateSubjectResponse struct {
	Subject Subject `json:"subject"`
	Error   string  `json:"error,omitempty"`
}

type DeleteSubjectResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
