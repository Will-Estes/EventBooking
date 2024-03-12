package models

import (
	"errors"
	"eventbooking/db"
	"eventbooking/utils"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	return err
}

func (u *User) ValidateUser() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	var storedPassword string
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &storedPassword)
	if err != nil {
		return errors.New("Invalid credentials")
	}
	isValid := utils.ValidatePassword(u.Password, storedPassword)
	if !isValid {
		return errors.New("Invalid credentials")
	}
	return nil
}
