package picker

import (
	. "github.com/worldiety/wtk"
	"strconv"
)

const Path = "/demo/picker"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Picker").Style(Font(Headline1)),
		NewText("A picker allows to select exactly one option from a fixed list.").Style(Font(Body)),
		NewPicker().
			SetCaption("select your meal").
			SetOptions("", "bread", "butter", "salt").
			SetSelected(1).
			SetSelectListener(func(v *Picker) {
				ShowMessage(view, "you selected index "+strconv.Itoa(v.Selected()))
			}),
		NewPicker().SetCaption("disabled").SetEnabled(false),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = ``
