package menu

import (
	"github.com/golangee/forms-example/www/forms/modal"
	. "github.com/golangee/forms-example/www/forms/view"
)

// Menu provides the basic frame and design for a material design popup menu.
type Menu struct {
	View
	content Renderable
}

func NewMenu(content Renderable) *Menu {
	return &Menu{content: content}
}

func (c *Menu) Render() Node {
	return Div(Class("w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5"),
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

// ShowPopup tries to popup the content intelligent around the given anchor, considering window size and other
// alignment rules. The Popup is closed automatically when clicked outside. The anchor must denote a valid ID.
func ShowPopup(anchorID string, menuContent Renderable) {
	modal.ShowOverlay(modal.NewOverlay().Put(anchorID,menuContent))
}
