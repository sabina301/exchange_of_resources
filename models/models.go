package models

import (
	"time"
)

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	GroupNumber string `json:"group_number"`
}

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
