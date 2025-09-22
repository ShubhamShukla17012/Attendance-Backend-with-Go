package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Employee struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Grade      string    `json:"grade"`
	LogInTime  time.Time `json:"log_out_time,omitempty"`
	LogOutTime time.Time `json:"log_in_time,omitempty"`
}

func (e *Employee) Jsonify() []byte {
	result, _ := json.Marshal(*e)
	return result
}

func (e *Employee) Validate(checkTime bool) error {

	if e.ID == "" {
		return errors.New("ID cannot be empty")
	}

	if e.Name == "" {
		return errors.New("name cannot be empty")
	}

	if e.Grade == "" {
		return errors.New("grade cannot be empty")
	}

	switch e.Grade {
	case "A4", "A5":
	default:
		return errors.New("grade should be one of : A4, A5")
	}

	if !checkTime {
		return nil
	}

	if e.LogInTime.IsZero() {
		return errors.New("login time cannot be zero")
	}

	if e.LogOutTime.IsZero() {
		return errors.New("logout time cannot be zero")
	}

	return nil

}
