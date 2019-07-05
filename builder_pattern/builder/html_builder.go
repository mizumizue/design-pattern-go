package builder

import (
	"fmt"
	"io/ioutil"
	"os"
)

type HtmlBuilder struct {
	filename string
	writeBuf string
}

func NewHtmlBuilder() *HtmlBuilder {
	return &HtmlBuilder{}
}

func (hb *HtmlBuilder) MakeTitle(title string) {
	hb.filename = title + ".html"
	hb.writeBuf += "<html><header><title>" + title + "</title></header><body>"
	hb.writeBuf += "<h1>" + title + "</h1>"
}

func (hb *HtmlBuilder) MakeString(str string) {
	hb.writeBuf += "<p>" + str + "</p>"
}

func (hb *HtmlBuilder) MakeItems(items []string) {
	for _, item := range items {
		hb.writeBuf += "<p>" + item + "</p>"
	}
}

func (hb *HtmlBuilder) Close() {
	hb.writeBuf += "</body></html>"
	if err := ioutil.WriteFile(hb.filename, []byte(hb.writeBuf), os.ModePerm); err != nil {
		_ = fmt.Errorf("write failure. filename: %s, content: %s, err: %v", hb.filename, hb.writeBuf, err)
	}
}

func (hb *HtmlBuilder) GetFilename() string {
	return hb.filename
}
