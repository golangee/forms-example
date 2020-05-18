// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package textarea

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/textarea"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("textarea").
			Style(Font(Headline1)),
		NewText("Text areas are like text fields but usually only used for "+
			"multiline input.").
			Style(Font(Body), Repel()),

		NewText("Default"),
		NewTextArea().
			SetLabel("your multiline input").
			Style(Repel()),

		NewText("Default prefilled"),
		NewTextArea().
			SetText("hello\nworld.").
			SetLabel("your multiline input").
			Style(Repel()),

		NewText("Default with max length"),
		NewTextArea().
			SetLabel("your multiline input").
			SetMaxLength(10).
			Style(Repel()),

		NewText("disabled"),
		NewTextArea().
			SetLabel("your multiline input").
			SetEnabled(false).
			Style(Repel()),

		NewText("invalid"),
		NewTextArea().
			SetLabel("your multiline input").
			SetInvalid(true).
			Style(Repel()),

		NewCode(GoSyntax, code),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package textarea

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/textarea"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("textarea").
			Style(Font(Headline1)),
		NewText("Text areas are like text fields but usually only used for "+
			"multiline input.").
			Style(Font(Body), Repel()),

		NewText("Default"),
		NewTextArea().
			SetLabel("your multiline input").
			Style(Repel()),

		NewText("Default prefilled"),
		NewTextArea().
			SetText("hello\nworld.").
			SetLabel("your multiline input").
			Style(Repel()),

		NewText("Default with max length"),
		NewTextArea().
			SetLabel("your multiline input").
			SetMaxLength(10).
			Style(Repel()),

		NewCode(GoSyntax, code),
	)}
}`
