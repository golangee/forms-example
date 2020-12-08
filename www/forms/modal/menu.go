package modal

import . "github.com/golangee/forms-example/www/forms/view"

// Menu provides the basic frame and design for a material design popup menu.
type Menu struct {
	View
	content Renderable
}

func NewMenu(content Renderable) *Menu {
	return &Menu{content: content}
}

func (c *Menu) Render() Node {
	return Div(Class("origin-top-right absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 m-4 md:m-0"),
		Style("max-width", "90vw"),
		Div(Class("py-1"), Role("menu"), AriaOrientation("vertical"),
			c.content,
		),
	)
}

// MenuItem provides the basic component for a simple entry in a Menu.
type MenuItem struct {
	View
	content Renderable
}

func NewMenuItem(content Renderable) *MenuItem {
	return &MenuItem{content: content}
}

func (c *MenuItem) Render() Node {
	return A(Class("block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 hover:text-gray-900"), Role("menuitem"),
		c.content,
	)
}

// PopupMenu takes a Renderable as a trigger and pops a menu up.
type PopupMenu struct {
	View
	parent Renderable
	menu   *Menu
}

func NewPopupMenu(parent, menuContent Renderable) *PopupMenu {
	return &PopupMenu{parent: parent, menu: NewMenu(menuContent)}
}

func (c *PopupMenu) Render() Node {
	return Div(Class("md:relative inline-block text-left"),
		Div(
			Span(c.parent, AddClickListener(func() {

			})),
			c.menu,

		),

	)
}
