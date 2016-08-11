package api

import (
	"testing"

	"github.com/microservices-demo/user/users"
)

var (
	TestService  Service
	TestCustomer = users.User{Username: "testuser", Password: ""}
)

func init() {
	TestService = NewFixedService()
}

func TestLogin(t *testing.T) {

}

func TestRegister(t *testing.T) {

}
