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

package link

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
