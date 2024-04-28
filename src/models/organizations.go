package models

type Organization struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Description string
}
