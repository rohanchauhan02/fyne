// // // package main

// // // import (
// // // 	"fyne.io/fyne/v2/app"
// // // 	"fyne.io/fyne/v2/container"
// // // 	// "fyne.io/fyne/v2/data/binding"
// // // 	"fyne.io/fyne/v2/widget"
// // // )

// // // func main() {
// // // 	myApp := app.New()
// // // 	w := myApp.NewWindow("Two Way")

// // // 	// str := binding.NewString()
// // // 	// str.Set("Hi!")

// // // 	w.SetContent(container.NewVBox(
// // // 		widget.NewLabelWithData("str"),
// // // 		widget.NewEntryWithData("str"),
// // // 	))

// // // 	w.ShowAndRun()
// // // }
// // // package main

// // // import (
// // // 	"fyne.io/fyne/v2/app"
// // // 	"fyne.io/fyne/v2/container"
// // // 	"fyne.io/fyne/v2/data/binding"
// // // 	"fyne.io/fyne/v2/widget"
// // // )

// // // func main() {
// // // 	myApp := app.New()
// // // 	w := myApp.NewWindow("Conversion")

// // // 	f := binding.NewFloat()
// // // 	str := binding.FloatToString(f)
// // // 	short := binding.FloatToStringWithFormat(f, "%0.0f%%")
// // // 	f.Set(25.0)

// // // 	w.SetContent(container.NewVBox(
// // // 		widget.NewSliderWithData(0, 100.0, f),
// // // 		widget.NewLabelWithData(str),
// // // 		widget.NewLabelWithData(short),
// // // 	))

// // // 	w.ShowAndRun()
// // // }
// // // package main

// // // import (
// // // 	"time"

// // // 	"fyne.io/fyne/v2/app"
// // // 	"fyne.io/fyne/v2/data/binding"
// // // 	"fyne.io/fyne/v2/widget"
// // // )

// // // func main() {
// // // 	a := app.New()
// // // 	w := a.NewWindow("Hello")

// // // 	str := binding.NewString()
// // // 	go func() {
// // // 		dots := "....."
// // // 		for i := 5; i >= 0; i-- {
// // // 			str.Set("Count down" + dots[:i])
// // // 			time.Sleep(time.Second)
// // // 		}
// // // 		str.Set("Blast off!")
// // // 	}()

// // // 	w.SetContent(widget.NewLabelWithData(str))
// // // 	w.ShowAndRun()
// // // }
// // package main

// // import (
// // 	// "image/color"
// // 	"fyne.io/fyne/v2"

// // 	"fyne.io/fyne/v2/layout"
// // 	"fyne.io/fyne/v2/app"
// // 	"fyne.io/fyne/v2/widget"
// // 	"fyne.io/fyne/v2/container"
// // )

// // func main() {
// // 	appClient := app.New()
// //     winWindow := appClient.NewWindow("Test")

// //     txtEntry := widget.NewEntry()
// //     txtStatus := widget.NewLabel("Status")
// //     txtResults := widget.NewTextGrid()
// //     txtResults.ShowLineNumbers = true

// //     btnQuit := widget.NewButton("Quit", func() {
// //         appClient.Quit()
// //     })
// //     btnQuit.Resize(fyne.NewSize(300, 300))

// //     cntScrolling := container.NewScroll(txtResults)
// //     cntButtons := container.NewGridWithColumns(4, layout.NewSpacer(), layout.NewSpacer(), layout.NewSpacer(), btnQuit)
// //     cntContent := container.NewGridWithRows(4, txtEntry, txtStatus, cntScrolling, cntButtons)

// //     winWindow.Resize(fyne.NewSize(1200, 600))
// //     winWindow.SetContent(cntContent)

// //     go func() {
// //         for {
// //             select {
// //             case strMessage := <-chnSendMessageToClientWindow:
// //                 txtResults.SetText(strings.TrimPrefix(txtResults.Text()+"\n"+strMessage, "\n"))
// //                 cntScrolling.Refresh()
// //             }
// //         }
// //     }()

// //     winWindow.ShowAndRun()

// // }

// // // package main
// // // // import fyne
// // // import (
// // //     "io/ioutil"
// // //     "fyne.io/fyne/v2"
// // //     "fyne.io/fyne/v2/app"
// // //     "fyne.io/fyne/v2/container"
// // //     "fyne.io/fyne/v2/dialog"
// // //     "fyne.io/fyne/v2/storage"
// // //     "fyne.io/fyne/v2/widget"
// // // )
// // // func main() {
// // //     // New app
// // //     a := app.New()
// // //     //New title and window
// // //     w := a.NewWindow("Open file in FYNE")
// // //     // resize window
// // //     w.Resize(fyne.NewSize(400, 400))
// // //     // New Buttton
// // //     btn := widget.NewButton("Open .txt files", func() {
// // //         // Using dialogs to open files
// // //         // first argument func(fyne.URIReadCloser, error)
// // //         // 2nd is parent window in our case "w"
// // //         // r for reader
// // //         // _ is ignore error
// // //         file_Dialog := dialog.NewFileOpen(
// // //             func(r fyne.URIReadCloser, _ error) {
// // //                 // read files
// // //                 data, _ := ioutil.ReadAll(r)
// // //                 // reader will read file and store data
// // //                 // now result
// // //                 result := fyne.NewStaticResource("name", data)
// // //                 // lets display our data in label or entry
// // //                 entry := widget.NewMultiLineEntry()
// // //                 // string() function convert byte to string
// // //                 entry.SetText(string(result.StaticContent))
// // //                 // Lets show and setup content
// // //                 // tile of our new window
// // //                 w := fyne.CurrentApp().NewWindow(
// // //                     string(result.StaticName)) // title/name
// // //                 w.SetContent(container.NewScroll(entry))
// // //                 w.Resize(fyne.NewSize(400, 400))
// // //                 // show/display content
// // //                 w.Show()
// // //                 // we are almost done
// // //             }, w)
// // //         // fiter to open .txt files only
// // //         // array/slice of strings/extensions
// // //         file_Dialog.SetFilter(
// // //             storage.NewExtensionFileFilter([]string{".txt"}))
// // //         file_Dialog.Show()
// // //         // Show file selection dialog.
// // //     })
// // //     // lets show button in parent window
// // //     w.SetContent(container.NewVBox(
// // //         btn,
// // //     ))
// // //     w.ShowAndRun()
// // // }

// package main

// import (
//     "fyne.io/fyne/v2"
//     "fyne.io/fyne/v2/app"
//     "fyne.io/fyne/v2/widget"
// 	"fyne.io/fyne/v2/container"
// )

// func runPopUp(w fyne.Window) {
//     widget.ShowModalPopUp(
//         container.NewVBox(
//             widget.NewLabel("bar"),
//         ),
//         w.Canvas(),
//     )
// }

// func main() {
//     a := app.New()
//     w := a.NewWindow("Test")
//     button := widget.NewButton("foo", func() { runPopUp(w) })
//     w.SetContent(button)
//     w.Resize(fyne.NewSize(1024, 768))
//     w.ShowAndRun()
// }

// func runPopUp(w fyne.Window) {
// 	var text *widget.PopUp
// 	hide:=widget.NewButton("Close", func() { text.Hide() })
// 	info := widget.NewLabel("")
// 	info.SetText(`

// 				Estimate the solar capacity of your rooftop:

// 				This tool allows you to estimate the amount of electricity generated on your rooftop should you install solar panels on it.
// 				The tool is defaulted to values that would work for locations based in Singapore. The measured area should be rooftops only
// 				where the assumption is to pack the number of solar panels in a small area. Note that this tool provides only a high level
// 				estimation.

// 				Electricity generated per year = rooftop area x module efficiency x average insolation x system performance ratio x 365

// 				Parameters:

// 					The following are the parameters used:
// 					1. Rooftop Area. This is the area of your rooftop.
// 					2. Panel efficiency. This is the efficiency of the solar panels you might deploy on your rooftop. The value comes from
// 						the solar panel manufacturer and is typically between 0.15 to 0.22. The value 0.17 is chosen arbitrarily and you
// 						can change it accordingly.
// 					3. Average insolation. This is the average amount of solar energy that hits an area each day, measured in kWh/sqm/day.
// 						The value 4.284 hours is the average of 10 years between Mar 2009 to Mar 2019 as derived from
// 						https://power.larc.nasa.gov/docs/v1/
// 					4. System performance ratio. This is the average performance of the system. The value 0.85 is from the upper value given
// 						by the National Solar Repository of India.

// 			`)

//     text=widget.NewPopUp(
// 		container.NewMax(
// 			container.NewVScroll(
// 				info,
// 			),
// 			hide,
// 		),
// 		w.Canvas(),
//     )
// 	text.Show()
// 	text.Resize(fyne.NewSize(600, 768))

// }