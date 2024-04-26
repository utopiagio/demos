// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoMainWindow/demomainwindow.go */

package main

import (
	"log"
	"os"
	"strconv"

	ui "github.com/utopiagio/utopia"
	"github.com/utopiagio/utopia/desktop"
    "github.com/utopiagio/utopia/metrics"
    pointer_gio "github.com/utopiagio/gio/io/pointer"
)

var mainwin *ui.GoWindowObj
var lblWindowProperties *ui.GoLabelObj

func main() {
	// create application instance before any other objects
	app := ui.GoApplication("GoMainWindowDemo")
	// create application window
	mainwin = ui.GoMainWindow("GoMainWindow Demo - UtopiaGio Package")
	// set the window layout style to stack widgets vertically
	mainwin.SetLayoutStyle(ui.VFlexBoxLayout)
	mainwin.SetPadding(10,10,10,10)
	mainwin.SetPos(0,0)
	mainwin.SetSize(800, 600)

	mainwin.SetOnPointerPress(UpdatePointerPress)
	mainwin.SetOnPointerRelease(UpdatePointerRelease)
	mainwin.SetOnPointerMove(UpdatePointerPosition)

	menuBar := mainwin.MenuBar()
	//menuBar.Show()

	mnuFile := menuBar.AddMenu("File")
	mnuEdit := menuBar.AddMenu("Edit")
	mnuHelp := menuBar.AddMenu("Help")
	mnuHelp.SetOnClick(ActionHelp_Clicked)

	mnuFile.AddAction("New", ActionFileNew_Clicked)
	mnuFile.AddAction("Open", ActionFileOpen_Clicked)
	mnuFile.AddAction("Save", ActionFileSave_Clicked)
	mnuFile.AddAction("Close", ActionFileClose_Clicked)
	mnuFile.AddAction("Exit", ActionExit_Clicked)

	mnuEdit.AddItem("Cut")
	mnuEdit.AddItem("Copy")
	mnuEdit.AddItem("Paste")
	mnuEdit.AddItem("Undo")
	mnuEdit.AddItem("Redo")
	
	mnuHelp.AddItem("Documentation")
	mnuHelp.AddItem("About UtopiaGio")

	layoutTop := ui.GoHFlexBoxLayout(mainwin.Layout())
	layoutTop.SetMargin(0,0,0,0)
	layoutTop.SetPadding(10,10,10,10)
	layoutTop.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

	lblWindowProperties = ui.GoLabel(layoutTop, "")
	lblWindowProperties.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
	lblWindowProperties.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_Blue)
	lblWindowProperties.SetMaxLines(100)
	lblWindowProperties.SetPadding(8,8,8,8)

	/*statusBar :=*/ mainwin.StatusBar()
	//statusbar.Show()

	// show the application window
	mainwin.SetOnConfig(UpdateWindowProperties)
	mainwin.Show()

	// run the application
	app.Run()
}

func ActionExit_Clicked() {
	log.Println("ActionExit_Clicked().......")
	//win.Close()
	os.Exit(0)
	
}

func ActionFileNew_Clicked() {
	log.Println("ActionFileNew_Clicked().......")
}

func ActionFileOpen_Clicked() {
	log.Println("ActionFileOpen_Clicked().......")
}

func ActionFileSave_Clicked() {
	log.Println("ActionFileSave_Clicked().......")
}

func ActionFileClose_Clicked() {
	log.Println("ActionFileClose_Clicked().......")
}

func ActionHelp_Clicked() {
	log.Println("ActionHelp_Clicked().......")
}

func GetWindowProperties() (text string) {
	text = "WINDOW PROPERTIES>\n\n"
    text += "Screen Geometry :" + "\n"
    text += "    ScreenWidth:       " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Width())) + " screen pixels\n"    // * ui.GoDpr)) + "\n"
    text += "    ScreenHeight:      " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Height())) + " screen pixels\n"    // * ui.GoDpr)) + "\n\n"
    text += "    HorizontalRes:       " + strconv.Itoa(desktop.HorizontalRes()) + " dpi\n"
    text += "    VerticalRes:           " + strconv.Itoa(desktop.VerticalRes()) + " dpi\n"
    text += "    TaskBarHeight:      " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.TaskBarHeight())) + " screen pixels\n\n"

    /*text += "Aspect Ratios :" + "\n"
    text += "   X: " + strconv.Itoa(desktop.AspectX())
    text += ",    Y: " + strconv.Itoa(desktop.AspectY())
    text += ",    XY: " + strconv.Itoa(desktop.AspectXY()) + "\n\n"*/

    text += "Screen Available :" + "\n"
    text += "    ClientWidth:        " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientWidth())) + " px\n"  // * ui.GoDpr)) + "\n"
    text += "    ClientHeight:         " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientHeight())) + " px\n\n"    // * ui.GoDpr)) + "\n"
    

    wX, wY := mainwin.Pos()
    wWidth, wHeight := mainwin.Size()
    text += "Window Geometry :" + "\n"
    text += "    WindowPos:     " + " (" + strconv.Itoa(wX) + ", " + strconv.Itoa(wY) + ")" + " dp\n"
	text += "    WindowSize:    " + " (" + strconv.Itoa(wWidth) + ", " + strconv.Itoa(wHeight) + ")" + " dp\n\n"
    
    cX, cY := mainwin.ClientPos()
    cWidth, cHeight := mainwin.ClientSize()
    text += "Window Client Geometry :" + "\n"
    text += "    ClientPos:     " + " (" + strconv.Itoa(cX) + ", " + strconv.Itoa(cY) + ")" + " dp\n"
    text += "    ClientSize:    " + " (" + strconv.Itoa(cWidth) + ", " + strconv.Itoa(cHeight) + ")" + " dp\n\n"

    text += "Window Geometry Screen Px:" + "\n"
	text += "    WindowPos:     " + " (" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wX)) + ", " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wY)) + ")" + " px\n"
    text += "    WindowSize:    " + " (" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wWidth)) + ", " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wHeight)) + ")" + " px\n\n"

    return text
}

func UpdateWindowProperties() {
    lblWindowProperties.SetText(GetWindowProperties())
}

func UpdatePointerPosition(e pointer_gio.Event) {
    log.Println("PointerPos")
    /*log.Println("GoCanvasObj::PointerPressed")
	log.Println("Type:", e.Type)
	log.Println("Source:", e.Source)
	log.Println("PointerID:", e.PointerID)
	log.Println("Priority:", e.Priority)
	log.Println("Time:", e.Time)
	log.Println("Buttons:", e.Buttons)
	log.Println("Position:", e.Position)
	log.Println("Scroll:", e.Scroll)
	log.Println("Modifiers:", e.Modifiers)*/
}

func UpdatePointerPress(e pointer_gio.Event) {
    log.Println("PointerPress")
    /*log.Println("GoCanvasObj::PointerPressed")
	log.Println("Type:", e.Type)
	log.Println("Source:", e.Source)
	log.Println("PointerID:", e.PointerID)
	log.Println("Priority:", e.Priority)
	log.Println("Time:", e.Time)
	log.Println("Buttons:", e.Buttons)
	log.Println("Position:", e.Position)
	log.Println("Scroll:", e.Scroll)
	log.Println("Modifiers:", e.Modifiers)*/
}

func UpdatePointerRelease(e pointer_gio.Event) {
    log.Println("PointerRelease")
    /*log.Println("GoCanvasObj::PointerPressed")
	log.Println("Type:", e.Type)
	log.Println("Source:", e.Source)
	log.Println("PointerID:", e.PointerID)
	log.Println("Priority:", e.Priority)
	log.Println("Time:", e.Time)
	log.Println("Buttons:", e.Buttons)
	log.Println("Position:", e.Position)
	log.Println("Scroll:", e.Scroll)
	log.Println("Modifiers:", e.Modifiers)*/
}