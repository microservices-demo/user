package db

import (
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
	Get(string) (users.User, error)
}

var (
	database           string
	DefaultDb          Database
	DBTypes            map[string]Database = map[string]Database{}
	ErrNoDatabaseFound                     = "No database with name %v registered"
)

func init() {
	flag.StringVar(&database, "database", os.Getenv("USER_DATABASE"), "Database to use, Mongolar or ...")
}

func Init() error {
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

func Get(n string) (users.User, error) {
	return DefaultDb.Get(n)
}
