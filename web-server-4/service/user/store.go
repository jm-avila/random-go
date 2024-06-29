package user

import (
	"database/sql"
	"fmt"

	"github.com/jmavila/golang/web-server-4/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*models.User, error) {
	rows, err := s.db.Query("SELECT I FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	user := new(models.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *Store) GetUserById(id int) (*models.User, error) {
	rows, err := s.db.Query("SELECT I FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	user := new(models.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *Store) CreateUser(user models.User) error {
	return nil
}

func scanRowIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return nil, err
}
