package entity

import "encoding/json"

type User struct {
	ID    uint
	Name  string
	Email string
}

func NewUser(id uint, name string, email string) *User {
	return &User{ID: id, Name: name, Email: email}
}

func (u User) ToJSON() (string, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
