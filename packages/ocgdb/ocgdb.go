package ocgdb

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

// Language is supported language at OCG DB.
type Language string

// List of supported languages at OCG DB.
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

type dom int

const (
	domNone dom = iota
	domName
	domSpec
	domAttribute
	domEffect
	domLevel
	domAttack
	domDefence
)

// Card is parameter set for OCG Card.
type Card struct {
	Name      string
	Attribute string
	Effect    string
	Level     int
	Attack    int
	Defence   int
}

// Scraping from OCG DB.
func Scraping(keyword string, lang Language) {
	res, _ := http.Get(apiURL(keyword, lang))
	defer res.Body.Close()

	scn := bufio.NewScanner(res.Body)

	var s string
	d := domNone
	for scn.Scan() {
		s, d = readLine(scn.Text(), d)

		if len(s) > 0 && d != domNone {
			fmt.Println(s)
		}
	}
}

func apiURL(keyword string, lang Language) string {
	url := "https://www.db.yugioh-card.com/yugiohdb/card_search.action"
	param := fmt.Sprintf("ope=1&sess=1&keyword=%s&stype=1&ctype=&starfr=&starto=&pscalefr=&pscaleto=&linkmarkerfr=&linkmarkerto=&link_m=2&atkfr=&atkto=&deffr=&defto=&othercon=1&request_locale=%s", keyword, lang)

	return fmt.Sprintf("%s?%s", url, param)
}

func readLine(s string, d dom) (string, dom) {
	s = strings.TrimSpace(s)
	{
		// card name
		if strings.Contains(s, "<dt class=\"box_card_name\">") {
			return "", domName
		}

		if d == domName && strings.HasPrefix(s, "<strong>") {
			return strings.Trim(s, "<strong>"), d
		}

		if d == domName && strings.Contains(s, "</dt>") {
			return "", domNone
		}
	}

	{
		// card spec
		if strings.Contains(s, "<dd class=\"box_card_spec\">") {
			return "", domSpec
		}

		if d == domSpec && strings.Contains(s, "</dd>") {
			return "", domNone
		}
	}

	/*
		if d != domNone {
			return s, d
		}
	*/

	return "", d
}
