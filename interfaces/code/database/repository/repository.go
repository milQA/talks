package repository

import (
	"code/database/structs"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	conn *sql.DB
}

// NewUserRepository return storage for user
func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{conn: conn}
}

// SaveUser save user to database
func (r *UserRepository) SaveUser(user *structs.User) error {
	if err := r.conn.QueryRow(`
		INSERT INTO users 
			(name, email, address)
			VALUES 
				($1, $2, $3)
		RETURNING id`,
		user.Name, user.Email, user.Address,
	).Scan(&user.ID); err != nil {
		return fmt.Errorf("cannot insert new user: %w", err)
	}

	return nil
}

// GetUserByID get user by id in database
func (r *UserRepository) GetUserByID(userID int64) (user *structs.User, _ error) {
	user = new(structs.User)

	if err := r.conn.QueryRow(`
		SELECT 
			id, name, email, address
		FROM 
			users 
		WHERE 
			id = $1`,
		userID,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Address); err != nil {
		return nil, fmt.Errorf("cannot select user by id: %w", err)
	}

	return user, nil
}
