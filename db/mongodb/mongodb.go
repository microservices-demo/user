package mongodb

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/microservices-demo/user/users"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	name                string
	password            string
	host                string
	db                  = "users"
	ErrInvalidHexID     = errors.New("Invalid Id Hex")
	ErrorSavingCardData = errors.New("There was a problem saving some card data")
	ErrorSavingAddrData = errors.New("There was a problem saving some address data")
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
	users.User `bson:",inline"`
	ID         bson.ObjectId   `bson:"_id"`
	AddressIDs []bson.ObjectId `bson:"addresses"`
	CardIDs    []bson.ObjectId `bson:"cards"`
}

func New() MongoUser {
	u := users.New()
	return MongoUser{
		User:       u,
		AddressIDs: make([]bson.ObjectId, 0),
		CardIDs:    make([]bson.ObjectId, 0),
	}
}

func (mu *MongoUser) AddUserIDs() {
	if mu.User.Addresses == nil {
		mu.User.Addresses = make([]users.Address, 0)
	}
	for _, id := range mu.AddressIDs {
		mu.User.Addresses = append(mu.User.Addresses, users.Address{
			ID: id.Hex(),
		})
	}
	if mu.User.Cards == nil {
		mu.User.Cards = make([]users.Card, 0)
	}
	for _, id := range mu.CardIDs {
		mu.User.Cards = append(mu.User.Cards, users.Card{ID: id.Hex()})
	}
	mu.User.UserID = mu.ID.Hex()
}

type MongoAddress struct {
	users.Address `bson:",inline"`
	ID            bson.ObjectId `bson:"_id"`
}
type MongoCard struct {
	users.Card `bson:",inline"`
	ID         bson.ObjectId `bson:"_id"`
}

func (m *Mongo) Init() error {
	u := getURL()
	var err error
	m.Session, err = mgo.Dial(u.String())
	if err != nil {
		return err
	}
	return m.EnsureIndexes()
}

func (m *Mongo) Create(u *users.User) error {
	s := m.Session.Copy()
	defer s.Close()
	id := bson.NewObjectId()
	mu := New()
	mu.User = *u
	mu.ID = id
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
	// Cheap err for attributes
	if carderr != nil || addrerr != nil {
		return fmt.Errorf("%v %v", carderr, addrerr)
	}
	u = &mu.User
	return nil
}

func (m *Mongo) createCards(cs []users.Card) ([]bson.ObjectId, error) {
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
		cs[k].ID = id.Hex()
	}
	return ids, nil
}

func (m *Mongo) createAddresses(as []users.Address) ([]bson.ObjectId, error) {
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
		as[k].ID = id.Hex()
	}
	return ids, nil
}

func (m *Mongo) cleanAttributes(mu MongoUser) error {
	s := m.Session.Copy()
	defer s.Close()
	c := s.DB("").C("addresses")
	_, err := c.RemoveAll(bson.M{"_id": bson.M{"$in": mu.AddressIDs}})
	return err
}

func (m *Mongo) GetByName(name string) (users.User, error) {
	s := m.Session.Copy()
	defer s.Close()
	c := s.DB("").C("customers")
	mu := New()
	err := c.Find(bson.M{"username": name}).One(&mu)
	mu.AddUserIDs()
	return mu.User, err
}

func (m *Mongo) GetByID(id string) (users.User, error) {
	s := m.Session.Copy()
	defer s.Close()
	if !bson.IsObjectIdHex(id) {
		return users.New(), errors.New("Invalid Id Hex")
	}
	c := s.DB("").C("customers")
	mu := New()
	err := c.FindId(bson.ObjectIdHex(id)).One(&mu)
	mu.AddUserIDs()
	return mu.User, err
}

func (m *Mongo) GetAttributes(u *users.User) error {
	s := m.Session.Copy()
	defer s.Close()
	ids := make([]bson.ObjectId, 0)
	for _, a := range u.Addresses {
		if !bson.IsObjectIdHex(a.ID) {
			return ErrInvalidHexID
		}
		ids = append(ids, bson.ObjectIdHex(a.ID))
	}
	var ma []MongoAddress
	c := s.DB("").C("addresses")
	err := c.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&ma)
	spew.Dump(ma)
	if err != nil {
		return err
	}
	na := make([]users.Address, 0)
	for _, a := range ma {
		a.Address.ID = a.ID.Hex()
		na = append(na, a.Address)
	}
	u.Addresses = na

	ids = make([]bson.ObjectId, 0)
	for _, c := range u.Cards {
		if !bson.IsObjectIdHex(c.ID) {
			return ErrInvalidHexID
		}
		ids = append(ids, bson.ObjectIdHex(c.ID))
	}
	var mc []MongoCard
	c = s.DB("").C("cards")
	err = c.Find(bson.M{"_id": bson.M{"$in": ids}}).All(&mc)
	if err != nil {
		return err
	}

	nc := make([]users.Card, 0)
	for _, ca := range mc {
		ca.Card.ID = ca.ID.Hex()
		nc = append(nc, ca.Card)
	}
	u.Cards = nc
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

func (m *Mongo) EnsureIndexes() error {
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
