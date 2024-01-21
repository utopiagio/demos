/* GoWidgets/widgets.go */

package main

import (
	"log"
	"os"
	ui "github.com/utopiagio/utopia"

	icon "golang.org/x/exp/shiny/materialdesign/icons"	// eg: icon.FileFolder
	
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
	layoutTop.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)

	layoutBottom := ui.GoHFlexBoxLayout(win.Layout())
	layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
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
	//btnSwitch.SetMargin(0,4,0,4)
	btnSwitch.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	btnSwitch.SetPadding(0,0,0,0)
	//btnSwitch.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
	btnSwitch.SetOnChange(Switch_OnChange)

	grpRadio := ui.GoButtonGroup()
	//grpRadio.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)

	radButton := ui.GoRadioButton(layoutTop, grpRadio, "rad1", "RadioButton1")
	radButton.SetMargin(0,0,0,0)
	//btnSwitch.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	radButton.SetPadding(0,0,0,0)
	//radButton.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
	radButton.SetOnChange(RadioButton1_OnChange)
	radButton.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)

	radButton2 := ui.GoRadioButton(layoutTop, grpRadio, "rad2", "RadioButton2")
	radButton2.SetMargin(0,0,0,0)
	//btnSwitch.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	radButton2.SetPadding(0,0,0,0)
	//radButton2.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
	radButton2.SetOnChange(RadioButton2_OnChange)
	radButton2.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)

	slider := ui.GoSlider(layoutTop, 0, 100)
	slider.SetWidth(100)
	slider.SetOnChange(Slider_OnChange)
	slider.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	slider.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	lblSlider := ui.GoLabel(layoutTop, "Slider Information anything longer")
	lblSlider.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)

	txtEdit := ui.GoTextEdit(layoutTop, "Hint.......")
	txtEdit.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
	txtEdit.SetWidth(200)
	txtEdit.SetMaxHeight(100)
	txtEdit.SetPadding(4,4,4,4)
	txtEdit.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	
	facebookIcon := ui.GoIconPNG("../../icons/facebook-icon-48x48.png", 24)

	facebook := ui.GoIconPNGButton(layoutTop, facebookIcon)
	facebook.SetSize(24)
	facebook.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)

	lblBackArrow := ui.GoIconLabel(layoutTop, icon.NavigationArrowBack, "Back Arrow")
	lblBackArrow.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)

	progressBar := ui.GoProgressBar(layoutTop, 100)
	progressBar.SetWidth(200)
	progressBar.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	progressBar.SetProgress(50)
	progressBar.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	progressCircle := ui.GoProgressCircle(layoutTop, 100)
	progressCircle.SetWidth(20)
	progressCircle.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	progressCircle.SetProgress(90)
	progressCircle.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	//loader := ui.GoLoader(layoutTop)
	//loader.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)

	expander := ui.GoSpacer(layoutBottom, 20)
	expander.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

	btnClose := ui.GoButton(layoutBottom, "Close")
	btnClose.SetWidth(160)
	btnClose.SetHeight(160)
	btnClose.SetMargin(10,10,10,10)
	btnClose.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)
	btnClose.SetPadding(4,4,4,4)
	btnClose.SetOnClick(ActionExit_Clicked)
	//btnClose.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	ui.GoSpacer(layoutBottom, 20)
	
	
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

func RadioButton1_OnChange(selected bool) {
	log.Println("RadioButton1_Selected(", selected, ").......")
}

func RadioButton2_OnChange(selected bool) {
	log.Println("RadioButton2_Selected(", selected, ").......")
}

func Slider_OnChange(value int) {
	log.Println("Slider_OnChange Value=(", value, ").......")
}

func Switch_OnChange(state bool) {
	log.Println("Switch_OnChange(", state, ").......")
}