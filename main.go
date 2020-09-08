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

var lang language

func init() {
	lang = langJA
}

func main() {
	fmt.Println(search(url.QueryEscape("レッドアイズ"), lang)) // htmlをstringで取得
}

func search(keyword string, lang language) string {
	url := "https://www.db.yugioh-card.com/yugiohdb/card_search.action"
	param := fmt.Sprintf("ope=1&sess=1&keyword=%s&stype=1&ctype=&starfr=&starto=&pscalefr=&pscaleto=&linkmarkerfr=&linkmarkerto=&link_m=2&atkfr=&atkto=&deffr=&defto=&othercon=1&request_locale=%s", keyword, lang)

	res, _ := http.Get(fmt.Sprintf("%s?%s", url, param))

	fmt.Printf("%s?%s\n", url, param)
	defer res.Body.Close()

	scn := bufio.NewScanner(res.Body)

	var b bool
	for scn.Scan() {
		if strings.Contains(scn.Text(), "<dt class=\"box_card_name\"") {
			b = true
		}
		if strings.Contains(scn.Text(), "</dl>") {
			b = false
		}

		if b {
			fmt.Println(scn.Text())
		}
	}

	return ""
}
