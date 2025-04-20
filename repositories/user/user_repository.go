package repository

import (
	model "api/models/user"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT id, name, email, age FROM users"
	rows, err := ur.connection.Query(query)
	if err != nil {
		fmt.Println("ERROR", err)
		return []model.User{}, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Age,
		)
		if err != nil {
			fmt.Println("ERROR", err)
			return []model.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) CreateUser(user model.User) (model.User, error) {
	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`

	err := ur.connection.QueryRow(query, user.Name, user.Email, user.Age).Scan(&user.ID)
	if err != nil {
		fmt.Println("ERROR", err)
		return model.User{}, err
	}

	return user, nil
}


func (ur *UserRepository) GetUserById(id int) (*model.User, error) {
	query := `SELECT id, name, email, age FROM users WHERE id = $1`

	var user model.User
	err := ur.connection.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Age,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Não é erro, apenas não existe
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}