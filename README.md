# ygoc
The tool set for Yu-Gi-Oh! CARD.  

English|[Japanese](README_JP.md)

[![build](https://github.com/kotaoue/ygoc/workflows/build/badge.svg)](https://github.com/kotaoue/ygoc/actions?query=workflow%3Abuild)
[![Coverage Status](https://coveralls.io/repos/github/kotaoue/ygoc/badge.svg?branch=master)](https://coveralls.io/github/kotaoue/ygoc?branch=master)
[![CodeScanning](https://github.com/kotaoue/ygoc/workflows/CodeScanning/badge.svg)](https://github.com/kotaoue/ygoc/actions?query=workflow%3ACodeScanning)

## Requirement
* [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
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

### mode=markdown
#### before
```bash
$ cat wishlist.md 
# wishlist.md
My wishlist for new deck.

* Masquerena
* Crusadia Avramax
* Ash Blossom
```

#### after
```bash
$ go run main.go -mode=markdown -lang=en -file="wishlist.md"
# wishlist.md
My wishlist for new deck.

* [I:P Masquerena](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=14676)
* [Mekk-Knight Crusadia Avramax](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=14297)
* [Ash Blossom & Joyous Spring](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=12950)
```

### mode=link
```bash
$ go run main.go -mode=link -lang=en -name="Utopic ZEXAL"
https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=11932
```

## Links
* [Yu-Gi-Oh! CARD DATABASE](https://www.db.yugioh-card.com/yugiohdb/)
* [Yu-Gi-Oh! Official Twitter](https://twitter.com/yugioh_ocg_info)