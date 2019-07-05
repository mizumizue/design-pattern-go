package main

import (
	"fmt"
	"github.com/trewanek/PracticeDesignPatternWithGo/builder_pattern/builder"
	"github.com/trewanek/PracticeDesignPatternWithGo/builder_pattern/director"
	"os"
)

func main() {
	if os.Args[len(os.Args)-1] == "plain" {
		tb := builder.NewTextBuilder()
		d := director.NewDirector(tb)
		d.Construct()
		fmt.Println("created text: " + tb.GetResult())
		os.Exit(0)
	}

	if os.Args[len(os.Args)-1] == "html" {
		hb := builder.NewHtmlBuilder()
		d := director.NewDirector(hb)
		d.Construct()
		fmt.Println("created file: " + hb.GetFilename())
		os.Exit(0)
	}

	fmt.Println("usage: go run main.go plain = create plain text")
	fmt.Println("usage: go run main.go html = create html file")
	os.Exit(0)
}
