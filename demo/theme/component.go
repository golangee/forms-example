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

package theme

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/theme"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	colorButtons := NewGroup()
	for name, color := range Colors {
		myColor := color
		colorButtons.AddViews(
			NewButton(name).
				AddClickListener(func(v View) {
					Theme().SetColor(myColor)
				}).Style(ForegroundColor(myColor)),
		)
	}

	view.VStack = NewVStack().AddViews(
		NewText("Theme").Style(Font(Headline1)),
		NewText("Colors are important to increase the recognition factor "+
			" and to focus on sensitive actions.").Style(Font(Body)),
		colorButtons,
		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package theme

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/theme"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	colorButtons := NewGroup()
	for name, color := range Colors {
		myColor := color
		colorButtons.AddViews(
			NewButton(name).
				AddClickListener(func(v View) {
					Theme().SetPrimaryColor(myColor)
				}).Style(ForegroundColor(myColor)),
		)
	}

	view.VStack = NewVStack().AddViews(
		NewText("Theme").Style(Font(Headline1)),
		NewText("Colors are important to increase the recognition factor "+
			" and to focus on sensitive actions.").Style(Font(Body)),
		colorButtons,
		NewCode(GoSyntax, code),
	)
	return view
}`
