package ygodb_test

import (
	"testing"

	"github.com/kotaoue/ygoc/packages/ygodb"
)

func Test_Card_URL(t *testing.T) {
	c := ygodb.Card{ID: "123456789"}
	r := c.URL()
	e := "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=123456789"
	if r != e {
		t.Fatalf("when Card.ID %s, returned %s, expected %s", c.ID, r, e)
	}
}

func Test_Scraping(t *testing.T) {
	k := "Harpie's Feather Duster"
	r, err := ygodb.Scraping(k, ygodb.LangEN)
	e := ygodb.Card{
		ID:        "4678",
		Name:      "Harpie's Feather Duster",
		Limited:   "Limited",
		Attribute: "SPELL",
		Text:      "Destroy all Spell and Trap Cards your opponent controls.",
	}
	if err != nil {
		t.Skipf("when keyword %s error occurred %s\n", k, err.Error())
	}
	if r != e {
		t.Fatalf("when keyword %s\nreturned %#v\nexpected %#v", k, r, e)
	}

	k = "Raigeki"
	r, err = ygodb.Scraping(k, ygodb.LangEN)
	e = ygodb.Card{
		ID:        "4343",
		Name:      "Raigeki",
		Limited:   "Limited",
		Attribute: "SPELL",
		Text:      "Destroy all monsters your opponent controls.",
	}
	if err != nil {
		t.Skipf("when keyword %s error occurred %s\n", k, err.Error())
	}
	if r != e {
		t.Fatalf("when keyword %s\nreturned %#v\nexpected %#v", k, r, e)
	}

	k = "Red-Eyes"
	r, err = ygodb.Scraping(k, ygodb.LangEN)
	s := "Error: Couldn't narrow down the cards to one."
	if err == nil {
		t.Fatalf("when keyword %s error not occurred\n", k)
	}
	if err.Error() != s {
		t.Fatalf("when keyword %s returned %s expected %s", k, err.Error(), s)
	}

	k = "Shivan Dragon"
	r, err = ygodb.Scraping(k, ygodb.LangEN)
	s = "Error: Card not found."
	if err == nil {
		t.Fatalf("when keyword %s error not occurred\n", k)
	}
	if err.Error() != s {
		t.Fatalf("when keyword %s returned %s expected %s", k, err.Error(), s)
	}
}

func Test_ExtractCardID(t *testing.T) {
	s := "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=13049"
	r := ygodb.ExtractCardID(s)
	e := "13049"
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}

	s = "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=1&sess=1&keyword=Gouki+Suprex&stype=1&ctype=&starfr=&starto=&pscalefr=&pscaleto=&linkmarkerfr=&linkmarkerto=&link_m=2&atkfr=&atkto=&deffr=&defto=&othercon=2&request_locale=en"
	r = ygodb.ExtractCardID(s)
	e = ""
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}
}

func Test_ExtractValue(t *testing.T) {
	s := "ATK 3000"
	r := ygodb.ExtractValue(s)
	e := "3000"
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}

	s = "DEF -"
	r = ygodb.ExtractValue(s)
	e = ""
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}
}
