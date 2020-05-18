# wtk-example
wtk kitchen sink demo. It is inspired by the [SwiftUI tutorial](https://developer.apple.com/tutorials/swiftui/creating-and-combining-views).
The only available implementation is [WTK](https://github.com/golangee/forms) with a WASM HTML core. 

## Section 1

A View must comply to the View interface. It cannot be implemented entirely by your own, because
it contains some hidden implementation specific contracts. You have to implement a View interface
by embedding an existing view or container. By convention, you should always provide a factory method
prefixed with *New* followed by your views name.

```go
package section1

import . "github.com/golangee/forms"

type ContentView struct {
	*Text
}

func NewContentView() *ContentView {
	return &ContentView{Text: NewText("hello world")}
}
```


## Section 2
You can style your View in a fluent way, using the predefined parameters.

```go
package section2

import . "github.com/golangee/forms"

type ContentView struct {
	*Text
}

func NewContentView() *ContentView {
	return &ContentView{
		Text: NewText("Turtle Rock").
			Style(
				ForegroundColor(Green),
				Font(Title),
			),
	}
}

```