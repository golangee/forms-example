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

package circularprogress

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/circularprogress"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Circular Progress").Style(Font(Headline1)),
		NewText("An entertaining and progress indicator component in a circle.").Style(Font(Body)),

		NewCircularProgress(),


		NewCode(GoSyntax, code),
	)

	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package circularprogress

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/circularprogress"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Circular Progress").Style(Font(Headline1)),
		NewText("An entertaining and progress indicator component in a circle.").Style(Font(Body)),

		NewCircularProgress(),


		NewCode(GoSyntax, code),
	)

	return view
}`
