package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/kotaoue/ygolinker/packages/ocgdb"
)

var lang ocgdb.Language
var opt Options

// Options is flag for this code.
type Options struct {
	Lang     ocgdb.Language
	CardName string
}

func init() {
	l := flag.String("lang", string(ocgdb.LangJA), "Language for selecting from the DB.")
	c := flag.String("name", "", "The card name you want to select.")
	flag.Parse()

	opt = Options{
		Lang:     ocgdb.Language(*l),
		CardName: *c,
	}
}

func main() {
	fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape(opt.CardName), opt.Lang))
	fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("ドラグーン・オブ・レッドアイズ"), opt.Lang))
	fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("ハリファイバー"), opt.Lang))
	fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("リビングデッド"), opt.Lang))
}
