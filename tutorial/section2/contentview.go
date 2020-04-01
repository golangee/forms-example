package section2

import . "github.com/worldiety/wtk"

type ContentView struct {
	*Text
}

func NewContentView() *ContentView {
	return &ContentView{
		Text: NewText("Turtle Rock").
			Style(
				ForegroundColor(Green),
				Font(Title),
			),
	}
}
