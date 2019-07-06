package banner

import "strings"

type Banner struct{}

func NewBanner() *Banner {
	return &Banner{}
}

func (b *Banner) ShowUpper(str string) string {
	return strings.ToUpper(str)
}

func (b *Banner) ShowLower(str string) string {
	return strings.ToLower(str)
}
