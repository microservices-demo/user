package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	TestService Service
	TestUser    User = User{
		Name:     "testuser",
		Password: "testpassword",
	}
	TestUsers            []User = []User{TestUser}
	TestDomain                  = "testdomain"
	TestCustomerResponse        = customerResponse{
		Embedded: Wrapper{
			Customers: []customer{TestCustomer},
		},
	}
	TestCustomer = customer{Username: "testuser"}
)

func init() {
	TestService = NewFixedService(TestUsers, TestDomain)
}

func TestLogin(t *testing.T) {
	customerLookUpMock := func(w http.ResponseWriter, r *http.Request) {
		bs, err := json.Marshal(TestCustomerResponse)
		if err != nil {
			t.Error(err)
		}
		w.Write(bs)
	}
	ts := httptest.NewServer(http.HandlerFunc(customerLookUpMock))
	customerHost = ts.Listener.Addr().String()
	defer ts.Close()

	u, err := TestService.Login("testuser", "testpassword")
	if err != err {
		t.Error(err)
	}
	if TestUser.Name != u.Name {
		t.Error("login user does not equal test user")
	}
}

func TestRegister(t *testing.T) {
	resp := TestService.Register("newuser", "newpassword")
	if !resp {
		t.Error("expected true all the time, got false")
	}
}
