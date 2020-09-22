package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kotaoue/ygoc/packages/ygodb"
)

var opt options

// options is flag for this code.
type options struct {
	executeMode mode
	lang        ygodb.Language
	cardName    string
}

// mode is a value that specifies the behavior of this code.
type mode int

// List of prepared mode.
const (
	modeUnknown mode = iota
	modeHelp
	modeSelect
	modeLink
)

type modeDetail struct {
	key         string
	description string
}

var modeMap = map[mode]modeDetail{
	modeUnknown: {key: "unknown", description: "Unknown mode."},
	modeHelp:    {key: "help", description: "Show selectable modes."},
	modeSelect:  {key: "select", description: "Select from DB with the specified card name."},
	modeLink:    {key: "insert", description: "Show card details url."},
}

func (m mode) String() string {
	if m, ok := modeMap[m]; ok {
		return m.key
	}
	return ""
}

// modes is returns all valid values ​​excluded 'Unknown'.
func modes() []string {
	var s []string
	for _, v := range modeMap {
		if !strings.EqualFold(v.key, fmt.Sprint(modeUnknown)) {
			s = append(s, v.key)
		}
	}
	return s
}

// atoMode is convert from string to mode.
func atoMode(s string) mode {
	for k, v := range modeMap {
		if strings.EqualFold(v.key, s) {
			return k
		}
	}
	return modeUnknown
}

func init() {
	m := flag.String("mode", fmt.Sprint(modeSelect), fmt.Sprintf("Specifies the behavior of this code. [%s]", strings.Join(modes(), "|")))
	l := flag.String("lang", string(ygodb.LangJA), "Language for selecting from the DB.")
	c := flag.String("name", "", "The card name you want to select.")
	flag.Parse()

	opt = options{
		executeMode: atoMode(*m),
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
	case modeLink:
		fmt.Println(link(opt.cardName, opt.lang))
	case modeHelp:
		help()
	}
}

// selectCard is scraping from DB with the specified card name.
func selectCard(cardName string, lang ygodb.Language) []string {
	c, _ := ygodb.Scraping(cardName, lang)

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

// link is get card detail pages url.
func link(cardName string, lang ygodb.Language) string {
	c, _ := ygodb.Scraping(cardName, lang)
	return c.URL()
}

// help is priting the mode options for this code.
func help() {
	fmt.Println("The execute modes of this code.")
	fmt.Println("You specify any one to '-mode'.")
	for _, v := range modeMap {
		fmt.Printf("  -%s\n\t%s\n", v.key, v.description)
	}
}
