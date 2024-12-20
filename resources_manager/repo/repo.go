package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sabina301/exchange_of_resources/auth/db"
	"github.com/sabina301/exchange_of_resources/auth/models"
)

func GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := db.Db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, password, role) VALUES ($1, $2, $3)"
	_, err := db.Db.Exec(query, user.Username, user.Password, user.Role)
	return err
}

func FindUserByUsername(username string) (*models.User, error) {
	query := "SELECT username, role FROM users WHERE username = $1"
	row := db.Db.QueryRowContext(context.Background(), query, username)

	var user models.User
	var role string

	if err := row.Scan(&user.Username, &role); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	user.Role = role

	return &user, nil
}
