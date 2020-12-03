package s01_00_setup

import (
	. "github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/forms-example/www/internal/tutorial/view"
)

func Step() Renderable {
	return view.NewStep("Step 1", P(
		Text("To set things up, first ensure that you have a recent Go SDK installation. This tutorial has been tested with "),
		A(Href("https://golang.org/dl/"), Text("Go 1.15")), Text("."),
		Em(Text("Go")), Text("development is really fun with a good IDE. We recommend "),
		A(Href("https://www.jetbrains.com/go/"), Text("Goland")), Text("to have a hassle-free and productive time."),
	))
}
