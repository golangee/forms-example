package notfound

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk-example/demo/index"
)

const Path = "/demo"

type ContentView struct {
	*VStack
}

func NewContentView(path string) *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("the route '"+path+"' is not available").Style(Font(Headline1)),
		NewButton("Index").AddClickListener(func(v View) {
			v.Context().Navigate(index.Path)
		}),
	)}
}

func FromQuery(q Query) View {
	return NewContentView(q.Path())
}
