package main

import (
	"fmt"
	"net/url"

	"github.com/kotaoue/ygolinker/packages/ocgdb"
)

var lang ocgdb.Language

func init() {
	lang = ocgdb.LangJA
}

func main() {
	fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("ラグーン・オブ・レッドアイズ"), lang))
	fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("ハリファイバー"), lang))
	fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("リビングデッド"), lang))
}
