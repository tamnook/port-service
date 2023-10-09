package app

import "port-service/internal/app/config"

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run() error {

	Config := config.New()

	Config.Read()
	return nil
}
