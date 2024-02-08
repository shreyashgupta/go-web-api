package models

import (
	"errors"

	"example.com/go_api/db"
	"example.com/go_api/utils"
)

type Author struct {
	Id       int64
	Name     string `binding: "required"`
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func AddAuthor(author *Author) (int64, error) {
	query := `INSERT INTO Authors(name, email, password)
	 VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(author.Password)
	if err != nil {
		return -1, err
	}
	result, err := stmt.Exec(author.Name, author.Email, hashedPassword)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ValidateAuthor(author *Author) error {
	query := `SELECT * from authors where email = ?`
	rows := db.DB.QueryRow(query, author.Email)
	queriedAuthor := Author{}

	err := rows.Scan(&queriedAuthor.Id, &queriedAuthor.Name, &queriedAuthor.Email, &queriedAuthor.Password)
	if err != nil {
		return errors.New("email not found")
	}

	isPasswordValid := utils.DoesPasswordMatch(author.Password, queriedAuthor.Password)
	if !isPasswordValid {
		return errors.New("invalid password")
	}

	return nil
}

// func GetAllAuthors() ([]User, error) {
// 	users := []User{}
// 	getAllUsers := `
// 	SELECT * FROM users;
// 	`
// 	rows, err := db.DB.Query(getAllUsers)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, err
// }

// func GetUserWithId(id int64) (*User, error) {
// 	query := `SELECT * from users where id = (?)`
// 	rows := db.DB.QueryRow(query, id)
// 	user := User{}

// 	err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
