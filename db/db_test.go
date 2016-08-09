package db

import (
	"errors"
	"testing"

	"github.com/microservices-demo/user/users"
)

var (
	TestDB       = fake{}
	ErrFakeError = errors.New("Fake error")
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
	err := Create(users.User{})
	if err != ErrFakeError {
		t.Error("expected fake db error from create")
	}
}

func TestGet(t *testing.T) {
	_, err := Get("test")
	if err != ErrFakeError {
		t.Error("expected fake db error from get")
	}
}

type fake struct{}

func (f fake) Init() error {
	return ErrFakeError
}
func (f fake) Get(name string) (users.User, error) {
	return users.User{}, ErrFakeError
}
func (f fake) Create(users.User) error {
	return ErrFakeError
}
