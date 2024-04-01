// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoListBox/listview.go */

package main

import (
    "log"
    "os"
    "strconv"
    
    ui "github.com/utopiagio/utopia"

    "github.com/utopiagio/utopia/desktop"
    "github.com/utopiagio/utopia/metrics"
    "golang.org/x/exp/shiny/materialdesign/icons"   // eg: icons.FileFolder
)

var mainwin *ui.GoWindowObj
var lblWindowProperties *ui.GoLabelObj
var lstView0 *ui.GoListViewObj
var lstView1 *ui.GoListViewObj
var lstView2 *ui.GoListViewObj

func main() {
    // create application instance before any other objects
    app := ui.GoApplication("GoListViewDemo")
    // create application window
    mainwin = ui.GoMainWindow("GoListView Demo - UtopiaGio Package")
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

    layoutLstSizing := ui.GoVFlexBoxLayout(layoutWinProperties)

    lstViewLayout0 := ui.GoHFlexBoxLayout(layoutLstSizing)

    lstView0 = ui.GoListView(lstViewLayout0)
    lstView0.SetLayoutMode(ui.Horizontal)
    lstView0.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
    lstView0.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
    lstView0.SetWidth(300)
    lstView0.SetHeight(100)
   
    lstView0.AddListItem(icons.FileFolder, "Item1")
    lstView0.AddListItem(icons.FileFolder, "Item2")
    lstView0.AddListItem(icons.FileFolder, "Item3")

    ui.GoSpacer(layoutLstSizing, 10)

    lstViewLayout1 := ui.GoHFlexBoxLayout(layoutLstSizing)

    lstView1 = ui.GoListView(lstViewLayout1)
    lstView1.SetLayoutMode(ui.Horizontal)
    lstView1.SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
    lstView1.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
    //lstView1.SetWidth(300)
    //lstView1.SetHeight(100)

    lstView1.AddListItem(icons.FileFolder, "Item1")
    lstView1.AddListItem(icons.FileFolder, "Item2")
    lstView1.AddListItem(icons.FileFolder, "Item3")

    ui.GoSpacer(layoutLstSizing, 10)
 
    lstViewLayout2 := ui.GoHFlexBoxLayout(layoutLstSizing)

    lstView2 = ui.GoListView(lstViewLayout2)
    lstView2.SetLayoutMode(ui.Vertical)
    lstView2.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
    lstView2.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)

    lstView2.AddListItem(icons.FileFolder, "Item1")
    lstView2.AddListItem(icons.FileFolder, "Item2")
    lstView2.AddListItem(icons.FileFolder, "Item3")
    
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