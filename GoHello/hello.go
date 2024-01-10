/* hello.go */

package hello

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
	mainwin.SetSize(640, 480)
	mainwin.SetPos(100,100)

	layoutTop := ui.GoHFlexBoxLayout(mainwin.Layout())
	layoutTop.SetMargin(0,0,0,0)
	layoutTop.SetPadding(10,10,10,10)
	layoutTop.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

	lblWindowProperties = ui.GoLabel(layoutTop, "")
	lblWindowProperties.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
	lblWindowProperties.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)
	lblWindowProperties.SetMaxLines(0)
	lblWindowProperties.SetPadding(8,8,8,8)

	// Action Bar to contain button controls
	layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
	layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)	// Note ui.FixedHeight
	layoutBottom.SetMargin(0,10,0,0)
	layoutBottom.SetPadding(0,0,0,0)
	layoutBottom.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

	btnMove := ui.GoButton(layoutBottom, "Move")
	btnMove.SetMargin(4,4,4,4)
	btnMove.SetPadding(4,4,4,4)
	btnMove.SetOnClick(ActionMove_Clicked)

	btnSize := ui.GoButton(layoutBottom, "Size")
	btnSize.SetMargin(4,4,4,4)
	btnSize.SetPadding(4,4,4,4)
	btnSize.SetOnClick(ActionSize_Clicked)

	btnRefresh := ui.GoButton(layoutBottom, "Refresh")
	btnRefresh.SetMargin(4,4,4,4)
	btnRefresh.SetPadding(4,4,4,4)
	btnRefresh.SetOnClick(ActionRefresh_Clicked)

	btnClose := ui.GoButton(layoutBottom, "Close")
	btnClose.SetWidth(260)
	btnClose.SetHeight(160)
	btnClose.SetMargin(4,4,4,4)
	//btnClose.SetBorder(ui.BorderSingleLine, 1, 2, ui.Color_Red)
	btnClose.SetPadding(4,4,4,4)
	btnClose.SetOnClick(ActionExit_Clicked)
	//btnClose.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	lblWindowProperties.SetText("Click the Refresh Button........\n\n   to see the window properties.")
	// show the application window
	mainwin.SetOnConfig(UpdateWindowProperties)
	mainwin.Show()

	
	// run the application
	app.Run()
}

func ActionExit_Clicked() {
	log.Println("ActionExit_Clicked().......")
	os.Exit(0)
}

func ActionMove_Clicked() {
	x, y := mainwin.Pos()
	mainwin.SetPos(x + 10, y + 10)
}

func ActionRefresh_Clicked() {
	lblWindowProperties.SetText(GetWindowProperties())
}

func ActionSize_Clicked() {
	width, height := mainwin.ClientSize()
	mainwin.SetSize(width + 10, height + 10)
}

func GetWindowProperties() (text string) {
	text = "WINDOW PROPERTIES>\n\n"
	text += "Position:" + "\n"
	text += "ClientWidth: " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientWidth())) + "\n"	// * ui.GoDpr)) + "\n"
	text += "ClientHeight: " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientHeight())) + "\n"	// * ui.GoDpr)) + "\n"
	text += "ScreenWidth: " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Width())) + "\n"	// * ui.GoDpr)) + "\n"
	text += "ScreenHeight: " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Height())) + "\n\n"	// * ui.GoDpr)) + "\n\n"
	//text += "HorizontalSize: " + strconv.Itoa(desktop.HorizontalSize()) + "\n"
	//text += "VerticalSize: " + strconv.Itoa(desktop.VerticalSize()) + "\n"
	//text += "HorizontalRes: " + strconv.Itoa(desktop.HorizontalRes()) + "\n"
	//text += "VerticalRes: " + strconv.Itoa(desktop.VerticalRes()) + "\n\n"

	//text += "AspectX: " + strconv.Itoa(desktop.AspectX()) + "\n"
	//text += "AspectY: " + strconv.Itoa(desktop.AspectY()) + "\n"
	//text += "AspectXY: " + strconv.Itoa(desktop.AspectXY()) + "\n\n"

	X, Y := mainwin.Pos()
	Width, Height := mainwin.ClientSize()

	text += "MainWindowPos:" + " (" + strconv.Itoa(X) + ", " + strconv.Itoa(Y) + ")" + "\n"
	text += "MainWindowSize:" + " (" + strconv.Itoa(Width) + ", " + strconv.Itoa(Height) + ")" + "\n"

	return text
}

func UpdateWindowProperties() {
	lblWindowProperties.SetText(GetWindowProperties())
}
