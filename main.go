package main

import (
	"fmt"
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
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
var player_one_fields [][]BingoField

var player_one_view_fields [][]*fyne.Container
var red_canvas = canvas.NewRectangle(
	color.NRGBA{R: 255, G: 0, B: 0, A: 255})
var green_canvas = canvas.NewRectangle(
	color.NRGBA{R: 0, G: 255, B: 0, A: 0})
var table_of_player_fields *fyne.Container

func main() {
	myApp := app.NewWithID("com.example.tutorial.preferences") //to unique store app data
	clock := widget.NewLabel("")
	updateTime(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
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
	setMenu(myApp, myWindowOne)

	// Define a welcome text centered
	text := canvas.NewText("Display a random Gopher!", color.White)
	text.Alignment = fyne.TextAlignCenter

	// Define a Gopher image
	var resource, _ = fyne.LoadResourceFromPath(cyclist[r.Intn(6)])
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 150, Height: 150}) // by default size is 0, 0

	// Display a vertical box containing text, image and button
	var moderator_fields = [][]*fyne.Container{
		{createModeratorRandomButton(r, gopherImg, &board[0][0]), createModeratorRandomButton(r, gopherImg, &board[0][1]), createModeratorRandomButton(r, gopherImg, &board[0][2])},
		{createModeratorRandomButton(r, gopherImg, &board[1][0]), createModeratorRandomButton(r, gopherImg, &board[1][1]), createModeratorRandomButton(r, gopherImg, &board[1][2])},
		{createModeratorRandomButton(r, gopherImg, &board[2][0]), createModeratorRandomButton(r, gopherImg, &board[2][1]), createModeratorRandomButton(r, gopherImg, &board[2][2])},
	}

	player_one_fields = board
	shuffleBingo(player_one_fields)
	player_one_view_fields = [][]*fyne.Container{
		{createPlayerField(&player_one_fields[0][0], red_canvas), createPlayerField(&player_one_fields[0][1], red_canvas), createPlayerField(&player_one_fields[0][2], red_canvas)},
		{createPlayerField(&player_one_fields[1][0], red_canvas), createPlayerField(&player_one_fields[1][1], red_canvas), createPlayerField(&player_one_fields[1][2], red_canvas)},
		{createPlayerField(&player_one_fields[2][0], red_canvas), createPlayerField(&player_one_fields[2][1], red_canvas), createPlayerField(&player_one_fields[2][2], red_canvas)},
	}
	var game_fields = [][]*fyne.Container{
		{createBoardField(), createBoardField(), createBoardField()},
		{createBoardField(), createBoardField(), createBoardField()},
		{createBoardField(), createBoardField(), createBoardField()},
	}
	table_of_buttons := makeTable(moderator_fields)
	table_of_fields := makeTable(game_fields)
	table_of_player_fields := makeTable(player_one_view_fields)
	heheszki := widget.NewButton("Boczek! Ty grubasie!", func() {
		myWindowThree := myApp.NewWindow("Kiepscy")
		myWindowThree.SetContent(gopherImg)
		myWindowThree.Show()
	})
	box := container.NewVBox(
		table_of_fields,
		table_of_player_fields,
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
	myWindowOne.Show()
	myWindowTwo.Show()
	myWindowFour.Show()

	// Close the App when Escape key is pressed
	closeAppOnEscape(myWindowOne, myApp)
	myWindowOne.Resize(fyne.Size{Height: 500, Width: 300})
	// Show window and run app
	myApp.Run()

	//go changeFieldsColor()
}

func changeFieldsColor() {
	for _, rows := range player_one_fields {
		for _, field := range rows {
			if field.activated == true {
				for index_one, view_rows := range player_one_view_fields {
					for index_two, view_field := range view_rows {
						//view_field.Remove(view_field.Objects[0])
						////field.Objects[0].FillColor = color.NRGBA{R: 0, G: 255, B: 0, A: 0}
						//view_field.Add(green_canvas)
						//view_field.Refresh()
						new_container := container.New(
							// layout of container
							layout.NewMaxLayout(),
							// first use btn color
							green_canvas,
							// 2nd btn widget
							view_field.Objects[1],
						)
						player_one_view_fields[index_one][index_two] = new_container
						player_one_view_fields[index_one][index_two].Refresh()
						table_of_player_fields.Refresh()
					}
				}
			}
			//field.Remove(field.Objects[0])
			//field.Objects[0].FillColor = color.NRGBA{R: 0, G: 255, B: 0, A: 0}
		}
	}
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

func createModeratorRandomButton(r *rand.Rand, gopherImg *canvas.Image, game_field *BingoField) *fyne.Container {
	randomBtn := widget.NewButton(game_field.name, func() {
		random := r.Intn(6)
		resource, _ := fyne.LoadResourceFromPath(cyclist[random])
		gopherImg.Resource = resource
		//Redrawn the image with the new path
		gopherImg.Refresh()
		for index_one, rows := range player_one_fields {
			for index_two, field := range rows {
				field.activated = true
				player_one_fields[index_one][index_two] = field
			}
		}
		changeFieldsColor()
		fmt.Println(player_one_fields)

	})

	randomBtn.Importance = widget.LowImportance

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
		red_canvas,
		// 2nd btn widget
		randomBtn,
	)
	return container1
}

func createPlayerField(game_field *BingoField, color_canvas *canvas.Rectangle) *fyne.Container {
	boundString := binding.NewString()
	err := boundString.Set(game_field.name)
	if err != nil {
		log.Fatal(err)
	}
	label := widget.NewLabelWithData(boundString)

	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		color_canvas,
		// 2nd btn widget
		label,
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

func shuffleBingo(board [][]BingoField) {
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
