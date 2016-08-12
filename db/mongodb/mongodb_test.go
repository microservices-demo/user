package mongodb

import (
	"testing"

	"github.com/microservices-demo/user/users"

	"gopkg.in/mgo.v2/dbtest"
)

var (
	TestMongo  = Mongo{}
	TestServer = dbtest.DBServer{}
	TestUser   = users.User{
		FirstName: "firstname",
		LastName:  "lastname",
		Username:  "username",
		Password:  "blahblah",
		Addresses: []users.Address{
			users.Address{
				Street: "street",
			},
		},
	}
)

func init() {
	TestServer.SetPath("/tmp")
}

func TestMain(m *testing.M) {
	TestMongo.Session = TestServer.Session()
	TestMongo.EnsureIndexes()
	TestMongo.Session.Close()
	defer exitTest()
	m.Run()
}

func exitTest() {
	TestServer.Wipe()
	TestServer.Stop()
}

func TestCreate(t *testing.T) {
	TestMongo.Session = TestServer.Session()
	defer TestMongo.Session.Close()
	err := TestMongo.CreateUser(&TestUser)
	if err != nil {
		t.Error(err)
	}
	err = TestMongo.CreateUser(&TestUser)
	if err == nil {
		t.Error("Expected duplicate key error")
	}
}

func TestGetUserByName(t *testing.T) {
	TestMongo.Session = TestServer.Session()
	defer TestMongo.Session.Close()
	u, err := TestMongo.GetUserByName(TestUser.Username)
	if err != nil {
		t.Error(err)
	}
	if u.Username != TestUser.Username {
		t.Error("expected equal usernames")
	}
	_, err = TestMongo.GetUserByName("bogususers")
	if err == nil {
		t.Error("expected not found error")
	}
}

func TestGetUser(t *testing.T) {
	TestMongo.Session = TestServer.Session()
	defer TestMongo.Session.Close()

}
func TestGetUserAttributes(t *testing.T) {
	TestMongo.Session = TestServer.Session()
	defer TestMongo.Session.Close()

}
func TestGetURL(t *testing.T) {
	name = "test"
	password = "password"
	host = "thishostshouldnotexist:3038"
	u := getURL()
	if u.String() != "mongodb://test:password@thishostshouldnotexist:3038/users" {
		t.Error("expected url mismatch")
	}
}

func TestInit(t *testing.T) {
	err := TestMongo.Init()
	if err.Error() != "no reachable servers" {
		t.Error("expecting no reachable servers error")
	}
}
