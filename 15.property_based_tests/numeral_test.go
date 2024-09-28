package main

import (
	"testing"
	"testing/quick"
)

func TestConvertingToRoman(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			got := ConvertToRoman(tt.Arabic)
			want := tt.Roman
			if got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			got := ConvertToArabic(tt.Roman)
			want := tt.Arabic
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}

type Case struct {
	Name   string
	Arabic uint16
	Roman  string
}

var cases = []Case{
	{Name: `1 gets converted to "I"`, Arabic: 1, Roman: "I"},
	{Name: `2 gets converted to "II"`, Arabic: 2, Roman: "II"},
	{Name: `3 gets converted to "III"`, Arabic: 3, Roman: "III"},
	{Name: `4 gets converted to "IV"`, Arabic: 4, Roman: "IV"},
	{Name: `5 gets converted to "V"`, Arabic: 5, Roman: "V"},
	{Name: `6 gets converted to "VI"`, Arabic: 6, Roman: "VI"},
	{Name: `7 gets converted to "VII"`, Arabic: 7, Roman: "VII"},
	{Name: `8 gets converted to "VIII"`, Arabic: 8, Roman: "VIII"},
	{Name: `9 gets converted to "IX"`, Arabic: 9, Roman: "IX"},
	{Name: `10 gets converted to "X"`, Arabic: 10, Roman: "X"},
	{Name: `14 gets converted to "XIV"`, Arabic: 14, Roman: "XIV"},
	{Name: `18 gets converted to "XVIII"`, Arabic: 18, Roman: "XVIII"},
	{Name: `20 gets converted to "XX"`, Arabic: 20, Roman: "XX"},
	{Name: `39 gets converted to "XXXIX"`, Arabic: 39, Roman: "XXXIX"},
	{Name: `40 gets converted to "XL"`, Arabic: 40, Roman: "XL"},
	{Name: `47 gets converted to "XLVII"`, Arabic: 47, Roman: "XLVII"},
	{Name: `49 gets converted to "XLIX"`, Arabic: 49, Roman: "XLIX"},
	{Name: `50 gets converted to "L"`, Arabic: 50, Roman: "L"},
	{Name: `100 gets converted to "C"`, Arabic: 100, Roman: "C"},
	{Name: `90 gets converted to "XC"`, Arabic: 90, Roman: "XC"},
	{Name: `900 gets converted to "CM"`, Arabic: 900, Roman: "CM"},
	{Name: `400 gets converted to "CD"`, Arabic: 400, Roman: "CD"},
	{Name: `500 gets converted to "D"`, Arabic: 500, Roman: "D"},
	{Name: `1000 gets converted to "M"`, Arabic: 1000, Roman: "M"},
	{Name: `1984 gets converted to "MCMLXXXIV"`, Arabic: 1984, Roman: "MCMLXXXIV"},
	{Name: `3999 gets converted to "MMMCMXCIX"`, Arabic: 3999, Roman: "MMMCMXCIX"},
	{Name: `2014 gets converted to "MMXIV"`, Arabic: 2014, Roman: "MMXIV"},
	{Name: `1006 gets converted to "MVI"`, Arabic: 1006, Roman: "MVI"},
	{Name: `798 gets converted to "DCCXCVIII"`, Arabic: 798, Roman: "DCCXCVIII"},
}
