package main

import (
	"log"

	"fyne.io/fyne/app"
)

func main() {
	configureFyneFont()

	application := app.NewWithID("com.huajian.stockcalculator")
	ui, err := NewUI(application)
	if err != nil {
		log.Fatal(err)
	}

	ui.Run()
}
