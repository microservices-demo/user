package api

/// needs actual tests

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/microservices-demo/user/tests"
)

var (
	TestService   = &tests.TestServiceStruct{}
	TestEndpoints Endpoints
	Ctx           = context.Background()
)

func TestMakeEndpoints(t *testing.T) {
	TestEndpoints = MakeEndpoints(NewFixedService(tests.FakeDB{}))
}

func TestLoginEndpointEndpoint(t *testing.T) {
	_, err := TestEndpoints.LoginEndpoint(Ctx, loginRequest{"test", "pass"})
	if err != ErrFakeError {
		t.Error("expected fake error for not found")
	}
	_, err = TestEndpoints.LoginEndpoint(Ctx, loginRequest{"user", "pass"})
	if err != ErrUnauthorized {
		t.Error("expected unauthorized")
	}
	_, err = TestEndpoints.LoginEndpoint(Ctx, loginRequest{"user", "test2"})
	if err != nil {
		t.Error(err)
	}
}

func TestMakeRegisterEndpointEndpoint(t *testing.T) {
	b, _ := TestEndpoints.RegisterEndpoint(Ctx, registerRequest{"test", "myemail@here.com", "password"})
	v := b.(statusResponse)
	if v.Status {
		t.Error("expected false for register")
	}
	b, _ = TestEndpoints.RegisterEndpoint(Ctx, registerRequest{"passtest", "myemail@here.com", "password"})
	v = b.(statusResponse)
	if !v.Status {
		t.Error("expected true for register")
	}
}

func TestMakeUserGetEndpoint(t *testing.T) {
}

func TestMakeUserPostEndpoint(t *testing.T) {
}

func TestMakeAddressGetEndpoint(t *testing.T) {
}

func TestMakeAddressPostEndpoint(t *testing.T) {
}

func TestMakeCardGetEndpoint(t *testing.T) {
}

func TestMakeCardPostEndpoint(t *testing.T) {
}
