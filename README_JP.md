# ygoc
遊戯王カード用ツールセット。

[English](README.md)|Japanese

[![build](https://github.com/kotaoue/ygoc/workflows/build/badge.svg)](https://github.com/kotaoue/ygoc/actions?query=workflow%3Abuild)
[![Coverage Status](https://coveralls.io/repos/github/kotaoue/ygoc/badge.svg?branch=master)](https://coveralls.io/github/kotaoue/ygoc?branch=master)
[![CodeScanning](https://github.com/kotaoue/ygoc/workflows/CodeScanning/badge.svg)](https://github.com/kotaoue/ygoc/actions?query=workflow%3ACodeScanning)

## 要件
* [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
  ```bash
  $ go get github.com/PuerkitoBio/goquery
  ```

## 使用方法
### mode=select
```bash
$ go run main.go -lang=ja -name="レッドアイズダークネスメタルドラゴン"
レッドアイズ・ダークネスメタルドラゴン
制限
闇属性
レベル 10
攻撃力 2800
守備力 2400
このカード名の、①の方法による特殊召喚は１ターンに１度しかできず、②の効果は１ターンに１度しか使用できない。①：このカードは自分フィールドの表側表示のドラゴン族モンスター１体を除外し、手札から特殊召喚できる。②：自分メインフェイズに発動できる。自分の手札・墓地から「レッドアイズ・ダークネスメタルドラゴン」以外のドラゴン族モンスター１体を選んで特殊召喚する。
```

### mode=markdown
#### before
```bash
$ cat wishlist.md 
# wishlist.md
欲しい物リスト(markdown形式)

* マスカレーナ
* アストラム
* うらら
```

#### after
```bash
$ go run main.go -mode=markdown -lang=ja -file="wishlist.md"
# wishlist.md
欲しい物リスト(markdown形式)

* [I：Pマスカレーナ](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=14676)
* [双穹の騎士アストラム](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=14297)
* [灰流うらら](https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=12950)
```

### mode=link
```bash
$ go run main.go -mode=link -lang=ja -name="ホープゼアル"
https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=11932
```

## Links
* [Yu-Gi-Oh! CARD DATABASE](https://www.db.yugioh-card.com/yugiohdb/)
* [Yu-Gi-Oh! Official Twitter](https://twitter.com/yugioh_ocg_info)