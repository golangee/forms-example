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

package button

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
)

const Path = "/demo/button"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("buttons").Style(Font(Headline1)),
		NewText("Buttons are used to allow users to make actions. Here we show how"+
			" to use them in different styles and with trailing and leading icons.").Style(Font(Body)),
		NewButton("default button").
			Style(Margin()),
		NewButton("outline button").
			SetStyleKind(Outlined).
			Style(Margin()),
		NewButton("raised button").
			SetStyleKind(Raised).
			Style(Margin()),
		NewButton("unelevated button").
			SetStyleKind(Unelevated).
			Style(Margin()),
		NewButton("button with trailing icon").
			AddIcon(icon.AccessAlarm, Trailing).
			Style(Margin()),
		NewButton("button with leading icon").
			AddIcon(icon.AccessAlarm, Leading).
			Style(Margin()),
		NewButton("button with both icons").
			AddIcon(icon.AccessAlarm, Leading).
			AddIcon(icon.Call, Trailing).
			Style(Margin()),
		NewButton("click action").
			AddClickListener(func(v View) {
				v.(*Button).SetEnabled(false).SetText("disabled")
			}).
			Style(Margin()),
		NewIconButton(icon.Add),
		NewIconButton(icon.Add).Style(
			BackgroundColor(Theme().Color()),
			ForegroundColor(Theme().ForegroundColor()),
		),
		NewIconButton(icon.Add).SetChar('1'),
		NewCode(GoSyntax, code),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package button

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
)

const Path = "/demo/button"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("buttons").Style(Font(Headline1)),
		NewText("Buttons are used to allow users to make actions. Here we show how"+
			" to use them in different styles and with trailing and leading icons.").Style(Font(Body)),
		NewButton("default button").
			Style(Margin()),
		NewButton("outline button").
			SetStyleKind(Outlined).
			Style(Margin()),
		NewButton("raised button").
			SetStyleKind(Raised).
			Style(Margin()),
		NewButton("unelevated button").
			SetStyleKind(Unelevated).
			Style(Margin()),
		NewButton("button with trailing icon").
			AddIcon(icon.AccessAlarm, Trailing).
			Style(Margin()),
		NewButton("button with leading icon").
			AddIcon(icon.AccessAlarm, Leading).
			Style(Margin()),
		NewButton("button with both icons").
			AddIcon(icon.AccessAlarm, Leading).
			AddIcon(icon.Call, Trailing).
			Style(Margin()),
		NewButton("click action").
			AddClickListener(func(v View) {
				v.(*Button).SetEnabled(false).SetText("disabled")
			}).
			Style(Margin()),
		NewIconButton(icon.Add),
		NewIconButton(icon.Add).Style(
			BackgroundColor(Theme().Color()),
			ForegroundColor(Theme().ForegroundColor()),
		),
		NewIconButton(icon.Add).SetChar('1'),
		NewCode(GoSyntax, code),
	)}
}`
