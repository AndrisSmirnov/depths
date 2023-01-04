package app

func (a *App) Stop() error {
	return a.dataGateway.Close()
}
