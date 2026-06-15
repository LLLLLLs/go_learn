package main

import (
	"log"

	stockcalculator "stockCalculator"

	"fyne.io/fyne/app"
)

func main() {
	stockcalculator.ConfigureFyneFont()

	application := app.NewWithID("com.huajian.stockcalculator")
	ui, err := stockcalculator.NewUI(application)
	if err != nil {
		log.Fatal(err)
	}

	ui.Run()
}
