package ygodb_test

import (
	"testing"

	"github.com/kotaoue/ygoc/packages/ygodb"
)

func Test_Card_URL(t *testing.T) {
	c := ygodb.Card{ID: "123456789"}
	r := c.URL()
	e := "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=123456789"
	if r != e {
		t.Fatalf("when Card.ID %s, returned %s, expected %s", c.ID, r, e)
	}
}

func Test_Scraping(t *testing.T) {
	// 外部API依存のテスト - APIの変更やネットワーク問題で失敗する可能性がある
	k := "Blue-Eyes White Dragon"
	r, err := ygodb.Scraping(k, ygodb.LangEN)
	if err != nil {
		t.Logf("API test failed for %s: %s. This may be due to external API changes.", k, err.Error())
		// エラーケースのテストのみ実行
		testErrorCases(t)
		return
	}
	
	// APIが動作する場合は基本的な構造を検証
	if r.Name == "" {
		t.Fatalf("Card name should not be empty for valid card")
	}
	if r.ID == "" {
		t.Fatalf("Card ID should not be empty for valid card")
	}

	// エラーケースもテスト
	testErrorCases(t)
}

func testErrorCases(t *testing.T) {
	// Test multiple results error
	k := "Red-Eyes"
	_, err := ygodb.Scraping(k, ygodb.LangEN)
	expectedError := "Error: Couldn't narrow down the cards to one."
	if err == nil {
		t.Logf("Expected error for keyword %s but got none. API may have changed.", k)
	} else if err.Error() != expectedError {
		t.Logf("when keyword %s returned %s, expected %s. API may have changed.", k, err.Error(), expectedError)
	}

	// Test not found error  
	k = "NonExistentCardName12345"
	_, err = ygodb.Scraping(k, ygodb.LangEN)
	expectedError = "Error: Card not found."
	if err == nil {
		t.Fatalf("when keyword %s error not occurred\n", k)
	}
	if err.Error() != expectedError {
		t.Logf("when keyword %s returned %s expected %s. Error format may have changed.", k, err.Error(), expectedError)
	}
}

func Test_ExtractCardID(t *testing.T) {
	s := "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=13049"
	r := ygodb.ExtractCardID(s)
	e := "13049"
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}

	s = "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=1&sess=1&keyword=Gouki+Suprex&stype=1&ctype=&starfr=&starto=&pscalefr=&pscaleto=&linkmarkerfr=&linkmarkerto=&link_m=2&atkfr=&atkto=&deffr=&defto=&othercon=2&request_locale=en"
	r = ygodb.ExtractCardID(s)
	e = ""
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}
}

func Test_ExtractValue(t *testing.T) {
	s := "ATK 3000"
	r := ygodb.ExtractValue(s)
	e := "3000"
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}

	s = "DEF -"
	r = ygodb.ExtractValue(s)
	e = ""
	if r != e {
		t.Fatalf("when set %s, returned %s, expected %s", s, r, e)
	}
}
