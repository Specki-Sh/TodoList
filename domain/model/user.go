package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Tasks        []Task `json:"tasks"`
	PasswordHash string `json:"password_hash"`
}
