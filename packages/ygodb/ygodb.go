package ygodb

import (
	"fmt"
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

// Scraping from YGO DB.
func Scraping(keyword string, lang Language) Card {
	c := Card{}

	doc, err := goquery.NewDocument(apiURL(keyword, lang))
	if err != nil {
		fmt.Printf("%v", err)
		return c
	}

	boxList := doc.Find("div#article_body > table > tbody > tr > td > div.list_style > ul.box_list")
	l := boxList.Children().Length()

	if l == 1 {
		if v, ok := boxList.Children().Find("input.link_value").Attr("value"); ok {
			c.ID = ExtractCardID(v)
		}
		c.Name = boxList.Children().Find("dt.box_card_name > span.card_status > strong").Text()
		if a, ok := boxList.Children().Find("dt.box_card_name > span.card_status > span.f_right > img").Attr("alt"); ok {
			c.Limited = a
		}
		c.Attribute = boxList.Children().Find("dd.box_card_spec > span.box_card_attribute > span").Text()
		c.Effect = boxList.Children().Find("dd.box_card_spec > span.box_card_effect > span").Text()
		c.Level = boxList.Children().Find("dd.box_card_spec > span.box_card_level_rank > span").Text()
		c.Link = boxList.Children().Find("dd.box_card_spec > span.box_card_linkmarker > span").Text()
		c.Attack = boxList.Children().Find("dd.box_card_spec > span.atk_power").Text()
		c.Defence = boxList.Children().Find("dd.box_card_spec > span.def_power").Text()
		c.Text = strings.TrimSpace(boxList.Children().Find("dd.box_card_text").Text())
	} else if l > 1 {
		fmt.Println("Couldn't narrow down the cards to one.")
	} else {
		fmt.Println("Card not found.")
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
//
// Some parameter  can't be expressed as numbers.
// For that reason returning as strings.
// Example: Link Monster's deffence.
func ExtractValue(s string) string {
	reg := regexp.MustCompile(`\d+`)
	return reg.FindString(s)
}
