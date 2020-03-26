package example

import "log"

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	log.Println("wasm done4")
	select {}
}
