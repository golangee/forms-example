package section1

import . "github.com/worldiety/wtk"


type ContentView struct {
	*Text
}

func NewContentView() *ContentView {
	return &ContentView{Text: NewText("hello world")}
}
