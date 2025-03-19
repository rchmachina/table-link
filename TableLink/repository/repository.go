package repository

import (
	"database/sql"
	"errors"
	models "tablelink/dto"
)

type AuthRepository interface {
	Login(email, password string) (*models.User, error)
}

type authRepo struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepo{db: db}
}

func (r *authRepo) Login(email, password string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1 AND password = $2 LIMIT 1`

	err := r.db.QueryRow(query, email, password).Scan(&user).Error
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

////////// user

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.CreateNewUser) error
	UpdateUser(userID, name string) error
	DeleteUser(userID string) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

// 🔹 GetAllUsers menggunakan raw query
func (r *userRepo) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, role_id, name, email, last_access, role_name FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.RoleID, &user.Name, &user.Email, &user.LastAccess, &user.RoleName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// 🔹 CreateUser menggunakan raw query
func (r *userRepo) CreateUser(user models.CreateNewUser) error {
	query := `INSERT INTO users (id, role_id, name, email, password) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, user.ID, user.RoleID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// 🔹 UpdateUser menggunakan raw query
func (r *userRepo) UpdateUser(userID, name string) error {
	query := `UPDATE users SET name = $1 WHERE id = $2`
	res, err := r.db.Exec(query, name, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

// 🔹 DeleteUser menggunakan raw query
func (r *userRepo) DeleteUser(userID string) error {
	query := `DELETE FROM users WHERE id = $1`
	res, err := r.db.Exec(query, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
