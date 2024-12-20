package repo

import (
	"database/sql"
	"github.com/sabina301/exchange_of_resources/models"
)

type SubjectRepository interface {
	GetSubjectByID(id int) (*models.Subject, error)
	GetAllSubjects() ([]models.Subject, error)
	CreateSubject(subject *models.Subject) error
	DeleteSubjectByID(id int) error
}

type subjectRepository struct {
	db *sql.DB
}

func NewSubjectRepository(db *sql.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) GetSubjectByID(id int) (*models.Subject, error) {
	subject := &models.Subject{}
	query := "SELECT id, name FROM subjects WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&subject.ID, &subject.Name)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

func (r *subjectRepository) GetAllSubjects() ([]models.Subject, error) {
	subjects := []models.Subject{}
	query := "SELECT id, name FROM subjects"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		subject := models.Subject{}
		err := rows.Scan(&subject.ID, &subject.Name)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}

func (r *subjectRepository) CreateSubject(subject *models.Subject) error {
	query := "INSERT INTO subjects (name) VALUES ($1) RETURNING id"
	err := r.db.QueryRow(query, subject.Name).Scan(&subject.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *subjectRepository) DeleteSubjectByID(id int) error {
	query := "DELETE FROM subjects WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
