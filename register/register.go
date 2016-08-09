package register

import (
	"encoding/json"
	"net/http"

	"../accounts"
	"../login"
	log "github.com/Sirupsen/logrus"
)

type Registration struct {
	Address  *accounts.Address  `json:"address"`
	Card     *accounts.Card     `json:"card"`
	Customer *accounts.Customer `json:"customer"`
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
	err = reg.Address.Create()
	return
	if err != nil {
		log.Error(err)
		w.WriteHeader(400)
		return
	}
	reg.Customer.Addresses = []string{reg.Address.Links.Self.Href}
	err = reg.Card.Create()
	if err != nil {
		log.Error(err)
		w.WriteHeader(400)
		return
	}
	reg.Customer.Cards = []string{reg.Card.Links.Self.Href}
	err = reg.Customer.Create()
	if err != nil {
		log.Error(err)
		w.WriteHeader(400)
		return
	}
	l := login.NewLoginClaims(*reg.Customer)
	log.Debug("Customer id: %s\n", l.Id)
	signed, err := l.GetToken()

	if err != nil {
		log.Error(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("WeaveToken", signed)
}
