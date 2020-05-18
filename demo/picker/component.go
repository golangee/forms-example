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

package picker

import (
	. "github.com/golangee/forms"
	"strconv"
)

const Path = "/demo/picker"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Picker").Style(Font(Headline1)),
		NewText("A picker allows to select exactly one option from a fixed list.").Style(Font(Body)),
		NewPicker().
			SetLabel("select your meal").
			SetOptions("", "bread", "butter", "salt").
			SetSelected(1).
			SetSelectListener(func(v *Picker) {
				ShowMessage(view, "you selected index "+strconv.Itoa(v.Selected()))
			}).
			Style(Repel()),

		NewPicker().SetLabel("disabled").SetHelper("you can't do anything").SetEnabled(false).
			Style(Repel()),

		NewPicker().
			SetLabel("select invalid").
			SetHelper("you got something wrong").
			SetInvalid(true).
			SetOptions("a", "b", "c").
			Style(Repel()),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package picker

import (
	. "github.com/golangee/forms"
	"strconv"
)

const Path = "/demo/picker"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Picker").Style(Font(Headline1)),
		NewText("A picker allows to select exactly one option from a fixed list.").Style(Font(Body)),
		NewPicker().
			SetLabel("select your meal").
			SetOptions("", "bread", "butter", "salt").
			SetSelected(1).
			SetSelectListener(func(v *Picker) {
				ShowMessage(view, "you selected index "+strconv.Itoa(v.Selected()))
			}).
			Style(Repel()),

		NewPicker().SetLabel("disabled").SetHelper("you can't do anything").SetEnabled(false).
			Style(Repel()),
			
		NewPicker().
			SetLabel("select invalid").
			SetHelper("you got something wrong").
			SetInvalid(true).
			SetOptions("a", "b", "c").
			Style(Repel()),

		NewCode(GoSyntax, code),
	)
	return view
}`
