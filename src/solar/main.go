package main

import (
	// "fmt"
	"strconv"

	// "image/color"
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Solar Energy Calculator")
	w.Resize(fyne.NewSize(600, 600))
	info := widget.NewLabel("")
	flag := false
	readMore := widget.NewButton("Read More", func() {
		if !flag {
			info.SetText(`
				Estimate the solar capacity of your rooftop:

				This tool allows you to estimate the amount of electricity generated on your rooftop
				should you install solar panels on it.The tool is defaulted to values that would
				work for locations based in Singapore. The measured area should be rooftops only
				where the assumption is to pack the number of solar panels in a small area.
				Note that this tool provides only a high level estimation.
				
				Electricity generated per year =
				rooftop area x module efficiency x average insolation x system performance ratio x 365
				
				

				Parameters:

					The following are the parameters used:

					1. Rooftop Area. This is the area of your rooftop.

					2. Panel efficiency. This is the efficiency of the solar panels you might deploy on your
						rooftop. The value comes from
						the solar panel manufacturer and is typically between 0.15 to 0.22. The value 0.17 is
						chosen arbitrarily and you
						can change it accordingly.

					3. Average insolation. This is the average amount of solar energy that hits an area each
						day, measured in kWh/sqm/day.
						The value 4.284 hours is the average of 10 years between Mar 2009 to Mar 2019 as
						derived from
						https://power.larc.nasa.gov/docs/v1/

					4. System performance ratio. This is the average performance of the system. The value
						0.85 is from the upper value given
						by the National Solar Repository of India.
			`)
			// info.Resize(fyne.NewSize(700, 700))
			info.Refresh()
			flag = true
		} else {
			info.SetText("")
			flag = false
		}
	})
	suggest:=widget.NewLabel("")
	var city = []string{"Bengluru", "Delhi", "Gurugram", "Mumbai", "Patna"}
	var stateMap = map[string]map[string]string{
		"spr":      map[string]string{"systemPerformanceRatio": "0.83"},
		"Delhi":    map[string]string{},
		"Mumbai":   map[string]string{},
		"Bengluru": map[string]string{},
		"Gurugram": map[string]string{},
		"Patna":    map[string]string{},
	}
	stateMap["Delhi"]["averageIsolation"] = "9.86"
	stateMap["Delhi"]["perUnitElectricCharge"] = "8"
	stateMap["Mumbai"]["averageIsolation"] = "9.00"
	stateMap["Mumbai"]["perUnitElectricCharge"] = "6"
	stateMap["Bengluru"]["averageIsolation"] = "9.86"
	stateMap["Bengluru"]["perUnitElectricCharge"] = "8"
	stateMap["Gurugram"]["averageIsolation"] = "9.00"
	stateMap["Gurugram"]["perUnitElectricCharge"] = "6"
	stateMap["Patna"]["averageIsolation"] = "9.00"
	stateMap["Patna"]["perUnitElectricCharge"] = "6"

	aI := binding.NewString()
	spr := binding.NewString()
	stateOps := widget.NewSelect(city, func(value string) {
		// fmt.Println(value," : ",stateMap[value])
		aI.Set(stateMap[value]["averageIsolation"])
		spr.Set(stateMap["spr"]["systemPerformanceRatio"])
	})
	var valSolar float64
	
	var valElectric float64
	
	output := ""
	indianState := widget.NewLabel("Choose your State ")
	result := widget.NewLabel(output)
	entryArea := widget.NewEntry()
	entryPanelEfficiency := widget.NewEntry()
	entryAverageIsolation := widget.NewEntryWithData(aI)
	entrySystemPerformanceRatio := widget.NewEntryWithData(spr)
	entryBillAmount := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Roof Top Area (sqm)", HintText: "Enter Area in Square meter eg. 100", Widget: entryArea},
			{Text: "Panel Efficiency", HintText: "Enter the solar panel efficency provided by your suplier eg. 0.17", Widget: entryPanelEfficiency},
			{Text: "Average Insolation (kWh/sqm/day)", HintText: "Enter Insulation of your city eg. 9.86", Widget: entryAverageIsolation},
			{Text: "System Performance ratio", HintText: "Enter the system Performance ratio eg. 0.85", Widget: entrySystemPerformanceRatio},
			{Text: "Monthly Electric Bill (₹)", HintText: "Enter monthly electric bill eg. 8000", Widget: entryBillAmount}},

		OnSubmit: func() {
			output = entryArea.Text + "*" + entryPanelEfficiency.Text + "*" + entryAverageIsolation.Text + "*" + entrySystemPerformanceRatio.Text + "*" + "365/1000"
			
			suggestion:=entryBillAmount.Text+"/"+stateMap["spr"]["systemPerformanceRatio"]+"365/1000"

			expression, err := govaluate.NewEvaluableExpression(output)
			if err == nil {
				result, err2 := expression.Evaluate(nil)
				if err2 == nil {
					ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
					val,_:=strconv.ParseFloat(ans,64)
					valSolar=val
					// fmt.Println(valSolar)
					output = ans
				} else {
					output = "Insufficient Information"

				}
			} else {
				output = "Insufficien Information"
			}

			expSuggestion, err := govaluate.NewEvaluableExpression(suggestion)
			if err == nil {
				result, err2 := expSuggestion.Evaluate(nil)
				if err2 == nil {
					ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
					val,_:=strconv.ParseFloat(ans,64)
					valElectric=val
					// fmt.Println(valElectric)
				} else {
					output = "Insufficient Information"

				}
			} else {
				output = "Insufficien Information"
			}
			out:=valSolar-valElectric
			help:=entryPanelEfficiency.Text + "*" + entryAverageIsolation.Text + "*" + entrySystemPerformanceRatio.Text

			expHelper1, err := govaluate.NewEvaluableExpression(help)
			if err == nil {
				result, err2 := expHelper1.Evaluate(nil)
				if err2 == nil {
					ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
					help=ans
				}
			}
			addArea:=entryArea.Text+"+"+"("+strconv.FormatFloat(out, 'f', -1, 64)+"/"+help+")"
			expHelper2, err := govaluate.NewEvaluableExpression(addArea)
			
			if err == nil {
				result, err2 := expHelper2.Evaluate(nil)
				if err2 == nil {
					ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
					addArea=ans
				} 
			} 
			if out>=0{
				str:="The grid is profitable and you produce "+strconv.FormatFloat(out*1000, 'f', -1, 64)+"unit energy extra per year \n\t\t\t\t\t\t\t\t\t\t\t\t\t\tOR\n"+ strconv.FormatFloat((out*1000*28)/365, 'f', -1, 64)+"unit extra energy per month."
				suggest.SetText(str)
			}else{
				str:="You have to install solar panel in atleast "+addArea+"meter square of area to make grid profitable."
				suggest.SetText(str)
			}
			result.SetText(output)
			
		},
		// OnCancel: func(){

		// },
	}
	w.SetContent(container.NewVScroll(
		// note,
		container.NewVBox(
			container.NewGridWithColumns(2,
				indianState,
				stateOps,
			),
			form,
			container.NewGridWithColumns(2,
				widget.NewLabel("Electricity generated per year (MWh/Year): "),
				result,
			),
			suggest,
			container.NewGridWithColumns(2,
				widget.NewButton("Refresh", func() {
					output = ""
					result.SetText(output)
					entryArea.Text = ""
					entryPanelEfficiency.Text = ""
					entrySystemPerformanceRatio.Text = ""
					form.Refresh()
				}),
				readMore,
			),
			container.NewMax(
				info,
			),
		),
	))
	w.ShowAndRun()
}
