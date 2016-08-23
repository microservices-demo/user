package users

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddLinksAdd(t *testing.T) {
	domain = "mydomain"
	a := Address{ID: "test"}
	a.AddLinks()
	h := Href{"http://mydomain/addresses/test"}
	if !reflect.DeepEqual(a.Links["address"], h) {
		t.Error("expected equal address links")
	}

}

func TestAddLinksCard(t *testing.T) {
	domain = "mydomain"
	c := Card{ID: "test"}
	c.AddLinks()
	h := Href{"http://mydomain/cards/test"}
	if !reflect.DeepEqual(c.Links["card"], h) {
		t.Error("expected equal address links")
	}

}

func TestMaskCC(t *testing.T) {
	test1 := "1234567890"
	c := Card{LongNum: test1}
	c.MaskCC()
	test1comp := "******7890"
	if c.LongNum != test1comp {
		t.Errorf("Expected matching CC number %v received %v", test1comp, test1)
	}
}

func TestNew(t *testing.T) {
	u := New()
	if len(u.Addresses) != 0 && len(u.Cards) != 0 {
		t.Error("Expected zero length addresses and cards")
	}
}

func TestValidate(t *testing.T) {
	u := New()
	err := u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "FirstName") {
		t.Error("Expected missing first name error")
	}
	u.FirstName = "test"
	err = u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "LastName") {
		t.Error("Expected missing last name error")
	}
	u.LastName = "test"
	err = u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "Username") {
		t.Error("Expected missing username error")
	}
	u.Username = "test"
	err = u.Validate()
	if err.Error() != fmt.Sprintf(ErrMissingField, "Password") {
		t.Error("Expected missing password error")
	}
	u.Password = "test"
	err = u.Validate()
	if err != nil {
		t.Error(err)
	}
}

func TestMaskCCs(t *testing.T) {
	u := New()
	u.Cards = append(u.Cards, Card{LongNum: "abcdefg"})
	u.Cards = append(u.Cards, Card{LongNum: "hijklmnopqrs"})
	u.MaskCCs()
	if u.Cards[0].LongNum != "***defg" {
		t.Error("Card one CC not masked")
	}
	if u.Cards[1].LongNum != "********pqrs" {
		t.Error("Card two CC not masked")
	}
}
