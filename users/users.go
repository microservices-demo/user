package users

import (
	"errors"
	"fmt"
)

var (
	ErrNoCustomerInResponse = errors.New("Response has no matching customer")
	ErrMissingField         = "Error missing %v"
)

type User struct {
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	Email     string    `json:"email" bson:"email"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"-" bson:"password,omitempty"`
	Addresses []Address `json:"addresses,omitempty" bson:"-"`
	Cards     []Card    `json:"cards,omitempty" bson:"-"`
	UserID    string    `json:"id" bson:"-"`
}

func New() User {
	return User{Addresses: make([]Address, 0), Cards: make([]Card, 0)}
}

func (u *User) Validate() error {
	if u.FirstName == "" {
		return fmt.Errorf(ErrMissingField, "FirstName")
	}
	if u.LastName == "" {
		return fmt.Errorf(ErrMissingField, "LastName")
	}
	if u.Username == "" {
		return fmt.Errorf(ErrMissingField, "Username")
	}
	if u.Password == "" {
		return fmt.Errorf(ErrMissingField, "Password")
	}
	return nil
}

func (u *User) MaskCCs() {
	for k, c := range u.Cards {
		c.MaskCC()
		u.Cards[k] = c
	}
}
