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

	shuffleBingo()

	myApp := app.New()
	r := rand.New(rand.NewSource(99))

	myWindow := myApp.NewWindow("Gopher")
	setMenu(myApp, myWindow)

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
		{createBoardField(r, gopherImg, &board[0][0]), createBoardField(r, gopherImg, &board[0][1]), createBoardField(r, gopherImg, &board[0][2])},
		{createBoardField(r, gopherImg, &board[1][0]), createBoardField(r, gopherImg, &board[1][1]), createBoardField(r, gopherImg, &board[1][2])},
		{createBoardField(r, gopherImg, &board[2][0]), createBoardField(r, gopherImg, &board[2][1]), createBoardField(r, gopherImg, &board[2][2])},
	}
	table_of_buttons := makeTable(button_fields)
	table_of_fields := makeTable(game_fields)
	box := container.NewVBox(
		table_of_buttons,
		gopherImg,
		table_of_fields,
	)
	// Display our content
	myWindow.SetContent(
		box,
	)

	// Close the App when Escape key is pressed
	closeAppOnEscape(myWindow, myApp)
	myWindow.Resize(fyne.Size{Height: 500, Width: 300})
	// Show window and run app
	myWindow.ShowAndRun()
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

func createBoardField(r *rand.Rand, gopherImg *canvas.Image, game_field *BingoField) *fyne.Container {
	//counter := true
	rect := canvas.NewRectangle(color.Black)
	rect.SetMinSize(fyne.Size{Height: 50, Width: 30})
	field_container := container.New(
		layout.NewMaxLayout(),
		rect,
	)
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
	return container.NewHBox(objects...)
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
