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

package dialog

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/dialog"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("dialogs").Style(Font(Headline1)),
		NewText("Dialogs interact with a user about a task and usually "+
			" require important or destructive decisions.").Style(Font(Body)),
		NewButton("show dialog").
			AddClickListener(func(v View) {
				NewDialog().
					SetTitle("are you sure?").
					SetBody(NewVStack().AddViews(
						NewText("this could go wrong"),
						NewTextField().SetLabel("enter something"),
					)).
					AddAction("perhaps", func(dlg *Dialog) {
						dlg.Close()
					}).
					AddAction("may be", func(dlg *Dialog) {
						dlg.Close()
					}).
					Show(v)
			}).
			Style(Margin()),

		NewCode(GoSyntax, code),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package dialog

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/dialog"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("dialogs").Style(Font(Headline1)),
		NewText("Dialogs interact with a user about a task and usually "+
			" require important or destructive decisions.").Style(Font(Body)),
		NewButton("show dialog").
			AddClickListener(func(v View) {
				NewDialog().
					SetTitle("are you sure?").
					SetBody(NewVStack().AddViews(
						NewText("this could go wrong"),
						NewTextField().SetLabel("enter something"),
					)).
					AddAction("perhaps", func(dlg *Dialog) {
						dlg.Close()
					}).
					AddAction("may be", func(dlg *Dialog) {
						dlg.Close()
					}).
					Show(v)
			}).
			Style(Margin()),

		NewCode(GoSyntax, code),
	)}
}`
