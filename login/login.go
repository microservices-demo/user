package login

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"strings"

	"../accounts"
	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey   = []byte("test")
	passwordSalt = "passwordsalt"
)

func SigningKey(key string) {
	signingKey = []byte(key)
}

func PasswordSalt(salt string) {
	passwordSalt = salt
}

type Login struct {
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	u, p, ok := r.BasicAuth()
	if !ok {
		log.Info("No Authorization header present.\n")
		w.WriteHeader(401)
		return
	}
	log.Debug("Lookup for user %s and password: %s.\n", u, p)

	c, err := accounts.GetCustomerByName(u)
	if err != nil {
		log.Error(err)
		w.WriteHeader(401)
		return
	}
	if c.Password != calculatePassHash(p) {
		log.Info("User not authorized.\n")
		w.WriteHeader(401)
		return
	}
	lc := NewLoginClaims(c)
	log.Debug("Customer id: %s\n", lc.Id)
	signed, err := lc.GetToken()
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("WeaveToken", signed)
}

type LoginClaims struct {
	Username string `json:"username"`
	Customer string `json:"customer"`
	Id       string `json:"id"`
	jwt.StandardClaims
}

func NewLoginClaims(c accounts.Customer) LoginClaims {
	custLink := c.Links.Customer.Href
	idSplit := strings.Split(custLink, "/")
	id := idSplit[len(idSplit)-1]
	return LoginClaims{
		Username: c.Username,
		Customer: custLink,
		Id:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "WeaveDemo",
		},
	}

}

func (lc *LoginClaims) GetToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, lc)
	return token.SignedString(signingKey)
}

func calculatePassHash(pass string) string {
	h := sha1.New()
	io.WriteString(h, passwordSalt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
