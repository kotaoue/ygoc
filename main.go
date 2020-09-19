package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/kotaoue/ygolinker/packages/ocgdb"
)

var opt options

// options is flag for this code.
type options struct {
	executeMode mode
	lang        ocgdb.Language
	cardName    string
}

// Mode is a value that specifies the behavior of this code.
type mode string

// List of prepared mode.
const (
	modeHelp   mode = "help"
	modeSelect      = "select"
	modeLink        = "link"
)

// help is priting the mode options for this code.
func help() {
	fmt.Println("The execute modes of this code.")
	fmt.Println("You specify any one to '-mode'.")
	fmt.Printf("  -%s\n\tShow mode options.\n", modeHelp)
	fmt.Printf("  -%s\n\tSelect from DB with the specified card name.\n", modeSelect)
}

func init() {
	m := flag.String("mode", string(modeSelect), "Specifies the behavior of this code.")
	l := flag.String("lang", string(ocgdb.LangJA), "Language for selecting from the DB.")
	c := flag.String("name", "", "The card name you want to select.")
	flag.Parse()

	opt = options{
		executeMode: mode(*m),
		lang:        ocgdb.Language(*l),
		cardName:    *c,
	}
}

func main() {
	switch opt.executeMode {
	case modeSelect:
		fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape(opt.cardName), opt.lang))
	case modeHelp:
		help()
	}
	// fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("ドラグーン・オブ・レッドアイズ"), opt.Lang))
	// fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("ハリファイバー"), opt.Lang))
	// fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("リビングデッド"), opt.Lang))
}
