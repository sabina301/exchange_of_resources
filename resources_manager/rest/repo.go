package rest

import (
	"database/sql"
	"log"
)

type ResourceRepository interface {
	GetResourceByID(id int) (*Resource, error)
	GetAllResourcesBySubjectID(subjectID int) ([]Resource, error)
	CreateResource(resource *Resource) error
	DeleteResourceByID(id int) error
}

type resourceRepository struct {
	db *sql.DB
}

func NewResourceRepository(db *sql.DB) *resourceRepository {
	return &resourceRepository{db: db}
}

func (r *resourceRepository) GetResourceByID(id int) (*Resource, error) {
	resource := &Resource{}
	query := "SELECT id, name, blob, upload_date, author_name, group_number, subject_id FROM resources WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&resource.ID, &resource.Name, &resource.Blob, &resource.UploadDate, &resource.AuthorName, &resource.GroupNumber, &resource.SubjectID)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (r *resourceRepository) GetAllResourcesBySubjectID(subjectID int) ([]Resource, error) {
	resources := []Resource{}
	query := "SELECT id, name, blob, upload_date, author_name, group_number, subject_id FROM resources WHERE subject_id = $1"
	rows, err := r.db.Query(query, subjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		resource := Resource{}
		err := rows.Scan(&resource.ID, &resource.Name, &resource.Blob, &resource.UploadDate, &resource.AuthorName, &resource.GroupNumber, &resource.SubjectID)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}

	return resources, nil
}

func (r *resourceRepository) CreateResource(resource *Resource) error {
	log.Println(resource.SubjectID)
	query := "INSERT INTO resources (name, blob, author_name, group_number, subject_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, upload_date"
	err := r.db.QueryRow(query, resource.Name, resource.Blob, resource.AuthorName, resource.GroupNumber, resource.SubjectID).Scan(&resource.ID, &resource.UploadDate)
	if err != nil {
		return err
	}
	return nil
}

func (r *resourceRepository) DeleteResourceByID(id int) error {
	query := "DELETE FROM resources WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
