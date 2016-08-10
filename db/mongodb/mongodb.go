package mongodb

import (
	"flag"
	"net/url"
	"os"

	"github.com/microservices-demo/user/users"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	name     string
	password string
	host     string
	db       = "users"
)

func init() {
	flag.StringVar(&name, "mongo-user", os.Getenv("MONGO_USER"), "Mongo user")
	flag.StringVar(&password, "mongo-password", os.Getenv("MONGO_PASS"), "Mongo password")
	flag.StringVar(&host, "mongo-host", os.Getenv("MONGO_HOST"), "Mongo host")
}

type Mongo struct {
	Session *mgo.Session
}

type MongoUser struct {
	users.User
	Id         bson.ObjectId   `bson:"_id"`
	AddressIds []bson.ObjectId `bson:"addresses"`
	CardIds    []bson.ObjectId `bson:"cards"`
}
type MongoAddress struct {
	users.Address
	Id bson.ObjectId `bson:"_id"`
}
type MongoCard struct {
	users.Card
	Id bson.ObjectId `bson:"_id"`
}

func (m Mongo) Init() error {
	u := getURL()
	var err error
	m.Session, err = mgo.Dial(u.String())
	return err
}

func (m Mongo) Create(u users.User) error {
	return nil
}

func (m Mongo) GetByName(name string) (users.User, error) {
	return users.User{}, nil
}

func (m Mongo) GetByID(name string) (users.User, error) {
	return users.User{}, nil
}

func (m Mongo) GetAttributes(u *users.User) error {
	return nil
}

func getURL() url.URL {
	u := url.UserPassword(name, password)
	return url.URL{
		Scheme: "mongodb",
		User:   u,
		Host:   host,
		Path:   db,
	}
}
