package button

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
)

const Path = "/demo/button"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("buttons").Style(Font(Headline1)),
		NewText("Buttons are used to allow users to make actions. Here we show how"+
			" to use them in different styles and with trailing and leading icons.").Style(Font(Body)),
		NewButton("default button").
			Style(Margin()),
		NewButton("outline button").
			SetStyleKind(Outlined).
			Style(Margin()),
		NewButton("raised button").
			SetStyleKind(Raised).
			Style(Margin()),
		NewButton("unelevated button").
			SetStyleKind(Unelevated).
			Style(Margin()),
		NewButton("button with trailing icon").
			AddIcon(icon.AccessAlarm, Trailing).
			Style(Margin()),
		NewButton("button with leading icon").
			AddIcon(icon.AccessAlarm, Leading).
			Style(Margin()),
		NewButton("button with both icons").
			AddIcon(icon.AccessAlarm, Leading).
			AddIcon(icon.Call, Trailing).
			Style(Margin()),
		NewButton("click action").
			AddClickListener(func(v View) {
				v.(*Button).SetEnabled(false).SetText("disabled")
			}).
			Style(Margin()),
		NewCode(GoSyntax, code),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package button

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
)

const Path = "/demo/button"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("button demo").Style(Font(Title)),
		NewText("Buttons are used to allow users to make actions. Here we show how"+
			" to use them in different styles and with trailing and leading icons.").Style(Font(Body)),
		NewButton("default button").
			Style(Margin()),
		NewButton("outline button").
			SetStyleKind(Outlined).
			Style(Margin()),
		NewButton("raised button").
			SetStyleKind(Raised).
			Style(Margin()),
		NewButton("unelevated button").
			SetStyleKind(Unelevated).
			Style(Margin()),
		NewButton("button with trailing icon").
			AddIcon(icon.AccessAlarm, Trailing).
			Style(Margin()),
		NewButton("button with leading icon").
			AddIcon(icon.AccessAlarm, Leading).
			Style(Margin()),
		NewButton("button with both icons").
			AddIcon(icon.AccessAlarm, Leading).
			AddIcon(icon.Call, Trailing).
			Style(Margin()),
		NewButton("click action").
			AddClickListener(func(v View) {
				v.(*Button).SetEnabled(false).SetText("disabled")
			}).
			Style(Margin()),
		NewCode(code),
	)}
}`
