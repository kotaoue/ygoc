package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var api = "https://www.db.yugioh-card.com/yugiohdb/card_search.action"

type language string

const (
	langJA language = "ja" // 日本語
	langEN          = "en" // English
	langDE          = "de" // Deutsch
	langFR          = "fr" // Français
	langIT          = "it" // Italiano
	langES          = "es" // Español
	langPT          = "pt" // Portugues
	langKO          = "ko" // 한글
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

var lang language

func init() {
	lang = langJA
}

func main() {
	fmt.Println(search(url.QueryEscape("レッドアイズ"), lang)) // htmlをstringで取得
}

func search(keyword string, lang language) string {
	res, _ := http.Get(apiURL(keyword, lang))
	defer res.Body.Close()

	scn := bufio.NewScanner(res.Body)

	var cardNameBlock bool
	var cardSpecBlock bool
	for scn.Scan() {
		if strings.Contains(scn.Text(), "<dt class=\"box_card_name\">") {
			cardNameBlock = true
		}
		if strings.Contains(scn.Text(), "</dt>") {
			cardNameBlock = false
		}
		if strings.Contains(scn.Text(), "<dd class=\"box_card_spec\">") {
			cardSpecBlock = true
		}
		if strings.Contains(scn.Text(), "</dd>") {
			cardSpecBlock = false
		}

		if cardNameBlock || cardSpecBlock {
			fmt.Println(scn.Text())
		}
	}

	return ""
}

func apiURL(keyword string, lang language) string {
	url := "https://www.db.yugioh-card.com/yugiohdb/card_search.action"
	param := fmt.Sprintf("ope=1&sess=1&keyword=%s&stype=1&ctype=&starfr=&starto=&pscalefr=&pscaleto=&linkmarkerfr=&linkmarkerto=&link_m=2&atkfr=&atkto=&deffr=&defto=&othercon=1&request_locale=%s", keyword, lang)

	return fmt.Sprintf("%s?%s", url, param)
}
