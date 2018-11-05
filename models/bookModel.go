package models

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
	Year int `json:"year"`
}