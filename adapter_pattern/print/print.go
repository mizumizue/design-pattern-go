package print

import (
	"fmt"
	"github.com/trewanek/PracticeDesignPatternWithGo/adapter_pattern/banner"
)

type Print interface {
	printWeak(str string) string
	printStrong(str string) string
}

type PrintBanner struct {
	banner *banner.Banner
}

func NewPrintBanner(banner *banner.Banner) *PrintBanner {
	return &PrintBanner{banner: banner}
}

func (pb *PrintBanner) PrintWeak(str string) {
	fmt.Println("[" + pb.banner.ShowLower(str) + "]")
}

func (pb *PrintBanner) PrintStrong(str string) {
	fmt.Println("[" + pb.banner.ShowUpper(str) + "]")
}
