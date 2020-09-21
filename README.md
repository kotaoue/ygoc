# ygoc
The tool set for Yu-Gi-Oh! CARD.  
Convert from card name to link text for Yu-Gi-Oh! CARD DATABASE.

![build](https://github.com/kotaoue/ygoc/workflows/build/badge.svg)

## Requirement
```bash
$ go get github.com/PuerkitoBio/goquery
```

## Usage
### mode=select
```bash
$ go run main.go -lang=en -name="Red-Eyes Darkness Metal Dragon"
Red-Eyes Darkness Metal Dragon
Limited
DARK
Level 10
ATK 2800
DEF 2400
You can Special Summon this card (from your hand) by banishing 1 face-up Dragon monster you control. You can only Special Summon "Red-Eyes Darkness Metal Dragon" once per turn this way. During your Main Phase: You can Special Summon 1 Dragon monster from your hand or GY, except "Red-Eyes Darkness Metal Dragon". You can only use this effect of "Red-Eyes Darkness Metal Dragon" once per turn.
```

### mode=link
```bash
$ go run main.go -mode=link -lang=en -name="Utopic ZEXAL"
https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=11932
```

## Links
* [Yu-Gi-Oh! CARD DATABASE](https://www.db.yugioh-card.com/yugiohdb/)
* [Yu-Gi-Oh! Official Twitter](https://twitter.com/yugioh_ocg_info)