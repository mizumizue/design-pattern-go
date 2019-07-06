package banner

import (
	"testing"
	"unicode"
)

func TestShowUpper(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{"LowerOnly", "hoge"},
		{"Mixed", "Hoge"},
		{"UpperOnly", "HOGE"},
	}
	for _, test := range tests {
		b := NewBanner()
		t.Run(test.name, func(t *testing.T) {
			result := b.ShowUpper(test.str)
			for _, char := range result {
				if !unicode.IsUpper(char) {
					t.Errorf("char is lower. something is wrong.")
				}
			}
		})
	}
}

func TestShowLower(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{"LowerOnly", "hoge"},
		{"Mixed", "Hoge"},
		{"UpperOnly", "HOGE"},
	}
	for _, test := range tests {
		b := NewBanner()
		t.Run(test.name, func(t *testing.T) {
			result := b.ShowLower(test.str)
			for _, char := range result {
				if !unicode.IsLower(char) {
					t.Errorf("char is upper. something is wrong.")
				}
			}
		})
	}
}
