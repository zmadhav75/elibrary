package repository

import (
	"database/sql"
	"elibrary/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		`SELECT id, email, subscribed, subscribed_at 
		FROM users WHERE email = $1`, email,
	).Scan(&user.ID, &user.Email, &user.Subscribed, &user.SubscribedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.QueryRow(
		`INSERT INTO users (email) 
		VALUES ($1) 
		RETURNING id`, user.Email,
	).Scan(&user.ID)
}

func (r *UserRepository) UpdateSubscription(userID int, subscribed bool) error {
	_, err := r.db.Exec(
		`UPDATE users 
		SET subscribed = $1, subscribed_at = CURRENT_TIMESTAMP 
		WHERE id = $2`, subscribed, userID,
	)
	return err
}
