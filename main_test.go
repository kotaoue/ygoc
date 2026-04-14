package main

import (
	"os"
	"testing"

	"github.com/kotaoue/ygoc/packages/ygodb"
)

func Test_mode_String(t *testing.T) {
	tests := []struct {
		name string
		m    mode
		want string
	}{
		{"unknown", modeUnknown, "unknown"},
		{"help", modeHelp, "help"},
		{"select", modeSelect, "select"},
		{"link", modeLink, "link"},
		{"markdown", modeMarkDown, "markdown"},
		{"invalid", mode(999), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.m.String()
			if r != tt.want {
				t.Errorf("mode(%d).String() got %q, want %q", tt.m, r, tt.want)
			}
		})
	}
}

func Test_atoMode(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want mode
	}{
		{"unknown", "unknown", modeUnknown},
		{"help", "help", modeHelp},
		{"select", "select", modeSelect},
		{"link", "link", modeLink},
		{"markdown", "markdown", modeMarkDown},
		{"upper-help", "HELP", modeHelp},
		{"invalid", "nonexistent", modeUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := atoMode(tt.s)
			if r != tt.want {
				t.Errorf("atoMode(%q) got %v, want %v", tt.s, r, tt.want)
			}
		})
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
	tests := []struct {
		name string
		card ygodb.Card
		want string
	}{
		{
			name: "normal",
			card: ygodb.Card{ID: "4007", Name: "Blue-Eyes White Dragon"},
			want: "[Blue-Eyes White Dragon](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=4007)",
		},
		{
			name: "empty-fields",
			card: ygodb.Card{},
			want: "[](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := linkMD(tt.card)
			if r != tt.want {
				t.Fatalf("linkMD got %q, want %q", r, tt.want)
			}
		})
	}
}

func Test_isPrettyTarget(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"dash-list", "- Blue-Eyes White Dragon", true},
		{"asterisk-list", "* Blue-Eyes White Dragon", true},
		{"checkbox-list", "- [x] Blue-Eyes White Dragon", true},
		{"already-linked", "- [Dragon](https://example.com/)", false},
		{"plain", "Blue-Eyes White Dragon", false},
		{"heading", "# Heading", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := isPrettyTarget(tt.s)
			if r != tt.want {
				t.Errorf("isPrettyTarget(%q) got %v, want %v", tt.s, r, tt.want)
			}
		})
	}
}

func Test_prettyList(t *testing.T) {
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

	tests := []struct {
		name     string
		fileName string
		wantLen  int
		wantHead string
	}{
		{"file-not-found", "/nonexistent/file.md", 0, ""},
		{"non-list-content", f.Name(), 3, "# Heading"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := prettyList(tt.fileName, ygodb.LangEN)
			if len(r) != tt.wantLen {
				t.Errorf("prettyList got %d lines, want %d", len(r), tt.wantLen)
			}
			if tt.wantLen > 0 && len(r) > 0 && r[0] != tt.wantHead {
				t.Errorf("line 0: got %q, want %q", r[0], tt.wantHead)
			}
		})
	}
}

func Test_help(t *testing.T) {
	help()
}
