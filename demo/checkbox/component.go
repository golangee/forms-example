package checkbox

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/checkbox"

type ContentView struct {
	*VStack
	checkbox1 *Checkbox
	btn       *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Checkbox").Style(Font(Headline1)),
		NewText("Selects or deselects one or more items. ").Style(Font(Body)),
		NewCheckbox().SetText("checkbox 1").Self(&view.checkbox1),
		NewCheckbox().SetText("checkbox 2").AddChangeListener(func(v *Checkbox) {
			view.checkbox1.SetEnabled(v.Checked())
		}).SetChecked(true),
		NewCheckbox().SetText("indeterminate").SetIndeterminate(true),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package checkbox

import (
	. "github.com/worldiety/wtk"
)

const Path = "/demo/checkbox"

type ContentView struct {
	*VStack
	checkbox1 *Checkbox
	btn       *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Checkbox").Style(Font(Headline1)),
		NewText("Selects or deselects one or more items. ").Style(Font(Body)),
		NewCheckbox().SetText("checkbox 1").Self(&view.checkbox1),
		NewCheckbox().SetText("checkbox 2").AddChangeListener(func(v *Checkbox) {
			view.checkbox1.SetEnabled(v.Checked())
		}).SetChecked(true),
		NewCheckbox().SetText("indeterminate").SetIndeterminate(true),

		NewCode(GoSyntax, code),
	)
	return view
}`
