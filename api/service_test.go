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
		t.Error("expected no error for get  users")
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
	a, err := TestFixedService.GetAddresses("")
	if len(a) > 0 {
		t.Error("expected 0 users")
	}
	a, err = TestFixedService.GetAddresses("fakeuser")
	if err != ErrFakeError {
		t.Error("expected fake error for get addresses")
	}
	a, err = TestFixedService.GetAddresses("realaddress")
	if err != nil {
		t.Error("expected no error for get addresses")
	}
	if len(a) != 1 {
		t.Error("expected one address returned")
	}
}

func TestFSPostAddress(t *testing.T) {
	a := users.Address{}
	a.Street = "fakeaddr"
	b := TestFixedService.PostAddress(a, "")
	if b {
		t.Error("expected false for post address")
	}
	a.Street = "passtest"
	b = TestFixedService.PostAddress(a, "")
	if !b {
		t.Error("expected true for post address")
	}
}

func TestFSGetCards(t *testing.T) {
	c, err := TestFixedService.GetCards("")
	if len(c) > 0 {
		t.Error("expected 0 cards")
	}
	c, err = TestFixedService.GetCards("fakecard")
	if err != ErrFakeError {
		t.Error("expected fake error for get cards")
	}
	c, err = TestFixedService.GetCards("realcard")
	if err != nil {
		t.Error("expected no error for get cards")
	}
	if len(c) != 1 {
		t.Error("expected one card returned")
	}
}

func TestFSPostCard(t *testing.T) {
	c := users.Card{}
	c.LongNum = "fakecard"
	b := TestFixedService.PostCard(c, "")
	if b {
		t.Error("expected false for post card")
	}
	c.LongNum = "passtest"
	b = TestFixedService.PostCard(c, "")
	if !b {
		t.Error("expected true for post card")
	}
}
