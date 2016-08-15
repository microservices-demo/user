package api

// transport.go contains the binding from endpoints to a concrete transport.
// In our case we just use a REST-y HTTP transport.

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type EmbedStruct struct {
	Embed interface{} `json:"_embedded"`
}

// MakeHTTPHandler mounts the endpoints into a REST-y HTTP handler.
func MakeHTTPHandler(ctx context.Context, e Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter().StrictSlash(false)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// GET /login       Login
	// GET /register    Register
	// GET /health      Health Check

	r.Methods("GET").Path("/login").Handler(httptransport.NewServer(
		ctx,
		e.LoginEndpoint,
		decodeLoginRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/register").Handler(httptransport.NewServer(
		ctx,
		e.RegisterEndpoint,
		decodeRegisterRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").PathPrefix("/customers").Handler(httptransport.NewServer(
		ctx,
		e.UserGetEndpoint,
		decodeGetRequest,
		encodeEmbedResponse,
		options...,
	))
	r.Methods("GET").PathPrefix("/cards").Handler(httptransport.NewServer(
		ctx,
		e.CardGetEndpoint,
		decodeGetRequest,
		encodeEmbedResponse,
		options...,
	))
	r.Methods("GET").PathPrefix("/addresses").Handler(httptransport.NewServer(
		ctx,
		e.AddressGetEndpoint,
		decodeGetRequest,
		encodeEmbedResponse,
		options...,
	))
	r.Methods("POST").Path("/customers").Handler(httptransport.NewServer(
		ctx,
		e.UserPostEndpoint,
		decodeUserRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/addresses").Handler(httptransport.NewServer(
		ctx,
		e.AddressPostEndpoint,
		decodeAddressRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/cards").Handler(httptransport.NewServer(
		ctx,
		e.CardPostEndpoint,
		decodeCardRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").PathPrefix("/health").Handler(httptransport.NewServer(
		ctx,
		e.HealthEndpoint,
		decodeHealthRequest,
		encodeHealthResponse,
		options...,
	))
	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	switch err {
	case ErrUnauthorized:
		code = http.StatusUnauthorized
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":       err.Error(),
		"status_code": code,
		"status_text": http.StatusText(code),
	})
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	u, p, ok := r.BasicAuth()
	if !ok {
		return nil, ErrUnauthorized
	}

	return loginRequest{
		Username: u,
		Password: p,
	}, nil
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	u := r.FormValue("username")
	p := r.FormValue("password")
	e := r.FormValue("email")

	return registerRequest{
		Username: u,
		Password: p,
		Email:    e,
	}, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	g := GetRequest{}
	u := strings.Split(r.URL.Path, "/")
	if len(u) == 3 {
		g.ID = u[2]
	}
	return g, nil
}

func decodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	u := userPostRequest{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func decodeAddressRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	a := addressPostRequest{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func decodeCardRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	c := cardPostRequest{}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func decodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func encodeHealthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return encodeResponse(ctx, w, response.(healthResponse))
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// All of our response objects are JSON serializable, so we just do that.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeEmbedResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	// All of our response objects are JSON serializable, so we just do that.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	er := EmbedStruct{Embed: response}
	return json.NewEncoder(w).Encode(er)
}
