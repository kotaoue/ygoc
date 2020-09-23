package md_test

import (
	"testing"

	"github.com/kotaoue/ygoc/packages/md"
)

func Test_IsList(t *testing.T) {
	c := "- Test"
	r := md.IsList(c)
	e := true
	if r != e {
		t.Fatalf("when %s, returned %t, expected %t", c, r, e)
	}

	c = "* Test"
	r = md.IsList(c)
	e = true
	if r != e {
		t.Fatalf("when %s, returned %t, expected %t", c, r, e)
	}

	c = "- [ ] Test"
	r = md.IsList(c)
	e = true
	if r != e {
		t.Fatalf("when %s, returned %t, expected %t", c, r, e)
	}

	c = "- [x] Test"
	r = md.IsList(c)
	e = true
	if r != e {
		t.Fatalf("when %s, returned %t, expected %t", c, r, e)
	}

	c = "Test"
	r = md.IsList(c)
	e = false
	if r != e {
		t.Fatalf("when %s, returned %t, expected %t", c, r, e)
	}
}

func Test_ListText(t *testing.T) {
	c := "- Crusadia Avramax"
	r := md.ListText(c)
	e := "Crusadia Avramax"
	if r != e {
		t.Fatalf("when %s, returned %s, expected %s", c, r, e)
	}

	c = "* Crusadia Avramax"
	r = md.ListText(c)
	e = "Crusadia Avramax"
	if r != e {
		t.Fatalf("when %s, returned %s, expected %s", c, r, e)
	}

	c = "- [ ] Crusadia Avramax"
	r = md.ListText(c)
	e = "Crusadia Avramax"
	if r != e {
		t.Fatalf("when %s, returned %s, expected %s", c, r, e)
	}

	c = "- [x] Crusadia Avramax"
	r = md.ListText(c)
	e = "Crusadia Avramax"
	if r != e {
		t.Fatalf("when %s, returned %s, expected %s", c, r, e)
	}

	c = "Crusadia Avramax"
	r = md.ListText(c)
	e = "Crusadia Avramax"
	if r != e {
		t.Fatalf("when %s, returned %s, expected %s", c, r, e)
	}
}
