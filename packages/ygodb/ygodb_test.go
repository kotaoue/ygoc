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
	r := ygodb.Scraping("Harpie's Feather Duster", ygodb.LangEN)
	e := ygodb.Card{
		ID:        "4678",
		Name:      "Harpie's Feather Duster",
		Limited:   "Limited",
		Attribute: "SPELL",
		Text:      "Destroy all Spell and Trap Cards your opponent controls.",
	}
	if r != e {
		t.Fatalf("when keyword %s\nreturned %#v\nexpected %#v", k, r, e)
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
