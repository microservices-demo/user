package api

// service.go contains the definition and implementation (business logic) of the
// login service. Everything here is agnostic to the transport (HTTP).

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/users"
)

var (
	passwordSalt    = "passwordsalt"
	ErrUnauthorized = errors.New("Unauthorized")
)

// Service is the login service, providing operations for users to login and register.
type Service interface {
	Login(username, password string) (users.User, error) // GET /login
	// Only used for testing at the moment
	Register(username, password, email string) bool
	GetUsers(id string) ([]users.User, error)
	PostUser(u users.User) bool
	GetAddresses(id string) ([]users.Address, error)
	PostAddress(u users.Address, userid string) bool
	GetCards(id string) ([]users.Card, error)
	PostCard(u users.Card, userid string) bool
}

// NewFixedService returns a simple implementation of the Service interface,
func NewFixedService() Service {
	return &fixedService{}
}

type fixedService struct{}

func (s *fixedService) Login(username, password string) (users.User, error) {
	u, err := db.GetByName(username)
	if err != nil {
		return users.New(), err
	}
	if u.Password != calculatePassHash(password) {
		return users.New(), ErrUnauthorized
	}
	db.GetAttributes(&u)
	u.MaskCCs()
	return u, nil

}

func (s *fixedService) Register(username, password, email string) bool {
	u := users.New()
	u.Username = username
	u.Password = calculatePassHash(password)
	u.Email = email
	err := db.Create(&u)
	if err != nil {
		return false
	}
	return true
}

func (s *fixedService) GetUsers(id string) ([]users.User, error) {
	return make([]users.User, 0), nil
}

func (s *fixedService) PostUser(users.User) bool {
	return true
}

func (s *fixedService) GetAddresses(id string) ([]users.Address, error) {
	return make([]users.Address, 0), nil
}

func (s *fixedService) PostAddress(add users.Address, userid string) bool {
	return true
}

func (s *fixedService) GetCards(id string) ([]users.Card, error) {
	return make([]users.Card, 0), nil
}

func (s *fixedService) PostCard(card users.Card, userid string) bool {
	return true
}

func calculatePassHash(pass string) string {
	h := sha1.New()
	io.WriteString(h, passwordSalt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
