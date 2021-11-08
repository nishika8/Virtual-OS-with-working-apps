package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
	//a := app.New()
	//w := a.NewWindow("Calculator")
	//w.Resize(fyne.NewSize(400, 280))

	w := myApp.NewWindow("Calc ")
	w.Resize(fyne.NewSize(500, 280))

	output := " "
	input := widget.NewLabel(output)
	isHistory := false
	HistoryStr := ""
	history := widget.NewLabel(HistoryStr)
	var HistoryArr []string

	History := widget.NewButton("History", func() {
		if isHistory {
			HistoryStr = ""
		} else {
			for i := len(HistoryArr) - 1; i >= 0; i-- {
				HistoryStr = HistoryStr + HistoryArr[i]
				HistoryStr += "\n"
			}
		}
		isHistory = !isHistory
		history.SetText(HistoryStr)
	})
	Clear := widget.NewButton("C", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}

	})
	AllClear := widget.NewButton("AC", func() {
		output = ""
		input.SetText(output)
	})
	OpenBracket := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	CloseBracket := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	Divide := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})
	Seven := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	Eight := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	Nine := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	Multiply := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	Four := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	Five := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	Six := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	Minus := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	One := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	Two := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	Three := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	Plus := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	Zero := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	Dot := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	Equal := widget.NewButton("=", func() {

		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + " = " + ans
				HistoryArr = append(HistoryArr, strToAppend)
				output = strconv.FormatFloat(result.(float64), 'f', -1, 64)
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}
		input.SetText(output)

	})

	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				History,
				Clear),
			container.NewGridWithColumns(4,
				AllClear,
				OpenBracket,
				CloseBracket,
				Divide),
			container.NewGridWithColumns(4,
				Seven,
				Eight,
				Nine,
				Multiply),
			container.NewGridWithColumns(4,
				Four,
				Five,
				Six,
				Minus),
			container.NewGridWithColumns(4,
				One,
				Two,
				Three,
				Plus),
			container.NewGridWithColumns(2,
				Zero,
				Dot),
			container.NewGridWithColumns(1,
				Equal),
		),
	))

	//w.ShowAndRun()

	// w.SetContent(
	// container.NewBorder(DeskBtn, nil, nil, nil, CalcContainer),
	// )
	w.Show()
}
