package models

type Team struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	League string `json:"league" db:"league"`
}
