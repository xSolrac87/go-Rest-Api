package models

import (
	"api/security"
	"time"
)

// User struct
type User struct {
	NickName  string    `bson:"nickname" json:"nickname"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// HashPassword encode user password
func (u *User) HashPassword(p string) error {
	hashedPassword, err := security.Hash(p)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}