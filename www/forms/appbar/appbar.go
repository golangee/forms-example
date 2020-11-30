package appbar

import (
	"github.com/golangee/forms-example/www/forms/property"
	. "github.com/golangee/forms-example/www/forms/view"
)

// The AppBar provides a drawer area (side menu) at the left and a toolbar area at the top. It has only limited
// capabilities for customization. If you are sure, that you really need a custom AppBar, feel free to
// copy-paste to start a new component for your specific project.
type AppBar struct {
	isOpen       property.Bool
	toolbarArea  Renderable
	icon         Renderable
	title        Renderable
	drawerHeader Renderable
	drawerMain   Renderable
	drawerFooter Renderable
	View
}

func NewAppBar() *AppBar {
	c := &AppBar{}
	return c
}

// SetToolbarArea updates the node for right side of the AppBar. Consider mobile devices and only offer a context
// menu for small screens.
func (c *AppBar) SetToolbarArea(node Renderable) *AppBar {
	c.toolbarArea = node
	c.Invalidate()
	return c
}

// SetIcon sets a Node as the first entry right of the hamburger menu.
func (c *AppBar) SetIcon(node Renderable) *AppBar {
	c.icon = node
	c.Invalidate()
	return c
}

// SetTitle sets a Node as the entry right of the Icon.
func (c *AppBar) SetTitle(node Renderable) *AppBar {
	c.title = node
	c.Invalidate()
	return c
}

// SetDrawerHeader sets a Node into the header section of the drawer. At least this should be the app icon.
func (c *AppBar) SetDrawerHeader(node Renderable) *AppBar {
	c.drawerHeader = node
	c.Invalidate()
	return c
}

// SetDrawerMain sets a Node as the drawers main content.
func (c *AppBar) SetDrawerMain(node Renderable) *AppBar {
	c.drawerMain = node
	c.Invalidate()
	return c
}

// SetDrawerFooter sets a Node into the bottom of the drawer.
func (c *AppBar) SetDrawerFooter(node Renderable) *AppBar {
	c.drawerFooter = node
	c.Invalidate()
	return c
}

func (c *AppBar) Close() *AppBar {
	c.isOpen.Set(false)
	return c
}

func (c *AppBar) Render() Node {
	return Div(
		Nav(Class("flex fixed w-full items-center justify-between px-6 h-16 bg-primary text-on-primary border-b border-gray-200 z-10"),

			// menu and logo
			Div(Class("flex items-center"),

				// burger menu button
				Button(Class("mr-2 focus:outline-none"), AriaLabel("Open Menu"),
					AddClickListener(c.isOpen.Toggle),
					Svg(
						Class("w-8 h-8"),
						Fill("none"),
						Stroke("currentColor"),
						StrokeLinecap("round"),
						StrokeLinejoin("round"),
						StrokeWidth("2"),
						ViewBox("0 0 24 24"),
						Path(D("M4 6h16M4 12h16M4 18h16")),
					),
				),

				// app logo in app bar
				c.icon,

				// app title
				c.title,
			),


			// button section in app bar
			Div(Class("flex items-center"),
				c.toolbarArea,
			),

			// semi-transparent content blocking layer
			Div(
				Class(" z-10 fixed ease-in-out inset-0 bg-black opacity-0 transition-all duration-500"),

				If(&c.isOpen,
					WithModifiers(
						Visibility("visible"),
						AddClass("opacity-50"),
					),
					WithModifiers(
						Visibility("hidden"),
						RemoveClass("opacity-50"),
					),
				),
				Div(
					Class("absolute inset-0"),
					AddClickListener(c.isOpen.Toggle),
				),


			),

			// Side menu
			Aside(
				Class("transform top-0 left-0 w-64 bg-white fixed h-full overflow-auto ease-in-out transition-all duration-500 z-30"),

				If(&c.isOpen,
					WithModifiers(
						AddClass("translate-x-0"),
						RemoveClass("-translate-x-full"),
					),
					WithModifiers(
						RemoveClass("translate-x-0"),
						AddClass("-translate-x-full"),
					),
				),

				// keep the logo in the menu
				Span(
					Class("flex w-full items-center p-4 border-b"),
					c.drawerHeader,
				),

				Div(
					c.drawerMain,
				),


				// button at the bottom in the side menu

				Div(
					c.drawerFooter,
				),

			),
		),
	)

}
