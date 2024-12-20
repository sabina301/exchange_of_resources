package repo

import (
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
