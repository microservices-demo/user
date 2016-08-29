package api

import (
	"testing"

	"github.com/microservices-demo/user/tests"
	"github.com/microservices-demo/user/users"
)

var (
	TestFixedService Service
	TestDB           = &tests.FakeDB{}
	ErrFakeError     = tests.ErrFakeError
)

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
	_, err := TestFixedService.Register("test", "myemail@here.com", "password")
	if err == nil {
		t.Error("expected err for register")
	}
	_, err = TestFixedService.Register("passtest", "myemail@here.com", "password")
	if err != nil {
		t.Error(err)
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
	_, err := TestFixedService.PostUser(u)
	if err == nil {
		t.Error("expected err for post user")
	}
	u.Username = "passtest"
	_, err = TestFixedService.PostUser(u)
	if err != nil {
		t.Error(err)
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
	_, err := TestFixedService.PostAddress(a, "")
	if err == nil {
		t.Error("expected err for post address")
	}
	a.Street = "passtest"
	_, err = TestFixedService.PostAddress(a, "")
	if err != nil {
		t.Error(err)
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
	_, err := TestFixedService.PostCard(c, "")
	if err == nil {
		t.Error("expected err for post card")
	}
	c.LongNum = "passtest"
	_, err = TestFixedService.PostCard(c, "")
	if err != nil {
		t.Error(err)
	}
}

func TestCalculatePassHash(t *testing.T) {
	hash1 := calculatePassHash("eve", "c748112bc027878aa62812ba1ae00e40ad46d497")
	if hash1 != "fec51acb3365747fc61247da5e249674cf8463c2" {
		t.Error("Eve's password failed hash test")
	}
	hash2 := calculatePassHash("password", "6c1c6176e8b455ef37da13d953df971c249d0d8e")
	if hash2 != "e2de7202bb2201842d041f6de201b10438369fb8" {
		t.Error("user's password failed hash test")
	}
	hash3 := calculatePassHash("password", "bd832b0e10c6882deabc5e8e60a37689e2b708c2")
	if hash3 != "8f31df4dcc25694aeb0c212118ae37bbd6e47bcd" {
		t.Error("user1's password failed hash test")
	}
}
