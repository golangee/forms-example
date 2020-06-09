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

package hstepper

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	. "github.com/golangee/forms/views/hstepper"
)

const Path = "/demo/horizontalstepper"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("horizontal stepper").Style(Font(Headline1)),
		NewText("A stepper visualizes the progress through a sequence "+
			"of logical and (numbered) steps or navigation in "+
			"assistants.").Style(Font(Body)),

		NewStepper(
			NewIconStep(icon.Assignment, "Step 1"),
			NewIconStep(icon.Storage, "Step 2"),
			NewIconStep(icon.Folder, "Step 3"),
			NewStep("last step"),
		).SetProgress(2).Style(Repel()),

		NewCode(GoSyntax, code),
	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package hstepper

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	. "github.com/golangee/forms/views/hstepper"
)

const Path = "/demo/horizontalstepper"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("horizontal stepper").Style(Font(Headline1)),
		NewText("A stepper visualizes the progress through a sequence "+
			"of logical and (numbered) steps or navigation in "+
			"assistants.").Style(Font(Body)),

		NewStepper(
			NewIconStep(icon.Assignment, "Step 1"),
			NewIconStep(icon.Storage, "Step 2"),
			NewIconStep(icon.Folder, "Step 3"),
			NewStep("last step"),
		).SetProgress(2).Style(Repel()),

		NewCode(GoSyntax, code),
	)
	return view
}`
