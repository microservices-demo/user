package register

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/microservices-demo/user/db"
	"github.com/microservices-demo/user/login"
	"github.com/microservices-demo/user/users"
)

type Registration struct {
	Address users.Address `json:"address"`
	Card    users.Card    `json:"card"`
	User    users.User    `json:"customer"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var reg Registration
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reg)
	if err != nil {
		log.Error(err)
		w.WriteHeader(400)
		return
	}
	reg.User.Addresses = []users.Address{reg.Address}
	reg.User.Cards = []users.Card{reg.Card}
	err = db.Create(reg.User)
	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	l := login.NewLoginClaims(reg.User)
	log.Debug("Customer id: %s\n", l.Id)
	signed, err := l.GetToken()

	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("WeaveToken", signed)
}
