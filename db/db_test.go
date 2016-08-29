package db

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/microservices-demo/user/tests"
	"github.com/microservices-demo/user/users"
)

var (
	ErrFakeError = tests.ErrFakeError
	TestAddress  = tests.TestAddress
	TestDB       = &tests.FakeDB{}
)

func init() {
	TestAddress.AddLinks()
}

func TestInit(t *testing.T) {
	err := Init()
	if err == nil {
		t.Error("Expected no registered db error")
	}
	database = "test"
	err = Init()
	if err.Error() != fmt.Sprintf(ErrNoDatabaseFound, "test") {
		t.Error("expected db not found from init")
	}
	Register("test", TestDB)
	database = "test"
	err = Init()
	if err != ErrFakeError {
		t.Error("expected FakeDB db error from init")
	}
	TestDB.PassInit = true
	err = Init()
	if err != nil {
		t.Error("expected no error from fake db")
	}
}

func TestSet(t *testing.T) {
	database = "nodb"
	err := Set()
	if err == nil {
		t.Error("Expecting error for no databade found")
	}
	Register("nodb2", TestDB)
	database = "nodb2"
	err = Set()
	if err != nil {
		t.Error(err)
	}
}

func TestRegister(t *testing.T) {
	l := len(DBTypes)
	Register("test2", TestDB)
	if len(DBTypes) != l+1 {
		t.Errorf("Expecting %v DB types received %v", l+1, len(DBTypes))
	}
	l = len(DBTypes)
	Register("test2", TestDB)
	if len(DBTypes) != l {
		t.Errorf("Expecting %v DB types received %v duplicate names", l, len(DBTypes))
	}
}

func TestCreateUser(t *testing.T) {
	err := CreateUser(&users.User{})
	if err != ErrFakeError {
		t.Error("expected FakeDB db error from create")
	}
}

func TestGetUser(t *testing.T) {
	u, err := GetUser("test")
	if err != ErrFakeError {
		t.Error("expected FakeDB db error from get")
	}
	u, err = GetUser("realuser")
	if err != nil {
		t.Error(err)
	}
	if len(u.Links) != 4 {
		t.Errorf("expected 4 links returned received %v", len(u.Links))
	}
}

func TestGetUserByName(t *testing.T) {
	u, err := GetUserByName("test")
	if err != ErrFakeError {
		t.Error("expected FakeDB db error from get")
	}
	u, err = GetUserByName("user")
	if err != nil {
		t.Error(err)
	}
	if len(u.Links) != 4 {
		t.Errorf("expected 4 links returned received %v", len(u.Links))
	}
}

func TestGetUserAttributes(t *testing.T) {
	u := users.New()
	GetUserAttributes(&u)
	if len(u.Addresses) != 1 {
		t.Error("expected one address added for GetUserAttributes")
	}
	if !reflect.DeepEqual(u.Addresses[0], TestAddress) {
		t.Error("expected matching addresses")
	}
}
