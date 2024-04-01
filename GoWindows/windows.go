// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoWindows/windows.go */

package main

import (
	"log"
	//"os"
	ui "github.com/utopiagio/utopia"

	//dialog "github.com/utopiagio/utopia-x/filedialog"
)

var mainwin *ui.GoWindowObj
var win *ui.GoWindowObj

func main() {
	// create application instance before any other objects
	app := ui.GoApplication("Windows")
	// create application window
	mainwin = ui.GoMainWindow("Windows Demo")
	//mainwin.SetBorder(ui.BorderSingleLine, 1, 10, ui.Color_Blue)
	mainwin.Layout().SetBackgroundColor(ui.Color_LightGray)
	mainwin.SetPos(100, 200)
	mainwin.SetSize(340, 280)

	menuBar := mainwin.MenuBar()
	menuBar.Show()

	mnuWindows := menuBar.AddMenu("Windows")
	mnuHelp := menuBar.AddMenu("Help")
	mnuHelp.SetOnClick(ActionHelp_Clicked)

	mnuWindows.AddAction("Add New Window", ActionWindowsNew_Clicked)
	mnuWindows.AddAction("Exit", ActionExit_Clicked)

	layoutTop := ui.GoVFlexBoxLayout(mainwin.Layout())
	layoutTop.SetMargin(0,0,0,0)
	layoutTop.SetBorder(ui.BorderSingleLine, 2, 0, ui.Color_White)
	layoutTop.SetPadding(0,0,0,0)

	lstBox := ui.GoListBox(layoutTop)
	lstBox.SetLayoutMode(ui.Vertical)
    lstBox.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
    lstBox.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)

	lblTop1 := ui.GoLabel(lstBox, text)
	lblTop1.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
	lblTop1.SetBorder(ui.BorderSingleLine, 1, 0, ui.Color_Black)

	lblTop2 := ui.GoLabel(lstBox, "A short subnote.")
	lblTop2.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)

	layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
	layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
	layoutBottom.SetMargin(0,0,0,0)
	layoutBottom.SetBorder(ui.BorderSingleLine, 2, 0, ui.Color_White)
	layoutBottom.SetPadding(0,0,0,0)

	btnClose := ui.GoButton(layoutBottom, "Close")
	btnClose.SetWidth(160)
	btnClose.SetHeight(160)
	btnClose.SetMargin(10,10,10,10)
	btnClose.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	btnClose.SetPadding(4,4,4,4)
	btnClose.SetOnClick(ActionExit_Clicked)
	btnClose.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	// show the application window
	mainwin.Show()
	// run the application
	app.Run()
}

func ActionExit_Clicked() {
	log.Println("ActionExit_Clicked().......")
	mainwin.Close()
	//os.Exit(0)
}

func ActionWindowsNew_Clicked() {
	// create application window
	win = ui.GoWindow("Windows Demo 2nd Window")
	// set the window layout style to stack widgets vertically
	//win.SetLayoutStyle(ui.VFlexBoxLayout)
	win.SetMargin(10,10,10,10)
	//win.SetMargin(15,15,15,15)
	win.SetBorder(ui.BorderSingleLine, 1, 10, ui.Color_Blue)
	//win.SetPadding(10,10,10,10)
	win.SetPadding(0,0,0,0)
	win.SetPos(440, 200)
	win.SetSize(700, 280)
	// show the second window
	win.Show()
}

func ActionHelp_Clicked() {
	
}

var text string = "Line 0\nAn even longer very very very very long text line with textflow\nLine 2\nLine 3\nLine 4\nLine 5\nLine 6\nLine 7\nLine 8\n" + 
				  "Line 9\nLine10\nLine11\nLine12\nLine13\nLine14\nLine15\nLine16\nLine17\n" + 
				  "Line18\nLine19\nLine20\nLine21\nLine22\nLine23\nLine24\nLine25\nLine26\n"


