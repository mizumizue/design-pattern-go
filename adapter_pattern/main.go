package main

import (
	"github.com/trewanek/PracticeDesignPatternWithGo/adapter_pattern/banner"
	"github.com/trewanek/PracticeDesignPatternWithGo/adapter_pattern/print"
)

func main() {
	b := banner.NewBanner()
	pb := print.NewPrintBanner(b)
	pb.PrintStrong("my job is programmer")
}
