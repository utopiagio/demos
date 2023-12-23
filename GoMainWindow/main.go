/* main.go */

package main

import (
	"log"
	"os"
	ui "github.com/utopiagio/utopia"
)

var mainwin *ui.GoWindowObj
var lblWindowProperties *ui.GoLabelObj

func main() {
	// create application instance before any other objects
	app := ui.GoApplication("GoMainWindow")
	// create application window
	mainwin = ui.GoMainWindow("GoMainWindow Demo")
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
	lblWindowProperties.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
	lblWindowProperties.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_Blue)
	lblWindowProperties.SetMaxLines(100)
	lblWindowProperties.SetPadding(8,8,8,8)

	layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
	layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)
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

	lblWindowProperties.SetText(GetWindowProperties())
	// show the application window
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
	text += "Position:"

	/*screenMetrics = "Screen Height ...... " + strconv.Itoa(desktop.Height()) + " pixels\r\n"
	screenMetrics += "Screen Width ....... " + strconv.Itoa(desktop.Width()) + " pixels\r\n"

	screenMetrics += "Screen ClientHeight ...... " + strconv.Itoa(desktop.ClientHeight()) + " pixels\r\n"
	screenMetrics += "Screen ClientWidth ....... " + strconv.Itoa(desktop.ClientWidth()) + " pixels\r\n"

	screenMetrics += "Screen Vertical Size ...... " + strconv.Itoa(desktop.VerticalSize()) + " mm\r\n"
	screenMetrics += "Screen Horizontal Size .... " + strconv.Itoa(desktop.HorizontalSize()) + " mm\r\n"

	screenMetrics += "Screen Vertical Res ...... " + strconv.Itoa(desktop.VerticalRes()) + " pixels/inch\r\n"
	screenMetrics += "Screen Horizontal Res .... " + strconv.Itoa(desktop.HorizontalRes()) + " pixels/inch\r\n"*/
	return text
}
