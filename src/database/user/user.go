package userModel

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint      `json:"id,omitempty" example:"1"`
	FirstName    string    `json:"first_name,omitempty" example:"Matias"`
	LastName     string    `json:"last_name,omitempty" example:"Lamela"`
	Username     string    `json:"username,omitempty" example:"escalera92"`
	Email        string    `json:"email,omitempty" example:"matiaslamela1992@gmail.com"`
	Picture      string    `json:"picture,omitempty" example:"https://media.revistagq.com/photos/5ca5fb2aeccc6ac985d559a9/1:1/w_707,h_707,c_limit/rickrolling_3543.png"`
	Password     string    `json:"password,omitempty" example:"1234"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func (u *User) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(passwordHash)

	return nil
}

func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}
