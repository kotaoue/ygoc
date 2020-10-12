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
	keyword = url.QueryEscape(keyword)
	c := Card{}

	doc, err := goquery.NewDocument(apiURL(keyword, lang))
	if err != nil {
		return c, err
	}

	boxList := doc.Find("div#article_body > table > tbody > tr > td > div.list_style > ul.box_list")
	l := boxList.Children().Length()

	if l == 1 {
		c = scrapingCard(boxList.Children())
	} else if l > 1 {
		boxList.Children().Each(func(index int, s *goquery.Selection) {
			if strings.EqualFold(keyword, s.Find("dt.box_card_name > span.card_status > strong").Text()) {
				c = scrapingCard(s)
			}
		})

		if len(c.ID) == 0 {
			return c, fmt.Errorf("Error: %s", "Couldn't narrow down the cards to one.")
		}
	} else {
		return c, fmt.Errorf("Error: %s", "Card not found.")
	}

	return c, nil
}

// scrapingCard is scraping card detail.
func scrapingCard(s *goquery.Selection) Card {
	c := Card{}

	if v, ok := s.Find("input.link_value").Attr("value"); ok {
		c.ID = ExtractCardID(v)
	}
	c.Name = s.Find("dt.box_card_name > span.card_status > strong").Text()
	if a, ok := s.Find("dt.box_card_name > span.card_status > span.f_right > img").Attr("alt"); ok {
		c.Limited = a
	}
	c.Attribute = s.Find("dd.box_card_spec > span.box_card_attribute > span").Text()
	c.Effect = s.Find("dd.box_card_spec > span.box_card_effect > span").Text()
	c.Level = s.Find("dd.box_card_spec > span.box_card_level_rank > span").Text()
	c.Link = s.Find("dd.box_card_spec > span.box_card_linkmarker > span").Text()
	c.Attack = s.Find("dd.box_card_spec > span.atk_power").Text()
	c.Defence = s.Find("dd.box_card_spec > span.def_power").Text()
	c.Text = strings.TrimSpace(s.Find("dd.box_card_text").Text())

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
//
// Some parameter  can't be expressed as numbers.
// For that reason returning as strings.
// Example: Link Monster's deffence.
func ExtractValue(s string) string {
	reg := regexp.MustCompile(`\d+`)
	return reg.FindString(s)
}
