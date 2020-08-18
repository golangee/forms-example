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

package autocomplete

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	"log"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

const Path = "/demo/autocomplete"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Autocomplete").Style(Font(Headline1)),
		NewText("An auto completion is important to improve usability.").Style(Font(Body)),
		NewAutocompleteView(),
		NewCode(GoSyntax, code),
	)
	return view
}

type AutocompleteView struct {
	*VStack
	textField             *TextField
	lastSearchRequest     int32 // we need to check if we surpass ourself
	autocompleteContainer *Frame
	onBuildPopup          func() View
}

func NewAutocompleteView() *AutocompleteView {
	view := &AutocompleteView{}
	view.VStack = NewVStack(
		NewTextField().Self(&view.textField).AddKeyUpListener(func(v View, keyCode int) {
			if keyCode == 27 { //escape key
				view.autocompleteContainer.ClearViews()
				return
			}

			view.triggerSearch()
		}).AddFocusChangeListener(func(v View, hasFocus bool) {
			if hasFocus {
				view.triggerSearch()
			} else {
				//view.autocompleteContainer.ClearViews() //TODO this does not work, because focusout is always earlier
			}
		}).SetLeadingIcon(icon.AccessAlarms).
			AddLeadingIconClickListener(func(v View) {
				ShowMessage(v, "leading icon click")
			}).SetTrailingIcon(icon.Domain).
			AddTrailingIconClickListener(func(v View) {
				ShowMessage(v, "trailing icon click")
			}),
		NewFrame().Self(&view.autocompleteContainer),
	)

	return view
}

func (c *AutocompleteView) triggerSearch() {
	if len(c.textField.Text()) < 3 {
		c.textField.SetHelper("not enough text for autocomplete")
		return
	}

	c.textField.SetHelper("")
	c.simulateSearch(c.textField.Text(), func(suggestions []string, outdated bool) {
		if outdated {
			log.Printf("result is outdated\n")
			return
		}

		table := NewTable().SetRowClickListener(func(v View, rowIdx int) {
			c.textField.SetText(suggestions[rowIdx])
			c.autocompleteContainer.ClearViews()
		}).Style(Position(PositionAbsolute))

		for i, suggestion := range suggestions {
			table.AddRow(NewText(strconv.Itoa(i)), NewText(suggestion))
		}

		c.autocompleteContainer.SetView(table)
		log.Printf("%v\n", suggestions)
	})
}

func (c *AutocompleteView) simulateSearch(text string, callback func(suggestions []string, outdated bool)) {
	generation := atomic.AddInt32(&c.lastSearchRequest, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // network requests may take a random time
		outdated := atomic.LoadInt32(&c.lastSearchRequest) != generation
		callback([]string{text + "-hello", text + "-world", text + "-gopher"}, outdated)
	}()
}

func FromQuery(Query) View {
	return NewContentView()
}

const code = `package autocomplete

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	"log"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

const Path = "/demo/autocomplete"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("Autocomplete").Style(Font(Headline1)),
		NewText("An auto completion is important to improve usability.").Style(Font(Body)),
		NewAutocompleteView(),
		NewCode(GoSyntax, code),
	)
	return view
}

type AutocompleteView struct {
	*VStack
	textField             *TextField
	lastSearchRequest     int32 // we need to check if we surpass ourself
	autocompleteContainer *Frame
	onBuildPopup          func() View
}

func NewAutocompleteView() *AutocompleteView {
	view := &AutocompleteView{}
	view.VStack = NewVStack(
		NewTextField().Self(&view.textField).AddKeyUpListener(func(v View, keyCode int) {
			if keyCode == 27 { //escape key
				view.autocompleteContainer.ClearViews()
				return
			}

			view.triggerSearch()
		}).AddFocusChangeListener(func(v View, hasFocus bool) {
			if hasFocus {
				view.triggerSearch()
			} else {
				//view.autocompleteContainer.ClearViews() //TODO this does not work, because focusout is always earlier
			}
		}).SetLeadingIcon(icon.AccessAlarms).
			AddLeadingIconClickListener(func(v View) {
				ShowMessage(v, "leading icon click")
			}).SetTrailingIcon(icon.Domain).
			AddTrailingIconClickListener(func(v View) {
				ShowMessage(v, "trailing icon click")
			}),
		NewFrame().Self(&view.autocompleteContainer),
	)

	return view
}

func (c *AutocompleteView) triggerSearch() {
	if len(c.textField.Text()) < 3 {
		c.textField.SetHelper("not enough text for autocomplete")
		return
	}

	c.textField.SetHelper("")
	c.simulateSearch(c.textField.Text(), func(suggestions []string, outdated bool) {
		if outdated {
			log.Printf("result is outdated\n")
			return
		}

		table := NewTable().SetRowClickListener(func(v View, rowIdx int) {
			c.textField.SetText(suggestions[rowIdx])
			c.autocompleteContainer.ClearViews()
		}).Style(Position(PositionAbsolute))

		for i, suggestion := range suggestions {
			table.AddRow(NewText(strconv.Itoa(i)), NewText(suggestion))
		}

		c.autocompleteContainer.SetView(table)
		log.Printf("%v\n", suggestions)
	})
}

func (c *AutocompleteView) simulateSearch(text string, callback func(suggestions []string, outdated bool)) {
	generation := atomic.AddInt32(&c.lastSearchRequest, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // network requests may take a random time
		outdated := atomic.LoadInt32(&c.lastSearchRequest) != generation
		callback([]string{text + "-hello", text + "-world", text + "-gopher"}, outdated)
	}()
}

func FromQuery(Query) View {
	return NewContentView()
}`
