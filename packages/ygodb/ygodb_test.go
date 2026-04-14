package ygodb_test

import (
	"testing"

	"github.com/kotaoue/ygoc/packages/ygodb"
)

func Test_Card_URL(t *testing.T) {
	tests := []struct {
		name string
		card ygodb.Card
		want string
	}{
		{
			name: "normal",
			card: ygodb.Card{ID: "123456789"},
			want: "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=123456789",
		},
		{
			name: "empty-id",
			card: ygodb.Card{},
			want: "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.card.URL()
			if r != tt.want {
				t.Fatalf("when Card.ID %s, returned %s, expected %s", tt.card.ID, r, tt.want)
			}
		})
	}
}

func Test_Scraping(t *testing.T) {
	// 外部API依存のテスト - APIの変更やネットワーク問題で失敗する可能性がある
	tests := []struct {
		name       string
		keyword    string
		wantCard   ygodb.Card
		wantErr    string
		errStrict  bool
		allowNoErr bool
	}{
		{
			name:    "blue-eyes",
			keyword: "Blue-Eyes White Dragon",
			wantCard: ygodb.Card{
				ID:        "4007",
				Name:      "Blue-Eyes White Dragon",
				Limited:   "",
				Attribute: "LIGHT",
				Effect:    "",
				Level:     "Level 8",
				Link:      "",
				Attack:    "ATK 3000",
				Defense:   "DEF 2500",
				Text:      "This legendary dragon is a powerful engine of destruction. Virtually invincible, very few have faced this awesome creature and lived to tell the tale.",
			},
		},
		{
			name:    "raigeki",
			keyword: "Raigeki",
			wantCard: ygodb.Card{
				ID:        "4343",
				Name:      "Raigeki",
				Limited:   "",
				Attribute: "SPELL",
				Text:      "Destroy all monsters your opponent controls.",
			},
		},
		{
			name:       "red-eyes-ambiguous",
			keyword:    "Red-Eyes",
			wantErr:    "Error: Couldn't narrow down the cards to one.",
			allowNoErr: true,
		},
		{
			name:      "not-found",
			keyword:   "NonExistentCardName12345",
			wantErr:   "Error: Card not found.",
			errStrict: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := ygodb.Scraping(tt.keyword, ygodb.LangEN)

			if tt.wantErr == "" {
				if err != nil {
					t.Skipf("when keyword %s error occurred %s", tt.keyword, err.Error())
				}
				if r != tt.wantCard {
					t.Fatalf("when keyword %s\nreturned %#v\nexpected %#v", tt.keyword, r, tt.wantCard)
				}
				return
			}

			if err == nil {
				if tt.allowNoErr {
					t.Logf("Expected error for keyword %s but got none. API may have changed.", tt.keyword)
					return
				}
				t.Fatalf("when keyword %s error not occurred", tt.keyword)
			}

			if err.Error() != tt.wantErr {
				if tt.errStrict {
					t.Fatalf("when keyword %s returned %s expected %s", tt.keyword, err.Error(), tt.wantErr)
				}
				t.Logf("when keyword %s returned %s, expected %s. API may have changed.", tt.keyword, err.Error(), tt.wantErr)
			}
		})
	}
}

func Test_ExtractCardID(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "with-cid",
			in:   "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=13049",
			want: "13049",
		},
		{
			name: "without-cid",
			in:   "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=1&sess=1&keyword=Gouki+Suprex&stype=1&ctype=&starfr=&starto=&pscalefr=&pscaleto=&linkmarkerfr=&linkmarkerto=&link_m=2&atkfr=&atkto=&deffr=&defto=&othercon=2&request_locale=en",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ygodb.ExtractCardID(tt.in)
			if r != tt.want {
				t.Fatalf("when set %s, returned %s, expected %s", tt.in, r, tt.want)
			}
		})
	}
}

func Test_ExtractValue(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"atk", "ATK 3000", "3000"},
		{"no-number", "DEF -", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ygodb.ExtractValue(tt.in)
			if r != tt.want {
				t.Fatalf("when set %s, returned %s, expected %s", tt.in, r, tt.want)
			}
		})
	}
}
