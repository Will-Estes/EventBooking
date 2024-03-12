package models

import (
	"database/sql"
	"eventbooking/db"
	"fmt"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, date_time, user_id) 
				VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetEventById(id int64) (*Event, error) {
	var event Event
	query := `SELECT * FROM events WHERE id = ?`
	err := db.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("event not found with ID %d", id)
		}
		return nil, err
	}
	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (e *Event) Update() error {
	query := `UPDATE events SET name=?, description=?, location=?, date_time=? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID)
	return err
}

func (e *Event) Register(userId int64) error {
	query := `INSERT INTO registrations(event_id, user_id) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) GetRegistrations() ([]Register, error) {
	query := `SELECT * FROM registrations WHERE event_id = ?`
	rows, err := db.DB.Query(query, e.ID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var registers []Register
	for rows.Next() {
		var r Register
		err := rows.Scan(&r.ID, &r.EventId, &r.UserId)
		if err != nil {
			return nil, err
		}
		registers = append(registers, r)
	}

	return registers, nil
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err
}
