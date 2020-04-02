package section3

import . "github.com/worldiety/wtk"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{
		VStack: NewVStack().AddViews(
			NewText("Turtle Rock").
				Style(Font(Title)),
			NewHStack().AddViews(
				NewText("Joshua Tree National Park").
					Style(Font(SubHeading)),
				NewSpacer(),
				NewText("California").
					Style(Font(SubHeading)),
			),

		),
	}
}
