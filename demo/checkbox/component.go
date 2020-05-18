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

package checkbox

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/checkbox"

type ContentView struct {
	*VStack
	checkbox1 *Checkbox
	btn       *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Checkbox").Style(Font(Headline1)),
		NewText("Selects or deselects one or more items. ").Style(Font(Body)),
		NewCheckbox().SetText("checkbox 1").Self(&view.checkbox1),
		NewCheckbox().SetText("checkbox 2").AddChangeListener(func(v *Checkbox) {
			view.checkbox1.SetEnabled(v.Checked())
		}).SetChecked(true),
		NewCheckbox().SetText("indeterminate").SetIndeterminate(true),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package checkbox

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/checkbox"

type ContentView struct {
	*VStack
	checkbox1 *Checkbox
	btn       *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Checkbox").Style(Font(Headline1)),
		NewText("Selects or deselects one or more items. ").Style(Font(Body)),
		NewCheckbox().SetText("checkbox 1").Self(&view.checkbox1),
		NewCheckbox().SetText("checkbox 2").AddChangeListener(func(v *Checkbox) {
			view.checkbox1.SetEnabled(v.Checked())
		}).SetChecked(true),
		NewCheckbox().SetText("indeterminate").SetIndeterminate(true),

		NewCode(GoSyntax, code),
	)
	return view
}`
