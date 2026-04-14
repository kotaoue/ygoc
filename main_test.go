package main

import (
	"os"
	"testing"

	"github.com/kotaoue/ygoc/packages/ygodb"
)

func Test_mode_String(t *testing.T) {
	tests := []struct {
		m    mode
		want string
	}{
		{modeUnknown, "unknown"},
		{modeHelp, "help"},
		{modeSelect, "select"},
		{modeLink, "link"},
		{modeMarkDown, "markdown"},
		{mode(999), ""},
	}
	for _, tt := range tests {
		r := tt.m.String()
		if r != tt.want {
			t.Errorf("mode(%d).String() got %q, want %q", tt.m, r, tt.want)
		}
	}
}

func Test_atoMode(t *testing.T) {
	tests := []struct {
		s    string
		want mode
	}{
		{"unknown", modeUnknown},
		{"help", modeHelp},
		{"select", modeSelect},
		{"link", modeLink},
		{"markdown", modeMarkDown},
		{"HELP", modeHelp},
		{"nonexistent", modeUnknown},
	}
	for _, tt := range tests {
		r := atoMode(tt.s)
		if r != tt.want {
			t.Errorf("atoMode(%q) got %v, want %v", tt.s, r, tt.want)
		}
	}
}

func Test_modes(t *testing.T) {
	r := modes()
	if len(r) == 0 {
		t.Fatal("modes() returned empty slice")
	}
	for _, v := range r {
		if v == modeUnknown.String() {
			t.Errorf("modes() contains unknown: %v", r)
		}
	}
}

func Test_linkMD(t *testing.T) {
	c := ygodb.Card{ID: "4007", Name: "Blue-Eyes White Dragon"}
	r := linkMD(c)
	e := "[Blue-Eyes White Dragon](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=4007)"
	if r != e {
		t.Fatalf("linkMD got %q, want %q", r, e)
	}
}

func Test_isPrettyTarget(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"- Blue-Eyes White Dragon", true},
		{"* Blue-Eyes White Dragon", true},
		{"- [x] Blue-Eyes White Dragon", true},
		{"- [Dragon](https://example.com/)", false},
		{"Blue-Eyes White Dragon", false},
		{"# Heading", false},
	}
	for _, tt := range tests {
		r := isPrettyTarget(tt.s)
		if r != tt.want {
			t.Errorf("isPrettyTarget(%q) got %v, want %v", tt.s, r, tt.want)
		}
	}
}

func Test_prettyList_file_not_found(t *testing.T) {
	r := prettyList("/nonexistent/file.md", ygodb.LangEN)
	if len(r) != 0 {
		t.Errorf("expected empty slice, got %v", r)
	}
}

func Test_prettyList_non_list_content(t *testing.T) {
	f, err := os.CreateTemp("", "ygoc_test_*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	content := "# Heading\nPlain text\n[Link](https://example.com/)"
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
	f.Close()

	r := prettyList(f.Name(), ygodb.LangEN)
	if len(r) != 3 {
		t.Errorf("prettyList got %d lines, want 3", len(r))
	}
	if len(r) > 0 && r[0] != "# Heading" {
		t.Errorf("line 0: got %q, want %q", r[0], "# Heading")
	}
}

func Test_help(t *testing.T) {
	help()
}
