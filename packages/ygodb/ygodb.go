package ygodb

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Language is supported language at YGO DB.
type Language string

// List of supported languages at YGO DB.
const (
	LangJA Language = "ja" // 日本語
	LangEN          = "en" // English
	LangDE          = "de" // Deutsch
	LangFR          = "fr" // Français
	LangIT          = "it" // Italiano
	LangES          = "es" // Español
	LangPT          = "pt" // Portugues
	LangKO          = "ko" // 한글
)

// Card is parameter set for YGO Card.
type Card struct {
	ID        string
	Name      string
	Limited   string
	Attribute string
	Effect    string
	Level     string
	Link      string
	Attack    string
	Defence   string
	Text      string
}

// URL is card detail pages url.
func (c Card) URL() string {
	return fmt.Sprintf("%s/yugiohdb/card_search.action?ope=2&cid=%s", siteURL(), c.ID)
}

// Scraping from YGO DB.
func Scraping(keyword string, lang Language) (Card, error) {
	debug := os.Getenv("YGOC_DEBUG") != ""
	keyword = url.QueryEscape(keyword)
	c := Card{}

	doc, err := goquery.NewDocument(apiURL(keyword, lang))
	if err != nil {
		return c, err
	}

	// Debug: print full HTML to stderr
	if debug {
		html, _ := doc.Html()
		fmt.Fprintln(os.Stderr, "=== DEBUG HTML ===")
		fmt.Fprintln(os.Stderr, html[:min(len(html), 10000)])
		fmt.Fprintln(os.Stderr, "=== END DEBUG ===")
	}

	// New HTML structure: cards are in div.t_row.c_normal
	cardRows := doc.Find("div#article_body div.t_row.c_normal")
	l := cardRows.Length()

	if debug {
		fmt.Fprintf(os.Stderr, "DEBUG: card rows found: %d\n", l)
	}

	if l == 0 {
		return c, fmt.Errorf("Error: %s", "Card not found.")
	}

	if l == 1 {
		c = scrapingCard(cardRows.First())
	} else if l > 1 {
		// Multiple cards found, try to match by name
		// URL decode the keyword for comparison
		keywordDecoded, _ := url.QueryUnescape(keyword)
		keywordLower := strings.ToLower(keywordDecoded)

		cardRows.Each(func(index int, s *goquery.Selection) {
			cardName := strings.ToLower(s.Find("span.card_name").Text())
			cardName = strings.TrimSpace(cardName)
			if debug {
				fmt.Fprintf(os.Stderr, "DEBUG: comparing %q with %q\n", keywordLower, cardName)
			}

			if strings.Contains(cardName, keywordLower) || strings.EqualFold(keywordLower, cardName) {
				if len(c.ID) == 0 {
					c = scrapingCard(s)
					if debug {
						fmt.Fprintf(os.Stderr, "DEBUG: matched card: %q\n", cardName)
					}
				}
			}
		})

		if len(c.ID) == 0 {
			// If exact match failed, return the first card
			if debug {
				fmt.Fprintf(os.Stderr, "DEBUG: no exact match, using first card\n")
			}
			c = scrapingCard(cardRows.First())
		}
	}

	return c, nil
}

// scrapingCard is scraping card detail.
func scrapingCard(s *goquery.Selection) Card {
	c := Card{}

	// Extract card ID from link_value input
	if v, ok := s.Find("input.link_value").Attr("value"); ok {
		if os.Getenv("YGOC_DEBUG") != "" {
			fmt.Fprintf(os.Stderr, "DEBUG: link_value: %q\n", v)
		}
		c.ID = ExtractCardID(v)
	}

	// If ID not found in link_value, try cid input
	if len(c.ID) == 0 {
		if cidVal, ok := s.Find("input.cid").Attr("value"); ok {
			if os.Getenv("YGOC_DEBUG") != "" {
				fmt.Fprintf(os.Stderr, "DEBUG: cid value: %q\n", cidVal)
			}
			c.ID = cidVal
		}
	}

	// Extract card name
	c.Name = s.Find("span.card_name").Text()
	c.Name = strings.TrimSpace(c.Name)

	// Extract limited/forbidden status
	if limitedImg, ok := s.Find("dd.remove_btn a img").Attr("alt"); ok {
		c.Limited = limitedImg
	}
	if len(c.Limited) == 0 {
		c.Limited = strings.TrimSpace(s.Find("div.lr_icon p").First().Text())
	}
	if len(c.Limited) == 0 {
		c.Limited = strings.TrimSpace(s.Find("div.lr_icon span").First().Text())
	}

	// Extract attribute (get the span inside span.box_card_attribute)
	c.Attribute = s.Find("span.box_card_attribute > span").Text()
	c.Attribute = strings.TrimSpace(c.Attribute)

	// Extract level/rank
	c.Level = s.Find("span.box_card_level_rank > span").Text()
	c.Level = strings.TrimSpace(c.Level)

	// Extract link marker if present
	c.Link = s.Find("span.box_card_linkmarker > span").Text()
	c.Link = strings.TrimSpace(c.Link)

	// Extract attack power
	c.Attack = s.Find("span.atk_power > span").Text()
	c.Attack = strings.TrimSpace(c.Attack)

	// Extract defense power
	c.Defence = s.Find("span.def_power > span").Text()
	c.Defence = strings.TrimSpace(c.Defence)

	// Keep legacy behavior: only extract explicit effect label if present.
	c.Effect = s.Find("span.box_card_effect > span").Text()
	c.Effect = strings.TrimSpace(c.Effect)

	// Extract full text description
	c.Text = s.Find("dd.box_card_text").Text()
	c.Text = strings.TrimSpace(c.Text)

	if os.Getenv("YGOC_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "DEBUG card: ID=%q, Name=%q, Attribute=%q, Limited=%q\n", c.ID, c.Name, c.Attribute, c.Limited)
	}

	return c
}

// siteURL is site url for YGO DB.
func siteURL() string {
	return "https://www.db.yugioh-card.com"
}

// apiURL is search url for YGO DB.
func apiURL(keyword string, lang Language) string {
	api := "/yugiohdb/card_search.action"
	param := fmt.Sprintf("ope=1&sess=1&keyword=%s&stype=1&ctype=&starfr=&starto=&pscalefr=&pscaleto=&linkmarkerfr=&linkmarkerto=&link_m=2&atkfr=&atkto=&deffr=&defto=&othercon=1&request_locale=%s", keyword, lang)

	return fmt.Sprintf("%s%s?%s", siteURL(), api, param)
}

// ExtractCardID is extract cid from link text.
func ExtractCardID(s string) string {
	reg := regexp.MustCompile(`cid=([\d]+)`)
	r := reg.FindStringSubmatch(s)

	if len(r) >= 2 {
		return r[1]
	}
	return ""
}

// ExtractValue is extract value of number from string.
func ExtractValue(s string) string {
	reg := regexp.MustCompile(`\d+`)
	return reg.FindString(s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
