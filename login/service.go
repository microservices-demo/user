package login

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
	Register(username, password string) bool
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
	return u, nil

}

func (s *fixedService) Register(username, password string) bool {
	u := users.New()
	u.Username = username
	u.Password = calculatePassHash(password)
	err := db.Create(&u)
	if err != nil {
		return false
	}
	return true
}

func calculatePassHash(pass string) string {
	h := sha1.New()
	io.WriteString(h, passwordSalt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
