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

func Test_extractID_from_link_value(t *testing.T) {
	html := `<div class="root">
		<input class="link_value" value="https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=4007" />
	</div>`
	s := newSelection(t, html)
	got := extractID(s)
	if got != "4007" {
		t.Errorf("extractID from link_value: got %q, want %q", got, "4007")
	}
}

func Test_extractID_fallback_to_cid(t *testing.T) {
	// link_value exists but value does not contain cid=, so fall back to input.cid
	html := `<div class="root">
		<input class="link_value" value="https://example.com/page?ope=1" />
		<input class="cid" value="9999" />
	</div>`
	s := newSelection(t, html)
	got := extractID(s)
	if got != "9999" {
		t.Errorf("extractID fallback to cid: got %q, want %q", got, "9999")
	}
}

func Test_extractID_empty(t *testing.T) {
	// No link_value, no cid inputs
	html := `<div class="root"></div>`
	s := newSelection(t, html)
	got := extractID(s)
	if got != "" {
		t.Errorf("extractID empty: got %q, want %q", got, "")
	}
}

func Test_extractLimited_from_img_alt(t *testing.T) {
	html := `<div class="root">
		<dd class="remove_btn"><a><img alt="Forbidden" /></a></dd>
	</div>`
	s := newSelection(t, html)
	got := extractLimited(s)
	if got != "Forbidden" {
		t.Errorf("extractLimited from img alt: got %q, want %q", got, "Forbidden")
	}
}

func Test_extractLimited_from_p(t *testing.T) {
	// No img alt; div.lr_icon p has text
	html := `<div class="root">
		<div class="lr_icon"><p>Limited</p></div>
	</div>`
	s := newSelection(t, html)
	got := extractLimited(s)
	if got != "Limited" {
		t.Errorf("extractLimited from p: got %q, want %q", got, "Limited")
	}
}

func Test_extractLimited_from_span(t *testing.T) {
	// No img alt, empty p; div.lr_icon span has text
	html := `<div class="root">
		<div class="lr_icon"><p></p><span>Semi-Limited</span></div>
	</div>`
	s := newSelection(t, html)
	got := extractLimited(s)
	if got != "Semi-Limited" {
		t.Errorf("extractLimited from span: got %q, want %q", got, "Semi-Limited")
	}
}

func Test_siteURL(t *testing.T) {
	got := siteURL()
	want := "https://www.db.yugioh-card.com"
	if got != want {
		t.Errorf("siteURL() = %q, want %q", got, want)
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
