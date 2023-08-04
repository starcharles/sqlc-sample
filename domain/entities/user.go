package entities

import (
	"time"

	"github.com/starcharles/sqlc-example/domain/value_object"
)

type User struct {
	id        value_object.ID
	name      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func (u *User) ID() value_object.ID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func NewUser(id value_object.ID, name string, email string, password string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
