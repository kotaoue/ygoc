package ygodb_test

import (
	"testing"

	"github.com/kotaoue/ygolinker/packages/ygodb"
)

func Test_ExtractValue(t *testing.T) {
	s := "ATK 3000"
	r := ygodb.ExtractValue(s)
	e := "3000"
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}

	s = "DEF -"
	r = ygodb.ExtractValue(s)
	e = "-"
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}
}
