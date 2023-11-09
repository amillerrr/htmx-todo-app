package models

type Todo struct {
	ID   string `json:"id" db:"id"`
	Todo string `json:"todo" db:"todo"`
	Done bool   `json:"done" db:"done"`
}
