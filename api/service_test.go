package api

import (
	"testing"

	"github.com/microservices-demo/user/tests"
	"github.com/microservices-demo/user/users"
)

var (
	TestService      Service
	TestFixedService Service
	TestDB           = tests.FakeDB{}
	ErrFakeError     = tests.ErrFakeError
)

type TestServiceStruct struct{}

func (s *TestServiceStruct) Login(username, password string) (users.User, error) {
	return users.User{}, nil
}

func (s *TestServiceStruct) Register(username, password, email string) bool {
	return false
}

func (s *TestServiceStruct) GetUsers(id string) ([]users.User, error) {
	return make([]users.User, 0), nil
}

func (s *TestServiceStruct) PostUser(user users.User) bool {
	return false
}

func (s *TestServiceStruct) GetAddresses(id string) ([]users.Address, error) {
	return make([]users.Address, 0), nil
}

func (s *TestServiceStruct) PostAddress(add users.Address, userid string) bool {
	return false
}

func (s *TestServiceStruct) GetCards(id string) ([]users.Card, error) {
	return make([]users.Card, 0), nil
}

func (s *TestServiceStruct) PostCard(card users.Card, userid string) bool {
	return false
}

func init() {
	TestService = &TestServiceStruct{}
}

func TestNewFixedService(t *testing.T) {
	TestFixedService = NewFixedService(TestDB)

}

func TestFSLogin(t *testing.T) {
	_, err := TestFixedService.Login("test", "pass")
	if err != ErrFakeError {
		t.Error("expected fake error for not found")
	}
	_, err = TestFixedService.Login("user", "pass")
	if err != ErrUnauthorized {
		t.Error("expected unauthorized")
	}
	_, err = TestFixedService.Login("user", "test2")
	if err != nil {
		t.Error(err)
	}

}

func TestFSRegister(t *testing.T) {
	b := TestFixedService.Register("test", "myemail@here.com", "password")
	if b {
		t.Error("expected false for register")
	}
	b = TestFixedService.Register("passtest", "myemail@here.com", "password")
	if !b {
		t.Error("expected true for register")
	}
}

func TestFSGetUsers(t *testing.T) {
	u, err := TestFixedService.GetUsers("")
	if len(u) > 0 {
		t.Error("expected 0 users")
	}
	u, err = TestFixedService.GetUsers("fakeuser")
	if err != ErrFakeError {
		t.Error("expected fake error for get  users")
	}
	u, err = TestFixedService.GetUsers("realuser")
	if err != nil {
		t.Error("expected fake error for get  users")
	}
	if len(u) != 1 {
		t.Error("expected one user returned")
	}
}

func TestFSPostUser(t *testing.T) {
	u := users.New()
	u.Username = "fakeuser"
	b := TestFixedService.PostUser(u)
	if b {
		t.Error("expected false for post user")
	}
	u.Username = "passtest"
	b = TestFixedService.PostUser(u)
	if !b {
		t.Error("expected true for post user")
	}
}

func TestFSGetAddresses(t *testing.T) {

}
