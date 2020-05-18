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

package topappbar

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/topappbar"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Top App Bar").Style(Font(Headline1)),
		NewText("The top app bar is a container for navigation (menu, back or up), "+
			" application title and actions.").Style(Font(Body)),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package topappbar

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	"log"
)

const Path = "/demo/topappbar"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Top App Bar").Style(Font(Headline1)),
		NewText("The top app bar is a container for navigation (menu, back or up), "+
			" application title and actions.").Style(Font(Body)),
		NewTopAppBar().
			SetNavigation(icon.Menu, func(view View) {
				log.Println("pressed the menu")
			}).
			SetTitle("my action bar").
			AddActions(
				NewIconItem(icon.FileDownload, "download", func(v View) {
					log.Println("download")
				}),
				NewIconItem(icon.Print, "print", func(v View) {
					log.Print("print")
				}),
			),

		NewCode(GoSyntax, code),
	)
	return view
}`
