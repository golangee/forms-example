package modal

import . "github.com/golangee/forms-example/www/forms/view"

// A DialogCard represents a material design-like card view, providing a minimal width, shadow and font.
type DialogCard struct {
	View
	title   string
	content Renderable
}

func NewDialogCard(title string, content Renderable) *DialogCard {
	return &DialogCard{title: title, content: content}
}

func (c *DialogCard) Render() Node {
	return Div(Class("rounded-md shadow-xl bg-white m-auto p-6 pb-2"),
		Style("min-width", "240px"),
		AddClickListener(nil), // intentionally block any click events
		If(c.title != "",
			P(Class("text-xl font-medium"),
				Text(c.title),
			),
			nil,
		),

		c.content,
	)
}
