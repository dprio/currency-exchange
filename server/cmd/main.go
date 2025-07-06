package main

import "github.com/dprio/currency-exchange/server/cmd/app"

func main() {

	app := app.New()
	app.Start()

}
