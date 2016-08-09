package mongodb

import (
	"testing"

	"gopkg.in/mgo.v2/dbtest"
)

var (
	TestMongo  = Mongo{}
	TestServer = dbtest.DBServer{}
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

func TestCreate(t *testing.T) {

}

func TestGet(t *testing.T) {

}
