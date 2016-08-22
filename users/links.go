package users

import (
	"flag"
	"fmt"
	"os"
)

var (
	domain    string
	entitymap = map[string]string{
		"customer": "customers",
		"address":  "addresses",
		"card":     "cards",
	}
)

func init() {
	flag.StringVar(&domain, "link-domain", os.Getenv("HATEAOS"), "HATEAOS link domain")
}

type Links map[string]Href

func (l *Links) AddLink(ent string, id string) {
	nl := make(Links)
	link := fmt.Sprintf("http://%v/%v/%v", domain, entitymap[ent], id)
	nl[ent] = Href{link}
	nl["self"] = Href{link}
	*l = nl

}

func (l *Links) AddCustomer(id string) {
	l.AddLink("customer", id)
}

func (l *Links) AddAddress(id string) {
	l.AddLink("address", id)
}

func (l *Links) AddCard(id string) {
	l.AddLink("card", id)
}

func (l *Links) AddAttrLink(attr string, id string) {
	link := fmt.Sprintf("http://%v/%v/%v", domain, entitymap[attr], id)
	nl := *l
	nl[entitymap[attr]] = Href{link}
	*l = nl
}

func (l *Links) AddAttrAddress(id string) {
	l.AddAttrLink("address", id)
}

func (l *Links) AddAttrCard(id string) {
	l.AddAttrLink("card", id)
}

type Href struct {
	string `json:"href"`
}
