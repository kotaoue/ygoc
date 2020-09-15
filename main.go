package main

import (
	"net/url"

	"github.com/kotaoue/ygolinker/packages/ocgdb"
)

var lang ocgdb.Language

func init() {
	lang = ocgdb.LangJA
}

func main() {
	ocgdb.Scraping(url.QueryEscape("レッドアイズ"), lang)
}
