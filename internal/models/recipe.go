package models

type Recipe struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Steps       string `json:"steps"`
	Author      string `json:"author"`
}
