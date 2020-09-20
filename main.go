package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/kotaoue/ygolinker/packages/ygodb"
)

var opt options

// options is flag for this code.
type options struct {
	executeMode mode
	lang        ygodb.Language
	cardName    string
}

// mode is a value that specifies the behavior of this code.
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
	l := flag.String("lang", string(ygodb.LangJA), "Language for selecting from the DB.")
	c := flag.String("name", "", "The card name you want to select.")
	flag.Parse()

	opt = options{
		executeMode: mode(*m),
		lang:        ygodb.Language(*l),
		cardName:    *c,
	}
}

func main() {
	switch opt.executeMode {
	case modeSelect:
		for _, v := range selectCard(opt.cardName, opt.lang) {
			fmt.Println(v)
		}
	case modeHelp:
		help()
	}
	// fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("ドラグーン・オブ・レッドアイズ"), opt.Lang))
	// fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("ハリファイバー"), opt.Lang))
	// fmt.Printf("%v\n", ocgdb.Scraping(url.QueryEscape("リビングデッド"), opt.Lang))
}

// selectCard is scraping from DB with the specified card name.
func selectCard(cardName string, lang ygodb.Language) []string {
	c := ygodb.Scraping(url.QueryEscape(cardName), lang)

	var s []string

	if len(c.Name) > 0 {
		s = append(s, c.Name)
	}
	if len(c.Limited) > 0 {
		s = append(s, c.Limited)
	}
	if len(c.Attribute) > 0 {
		s = append(s, c.Attribute)
	}
	if len(c.Effect) > 0 {
		s = append(s, c.Effect)
	}
	if len(c.Level) > 0 {
		s = append(s, c.Level)
	}
	if len(c.Link) > 0 {
		s = append(s, c.Link)
	}
	if len(c.Attack) > 0 {
		s = append(s, c.Attack)
	}
	if len(c.Defence) > 0 {
		s = append(s, c.Defence)
	}
	if len(c.Text) > 0 {
		s = append(s, c.Text)
	}

	return s
}
