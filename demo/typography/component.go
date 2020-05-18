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

package typography

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/typography"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("typography").Style(Font(Headline1)),
		NewText("The material design text sizes and styles are well thought"+
			" and should be nice to read under typical screen conditions. These are "+
			"the available styles:").Style(Font(Body)),
		NewText("Headline1").Style(Font(Headline1), Margin()),
		NewText("Headline2").Style(Font(Headline2), Margin()),
		NewText("Headline3").Style(Font(Headline3), Margin()),
		NewText("Headline4").Style(Font(Headline4), Margin()),
		NewText("Headline5").Style(Font(Headline5), Margin()),
		NewText("Headline6").Style(Font(Headline6), Margin()),
		NewText("Subtitle1").Style(Font(Subtitle1), Margin()),
		NewText("Subtitle2").Style(Font(Subtitle2), Margin()),
		NewText("Body").Style(Font(Body), Margin()),
		NewText("Body2").Style(Font(Body2), Margin()),
		NewText("Caption").Style(Font(Caption), Margin()),
		NewText("Button").Style(Font(Btn), Margin()),
		NewText("Overline").Style(Font(Overline), Margin()),
		NewText("Default").Style(Margin()),


		NewCode(GoSyntax, code),
	)}
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package typography

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/typography"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("typographie").Style(Font(Headline1)),
		NewText("The material design text sizes and styles are well thought"+
			" and should be nice to read under typical screen conditions. These are "+
			"the available styles:").Style(Font(Body)),
		NewText("Headline1").Style(Font(Headline1), Margin()),
		NewText("Headline2").Style(Font(Headline2), Margin()),
		NewText("Headline3").Style(Font(Headline3), Margin()),
		NewText("Headline4").Style(Font(Headline4), Margin()),
		NewText("Headline5").Style(Font(Headline5), Margin()),
		NewText("Headline6").Style(Font(Headline6), Margin()),
		NewText("Subtitle1").Style(Font(Subtitle1), Margin()),
		NewText("Subtitle2").Style(Font(Subtitle2), Margin()),
		NewText("Body").Style(Font(Body), Margin()),
		NewText("Body2").Style(Font(Body2), Margin()),
		NewText("Caption").Style(Font(Caption), Margin()),
		NewText("Button").Style(Font(Btn), Margin()),
		NewText("Overline").Style(Font(Overline), Margin()),
		NewText("Default").Style(Margin()),


		NewCode(GoSyntax, code),
	)}
}`
