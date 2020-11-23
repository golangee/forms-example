package main

import (
	"flag"
	"github.com/golangee/forms-example/server/internal/app"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	host := flag.String("host", "localhost", "the host to bind on.")
	port := flag.Int("port", 8081, "the port to bind to.")
	wwwDir := flag.String("www", "", "the directory which contains the wasm module.")

	flag.Parse()

	app, err := app.NewApplication(*host, *port, *wwwDir)
	if err != nil {
		return err
	}

	return app.Run()
}
