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

package grid

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/grid"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Grid").Style(Font(Headline1)),
		NewText("A grid allows complex grid based layouts.").Style(Font(Body)),

		NewGrid().
			SetAreas([][]string{
				{"header", "header", "header"},
				{"menu", "main", "main"},
				{"menu", "footer", "footer"},
			}).
			AddView(NewText("header").Style(BackgroundColor(Red50)), GridLayoutParams{Area: "header"}).
			AddView(NewText("menu").Style(BackgroundColor(Blue50)), GridLayoutParams{Area: "menu"}).
			AddView(NewText("main").Style(BackgroundColor(Yellow50)), GridLayoutParams{Area: "main"}).
			AddView(NewText("footer").Style(BackgroundColor(Green50)), GridLayoutParams{Area: "footer"}).
			Style(BackgroundColor(BlueGray50),Padding()).
			SetGap(Pixel(DefaultPadding)),
		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package grid

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/grid"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Grid").Style(Font(Headline1)),
		NewText("A grid allows complex grid based layouts.").Style(Font(Body)),

		NewGrid().
			SetAreas([][]string{
				{"header", "header", "header"},
				{"menu", "main", "main"},
				{"menu", "footer", "footer"},
			}).
			AddView(NewText("header").Style(BackgroundColor(Red50)), GridLayoutParams{Area: "header"}).
			AddView(NewText("menu").Style(BackgroundColor(Blue50)), GridLayoutParams{Area: "menu"}).
			AddView(NewText("main").Style(BackgroundColor(Yellow50)), GridLayoutParams{Area: "main"}).
			AddView(NewText("footer").Style(BackgroundColor(Green50)), GridLayoutParams{Area: "footer"}).
			Style(BackgroundColor(BlueGray50),Padding()).
			SetGap(Pixel(DefaultPadding)),
		NewCode(GoSyntax, code),
	)
	return view
}`
