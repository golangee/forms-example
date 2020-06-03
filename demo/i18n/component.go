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
	"github.com/golangee/forms/locale"
	"github.com/golangee/i18n"
)

const Path = "/demo/i18n"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	res := NewResources(locale.Language())

	view := &ContentView{VStack: NewVStack()}
	view.VStack.AddViews(
		NewText(res.DemoI18NCaption()).Style(Font(Headline1)),
		NewText(res.DemoI18NText()).Style(Font(Body), Repel()),

		NewText(res.DemoI18NBrowserLocale(locale.Language())),
		NewText(res.DemoI18NActiveLocale(res.Locale())).Style(Repel()),
		NewPicker().
			SetLabel(res.DemoI18NSelectLanguage()).
			SetOptions(i18n.Locales()...).
			SetSelectedString(locale.Language()).
			SetSelectListener(func(v *Picker) {
				locale.SetLanguages(v.SelectedString())
				v.Context().Invalidate()
			}),

		NewText(res.HelloWorld()),
		NewText(res.HelloX("Dude")),
		NewText(res.XCats(1, 1)),
		NewText(res.XCats(5, 5)),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package i18n

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/locale"
	"github.com/golangee/i18n"
)

const Path = "/demo/i18n"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	res := NewResources(locale.Language())

	view := &ContentView{VStack: NewVStack()}
	view.VStack.AddViews(
		NewText(res.DemoI18NCaption()).Style(Font(Headline1)),
		NewText(res.DemoI18NText()).Style(Font(Body), Repel()),

		NewText(res.DemoI18NBrowserLocale(locale.Language())),
		NewText(res.DemoI18NActiveLocale(res.Locale())).Style(Repel()),
		NewPicker().
			SetLabel(res.DemoI18NSelectLanguage()).
			SetOptions(i18n.Locales()...).
			SetSelectedString(locale.Language()).
			SetSelectListener(func(v *Picker) {
				locale.SetLanguages(v.SelectedString())
				v.Context().Invalidate()
			}),

		NewText(res.HelloWorld()),
		NewText(res.HelloX("Dude")),
		NewText(res.XCats(1, 1)),
		NewText(res.XCats(5, 5)),

		NewCode(GoSyntax, code),
	)
	return view
}`
