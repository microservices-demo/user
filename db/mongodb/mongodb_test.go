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

func TestMain(m *testing.M) {
	TestServer.SetPath("/tmp")
	TestMongo.Session = TestServer.Session()
	defer exitTest()
	m.Run()
}

func exitTest() {
	TestMongo.Session.Close()
	TestServer.Stop()
}

func TestCreate(t *testing.T) {
	err := TestMongo.Create(&TestUser)
	if err != nil {
		t.Error(err)
	}
}

func TestGetByName(t *testing.T) {

}
func TestGetByID(t *testing.T) {

}
func TestGetAttributes(t *testing.T) {

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
