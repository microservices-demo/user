package api

// endpoints.go contains the endpoint definitions, including per-method request
// and response structs. Endpoints are the binding between the service and
// transport.

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/microservices-demo/user/users"
	"golang.org/x/net/context"
)

// Endpoints collects the endpoints that comprise the Service.
type Endpoints struct {
	LoginEndpoint       endpoint.Endpoint
	RegisterEndpoint    endpoint.Endpoint
	UserGetEndpoint     endpoint.Endpoint
	UserPostEndpoint    endpoint.Endpoint
	AddressGetEndpoint  endpoint.Endpoint
	AddressPostEndpoint endpoint.Endpoint
	CardGetEndpoint     endpoint.Endpoint
	CardPostEndpoint    endpoint.Endpoint
	HealthEndpoint      endpoint.Endpoint
}

// MakeEndpoints returns an Endpoints structure, where each endpoint is
// backed by the given service.
func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		LoginEndpoint:       MakeLoginEndpoint(s),
		RegisterEndpoint:    MakeRegisterEndpoint(s),
		HealthEndpoint:      MakeHealthEndpoint(s),
		UserGetEndpoint:     MakeUserGetEndpoint(s),
		UserPostEndpoint:    MakeUserPostEndpoint(s),
		AddressGetEndpoint:  MakeAddressGetEndpoint(s),
		AddressPostEndpoint: MakeAddressPostEndpoint(s),
		CardGetEndpoint:     MakeCardGetEndpoint(s),
		CardPostEndpoint:    MakeCardPostEndpoint(s),
	}
}

// MakeLoginEndpoint returns an endpoint via the given service.
func MakeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(loginRequest)
		u, err := s.Login(req.Username, req.Password)
		return userResponse{User: u}, err
	}
}

// MakeRegisterEndpoint returns an endpoint via the given service.
func MakeRegisterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(registerRequest)
		status := s.Register(req.Username, req.Password, req.Email)
		return statusResponse{Status: status}, err
	}
}

// MakeUserGetEndpoint returns an endpoint via the given service.
func MakeUserGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		users, err := s.GetUsers(req.ID)
		return usersResponse{Users: users}, err
	}
}

// MakeUserPostEndpoint returns an endpoint via the given service.
func MakeUserPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(userPostRequest)
		status := s.PostUser(req.User)
		return statusResponse{Status: status}, err
	}
}

// MakeAddressGetEndpoint returns an endpoint via the given service.
func MakeAddressGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		adds, err := s.GetAddresses(req.ID)
		return addressesResponse{Addresses: adds}, err
	}
}

// MakeAddressPostEndpoint returns an endpoint via the given service.
func MakeAddressPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(addressPostRequest)
		status := s.PostAddress(req.Address, req.UserID)
		return statusResponse{Status: status}, err
	}
}

// MakeUserGetEndpoint returns an endpoint via the given service.
func MakeCardGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRequest)
		cards, err := s.GetCards(req.ID)
		return cardsResponse{Cards: cards}, err
	}
}

// MakeCardPostEndpoint returns an endpoint via the given service.
func MakeCardPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(cardPostRequest)
		status := s.PostCard(req.Card, req.UserID)
		return statusResponse{Status: status}, err
	}
}

// MakeHealthEndpoint returns current health of the given service.
func MakeHealthEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return healthResponse{Status: "OK", Time: time.Now().String()}, nil
	}
}

type GetRequest struct {
	ID string
}

type loginRequest struct {
	Username string
	Password string
}

type userPostRequest struct {
	User users.User `json:"user"`
}

type userResponse struct {
	User users.User `json:"user"`
}

type usersResponse struct {
	Users []users.User `json:"customer"`
}

type addressPostRequest struct {
	Address users.Address `json:"address"`
	UserID  string        `json:"userID"`
}

type addressesResponse struct {
	Addresses []users.Address `json:"address"`
}

type cardPostRequest struct {
	Card   users.Card `json:"card"`
	UserID string     `json:"userID"`
}

type cardsResponse struct {
	Cards []users.Card `json:"card"`
}

type registerRequest struct {
	Username string
	Password string
	Email    string
}

type statusResponse struct {
	Status bool `json:"status"`
}

type healthRequest struct {
	//
}

type healthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}
