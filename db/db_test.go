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
	_, err := GetByID("test")
	if err != ErrFakeError {
		t.Error("expected fake db error from get")
	}
}

func TestGetByName(t *testing.T) {
	_, err := GetByName("test")
	if err != ErrFakeError {
		t.Error("expected fake db error from get")
	}
}

func TestGetAttributes(t *testing.T) {
	u := users.New()
	GetAttributes(&u)
	if len(u.Addresses) != 1 {
		t.Error("expected one address added for GetAttributes")
	}
	if !reflect.DeepEqual(u.Addresses[0], TestAddress) {
		t.Error("expected matching addresses")
	}
}

type fake struct{}

func (f fake) Init() error {
	return ErrFakeError
}
func (f fake) GetByName(name string) (users.User, error) {
	return users.User{}, ErrFakeError
}
func (f fake) GetByID(name string) (users.User, error) {
	return users.User{}, ErrFakeError
}
func (f fake) GetAttributes(u *users.User) error {
	u.Addresses = append(u.Addresses, TestAddress)
	return nil
}
func (f fake) Create(*users.User) error {
	return ErrFakeError
}
