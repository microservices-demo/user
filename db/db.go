package db

import (
	"context"
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
	Delete(string, string) error
	CreateCard(*users.Card, string) error
	Ping() error
}

var (
	database string
	//DefaultDb is the database set for the microservice
	DefaultDb TracingMiddleware
	//DBTypes is a map of DB interfaces that can be used for this service
	DBTypes = map[string]Database{}
	//ErrNoDatabaseFound error returnes when database interface does not exists in DBTypes
	ErrNoDatabaseFound = "No database with name %v registered"
	//ErrNoDatabaseSelected is returned when no database was designated in the flag or env
	ErrNoDatabaseSelected = errors.New("No DB selected")
)

func init() {
	flag.StringVar(&database, "database", os.Getenv("USER_DATABASE"), "Database to use, Mongodb or ...")

}

//Init inits the selected DB in DefaultDb
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

//Set the DefaultDb
func Set() error {
	if v, ok := DBTypes[database]; ok {
		DefaultDb = DbTracingMiddleware()(v)
		return nil
	}
	return fmt.Errorf(ErrNoDatabaseFound, database)
}

//Register registers the database interface in the DBTypes
func Register(name string, db Database) {
	DBTypes[name] = db
}

//CreateUser invokes DefaultDb method
func CreateUser(ctx context.Context, u *users.User) error {
	return DefaultDb.CreateUser(ctx, u)
}

//GetUserByName invokes DefaultDb method
func GetUserByName(ctx context.Context, n string) (users.User, error) {
	u, err := DefaultDb.GetUserByName(ctx, n)
	if err == nil {
		u.AddLinks()
	}
	return u, err
}

//GetUser invokes DefaultDb method
func GetUser(ctx context.Context, n string) (users.User, error) {
	u, err := DefaultDb.GetUser(ctx, n)
	if err == nil {
		u.AddLinks()
	}
	return u, err
}

//GetUsers invokes DefaultDb method
func GetUsers(ctx context.Context) ([]users.User, error) {
	us, err := DefaultDb.GetUsers(ctx)
	for k, _ := range us {
		us[k].AddLinks()
	}
	return us, err
}

//GetUserAttributes invokes DefaultDb method
func GetUserAttributes(ctx context.Context, u *users.User) error {
	err := DefaultDb.GetUserAttributes(ctx, u)
	if err != nil {
		return err
	}
	for k, _ := range u.Addresses {
		u.Addresses[k].AddLinks()
	}
	for k, _ := range u.Cards {
		u.Cards[k].AddLinks()
	}
	return nil
}

//CreateAddress invokes DefaultDb method
func CreateAddress(ctx context.Context, a *users.Address, userid string) error {
	return DefaultDb.CreateAddress(ctx, a, userid)
}

//GetAddress invokes DefaultDb method
func GetAddress(ctx context.Context, n string) (users.Address, error) {
	a, err := DefaultDb.GetAddress(ctx, n)
	if err == nil {
		a.AddLinks()
	}
	return a, err
}

//GetAddresses invokes DefaultDb method
func GetAddresses(ctx context.Context) ([]users.Address, error) {
	as, err := DefaultDb.GetAddresses(ctx)
	for k, _ := range as {
		as[k].AddLinks()
	}
	return as, err
}

//CreateCard invokes DefaultDb method
func CreateCard(ctx context.Context, c *users.Card, userid string) error {
	return DefaultDb.CreateCard(ctx, c, userid)
}

//GetCard invokes DefaultDb method
func GetCard(ctx context.Context, n string) (users.Card, error) {
	return DefaultDb.GetCard(ctx, n)
}

//GetCards invokes DefaultDb method
func GetCards(ctx context.Context) ([]users.Card, error) {
	cs, err := DefaultDb.GetCards(ctx)
	for k, _ := range cs {
		cs[k].AddLinks()
	}
	return cs, err
}

//Delete invokes DefaultDb method
func Delete(ctx context.Context, entity, id string) error {
	return DefaultDb.Delete(ctx, entity, id)
}

//Ping invokes DefaultDB method
func Ping() error {
	return DefaultDb.Ping()
}
