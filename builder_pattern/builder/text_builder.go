package builder

type TextBuilder struct {
	buffer string
}

func NewTextBuilder() *TextBuilder {
	return &TextBuilder{}
}

func (tb *TextBuilder) MakeTitle(title string) {
	tb.buffer = tb.buffer + "=================================\n"
	tb.buffer = tb.buffer + "「" + title + "」"
	tb.buffer = tb.buffer + "\n"
}

func (tb *TextBuilder) MakeString(str string) {
	tb.buffer = tb.buffer + "■" + str
	tb.buffer = tb.buffer + "\n"
}

func (tb *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		tb.buffer = tb.buffer + "・" + item
		tb.buffer = tb.buffer + "\n"
	}
}

func (tb *TextBuilder) Close() {
	tb.buffer = tb.buffer + "=================================\n"
}

func (tb *TextBuilder) GetResult() string {
	return tb.buffer
}
