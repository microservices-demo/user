package db

import (
	"errors"
	"reflect"
	"testing"

	"github.com/microservices-demo/user/users"
)

var (
	TestDB       = fake{}
	ErrFakeError = errors.New("Fake error")
	TestAddress  = users.Address{
		Street:  "street",
		Number:  "51b",
		Country: "Netherlands",
		City:    "Amsterdam",
		ID:      "000056",
	}
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
		t.Error("expected fake db error from init")
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

func TestCreate(t *testing.T) {
	err := Create(&users.User{})
	if err != ErrFakeError {
		t.Error("expected fake db error from create")
	}
}

func TestGetById(t *testing.T) {
	_, err := GetUser("test")
	if err != ErrFakeError {
		t.Error("expected fake db error from get")
	}
}

func TestGetUserName(t *testing.T) {
	_, err := GetUserName("test")
	if err != ErrFakeError {
		t.Error("expected fake db error from get")
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

type fake struct{}

func (f fake) Init() error {
	return ErrFakeError
}
func (f fake) GetUserByName(name string) (users.User, error) {
	return users.User{}, ErrFakeError
}
func (f fake) GetUser(name string) (users.User, error) {
	return users.User{}, ErrFakeError
}
func (f fake) GetUserAttributes(u *users.User) error {
	u.Addresses = append(u.Addresses, TestAddress)
	return nil
}
func (f fake) CreateUser(*users.User) error {
	return ErrFakeError
}
