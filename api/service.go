package api

// service.go contains the definition and implementation (business logic) of the
// user service. Everything here is agnostic to the transport (HTTP).

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/users"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

// Service is the user service, providing operations for users to login, register, and retrieve customer information.
type Service interface {
	Login(ctx context.Context, username, password string) (users.User, error) // GET /login
	Register(ctx context.Context, username, password, email, first, last string) (string, error)
	GetUsers(ctx context.Context, id string) ([]users.User, error)
	PostUser(ctx context.Context, u users.User) (string, error)
	GetAddresses(ctx context.Context, id string) ([]users.Address, error)
	PostAddress(ctx context.Context, u users.Address, userid string) (string, error)
	GetCards(ctx context.Context, id string) ([]users.Card, error)
	PostCard(ctx context.Context, u users.Card, userid string) (string, error)
	Delete(ctx context.Context, entity, id string) error
	Health() []Health // GET /health
}

// NewFixedService returns a simple implementation of the Service interface,
func NewFixedService() Service {
	return &fixedService{}
}

type fixedService struct{}

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

func (s *fixedService) Login(ctx context.Context, username, password string) (users.User, error) {
	u, err := db.GetUserByName(ctx, username)
	if err != nil {
		return users.New(), err
	}
	if u.Password != calculatePassHash(password, u.Salt) {
		return users.New(), ErrUnauthorized
	}
	db.GetUserAttributes(ctx, &u)
	u.MaskCCs()
	return u, nil

}

func (s *fixedService) Register(ctx context.Context, username, password, email, first, last string) (string, error) {
	u := users.New()
	u.Username = username
	u.Password = calculatePassHash(password, u.Salt)
	u.Email = email
	u.FirstName = first
	u.LastName = last
	err := db.CreateUser(ctx, &u)
	return u.UserID, err
}

func (s *fixedService) GetUsers(ctx context.Context, id string) ([]users.User, error) {
	if id == "" {
		us, err := db.GetUsers(ctx)
		for k, u := range us {
			u.AddLinks()
			us[k] = u
		}
		return us, err
	}
	u, err := db.GetUser(ctx, id)
	u.AddLinks()
	return []users.User{u}, err
}

func (s *fixedService) PostUser(ctx context.Context, u users.User) (string, error) {
	u.NewSalt()
	u.Password = calculatePassHash(u.Password, u.Salt)
	err := db.CreateUser(ctx, &u)
	return u.UserID, err
}

func (s *fixedService) GetAddresses(ctx context.Context, id string) ([]users.Address, error) {
	if id == "" {
		as, err := db.GetAddresses(ctx)
		for k, a := range as {
			a.AddLinks()
			as[k] = a
		}
		return as, err
	}
	a, err := db.GetAddress(ctx, id)
	a.AddLinks()
	return []users.Address{a}, err
}

func (s *fixedService) PostAddress(ctx context.Context, add users.Address, userid string) (string, error) {
	err := db.CreateAddress(ctx, &add, userid)
	return add.ID, err
}

func (s *fixedService) GetCards(ctx context.Context, id string) ([]users.Card, error) {
	if id == "" {
		cs, err := db.GetCards(ctx)
		for k, c := range cs {
			c.AddLinks()
			cs[k] = c
		}
		return cs, err
	}
	c, err := db.GetCard(ctx, id)
	c.AddLinks()
	return []users.Card{c}, err
}

func (s *fixedService) PostCard(ctx context.Context, card users.Card, userid string) (string, error) {
	err := db.CreateCard(ctx, &card, userid)
	return card.ID, err
}

func (s *fixedService) Delete(ctx context.Context, entity, id string) error {
	return db.Delete(ctx, entity, id)
}

func (s *fixedService) Health() []Health {
	var health []Health
	dbstatus := "OK"

	err := db.Ping()
	if err != nil {
		dbstatus = "err"
	}

	app := Health{"user", "OK", time.Now().String()}
	db := Health{"user-db", dbstatus, time.Now().String()}

	health = append(health, app)
	health = append(health, db)

	return health
}

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
