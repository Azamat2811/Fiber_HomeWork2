package services

import (
	"database/sql"
	"errors"
	"fiber_hw_2/db"
	"fiber_hw_2/models"
)

func CreateUsers(users models.Users) (models.Users, error) {
	err := db.DB.QueryRow("INSERT INTO users (name,  age,) VALUES ($1, $2) RETURNING id",
		users.Name,
		users.Age,
	).Scan(&users.ID)

	if err != nil {
		return models.Users{}, err
	}
	return users, nil
}

func GetUsersByID(id int) (models.Users, error) {
	var user models.Users
	err := db.DB.QueryRow("SELECT id, name, age, FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Name, &user.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Users{}, errors.New("user not found")
		}
		return models.Users{}, err
	}
	return user, nil
}

func GetAllUsers() ([]models.Users, error) {
	rows, err := db.DB.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.Users
	for rows.Next() {
		var user models.Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func DeleteUsers(id int) error {

	result, err := db.DB.Exec(
		"DELETE FROM users WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func UpdateUser(id int, data map[string]interface{}) (models.Users, error) {
	user, err := GetUsersByID(id)
	if err != nil {
		return models.Users{}, err
	}
	if name, ok := data["username"].(string); ok {
		user.Name = name
	}
	if age, ok := data["age"].(float64); ok {
		user.Age = int(age)
	}

	_, err = db.DB.Exec(

		"UPDATE users SET username = $1, phone = $2, email = $3, age = $4, city = $5 WHERE id = $6",
		user.Username,
		user.Phone,
		user.Email,
		user.Age,
		user.City,
		id,
	)
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func FullUpdateUsers(id int, updated models.Users) (models.Users, error) {
	result, err := db.DB.Exec(
		"UPDATE users SET username = $1, phone = $2, email = $3, age = $4, city = $5 WHERE id = $6",
		updated.Username,
		updated.Phone,
		updated.Email,
		updated.Age,
		updated.City,
		id,
	)

	if err != nil {
		return models.Users{}, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return models.Users{}, err
	}
	if rows == 0 {
		return models.Users{}, errors.New("user not found")
	}

	updated.ID = id
	return updated, nil
}
