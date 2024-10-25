package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Converter")

	// Input Label and Entry
	inputLabel := widget.NewLabel("Enter Temperature:")
	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Temperature")

	// Unit Selection
	unitLabel := widget.NewLabel("Select Unit:")
	unitGroup := widget.NewRadioGroup([]string{"Celsius", "Fahrenheit", "Kelvin"}, nil)

	// Output Labels
	celsiusOutput := widget.NewLabel("")
	fahrenheitOutput := widget.NewLabel("")
	kelvinOutput := widget.NewLabel("")

	// Conversion Button
	convertButton := widget.NewButton("Convert", func() {
		input, err := strconv.ParseFloat(inputEntry.Text, 64)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			return
		}

		inputUnit := unitGroup.Selected
		var celsius, fahrenheit, kelvin float64

		// Perform conversion based on input unit
		switch inputUnit {
		case "Celsius":
			celsius = input
			fahrenheit = celsius*9/5 + 32
			kelvin = celsius + 273.15
		case "Fahrenheit":
			fahrenheit = input
			celsius = (fahrenheit - 32) * 5 / 9
			kelvin = celsius + 273.15
		case "Kelvin":
			kelvin = input
			celsius = kelvin - 273.15
			fahrenheit = celsius*9/5 + 32
		}

		// Update output labels
		celsiusOutput.SetText(fmt.Sprintf("Celsius: %.2f", celsius))
		fahrenheitOutput.SetText(fmt.Sprintf("Fahrenheit: %.2f", fahrenheit))
		kelvinOutput.SetText(fmt.Sprintf("Kelvin: %.2f", kelvin))
	})

	// Layout
	content := container.NewVBox(
		inputLabel, inputEntry,
		unitLabel, unitGroup,
		widget.NewLabel("Converted Values:"),
		celsiusOutput, fahrenheitOutput, kelvinOutput,
		convertButton,
	)

	// Set content and run the application
	w.SetContent(content)
	w.ShowAndRun()
}
