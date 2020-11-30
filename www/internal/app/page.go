package app

import (
	"github.com/golangee/forms-example/www/forms/appbar"
	"github.com/golangee/forms-example/www/forms/button"
	"github.com/golangee/forms-example/www/forms/dom"
	"github.com/golangee/forms-example/www/forms/layout"
	"github.com/golangee/forms-example/www/forms/material/icon"
	"github.com/golangee/forms-example/www/forms/property"
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/forms-example/www/internal/index"
	"strings"
)

func (a *Application) page(q router.Query, content Renderable) Renderable {
	routes := a.router.Routes()

	bar := appbar.NewAppBar()
	bar.SetToolbarArea(Div(Class("hidden md:block md:flex md:justify-between md:bg-transparent"),

		//
		Button(
			Class("flex items-center p-3 font-medium mr-2 text-center bg-gray-300 rounded  hover:bg-gray-400 focus:outline-none focus:bg-gray-400"),
			Title("Wishlist"),

			Svg(
				Class("w-6 h-6 mr-2"),
				Fill("none"),
				Stroke("currentColor"),
				StrokeLinecap("round"),
				StrokeLinejoin("round"),
				StrokeWidth("2"),
				ViewBox("0 0 24 24"),
				Path(D("M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z")),
			),

			Span(Text("Wishlist")),
		),
	))

	bar.SetIcon(Img(Class("h-12 w-auto pr-3"), Src("https://www.worldiety.de/_nuxt/img/wdy_logo_white.8bf54.svg"), Alt("Logo")))
	bar.SetTitle(Span(Text("golangee/forms: " + q.Path())))
	bar.SetDrawerHeader(Img(
		Class("h-auto w-32 mx-auto"),
		Alt("Logo"),
		Src("asdf"),
	))

	bar.SetDrawerMain( // a side menu button
		/*Span(
			Class("flex items-center p-4 hover:bg-indigo-500 hover:text-white"),
			Span(
				Class("mr-2"),
				Svg(
					Class("w-6 h-6 mr-2"),
					Fill("none"),
					Stroke("currentColor"),
					StrokeLinecap("round"),
					StrokeLinejoin("round"),
					StrokeWidth("2"),
					ViewBox("0 0 24 24"),
					Path(D("M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6")),
				),
			),
			Span(Text("Home")),
		)*/
		layout.NewVStack(
			ForEach(len(routes), func(i int) Renderable {
				return button.NewIconTextButton(icon.OpenInBrowser, routes[i].Path, func() {
					bar.Close()
					router.Navigate("#" + routes[i].Path)
				})
			}),

		),
	)

	bar.SetDrawerFooter(Div(
		Class("fixed bottom-0 w-full"),
		Button(
			Class("flex items-center p-4 text-white bg-blue-500 hover:bg-blue-600 w-full"),
			Text("hey ho"),
		),
	))
	return bar
}

func (a *Application) page3(q router.Query, content Renderable) Renderable {
	var isOpen property.Bool

	// app bar
	return Div(
		Nav(Class("flex fixed w-full items-center justify-between px-6 h-16 bg-white text-gray-700 border-b border-gray-200 z-10 bg-primary text-primary"),


			// menu and logo
			Div(Class("flex items-center"),

				// burger menu button
				Button(Class("mr-2"), AriaLabel("Open Menu"),
					AddClickListener(isOpen.Toggle),
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
				Img(Class("h-14 w-auto"), Src("https://content-prod.worldiety.de/v2/media/images/wdy_bewerbung.537ee587.svg"), Alt("Logo")),
			),


			// button section in app bar
			Div(Class("flex items-center"),
				Div(Class("hidden md:block md:flex md:justify-between md:bg-transparent"),

					//
					Button(
						Class("flex items-center p-3 font-medium mr-2 text-center bg-gray-300 rounded  hover:bg-gray-400 focus:outline-none focus:bg-gray-400"),
						Title("Wishlist"),

						Svg(
							Class("w-6 h-6 mr-2"),
							Fill("none"),
							Stroke("currentColor"),
							StrokeLinecap("round"),
							StrokeLinejoin("round"),
							StrokeWidth("2"),
							ViewBox("0 0 24 24"),
							Path(D("M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z")),
						),

						Span(Text("Wishlist")),
					),
				),
			),

			// semi-transparent content blocking layer
			Div(
				Class(" z-10 fixed ease-in-out inset-0 bg-black opacity-0 transition-all duration-500"),

				If(&isOpen,
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
					//Class("absolute inset-0 bg-black opacity-50"),
					Class("absolute inset-0"),
					AddClickListener(isOpen.Toggle),
				),


			),

			// Side menu
			Aside(
				Class("transform top-0 left-0 w-64 bg-white fixed h-full overflow-auto ease-in-out transition-all duration-500 z-30"),

				If(&isOpen,
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
					Img(
						Class("h-auto w-32 mx-auto"),
						Alt("Logo"),
						Src("asf"),
					),
				),

				// a side menu button
				Span(
					Class("flex items-center p-4 hover:bg-indigo-500 hover:text-white"),
					Span(
						Class("mr-2"),
						Svg(
							Class("w-6 h-6 mr-2"),
							Fill("none"),
							Stroke("currentColor"),
							StrokeLinecap("round"),
							StrokeLinejoin("round"),
							StrokeWidth("2"),
							ViewBox("0 0 24 24"),
							Path(D("M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6")),
						),
					),
					Span(Text("Home")),
				),

				// button at the bottom in the side menu

				Div(
					Class("fixed bottom-0 w-full"),
					Button(
						Class("flex items-center p-4 text-white bg-blue-500 hover:bg-blue-600 w-full"),
						Text("hey ho"),
					),
				),

			),


		),

		Div(
			Class("pt-16 px-6"),
			Text("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."),
		),
	)

}

func (a *Application) page2(q router.Query, content Renderable) Renderable {
	routes := a.router.Routes()

	// global margin
	return Div(Class("m-10"),

		// global frame
		Div(Class("flex flex-col border border-gray-200 rounded-lg overflow-hidden bg-gray-100"),

			// top bar
			Div(Class("bg-white p-2 flex flex-col sm:flex-row justify-between items-center border-b border-gray-200 pr-4"),
				Span(Class("sm:block bg-transparent text-gray-600 py-2 sm:py-3 px-4 ml-2 sm:ml-4 text-xs sm:text-sm"),
					Text(q.Path()),
				),
			),
			Div(Class("md:flex flex-col md:flex-row w-full"),

				// side menu
				Div(Class("flex flex-col w-full md:w-64 text-gray-700 bg-white dark-mode:text-gray-200 dark-mode:bg-gray-800 flex-shrink-0"),
					Div(Class("flex-shrink-0 px-8 py-4 flex flex-row items-center justify-between"),
						A(Class("text-lg font-semibold tracking-widest text-gray-900 uppercase rounded-lg dark-mode:text-white focus:outline-none focus:shadow-outline"),
							Href("#/"),
							Text("forms"),
						),
					),
					Nav(Class("flex-grow md:block px-4 pb-4 md:pb-0 md:overflow-y-auto hidden"),
						Div(ForEach(len(routes), func(i int) Renderable {
							path := routes[i].Path
							idx := strings.LastIndexByte(path, '/') + 1
							name := path[idx:]
							if name == "" {
								name = "Home"
							}

							selected := q.Path() == path
							myClass := "block px-4 py-2 mt-2 text-sm font-semibold text-gray-900 bg-transparent rounded-lg dark-mode:bg-transparent dark-mode:hover:bg-gray-600 dark-mode:focus:bg-gray-600 dark-mode:focus:text-white dark-mode:hover:text-white dark-mode:text-gray-200 hover:text-gray-900 focus:text-gray-900 hover:bg-gray-200 focus:bg-gray-200 focus:outline-none focus:shadow-outline"
							if selected {
								myClass = "block px-4 py-2 mt-2 text-sm font-semibold text-gray-900 bg-gray-200 rounded-lg dark-mode:bg-gray-700 dark-mode:hover:bg-gray-600 dark-mode:focus:bg-gray-600 dark-mode:focus:text-white dark-mode:hover:text-white dark-mode:text-gray-200 hover:text-gray-900 focus:text-gray-900 hover:bg-gray-200 focus:bg-gray-200 focus:outline-none focus:shadow-outline"
							}

							return P(A(Class(myClass),
								Href("#"+path), Text(name)),
							)
						})),
					),
				),

				// main content area
				Div(
					// the documentation section
					Div(Class("bg-white p-1 border-b border-gray-200"),
						With(func() Renderable {
							tutorial := index.Tutorials.Find(q.Path())
							return P(Text(tutorial.Doc))
						}),
					),

					// our component
					content,

					// source code
					Div(Class("w-2/4 bg-white p-1 border-t border-gray-200 overflow-x-auto"),
						Pre(
							Code(Class("language-go"),
								Text(index.Tutorials.Find(q.Path()).Code),
								InsideDom(func(e dom.Element) {
									dom.GetGlobal().Get("hljs").Call("highlightBlock", e)
								}),
							),
						),
					),
				),


			),

		),
	)
}
