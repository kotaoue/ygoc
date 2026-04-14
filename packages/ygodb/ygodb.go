package ygodb

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Language is supported language at YGO DB.
type Language string

// List of supported languages at YGO DB.
const (
	LangJA Language = "ja" // 日本語
	LangEN Language = "en" // English
	LangDE Language = "de" // Deutsch
	LangFR Language = "fr" // Français
	LangIT Language = "it" // Italiano
	LangES Language = "es" // Español
	LangPT Language = "pt" // Portugues
	LangKO Language = "ko" // 한글
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
	Defense   string
	Text      string
}

// URL is card detail pages url.
func (c Card) URL() string {
	return fmt.Sprintf("%s/yugiohdb/card_search.action?ope=2&cid=%s", siteURL(), c.ID)
}

// Scraping from YGO DB.
func Scraping(keyword string, lang Language) (Card, error) {
	keyword = url.QueryEscape(keyword)
	c := Card{}

	doc, err := goquery.NewDocument(apiURL(keyword, lang))
	if err != nil {
		return c, err
	}

	cardRows := doc.Find("div#article_body div.t_row.c_normal")
	l := cardRows.Length()

	if l == 0 {
		return c, fmt.Errorf("Error: %s", "Card not found.")
	}

	if l == 1 {
		c = scrapingCard(cardRows.First())
	} else if l > 1 {
		keywordDecoded, _ := url.QueryUnescape(keyword)
		keywordLower := strings.ToLower(keywordDecoded)
		partial := Card{}

		cardRows.Each(func(index int, s *goquery.Selection) {
			cardName := strings.ToLower(s.Find("span.card_name").Text())
			cardName = strings.TrimSpace(cardName)

			if strings.EqualFold(keywordLower, cardName) {
				if len(c.ID) == 0 {
					c = scrapingCard(s)
				}
				return
			}

			if strings.Contains(cardName, keywordLower) && len(partial.ID) == 0 {
				partial = scrapingCard(s)
			}
		})

		if len(c.ID) == 0 {
			if len(partial.ID) > 0 {
				c = partial
			} else {
				// If no match found, return the first card
				c = scrapingCard(cardRows.First())
			}
		}
	}

	return c, nil
}

// scrapingCard is scraping card detail.
func scrapingCard(s *goquery.Selection) Card {
	c := Card{}

	c.ID = extractID(s)
	c.Name = extractName(s)
	c.Limited = extractLimited(s)
	c.Attribute = extractAttribute(s)
	c.Level = extractLevel(s)
	c.Link = extractLink(s)
	c.Attack = extractAttack(s)
	c.Defense = extractDefense(s)
	c.Effect = extractEffect(s)
	c.Text = extractText(s)

	return c
}

func extractID(s *goquery.Selection) string {
	if v, ok := s.Find("input.link_value").Attr("value"); ok {
		id := ExtractCardID(v)
		if len(id) > 0 {
			return id
		}
	}

	if cidVal, ok := s.Find("input.cid").Attr("value"); ok {
		return cidVal
	}

	return ""
}

func extractName(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span.card_name").Text())
}

func extractLimited(s *goquery.Selection) string {
	if limitedImg, ok := s.Find("dd.remove_btn a img").Attr("alt"); ok {
		return limitedImg
	}

	limited := strings.TrimSpace(s.Find("div.lr_icon p").First().Text())
	if len(limited) > 0 {
		return limited
	}

	return strings.TrimSpace(s.Find("div.lr_icon span").First().Text())
}

func extractAttribute(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span.box_card_attribute > span").Text())
}

func extractLevel(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span.box_card_level_rank > span").Text())
}

func extractLink(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span.box_card_linkmarker > span").Text())
}

func extractAttack(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span.atk_power > span").Text())
}

func extractDefense(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span.def_power > span").Text())
}

func extractEffect(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("span.box_card_effect > span").Text())
}

func extractText(s *goquery.Selection) string {
	return strings.TrimSpace(s.Find("dd.box_card_text").Text())
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

