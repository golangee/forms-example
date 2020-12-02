package app

import (
	"github.com/golangee/forms-example/www/forms/ico"
	"github.com/golangee/forms-example/www/forms/material/icon"
	"github.com/golangee/forms-example/www/forms/router"
	. "github.com/golangee/forms-example/www/forms/view"
	"github.com/golangee/forms-example/www/internal/index"
	"strconv"
)

func tutorialOverview(q router.Query) Renderable {
	return Div(Class("pt-20"), BackgroundColor("#191919"),
		Div(Class("container mx-auto max-w-4xl text-center text-white p-4"),
			P(Class("text-5xl"), Text(index.Tutorial.Title)),
			P(Class("text-xl pt-6"), Text(index.Tutorial.Body)),
			P(Class("text-xl pt-6 text-gray-400"), ico.NewIcon(icon.Timer), Span(Text("0hr 00min")), Span(Text(" ")), Span(Class("text-sm"), Text("Estimated Time"))),
			Img(Class("shadow-lg rounded-xl max-w-full h-auto align-middle border-none my-20"), Src(index.Tutorial.Teaser[0].File)),


		),
		Div(Class("bg-black text-white pt-1 md:pt-20"),
			Div(Class("container mx-auto flex max-w-4xl md:p-4"),
				// menu
				Div(Class("flex-none hidden md:block max-w-sm pr-12"),
					ForEach(len(index.Tutorial.Fragments), func(i int) Renderable {
						return Div(
							P(Class("text-lg text-gray-400"), Text(index.Tutorial.Fragments[i].Title)),
						)
					}),
				),

				// content
				Div(Class("flex-1 md:rounded-t-xl p-4 md:p-12 mb-20"), BackgroundColor("#191919"),
					ForEach(len(index.Tutorial.Fragments), func(i int) Renderable {
						chapter := index.Tutorial.Fragments[i]
						return Div(
							Div(Class("grid md:grid-cols-3 grid-cols-1 gap-12 mb-12"),
								Img(Class("md:col-span-1 md:object-cover md:h-48"), Src(chapter.Teaser[0].File)),
								Div(Class("md:col-span-2"),
									P(Class("text-lg text-gray-400"), Text("Chapter "+strconv.Itoa(i+1))),
									P(Class("text-xl"), Text(chapter.Title)),
									P(Class("md:col-span-2 text-gray-400"), Text(chapter.Body)),

								),
							),

							ForEach(len(chapter.Fragments), func(i int) Renderable {
								section := chapter.Fragments[i]
								return Div(Class(),
									P(Class("text-gray-400 flex"),
										Span(ico.NewIcon(icon.GroupWork)),
										Span(Class("text-blue-400 pl-6 text-lg"), Text(section.Title)),
										Span(Class("float-right flex-auto text-right"), ico.NewIcon(icon.Timer), Text("00min")),
									),
								)
							}),
						)
					}),
				),

			),
		),


	)

}
