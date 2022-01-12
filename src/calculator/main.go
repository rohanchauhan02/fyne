package main

import (
	// fyne "fyne.io/fyne/v2"
	// "image/color"
	// "fyne.io/fyne/v2/canvas"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	// w.Resize(fyne.NewSize(600, 600))
	output := ""
	input := widget.NewLabel(output)
	var historyArr []string
	var flag bool=false
	historyBtn := widget.NewButton("History", func() {
		var str string
		if flag==false{
			for i:=len(historyArr)-1;i>=0;i--{
				str+=historyArr[i]+"\n"
			}
			input.SetText(str)
			flag=true
		}else{
			input.SetText("")
			flag=false
		}
		
	})
	backBtn := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	clearBtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})
	openBtn := widget.NewButton("(", func() {
		output += "("
		input.SetText(output)
	})
	closeBtn := widget.NewButton(")", func() {
		output += ")"
		input.SetText(output)
	})
	divideBtn := widget.NewButton("/", func() {
		output += "/"
		input.SetText(output)
	})
	sevenBtn := widget.NewButton("7", func() {
		output += "7"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8", func() {
		output += "8"
		input.SetText(output)
	})
	nineBtn := widget.NewButton("9", func() {
		output += "9"
		input.SetText(output)
	})
	multiplyBtn := widget.NewButton("*", func() {
		output += "*"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4", func() {
		output += "4"
		input.SetText(output)
	})
	fiveBtn := widget.NewButton("5", func() {
		output += "5"
		input.SetText(output)
	})
	sixBtn := widget.NewButton("6", func() {
		output += "6"
		input.SetText(output)
	})
	minusBtn := widget.NewButton("-", func() {
		output += "-"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output += "1"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2", func() {
		output += "2"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3", func() {
		output += "3"
		input.SetText(output)
	})
	plusBtn := widget.NewButton("+", func() {
		output += "+"
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0", func() {
		output += "0"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".", func() {
		output += "."
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=", func() {
		expression,err:= govaluate.NewEvaluableExpression(output);
		if err==nil{
			result, err2 := expression.Evaluate(nil);
			if err2==nil{
				ans:=strconv.FormatFloat(result.(float64),'f',-1,64)
				str:=output+"="+ans
				historyArr=append(historyArr, str)
				output=str
			}else{
				output="error"
			}
		}else{
			output="error"
		}
		input.SetText(output)
	})

	w.SetContent(container.NewVBox(
		input,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
			container.NewGridWithColumns(4,
				clearBtn,
				openBtn,
				closeBtn,
				divideBtn,
			),
			container.NewGridWithColumns(4,
				nineBtn,
				eightBtn,
				sevenBtn,
				multiplyBtn,
			),
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				minusBtn,
			),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn,
				),
				container.NewGridWithColumns(1,
					equalBtn,
				),
			),
		),
	))
	w.ShowAndRun()
}
