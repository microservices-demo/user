package mongodb

import (
	"errors"
	"flag"
	"net/url"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/microservices-demo/user/users"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	name            string
	password        string
	host            string
	db              = "users"
	ErrInvalidHexID = errors.New("Invalid Id Hex")
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
	*users.User
	ID         bson.ObjectId   `bson:"_id"`
	AddressIDs []bson.ObjectId `bson:"addresses"`
	CardIDs    []bson.ObjectId `bson:"cards"`
}
type MongoAddress struct {
	users.Address
	ID bson.ObjectId `bson:"_id"`
}
type MongoCard struct {
	users.Card
	ID bson.ObjectId `bson:"_id"`
}

func (m Mongo) Init() error {
	u := getURL()
	var err error
	m.Session, err = mgo.Dial(u.String())
	if err != nil {
		return err
	}
	return m.EnsureIndexes()
}

func (m Mongo) Create(u *users.User) error {
	s := m.Session.Copy()
	defer s.Close()
	id := bson.NewObjectId()
	mu := MongoUser{User: u, ID: id}
	var carderr error
	var addrerr error
	mu.CardIDs, carderr = m.createCards(u.Cards)
	mu.AddressIDs, addrerr = m.createAddresses(u.Addresses)
	c := s.DB("").C("customers")
	_, err := c.UpsertId(mu.ID, mu)
	if err != nil {
		// Gonna clean up if we can, ignore error
		// because the user save error takes precedence.
		m.cleanAttributes(mu)
		return err
	}
	mu.User.UserID = mu.ID.Hex()
	spew.Dump(carderr)
	spew.Dump(addrerr)
	return nil
}

func (m Mongo) createCards(cs []users.Card) ([]bson.ObjectId, error) {
	s := m.Session.Copy()
	ids := make([]bson.ObjectId, 0)
	defer s.Close()
	for k, ca := range cs {
		id := bson.NewObjectId()
		mc := MongoCard{Card: ca, ID: id}
		c := s.DB("").C("cards")
		_, err := c.UpsertId(mc.ID, mc)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
		cs[k].ID = id.String()
	}
	return ids, nil
}

func (m Mongo) createAddresses(as []users.Address) ([]bson.ObjectId, error) {
	ids := make([]bson.ObjectId, 0)
	s := m.Session.Copy()
	defer s.Close()
	for k, a := range as {
		id := bson.NewObjectId()
		ma := MongoAddress{Address: a, ID: id}
		c := s.DB("").C("addresses")
		_, err := c.UpsertId(ma.ID, ma)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
		as[k].ID = id.String()
	}
	return ids, nil
}

func (m Mongo) cleanAttributes(mu MongoUser) error {
	s := m.Session.Copy()
	defer s.Close()
	c := s.DB("").C("addresses")
	_, err := c.RemoveAll(bson.M{"_id": bson.M{"$in": mu.AddressIDs}})
	return err
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

func (m Mongo) EnsureIndexes() error {
	s := m.Session.Copy()
	defer s.Close()
	i := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     false,
	}
	c := s.DB("").C("users")
	return c.EnsureIndex(i)
}
