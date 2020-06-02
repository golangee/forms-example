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

package i18n

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/i18n"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{VStack: NewVStack()}
	view.VStack.AddViews(
		NewText("Localization").Style(Font(Headline1)),
		NewText("This is not a core feature by forms itself, but from the ee/i18n module. "+
			"Using the embedding code generator, makes everything typesafe but requires space in wasm binary and "+
			"memory. The recommendation is to just use a base language and load a unified translation file at "+
			"runtime.").Style(Font(Body), Repel()),

		NewText(NewResources(view.Context().Languages()[0]).HelloWorld()),
		NewText(NewResources(view.Context().Languages()[0]).XHasYCats2(1, "Peter", 1)),
		NewText(NewResources(view.Context().Languages()[0]).XHasYCats2(2, "Peter", 2)),
		NewText(NewResources(view.Context().Languages()[0]).HelloX("dude")).Style(Repel()),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package link

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/link"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Link").Style(Font(Headline1)),
		NewText("A simple inline link "+
			" for text based non-button and non-icon navigation.").Style(Font(Body)),
		NewGroup(
			NewText("hello "),
			NewLink("world", "http://www.worldiety.de").SetTarget(TargetBlank),
		),

		NewCode(GoSyntax, code),
	)
	return view
}`
