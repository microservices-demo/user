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
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password,omitempty" bson:"password,omitempty"`
	Addresses []Address `json:"addresses,omitempty" bson:"-"`
	Cards     []Card    `json:"cards,omitempty" bson:"-"`
	UserID    string    `bson:"-"`
}

func New() User {
	return User{Addresses: make([]Address, 0), Cards: make([]Card, 0)}
}

func (c *User) Validate() error {
	if c.FirstName == "" {
		return fmt.Errorf(ErrMissingField, "FirstName")
	}
	if c.LastName == "" {
		return fmt.Errorf(ErrMissingField, "LastName")
	}
	if c.Username == "" {
		return fmt.Errorf(ErrMissingField, "Username")
	}
	if c.Password == "" {
		return fmt.Errorf(ErrMissingField, "Password")
	}
	return nil
}
