package main

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"time"
)

var cyclist = []string{
	"source/images/rowerzystka1.jpg",
	"source/images/rowerzystka2.jpg",
	"source/images/rowerzystka3.jpg",
	"source/images/rowerzystka4.jpg",
	"source/images/rowerzystka5.jpg",
	"source/images/rowerzystka7.jpg",
}

type BingoField struct {
	name      string
	activated bool
}

var board = [][]BingoField{
	{BingoField{"kod Witka", false}, BingoField{"Teamsy", false}, BingoField{"Unitree", false}},
	{BingoField{"pogoda", false}, BingoField{"AgileBoard", false}, BingoField{"TimeTracker", false}},
	{BingoField{"BitBucket", false}, BingoField{"syf w kiblu", false}, BingoField{"syf w kuchni", false}},
}

func main() {
	myApp := app.NewWithID("com.example.tutorial.preferences") //to unique store app data
	clock := widget.NewLabel("")
	updateTime(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	shuffleBingo()

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	//content := container.NewWithoutLayout(text1, text2)
	content := container.New(layout.NewGridLayout(2), text1, text2)

	r := rand.New(rand.NewSource(99))

	myWindowOne := myApp.NewWindow("Gopher")
	myWindowTwo := myApp.NewWindow("Moderator")
	myWindowFour := myApp.NewWindow("Layouty")
	myWindowFive := myApp.NewWindow("Timeout")
	timeOutSelectorWidget := timeoutWindow(myApp)
	setMenu(myApp, myWindowOne)

	// Define a welcome text centered
	text := canvas.NewText("Display a random Gopher!", color.White)
	text.Alignment = fyne.TextAlignCenter

	// Define a Gopher image
	var resource, _ = fyne.LoadResourceFromPath(cyclist[r.Intn(6)])
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 150, Height: 150}) // by default size is 0, 0

	// Display a vertical box containing text, image and button
	var button_fields = [][]*fyne.Container{
		{createRandomButton(r, gopherImg, &board[0][0]), createRandomButton(r, gopherImg, &board[0][1]), createRandomButton(r, gopherImg, &board[0][2])},
		{createRandomButton(r, gopherImg, &board[1][0]), createRandomButton(r, gopherImg, &board[1][1]), createRandomButton(r, gopherImg, &board[1][2])},
		{createRandomButton(r, gopherImg, &board[2][0]), createRandomButton(r, gopherImg, &board[2][1]), createRandomButton(r, gopherImg, &board[2][2])},
	}
	var game_fields = [][]*fyne.Container{
		{createBoardField(), createBoardField(), createBoardField()},
		{createBoardField(), createBoardField(), createBoardField()},
		{createBoardField(), createBoardField(), createBoardField()},
	}
	table_of_buttons := makeTable(button_fields)
	table_of_fields := makeTable(game_fields)
	heheszki := widget.NewButton("Boczek! Ty grubasie!", func() {
		myWindowThree := myApp.NewWindow("Kiepscy")
		myWindowThree.SetContent(gopherImg)
		myWindowThree.Show()
	})
	box := container.NewVBox(
		table_of_fields,
		clock,
		table_of_buttons,
		heheszki,
	)
	box_two := container.NewVBox(
		makeUI(),
	)
	// Display our content
	myWindowOne.SetContent(box)
	myWindowTwo.SetContent(box_two)
	myWindowFour.SetContent(content)
	myWindowFive.SetContent(timeOutSelectorWidget)
	myWindowOne.Show()
	myWindowTwo.Show()
	myWindowFour.Show()
	myWindowFive.Show()

	// Close the App when Escape key is pressed
	closeAppOnEscape(myWindowOne, myApp)
	myWindowOne.Resize(fyne.Size{Height: 500, Width: 300})
	// Show window and run app
	myApp.Run()
}

func timeoutWindow(myApp fyne.App) *widget.Select {
	var timeout time.Duration
	timeoutSelector := widget.NewSelect([]string{"10 seconds", "30 seconds", "1 minute"}, func(selected string) {
		switch selected {
		case "10 seconds":
			timeout = 10 * time.Second
		case "30 seconds":
			timeout = 30 * time.Second
		case "1 minute":
			timeout = time.Minute
		}

		myApp.Preferences().SetString("AppTimeout", selected)
	})
	timeoutSelector.SetSelected(myApp.Preferences().StringWithFallback("AppTimeout", "10 seconds"))
	go func() {
		time.Sleep(timeout)
		myApp.Quit()
	}()

	return timeoutSelector
}

func setMenu(myApp fyne.App, myWindow fyne.Window) {
	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { myApp.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Aurélie Vache"),
			), myWindow)
		}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	myWindow.SetMainMenu(mainMenu)
}

func createRandomButton(r *rand.Rand, gopherImg *canvas.Image, game_field *BingoField) *fyne.Container {
	//counter := true
	randomBtn := widget.NewButton(game_field.name, func() {
		random := r.Intn(6)
		resource, _ := fyne.LoadResourceFromPath(cyclist[random])
		gopherImg.Resource = resource
		//Redrawn the image with the new path
		gopherImg.Refresh()
	})

	randomBtn.Importance = widget.LowImportance
	btn_color := canvas.NewRectangle(
		color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	//if game_field.activated == true {
	//
	//	btn_color = canvas.NewRectangle(
	//		color.NRGBA{R: 46, G: 204, B: 113, A: 0})
	//	randomBtn.Refresh()
	//}
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		randomBtn,
	)
	return container1
}

func createBoardField() *fyne.Container {
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	rect.SetMinSize(fyne.Size{Height: 50, Width: 30})
	field_container := container.New(
		layout.NewMaxLayout(),
		rect,
	)
	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		rect.FillColor = green
		rect.Refresh()
	}()

	return field_container
}
func closeAppOnEscape(myWindow fyne.Window, myApp fyne.App) {
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})
}

func shuffleBingo() {
	rand.Seed(time.Now().UnixNano())
	for index := 0; index < 3; index++ {
		rand.Shuffle(len(board), func(i, j int) { board[i], board[j] = board[j], board[i] })
		for j := 0; j < 3; j++ {
			rand.Shuffle(len(board[index]), func(i, j int) {
				board[i][index], board[j][index] = board[j][index], board[i][index]
			})
		}
	}
}
func makeTable(rows [][]*fyne.Container) *fyne.Container {

	columns := rowsToColumns(rows)

	objects := make([]fyne.CanvasObject, len(columns))
	for k, col := range columns {
		box := container.NewVBox()
		for _, val := range col {
			box.Add(val)
		}
		objects[k] = box
	}
	return container.New(layout.NewGridLayout(3), objects...)
}

func rowsToColumns(rows [][]*fyne.Container) [][]*fyne.Container {
	columns := make([][]*fyne.Container, len(rows[0]))
	for _, row := range rows {
		for colK := range row {
			columns[colK] = append(columns[colK], row[colK])
		}
	}
	return columns
}
func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello championie!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello championie " + content + "!")
	}
	return out, in
}
