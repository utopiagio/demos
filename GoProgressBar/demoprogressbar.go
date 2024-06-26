// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoProgressBar/demoprogressbar.go */

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
    app := ui.GoApplication("GoProgressBarDemo")
    // create application window
    mainwin = ui.GoMainWindow("GoProgressBar Demo - UtopiaGio Package")
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
    lblWindowProperties.SetPadding(8,8,8,8)

    ui.GoSpacer(layoutWinProperties, 10)

    layoutLblSizing := ui.GoVFlexBoxLayout(layoutWinProperties)

    prgBar0 := ui.GoProgressBar(layoutLblSizing, 100)
    prgBar0.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
    prgBar0.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
    prgBar0.SetWidth(300)
    prgBar0.SetHeight(50)
    prgBar0.SetProgress(40)
   
    ui.GoSpacer(layoutLblSizing, 10)

    prgBar1 := ui.GoProgressBar(layoutLblSizing, 100)
    prgBar1.SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
    prgBar1.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
    prgBar1.SetProgress(50)
 
    ui.GoSpacer(layoutLblSizing, 10)

    prgBar2 := ui.GoProgressBar(layoutLblSizing, 100)
    prgBar2.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
    prgBar2.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
    prgBar2.SetProgress(60)

    ui.GoSpacer(mainwin.Layout(), 10)

    // Action Bar to contain button controls
    layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)   // Note: ui.FixedHeight
    //layoutBottom.SetMargin(0,10,0,0)
    //layoutBottom.SetPadding(0,0,0,0)
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
    text += "    ScreenWidth:       " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Width())) + " px, " + strconv.Itoa(desktop.Width()) + " dp\n"    // * ui.GoDpr)) + "\n"
    text += "    ScreenHeight:      " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Height())) + " px, " + strconv.Itoa(desktop.Height()) + " dp\n"    // * ui.GoDpr)) + "\n\n"
    text += "    HorizontalRes:       " + strconv.Itoa(desktop.HorizontalRes()) + " dpi\n"
    text += "    VerticalRes:           " + strconv.Itoa(desktop.VerticalRes()) + " dpi\n\n"

    /*text += "Aspect Ratios :" + "\n"
    text += "   X: " + strconv.Itoa(desktop.AspectX())
    text += ",    Y: " + strconv.Itoa(desktop.AspectY())
    text += ",    XY: " + strconv.Itoa(desktop.AspectXY()) + "\n\n"*/

    text += "Screen Available :" + "\n"
    text += "    ClientWidth:        " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientWidth())) + " px, " + strconv.Itoa(desktop.ClientWidth()) + " dp\n"  // * ui.GoDpr)) + "\n"
    text += "    ClientHeight:         " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientHeight())) + " px, "  + strconv.Itoa(desktop.ClientHeight()) + " dp\n\n"    // * ui.GoDpr)) + "\n"
    
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

    text += "Window Geometry Screen Pixels:" + "\n"
    text += "    WindowPos:     " + " (" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wX)) + ", " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wY)) + ")" + " px\n"
    text += "    WindowSize:    " + " (" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wWidth)) + ", " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wHeight)) + ")" + " px\n"

    return text
}

func UpdateWindowProperties() {
    lblWindowProperties.SetText(GetWindowProperties())
}