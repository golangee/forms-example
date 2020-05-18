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

package linearprogress

import (
	. "github.com/golangee/forms"
	"time"
)

const Path = "/demo/linearprogress"

type ContentView struct {
	*VStack
	btn         *Button
	progressBar *LinearProgress
	released    bool
	progress    float64
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Linear Progress").Style(Font(Headline1)),
		NewText("An entertaining and progress indicator component.").Style(Font(Body)),

		NewText("normal:"),
		NewLinearProgress().Self(&view.progressBar).Style(Repel()),

		NewText("indeterminate:"),
		NewLinearProgress().SetIndeterminate(true).Style(Repel()),

		NewText("secondary:"),
		NewLinearProgress().SetProgress(0.3).SetSecondaryProgress(0.7).Style(Repel()),

		NewCode(GoSyntax, code),
	)

	go func() {
		for !view.released {
			time.Sleep(500 * time.Millisecond)
			view.progress += 0.01
			if view.progress > 1 {
				view.progress = 0
			}
			view.progressBar.SetProgress(view.progress)
		}
	}()
	return view
}

func (t *ContentView) Release() {
	t.released = true
	t.VStack.Release()
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package linearprogress

import (
	. "github.com/golangee/forms"
	"time"
)

const Path = "/demo/linearprogress"

type ContentView struct {
	*VStack
	btn         *Button
	progressBar *LinearProgress
	released    bool
	progress    float64
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Linear Progress").Style(Font(Headline1)),
		NewText("An entertaining and progress indicator component.").Style(Font(Body)),

		NewText("normal:"),
		NewLinearProgress().Self(&view.progressBar).Style(Repel()),

		NewText("indeterminate:"),
		NewLinearProgress().SetIndeterminate(true).Style(Repel()),

		NewText("secondary:"),
		NewLinearProgress().SetProgress(0.3).SetSecondaryProgress(0.7).Style(Repel()),

		NewCode(GoSyntax, code),
	)

	go func() {
		for !view.released {
			time.Sleep(500 * time.Millisecond)
			view.progress += 0.01
			if view.progress > 1 {
				view.progress = 0
			}
			view.progressBar.SetProgress(view.progress)
		}
	}()
	return view
}

func (t *ContentView) Release() {
	t.released = true
	t.VStack.Release()
}`
