/* GoWidgets/widgets.go */

package main

import (
	"log"
	"os"
	ui "github.com/utopiagio/utopia"

	//icon "golang.org/x/exp/shiny/materialdesign/icons"	// eg: icon.FileFolder
	
)

var win *ui.GoWindowObj


func main() {
	// create application instance before any other objects
	app := ui.GoApplication("TestWidget")
	// create application window
	win = ui.GoMainWindow("TestWidget Demo")
	// set the window layout style to stack widgets vertically
	win.SetLayoutStyle(ui.VFlexBoxLayout)
	win.SetMargin(10,10,10,10)
	//win.SetMargin(15,15,15,15)
	win.SetBorder(ui.BorderSingleLine, 10, 10, ui.Color_Blue)
	//win.SetPadding(10,10,10,10)
	win.SetPadding(0,0,0,0)
	win.SetPos(40, 40)
	win.SetSize(600, 600)

	layoutTop := ui.GoVFlexBoxLayout(win.Layout())
	layoutTop.SetMargin(0,0,0,0)
	layoutTop.SetBorder(ui.BorderSingleLine, 10, 5, ui.Color_Red)
	layoutTop.SetPadding(10,10,10,10)

	layoutBottom := ui.GoHFlexBoxLayout(win.Layout())
	layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)
	//layoutTop.SetMargin(10,10,10,10)
	layoutBottom.SetMargin(0,0,0,0)
	layoutBottom.SetBorder(ui.BorderSingleLine, 10, 5, ui.Color_Red)
	layoutBottom.SetPadding(0,0,0,0)

	chkBox := ui.GoCheckBox(layoutTop, "CheckBox")
	chkBox.SetMargin(0,0,0,0)
	chkBox.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	chkBox.SetPadding(0,0,0,0)
	chkBox.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	/*spacer := ui.GoSpacer(layoutTop, 20)*/

	btnSwitch := ui.GoSwitch(layoutTop, "Switch")
	btnSwitch.SetMargin(0,0,0,0)
	//btnSwitch.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	btnSwitch.SetPadding(8,8,8,8)
	btnSwitch.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
	btnSwitch.SetOnChange(Switch_OnChange)

	btnClose := ui.GoButton(layoutBottom, "Close")
	btnClose.SetWidth(160)
	btnClose.SetHeight(160)
	btnClose.SetMargin(10,10,10,10)
	btnClose.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	//btnClose.SetPadding(1,1,1,1)
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

func Switch_OnChange(state bool) {
	log.Println("Switch_OnChange(", state, ").......")
	win.SetPos(200, 40)
}