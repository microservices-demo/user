package api

// service.go contains the definition and implementation (business logic) of the
// user service. Everything here is agnostic to the transport (HTTP).

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/users"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

// Service is the user service, providing operations for users to login, register, and retrieve customer information.
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
	u, err := db.GetUserByName(username)
	if err != nil {
		return users.New(), err
	}
	if u.Password != calculatePassHash(password, u.Salt) {
		return users.New(), ErrUnauthorized
	}
	db.GetUserAttributes(&u)
	u.MaskCCs()
	return u, nil

}

func (s *fixedService) Register(username, password, email string) bool {
	u := users.New()
	u.Username = username
	u.Password = calculatePassHash(password, u.Salt)
	u.Email = email
	err := db.CreateUser(&u)
	if err != nil {
		return false
	}
	return true
}

func (s *fixedService) GetUsers(id string) ([]users.User, error) {
	if id == "" {
		us, err := db.GetUsers()
		for k, u := range us {
			u.AddLinks()
			us[k] = u
		}
		return us, err
	}
	u, err := db.GetUser(id)
	u.AddLinks()
	return []users.User{u}, err
}

func (s *fixedService) PostUser(user users.User) bool {
	err := db.CreateUser(&user)
	if err != nil {
		return false
	}
	return true
}

func (s *fixedService) GetAddresses(id string) ([]users.Address, error) {
	if id == "" {
		as, err := db.GetAddresses()
		for k, a := range as {
			a.AddLinks()
			as[k] = a
		}
		return as, err
	}
	a, err := db.GetAddress(id)
	a.AddLinks()
	return []users.Address{a}, err
}

func (s *fixedService) PostAddress(add users.Address, userid string) bool {
	err := db.CreateAddress(&add, userid)
	if err != nil {
		return false
	}
	return true
}

func (s *fixedService) GetCards(id string) ([]users.Card, error) {
	if id == "" {
		cs, err := db.GetCards()
		for k, c := range cs {
			c.AddLinks()
			cs[k] = c
		}
		return cs, err
	}
	c, err := db.GetCard(id)
	return []users.Card{c}, err
}

func (s *fixedService) PostCard(card users.Card, userid string) bool {
	err := db.CreateCard(&card, userid)
	if err != nil {
		return false
	}
	return true
}

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
