package db

import (
	"reflect"
	"testing"

	"github.com/microservices-demo/user/tests"
	"github.com/microservices-demo/user/users"
)

var (
	ErrFakeError = tests.ErrFakeError
	TestAddress  = tests.TestAddress
	TestDB       = tests.FakeDB{}
)

func TestInit(t *testing.T) {
	err := Init()
	if err == nil {
		t.Error("Expected no registered db error")
	}
	Register("test", TestDB)
	database = "test"
	err = Init()
	if err != ErrFakeError {
		t.Error("expected FakeDB db error from init")
	}
	TestAddress.AddLinks()
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
	_, err := GetUser("test")
	if err != ErrFakeError {
		t.Error("expected FakeDB db error from get")
	}
}

func TestGetUserByName(t *testing.T) {
	_, err := GetUserByName("test")
	if err != ErrFakeError {
		t.Error("expected FakeDB db error from get")
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
