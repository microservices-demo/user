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
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Username  string   `json:"username"`
	Password  string   `json:"password,omitempty"`
	Addresses []string `json:"addresses,omitempty"`
	Cards     []string `json:"cards,omitempty"`
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
