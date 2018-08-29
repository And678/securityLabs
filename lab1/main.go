package main

import (
	"github.com/andlabs/ui"
    "strconv"
    "math"
)


func startUi() {
	err := ui.Main(func() {
		
		inputM := ui.NewSpinbox(0, math.MaxInt32)
		inputA := ui.NewSpinbox(0, math.MaxInt32)
		inputC := ui.NewSpinbox(0, math.MaxInt32)
		inputX := ui.NewSpinbox(0, math.MaxInt32)
		inputR := ui.NewSpinbox(0, math.MaxInt32)

		boxM := ui.NewHorizontalBox()
		boxM.Append(ui.NewLabel("m"), true)
		boxM.Append(inputM, true)
		
		boxA := ui.NewHorizontalBox()
		boxA.Append(ui.NewLabel("a"), true)
		boxA.Append(inputA, true)
		
		boxC := ui.NewHorizontalBox()
		boxC.Append(ui.NewLabel("c"), true)
		boxC.Append(inputC, true)
		
		boxX := ui.NewHorizontalBox()
		boxX.Append(ui.NewLabel("xxx"), true)
		boxX.Append(inputX, true)
		
		boxR := ui.NewHorizontalBox()
		boxR.Append(ui.NewLabel("rrrrrrr"), true)
		boxR.Append(inputR, true)

		button := ui.NewButton("Отримати число")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(boxM, false)
		box.Append(boxA, false)
		box.Append(boxC, false)
		box.Append(boxX, false)
		box.Append(boxR, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 200, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText(strconv.Itoa(int(getRandomNumber(inputM.Value(), inputA.Value(), inputC.Value(), inputX.Value(), inputR.Value()))))
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func getRandomNumber(m, a, c, x, r int) int {
	result := x
	for i := 0; i < r; i++ {
		result = (a * result + c) % m; 
	}
	return result
}


func main() {
	startUi()
}