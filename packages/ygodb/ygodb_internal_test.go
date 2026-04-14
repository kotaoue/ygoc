package ygodb

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func newSelection(t *testing.T, html string) *goquery.Selection {
	t.Helper()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("failed to parse html: %v", err)
	}
	return doc.Find("div.root")
}

func Test_extractID(t *testing.T) {
	tests := []struct {
		name string
		html string
		want string
	}{
		{
			name: "from-link-value",
			html: `<div class="root">
				<input class="link_value" value="https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=4007" />
			</div>`,
			want: "4007",
		},
		{
			name: "fallback-to-cid",
			html: `<div class="root">
				<input class="link_value" value="https://example.com/page?ope=1" />
				<input class="cid" value="9999" />
			</div>`,
			want: "9999",
		},
		{
			name: "empty",
			html: `<div class="root"></div>`,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSelection(t, tt.html)
			got := extractID(s)
			if got != tt.want {
				t.Errorf("extractID() got %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_extractLimited(t *testing.T) {
	tests := []struct {
		name string
		html string
		want string
	}{
		{
			name: "from-img-alt",
			html: `<div class="root">
				<dd class="remove_btn"><a><img alt="Forbidden" /></a></dd>
			</div>`,
			want: "Forbidden",
		},
		{
			name: "from-p",
			html: `<div class="root">
				<div class="lr_icon"><p>Limited</p></div>
			</div>`,
			want: "Limited",
		},
		{
			name: "from-span",
			html: `<div class="root">
				<div class="lr_icon"><p></p><span>Semi-Limited</span></div>
			</div>`,
			want: "Semi-Limited",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSelection(t, tt.html)
			got := extractLimited(s)
			if got != tt.want {
				t.Errorf("extractLimited() got %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_siteURL(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"default", "https://www.db.yugioh-card.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := siteURL()
			if got != tt.want {
				t.Errorf("siteURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_apiURL(t *testing.T) {
	got := apiURL("test", LangEN)
	if !strings.Contains(got, "https://www.db.yugioh-card.com") {
		t.Errorf("apiURL() should contain site URL, got %q", got)
	}
	if !strings.Contains(got, "keyword=test") {
		t.Errorf("apiURL() should contain keyword, got %q", got)
	}
	if !strings.Contains(got, "request_locale=en") {
		t.Errorf("apiURL() should contain locale, got %q", got)
	}
}
