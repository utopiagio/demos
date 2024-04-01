// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoScrollView/goscrollview.go */

package main

import (
	//"log"
	
	ui "github.com/utopiagio/utopia"
)

var win *ui.GoWindowObj


func main() {
	// create application instance before any other objects
	app := ui.GoApplication("GoScrollView")
	// create application window
	win = ui.GoMainWindow("GoScrollView Demo")
	// set the window layout style to stack widgets vertically
	win.SetLayoutStyle(ui.VFlexBoxLayout)
	win.SetMargin(10,10,10,10)
	win.SetBorder(ui.BorderSingleLine, 1, 10, ui.Color_Blue)
	win.SetPadding(10,10,10,10)
	win.SetPos(40, 40)
	win.SetSize(600, 600)

	mainview := ui.GoVBoxLayout(win.Layout())
	
		txtPad := ui.GoTextEdit(mainview, "Enter text here.")
		txtPad.SetBorder(ui.BorderSingleLine, 1, 10, ui.Color_Blue)
		txtPad.SetPadding(10,10,10,10)
		txtPad.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
		txtPad.SetFont("Go", ui.Regular, ui.Bold)

	// show the application window
	win.Show()
	// run the application
	app.Run()
}