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
	Create(users.User) error
	GetByName(string) (users.User, error)
	GetByID(string) (users.User, error)
	GetAttributes(*users.User) error
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

func Create(u users.User) error {
	return DefaultDb.Create(u)
}

func GetByName(n string) (users.User, error) {
	return DefaultDb.GetByName(n)
}

func GetByID(n string) (users.User, error) {
	return DefaultDb.GetByID(n)
}

func GetAttributes(u *users.User) error {
	return DefaultDb.GetAttributes(u)
}
