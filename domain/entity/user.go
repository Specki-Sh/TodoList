package entity

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Tasks    []Task `json:"tasks"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
