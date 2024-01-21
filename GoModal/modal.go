// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoModal/modal.go */

package main

import (
	"log"
	"os"
	ui "github.com/utopiagio/utopia"

	dialog "github.com/utopiagio/utopia-x/filedialog"
)

var win *ui.GoWindowObj
var modal *ui.GoWindowObj

func main() {
	// create application instance before any other objects
	app := ui.GoApplication("ModalWindow")
	// create application window
	win = ui.GoMainWindow("ModalWindow Demo")
	// set the window layout style to stack widgets vertically
	win.SetLayoutStyle(ui.VFlexBoxLayout)
	win.SetMargin(10,10,10,10)
	//win.SetMargin(15,15,15,15)
	win.SetBorder(ui.BorderSingleLine, 1, 10, ui.Color_Blue)
	//win.SetPadding(10,10,10,10)
	win.SetPadding(0,0,0,0)

	layoutTop := ui.GoHFlexBoxLayout(win.Layout())
	layoutTop.SetMargin(0,0,0,0)
	layoutTop.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	layoutTop.SetPadding(0,0,0,0)

	layoutBottom := ui.GoHFlexBoxLayout(win.Layout())
	layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
	layoutBottom.SetMargin(0,0,0,0)
	layoutBottom.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	layoutBottom.SetPadding(0,0,0,0)

	btnSwitch := ui.GoButton(layoutTop, "ShowModal")
	btnSwitch.SetMargin(0,0,0,0)
	btnSwitch.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	btnSwitch.SetPadding(6,6,6,6)
	btnSwitch.SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
	btnSwitch.SetOnClick(ActionShowModal_Clicked)

	btnClose := ui.GoButton(layoutBottom, "Close")
	btnClose.SetWidth(160)
	btnClose.SetHeight(160)
	btnClose.SetMargin(10,10,10,10)
	btnClose.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	btnClose.SetPadding(4,4,4,4)
	btnClose.SetOnClick(ActionExit_Clicked)
	btnClose.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	// show the application window
	win.Show()
	// run the application
	app.Run()
}

func ActionExit_Clicked() {
	log.Println("ActionExit_Clicked().......")
	//win.Close()
	os.Exit(0)
}

func ActionShowModal_Clicked() {
	log.Println("ActionShowModal_Clicked().......")
	modal = ui.GoModalWindow("GoFileDialog", "Open")
	modal.SetSize(600, 400)
	/*openFile :=*/ dialog.GoOpenFile(modal.Layout(), "/", "Open File Dialog", "")
	btnClose := ui.GoButton(modal.Layout(), "Close")
	btnClose.SetMargin(10,10,10,10)
	btnClose.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	btnClose.SetPadding(4,4,4,4)
	btnClose.SetOnClick(ActionModalClose_Clicked)
	btnClose.SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)

	modal.Show()
	return
}

func ActionModalClose_Clicked() {
	action := modal.ModalAction
	info := modal.ModalInfo
	log.Println("ModalAction:", action)
	log.Println("ModalInfo:", info)
	modal.Close()
}