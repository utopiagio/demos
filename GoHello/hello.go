// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoHello/hello.go */

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
    app := ui.GoApplication("GoHelloDemo")
    // create application window
    mainwin = ui.GoMainWindow("GoHello Demo - UtopiaGio Package")
    // set the window layout style to stack widgets vertically
    mainwin.SetLayoutStyle(ui.VFlexBoxLayout)
    mainwin.SetMargin(10,10,10,10)
    mainwin.SetPadding(0,0,0,0)
    //mainwin.SetMinSize(320, 480)
    mainwin.SetSize(640, 480)
    mainwin.SetPos(100,100)

    layoutWinProperties := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutWinProperties.SetMargin(0,0,0,0)
    layoutWinProperties.SetPadding(10,10,10,10)
    layoutWinProperties.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

    lblWindowProperties = ui.GoLabel(layoutWinProperties, "")
    lblWindowProperties.SetWrap(false)
    lblWindowProperties.SetSizePolicy(ui.PreferredWidth, ui.ExpandingHeight)
    lblWindowProperties.SetMinWidth(260)
    lblWindowProperties.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)
    lblWindowProperties.SetMaxLines(19)
    lblWindowProperties.SetPadding(8,8,8,8)

    ui.GoSpacer(layoutWinProperties, 10)

    layoutLblSizing := ui.GoVFlexBoxLayout(layoutWinProperties)

    lblLabel0 := ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: FixedWidth, Vert: FixedHeight}")
    lblLabel0.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
    lblLabel0.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)
    lblLabel0.SetWidth(500)
    lblLabel0.SetMaxLines(1)
   
    ui.GoSpacer(layoutLblSizing, 10)

    lblLabel1 := ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: PreferredWidth, Vert: PreferredHeight}")
    lblLabel1.SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
    lblLabel1.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)
    lblLabel1.SetMaxLines(0)
 
    ui.GoSpacer(layoutLblSizing, 10)
 
    lblLabel2 := ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: ExpandingWidth, Vert: ExpandingHeight}")
    lblLabel2.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
    lblLabel2.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)
    lblLabel2.SetMaxLines(0)
    

    // Action Bar to contain button controls
    layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)   // Note: ui.FixedHeight
    layoutBottom.SetMargin(0,10,0,0)
    layoutBottom.SetPadding(0,0,0,0)
    layoutBottom.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

    leftpadding := ui.GoSpacer(layoutBottom, 0)
    leftpadding.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

    btnMove := ui.GoButton(layoutBottom, "Move")
    btnMove.SetMargin(4,4,4,4)
    btnMove.SetPadding(4,4,4,4)
    btnMove.SetOnClick(ActionMove_Clicked)

    btnSize := ui.GoButton(layoutBottom, "Size")
    btnSize.SetMargin(4,4,4,4)
    btnSize.SetPadding(4,4,4,4)
    btnSize.SetOnClick(ActionSize_Clicked)

    padding := ui.GoSpacer(layoutBottom, 0)
    padding.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

    btnClose := ui.GoButton(layoutBottom, "Close")
    btnClose.SetWidth(260)
    btnClose.SetHeight(160)
    btnClose.SetMargin(4,4,4,4)
    btnClose.SetPadding(4,4,4,4)
    btnClose.SetOnClick(ActionExit_Clicked)
    
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

func ActionSize_Clicked() {
    width, height := mainwin.ClientSize()
    mainwin.SetSize(width + 10, height + 10)
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