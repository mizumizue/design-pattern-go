package director

import (
	"github.com/trewanek/PracticeDesignPatternWithGo/builder_pattern/builder"
)

type Director struct {
	builder builder.Builder
}

func NewDirector(builder builder.Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.MakeTitle("Greeting")
	d.builder.MakeString("From morning to noon")
	d.builder.MakeItems([]string{
		"Good morning",
		"Hello",
	})
	d.builder.MakeString("In night")
	d.builder.MakeItems([]string{
		"Good evening",
		"Good night",
		"Goodbye",
	})
	d.builder.Close()
}
