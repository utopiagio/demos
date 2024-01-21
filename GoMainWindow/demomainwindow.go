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
	mainwin.SetMargin(10,10,10,10)
	//win.SetMargin(15,15,15,15)
	//win.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)
	//win.SetPadding(10,10,10,10)
	mainwin.SetPadding(0,0,0,0)


	menuBar := mainwin.MenuBar()
	menuBar.Show()

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

	layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
	layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
	layoutBottom.SetMargin(0,10,0,0)
	layoutBottom.SetPadding(0,0,0,0)
	layoutBottom.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

	btnClose := ui.GoButton(layoutBottom, "Close")
	btnClose.SetWidth(260)
	btnClose.SetHeight(160)
	btnClose.SetMargin(4,4,4,4)
	//btnClose.SetBorder(ui.BorderSingleLine, 1, 2, ui.Color_Red)
	btnClose.SetPadding(4,4,4,4)
	btnClose.SetOnClick(ActionExit_Clicked)
	//btnClose.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

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
    text += "    ScreenWidth:       " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Width())) + " px\n"    // * ui.GoDpr)) + "\n"
    text += "    ScreenHeight:      " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Height())) + " px\n"    // * ui.GoDpr)) + "\n\n"
    text += "    HorizontalRes:       " + strconv.Itoa(desktop.HorizontalRes()) + " dpi\n"
    text += "    VerticalRes:           " + strconv.Itoa(desktop.VerticalRes()) + " dpi\n\n"

    /*text += "Aspect Ratios :" + "\n"
    text += "   X: " + strconv.Itoa(desktop.AspectX())
    text += ",    Y: " + strconv.Itoa(desktop.AspectY())
    text += ",    XY: " + strconv.Itoa(desktop.AspectXY()) + "\n\n"*/

    text += "Screen Available :" + "\n"
    text += "    ClientWidth:        " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientWidth())) + " px\n"  // * ui.GoDpr)) + "\n"
    text += "    ClientHeight:         " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientHeight())) + " px\n\n"    // * ui.GoDpr)) + "\n"
    
    X, Y := mainwin.Pos()
    Width, Height := mainwin.ClientSize()
    text += "Window Geometry :" + "\n"
    text += "    WindowPos:     " + " (" + strconv.Itoa(X) + ", " + strconv.Itoa(Y) + ")" + " px\n"
    text += "    WindowSize:    " + " (" + strconv.Itoa(Width) + ", " + strconv.Itoa(Height) + ")" + " px\n\n"

    //text += "Window Client Geometry :" + "\n"
    //text += "    ClientSize: " + " (" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.HorizontalSize())) + ", " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.VerticalSize())) + ")" + "\n"

    return text
}

func UpdateWindowProperties() {
    lblWindowProperties.SetText(GetWindowProperties())
}
