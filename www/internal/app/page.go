package app

import (
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
	"strings"
)

func (a *Application) page(q router.Query, content Renderable) Renderable {
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
				content,
			),

		),
	)
}
