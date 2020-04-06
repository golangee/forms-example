package textfield

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
)

const Path = "/demo/textfield"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("textfields").
			Style(Font(Headline1)),
		NewText("Text fields allow users to insert, edit and select text content. "+
			"They are surprisingly complex, to allow a lot of user help and feedback.").
			Style(Font(Body), Repel()),

		NewText("Default"),
		NewTextField().Style(Repel()),

		NewText("Default full width"),
		NewTextField().
			InputStyle(Width(Percent(100))).
			Style(Repel()),

		NewText("Password"),
		NewTextField().
			SetInputType(Password).
			Style(Repel()),

		NewText("Number"),
		NewTextField().
			SetInputType(Number).
			Style(Repel()),

		NewText("Range"),
		NewTextField().
			SetRange(2, 5).
			Style(Repel()),

		NewText("Date"),
		NewTextField().
			SetInputType(Date).
			Style(Repel()),

		NewText("Default prefilled"),
		NewTextField().
			SetText("my content").
			Style(Repel()),

		NewText("Default with label"),
		NewTextField().
			SetLabel("my label").
			Style(Repel()),

		NewText("Default with label and leading icon"),
		NewTextField().
			SetLabel("my label").
			SetLeadingIcon(icon.Favorite).
			Style(Repel()),

		NewText("Default with label and trailing icon"),
		NewTextField().
			SetLabel("my label").
			SetTrailingIcon(icon.Visibility).
			Style(Repel()),

		NewText("Default with label and leading trailing icon"),
		NewTextField().
			SetLabel("my label").
			SetLeadingIcon(icon.Favorite).
			SetTrailingIcon(icon.Visibility).
			Style(Repel()),

		NewText("Default with label prefilled"),
		NewTextField().
			SetText("my content").
			SetLabel("my label").
			Style(Repel()),

		NewText("Default with label prefilled disabled"),
		NewTextField().
			SetText("my content").
			SetLabel("my label").
			SetEnabled(false).
			Style(Repel()),

		NewText("Default with helper"),
		NewTextField().
			SetLabel("my label").
			SetHelper("more text to help you filling out the form").
			Style(Repel()),

		NewText("Default with max length"),
		NewTextField().
			SetLabel("my label").
			SetHelper("more text to help you filling out the form").
			SetMaxLength(5).
			Style(Repel()),

		NewText("Default with required"),
		NewTextField().
			SetLabel("my label").
			SetHelper("more text to help you filling out the form").
			SetRequired(true).
			Style(Repel()),

		NewText("Default with invalid flag"),
		NewTextField().
			SetLabel("my label").
			SetHelper("this is not correct").
			SetInvalid(true).
			Style(Repel()),


		NewCode(GoSyntax, code),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package textfield

import (
	. "github.com/worldiety/wtk"
	"github.com/worldiety/wtk/theme/material/icon"
)

const Path = "/demo/textfield"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("textfields").
			Style(Font(Headline1)),
		NewText("Text fields allow users to insert, edit and select text content. "+
			"They are surprisingly complex, to allow a lot of user help and feedback.").
			Style(Font(Body), Repel()),

		NewText("Default"),
		NewTextField().Style(Repel()),

		NewText("Default full width"),
		NewTextField().
			InputStyle(Width(Percent(100))).
			Style(Repel()),

		NewText("Password"),
		NewTextField().
			SetInputType(Password).
			Style(Repel()),

		NewText("Number"),
		NewTextField().
			SetInputType(Number).
			Style(Repel()),

		NewText("Range"),
		NewTextField().
			SetRange(2, 5).
			Style(Repel()),

		NewText("Date"),
		NewTextField().
			SetInputType(Date).
			Style(Repel()),

		NewText("Default prefilled"),
		NewTextField().
			SetText("my content").
			Style(Repel()),

		NewText("Default with label"),
		NewTextField().
			SetLabel("my label").
			Style(Repel()),

		NewText("Default with label and leading icon"),
		NewTextField().
			SetLabel("my label").
			SetLeadingIcon(icon.Favorite).
			Style(Repel()),

		NewText("Default with label and trailing icon"),
		NewTextField().
			SetLabel("my label").
			SetTrailingIcon(icon.Visibility).
			Style(Repel()),

		NewText("Default with label and leading trailing icon"),
		NewTextField().
			SetLabel("my label").
			SetLeadingIcon(icon.Favorite).
			SetTrailingIcon(icon.Visibility).
			Style(Repel()),

		NewText("Default with label prefilled"),
		NewTextField().
			SetText("my content").
			SetLabel("my label").
			Style(Repel()),

		NewText("Default with label prefilled disabled"),
		NewTextField().
			SetText("my content").
			SetLabel("my label").
			SetEnabled(false).
			Style(Repel()),

		NewText("Default with helper"),
		NewTextField().
			SetLabel("my label").
			SetHelper("more text to help you filling out the form").
			Style(Repel()),

		NewText("Default with max length"),
		NewTextField().
			SetLabel("my label").
			SetHelper("more text to help you filling out the form").
			SetMaxLength(5).
			Style(Repel()),

		NewText("Default with required"),
		NewTextField().
			SetLabel("my label").
			SetHelper("more text to help you filling out the form").
			SetRequired(true).
			Style(Repel()),

		NewText("Default with invalid flag"),
		NewTextField().
			SetLabel("my label").
			SetHelper("this is not correct").
			SetInvalid(true).
			Style(Repel()),


		NewCode(GoSyntax, code),
	)}
}`
