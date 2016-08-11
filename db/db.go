package db

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/microservices-demo/user/users"
)

// Database represents a simple interface so we can switch to a new system easily
// this is just basic and specific to this microservice
type Database interface {
	Init() error
	GetUserByName(string) (users.User, error)
	GetUser(string) (users.User, error)
	GetUsers() ([]users.User, error)
	CreateUser(*users.User) error
	GetUserAttributes(*users.User) error
	GetAddress(string) (users.Address, error)
	GetAddresses() ([]users.Address, error)
	CreateAddress(*users.Address, string) error
	GetCard(string) (users.Card, error)
	GetCards() ([]users.Card, error)
	CreateCard(*users.Card, string) error
}

var (
	database              string
	DefaultDb             Database
	DBTypes               map[string]Database = map[string]Database{}
	ErrNoDatabaseFound                        = "No database with name %v registered"
	ErrNoDatabaseSelected                     = errors.New("No DB selected")
)

func init() {
	flag.StringVar(&database, "database", os.Getenv("USER_DATABASE"), "Database to use, Mongodb or ...")

}

func Init() error {
	if database == "" {
		return ErrNoDatabaseSelected
	}
	err := Set()
	if err != nil {
		return err
	}
	return DefaultDb.Init()
}

func Set() error {
	if v, ok := DBTypes[database]; ok {
		DefaultDb = v
		return nil
	}
	return fmt.Errorf(ErrNoDatabaseFound, database)
}

func Register(name string, db Database) {
	DBTypes[name] = db
}

func CreateUser(u *users.User) error {
	return DefaultDb.CreateUser(u)
}

func GetUserByName(n string) (users.User, error) {
	return DefaultDb.GetUserByName(n)
}

func GetUser(n string) (users.User, error) {
	return DefaultDb.GetUser(n)
}

func GetUsers() ([]users.User, error) {
	return DefaultDb.GetUsers()
}

func GetUserAttributes(u *users.User) error {
	return DefaultDb.GetUserAttributes(u)
}

func CreateAddress(a *users.Address, userid string) error {
	return DefaultDb.CreateAddress(a, userid)
}

func GetAddress(n string) (users.Address, error) {
	return DefaultDb.GetAddress(n)
}

func GetAddresses() ([]users.Address, error) {
	return DefaultDb.GetAddresses()
}

func CreateCard(c *users.Card, userid string) error {
	return DefaultDb.CreateCard(c, userid)
}

func GetCard(n string) (users.Card, error) {
	return DefaultDb.GetCard(n)
}

func GetCards() ([]users.Card, error) {
	return DefaultDb.GetCards()
}
