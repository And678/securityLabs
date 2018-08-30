package main

import (
  "github.com/andlabs/ui"
  "strconv"
  "math"
  "os"
)


func startUi() {
	err := ui.Main(func() {
		inputM := ui.NewSpinbox(0, math.MaxInt32)
		inputM.SetValue(1 << 15 - 1)
		inputA := ui.NewSpinbox(0, math.MaxInt32)
		inputA.SetValue(1 << 3)
		inputC := ui.NewSpinbox(0, math.MaxInt32)
		inputC.SetValue(8)
		inputX := ui.NewSpinbox(0, math.MaxInt32)
		inputX.SetValue(64)
		inputR := ui.NewSpinbox(0, math.MaxInt32)

		boxM := ui.NewHorizontalBox()
		boxM.Append(ui.NewLabel("Модуль порівняння (m): "), true)
		boxM.Append(inputM, true)

		boxA := ui.NewHorizontalBox()
		boxA.Append(ui.NewLabel("Множник (a):"), true)
		boxA.Append(inputA, true)

		boxC := ui.NewHorizontalBox()
		boxC.Append(ui.NewLabel("Приріст (c):"), true)
		boxC.Append(inputC, true)

		boxX := ui.NewHorizontalBox()
		boxX.Append(ui.NewLabel("Початкове число (X0):"), true)
		boxX.Append(inputX, true)

		boxR := ui.NewHorizontalBox()
		boxR.Append(ui.NewLabel("Номер ітерації:"), true)
		boxR.Append(inputR, true)

		button := ui.NewButton("Отримати число")
		buttonFile := ui.NewButton("Записати період у файл")
		resultLabel := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.SetPadded(true)
		box.Append(ui.NewLabel("Введіть параметри генерації:"), false)
		box.Append(boxM, false)
		box.Append(boxA, false)
		box.Append(boxC, false)
		box.Append(boxX, false)
		box.Append(boxR, false)
		box.Append(button, false)
		box.Append(buttonFile, false)
		box.Append(resultLabel, false)
		window := ui.NewWindow("Лабораторна робота №1", 200, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		buttonFile.OnClicked(func(*ui.Button) {
			i := toFile(inputM.Value(), inputA.Value(), inputC.Value(), inputX.Value())
			resultLabel.SetText("Розмір періоду: " + strconv.Itoa(i))
		})
		button.OnClicked(func(*ui.Button) {
			resultLabel.SetText(strconv.Itoa(int(getRandomNumber(inputM.Value(), inputA.Value(), inputC.Value(), inputX.Value(), inputR.Value()))))
			inputR.SetValue(inputR.Value() + 1);
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

func toFile(m, a, c, x int) int {
	file, err := os.Create("period.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(strconv.Itoa(x))
	result := (a * x + c) % m;
	firstValue := result
	i := 0
	for ;result != firstValue || i == 0; i++ {
		file.WriteString(", " + strconv.Itoa(result))
		result = (a * result + c) % m;
	}
	return i
}

func main() {
	startUi()
}
