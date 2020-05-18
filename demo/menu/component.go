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

package menu

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/menu"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("menu").Style(Font(Headline1)),
		NewText("Menu shows a list of actions in a dialog-like "+
			" container.").Style(Font(Body)),
		NewButton("show menu").Self(&view.btn).
			AddClickListener(func(v View) {
				ShowMenu(v,
					NewMenuItem("first item", func(menu *MenuItem) {
						view.btn.SetText("you clicked the first")
					}),
					NewMenuItem("second item", func(menu *MenuItem) {
						view.btn.SetText("you clicked the second")
					}),
					NewMenuDivider(),
					NewMenuItem("third item", nil).SetEnabled(false),
				)
			}).
			Style(Margin()),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package menu

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/menu"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("menu").Style(Font(Headline1)),
		NewText("Menu shows a list of actions in a dialog-like "+
			" container.").Style(Font(Body)),
		NewButton("show menu").Self(&view.btn).
			AddClickListener(func(v View) {
				ShowMenu(v,
					NewMenuItem("first item", func(menu *MenuItem) {
						view.btn.SetText("you clicked the first")
					}),
					NewMenuItem("second item", func(menu *MenuItem) {
						view.btn.SetText("you clicked the second")
					}),
					NewMenuDivider(),
					NewMenuItem("third item", nil).SetEnabled(false),
				)
			}).
			Style(Margin()),

		NewCode(GoSyntax, code),
	)
	return view
}`
