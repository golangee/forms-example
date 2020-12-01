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

package forms

import (
	"github.com/golangee/forms-example/www/forms/dom"
)

// SetPrimaryColor sets the current primary theme color
func SetPrimaryColor(color string) {
	style := dom.GetWindow().Document().DocumentElement().Style()
	style.SetProperty("--color-primary", color)
}

// SetSecondaryColor sets the current secondary theme color
func SetSecondaryColor(color string) {
	style := dom.GetWindow().Document().DocumentElement().Style()
	style.SetProperty("--color-secondary", color)
}