package app

import (
	"fmt"
	"github.com/golangee/forms-example/www/dom"
	. "github.com/golangee/forms-example/www/html"
	"github.com/golangee/forms-example/www/internal/build"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"runtime/debug"
	"strings"
)

type Application struct {
	logger log.Logger
}

func NewApplication() *Application {
	return &Application{
		logger: log.NewLogger(ecs.Log("application")),
	}
}

func (a *Application) Run() {

	a.execProtect(a.doStuff)

	// keep alive
	select {}
}

func (a *Application) execProtect(f func()) {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Print(ecs.Msg(fmt.Sprint(r)), ecs.ErrStack())
			body := dom.GetWindow().Document().Body()
			body.Clear()
			formatPanic(r)(body)
		}
	}()

	f()
}

func formatPanic(p interface{}) Modifier {
	msg := fmt.Sprint(p)
	lines := strings.Split(string(debug.Stack()), "\n")
	return Div(Class("rounded overflow-hidden shadow-lg dark:bg-gray-800"),
		P(Class("text-red-500"), Text(msg)),
		ForEach(len(lines), func(i int, e dom.Element) {
			P(Class("text-gray-500"), Text(lines[i]))(e)
		}),
	)
}

func (a *Application) doStuff() {
	a.logger.Print(ecs.Msg("application is running30"), log.V("build.commit", build.Commit))

	body := dom.GetWindow().Document().Body()
	body.Clear()

	var myImg dom.Element
	content :=
		Div(Class("rounded overflow-hidden shadow-lg dark:bg-gray-800"),
			Figure(Class("md:flex bg-gray-100 rounded-xl p-8 md:p-0"),
				Img(
					Class("w-32 h-32 md:w-48 md:h-auto md:rounded-none rounded-full mx-auto"),
					Src("https://tailwindcss.com/_next/static/media/sarah-dayan.a8ff3f1095a58085a82e3bb6aab12eb2.jpg"),
					Width("384"),
					Height("512"),
					Self(&myImg),
					AddEventListener("click", false, func(e dom.Element) {
						a.logger.Print(ecs.Msg("clicked it"))
						myImg.SetClassName("rounded-xl")
					}),
				),
				Div(
					Class("pt-6 md:p-8 text-center md:text-left space-y-4"),
					Blockquote(
						P(
							Class("text-lg font-semibold"),
							Text("“Tailwind CSS is the only framework that I've seen scale on large teams. It’s easy to customize, adapts to any design,and the build size is tiny.”"),
						),
					),
					Figcaption(Class("font-medium"),
						Div(Class("text-yellow-400"),
							Text("Sarah Dayan"),
							AddEventListener("click", true, func(e dom.Element) {
								a.logger.Print(ecs.Msg("only once clicked it"))
								myImg.SetClassName("rounded-xl")
							}),
						),
						Div(Class("text-gray-500"),
							Text(" Staff Engineer, Algolia"),
						),
						Ul(
							Li(Text("hey")),
							Li(Text("ho")),
						),
					),
				),

			),
		)

	content(body)

	myImg.Release()
}
