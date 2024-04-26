// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/gohello/hello.go */

package main

import (
    "strconv"
    
    ui "github.com/utopiagio/utopia"
    "github.com/utopiagio/utopia-x/uireference"
    "github.com/utopiagio/utopia/desktop"
    "github.com/utopiagio/utopia/metrics"
)

var mainwin *ui.GoWindowObj
var viewer *ui.GoWindowObj
var popup *ui.GoPopupWindowObj
var lblWindowProperties *ui.GoLabelObj
var richText *ui.GoRichTextObj
var navigator *ui.GoListViewObj
var hdrSection *ui.GoLabelObj
var layoutWindowProperties *ui.GoLayoutObj

func main() {
    // create application instance before any other objects
    app := ui.GoApplication("GoHelloDemo")
    // create application window
    mainwin = ui.GoMainWindow("GoHello Demo - UtopiaGio Package")
    // set the window layout style to stack widgets vertically
    mainwin.SetLayoutStyle(ui.VFlexBoxLayout)
    mainwin.SetMargin(10,10,10,10)
    mainwin.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)
    mainwin.SetPadding(10,10,10,10)
    mainwin.SetSize(900, 640)
    mainwin.SetPos(100,100)

    // setup the main window MenuBar
    menuBar := mainwin.MenuBar()
    menuBar.Show()
    mnuCode := menuBar.AddMenu("Documentation")
    mnuCode.AddAction("Overview", LoadOverview)
    //mnuCode.AddAction("API Reference", ShowAPIReference_Clicked)
    //mnuCode.AddAction("Main Code", ShowMainCode_Clicked)
    
    // add the header layout to align widgets horizontally
    layoutHeader := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutHeader.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
    layoutHeader.SetPadding(0,0,0,5)

        lblVBox := ui.GoLabel(layoutHeader, "GoVBoxLayout")
        lblVBox.SetSizePolicy(ui.FixedWidth, ui.PreferredHeight)
        lblVBox.SetWidth(270)
        lblVBox.SetBackgroundColor(ui.Color_LightBlue)
        lblVBox.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_LightGray)

        ui.GoSpacer(layoutHeader, 10)

        lblVFlexBox := ui.GoLabel(layoutHeader, "GoVFlexBoxLayout")
        lblVFlexBox.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
        lblVFlexBox.SetBackgroundColor(ui.Color_LightBlue)
        lblVFlexBox.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_LightGray)

    // add the Content layout to align the content horizontally
    layoutContent := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutContent.SetMargin(0,0,0,0)
    layoutContent.SetPadding(0,0,0,0)
    
        // add the VBox layout to contain the Label text providing scrollbars
        layoutWindowProperties = ui.GoVBoxLayout(layoutContent)
        layoutWindowProperties.SetSizePolicy(ui.FixedWidth, ui.ExpandingHeight)
        layoutWindowProperties.SetWidth(270)
        layoutWindowProperties.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)

            lblWindowProperties = ui.GoLabel(layoutWindowProperties, "")
            lblWindowProperties.SetWrap(false)
            lblWindowProperties.SetPadding(8,8,8,8)

        // add fixed spacer between content
        ui.GoSpacer(layoutContent, 10)

        // add the VFlexBox layout to contain the Label layout options
        layoutLblSizing := ui.GoVFlexBoxLayout(layoutContent)
        layoutLblSizing.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)
        layoutLblSizing.SetPadding(10,10,10,10)

            ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: FixedWidth, Vert: FixedHeight}")
            lblLabel0 := ui.GoLabel(layoutLblSizing, "Hello from UtopiaGio")
            lblLabel0.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
            lblLabel0.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
            lblLabel0.SetHeight(50)
            lblLabel0.SetWidth(500)
            lblLabel0.SetMaxLines(1)
            lblLabel0.SetFontSize(24)
            lblLabel0.SetFontBold(true)
            lblLabel0.SetTextColor(ui.Color_Blue)
           
            ui.GoSpacer(layoutLblSizing, 10)

            ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: PreferredWidth, Vert: PreferredHeight}")
            lblLabel1 := ui.GoLabel(layoutLblSizing, "Hello from UtopiaGio")
            lblLabel1.SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
            lblLabel1.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
            lblLabel1.SetMaxLines(0)
            lblLabel1.SetFontSize(36)
            lblLabel1.SetFontBold(true)
            lblLabel1.SetTextColor(ui.Color_Blue)
         
            ui.GoSpacer(layoutLblSizing, 10)
         
            ui.GoLabel(layoutLblSizing, "GoLabel...GoSizePolicy{Horiz: ExpandingWidth, Vert: ExpandingHeight}")
            lblLabel2 := ui.GoLabel(layoutLblSizing, "Hello from UtopiaGio")
            lblLabel2.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
            lblLabel2.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)
            lblLabel2.SetMaxLines(0)
            lblLabel2.SetFontSize(48)
            lblLabel2.SetFontBold(true)
            lblLabel2.SetTextColor(ui.Color_Blue)
    
    // add the Action Bar layout to contain button controls
    layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)   // Note: ui.FixedHeight
    layoutBottom.SetMargin(0,10,0,0)
    layoutBottom.SetPadding(0,0,0,0)
    layoutBottom.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

        // add expanding spacer
        //leftpadding := ui.GoSpacer(layoutBottom, 0)
        //leftpadding.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

        lblHint := ui.GoLabel(layoutBottom, "Try Resizing the Window...")
        lblHint.SetMargin(10, 6, 0, 0)
        lblHint.SetMaxLines(1)
        lblHint.SetFontSize(24)
        lblHint.SetFontBold(true)
        /*btnMoveCode := ui.GoButton(layoutBottom, "?")
        btnMoveCode.SetMargin(0,4,0,4)
        btnMoveCode.SetPadding(4,4,2,4)
        btnMoveCode.SetFaceColor(ui.Color_LightGrey)
        btnMoveCode.SetOnClick(ShowMoveCode_Clicked)

        btnMove := ui.GoButton(layoutBottom, "Move")
        btnMove.SetMargin(0,4,4,4)
        btnMove.SetPadding(2,4,4,4)
        btnMove.SetOnClick(ActionMove_Clicked)

        btnSize := ui.GoButton(layoutBottom, "Size")
        btnSize.SetMargin(4,4,0,4)
        btnSize.SetPadding(4,4,2,4)
        btnSize.SetOnClick(ActionSize_Clicked)

        btnSizeCode := ui.GoButton(layoutBottom, "?")
        btnSizeCode.SetMargin(0,4,0,4)
        btnSizeCode.SetPadding(2,4,4,4)
        btnSizeCode.SetFaceColor(ui.Color_LightGrey)
        btnSizeCode.SetOnClick(ShowSizeCode_Clicked)*/

        // add expanding spacer
        padding := ui.GoSpacer(layoutBottom, 0)
        padding.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

        btnClose := ui.GoButton(layoutBottom, "Close")
        btnClose.SetWidth(260)
        btnClose.SetHeight(160)
        btnClose.SetMargin(4,4,4,4)
        btnClose.SetPadding(4,4,4,4)
        btnClose.SetOnClick(ActionExit_Clicked)
    
    lblWindowProperties.SetText("Click the Refresh Button........\n\n   to see the window properties.")
    mainwin.SetOnConfig(UpdateWindowProperties)
    // show the application window
    mainwin.Show()

    
    // run the application
    app.Run()
}

func ActionExit_Clicked() {
    mainwin.Close()
}

/*func ActionMove_Clicked() {
    x, y := mainwin.Pos()
    mainwin.SetPos(x + 10, y + 10)
}

func ActionSize_Clicked() {
    width, height := mainwin.Size()
    mainwin.SetSize(width + 10, height + 10)
}*/

func GetWindowProperties() (text string) {
    text = "WINDOW PROPERTIES>\n\n"
    text += "Screen Geometry :" + "\n"
    text += "    ScreenWidth:       " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Width())) + " px\n"    // * ui.GoDpr)) + "\n"
    text += "    ScreenHeight:      " + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Height())) + " px\n"    // * ui.GoDpr)) + "\n\n"
    text += "    HorizontalRes:       " + strconv.Itoa(desktop.HorizontalRes()) + " dpi\n"
    text += "    VerticalRes:           " + strconv.Itoa(desktop.VerticalRes()) + " dpi\n\n"

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

func LaunchViewer(section string, content string) {
    // create viewer window
    viewer = ui.GoWindow("UtopiaGio: Reference Documentation")
    //viewer.SetPos(440, 200)
    viewer.SetSize(1000, 600)
    viewer.SetLayoutStyle(ui.VFlexBoxLayout)
    viewer.Layout().SetPadding(10,10,10,10)
    /*page := */uireference.Page(viewer.Layout(), "UtopiaGio", section, content)
    // show the viewer window
    viewer.Show()
}

func LoadOverview() {
    LaunchViewer("GoWindowObj", "")
}

/*func ShowMoveCode_Clicked() {
    popup = mainwin.PopupWindow()
    popup.Layout().SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
    popup.Layout().SetPadding(4,4,4,4)
    popup.Layout().SetBackgroundColor(ui.Color_White)
    popup.Layout().SetMargin(50,340,0,0)
    popup.Layout().Clear()
    ui.GoLabel(popup.Layout(), codeActionMove_Clicked)
    popup.Show()
}

func ShowSizeCode_Clicked() {
    popup = mainwin.PopupWindow()
    popup.Layout().SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
    popup.Layout().SetMargin(285,340,0,0)
    popup.Layout().SetPadding(4,4,4,4)
    popup.Layout().SetBackgroundColor(ui.Color_White)
    popup.Layout().Clear()
    ui.GoLabel(popup.Layout(), codeActionSize_Clicked)
    popup.Show()
}*/

func UpdateWindowProperties() {
    lblWindowProperties.SetText(GetWindowProperties())
    //height := metrics.DpToPx(ui.GoDpr, lblWindowProperties.AbsHeight)
    mainwin.Refresh()
}

//------------------------------------------------------------------------------------------------------------------------------------------------------------------

var codeActionMove_Clicked string = "func ActionMove_Clicked() {\n" + 
    "    x, y := mainwin.Pos()\n" + 
    "    mainwin.SetPos(x + 10, y + 10)\n" + 
    "}"

var codeActionSize_Clicked string = "func ActionSize_Clicked() {\n" + 
    "    width, height := mainwin.Size()\n" + 
    "    mainwin.SetSize(width + 10, height + 10)\n" + 
    "}"

/*var codeGetWindowProperties string = "func GetWindowProperties() (text string) {\n" + 
    "    text = \"WINDOW PROPERTIES>\\n\\n\"\n" + 
    "\n" + 
    "    text += \"Screen Geometry :\\n\"\n" + 
    "    text += \"    ScreenWidth:       \" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Width())) + \" px\\n\"\n" + 
    "    text += \"    ScreenHeight:      \" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.Height())) + \" px\\n\"\n" + 
    "    text += \"    HorizontalRes:       \" + strconv.Itoa(desktop.HorizontalRes()) + \" dpi\\n\"\n" + 
    "    text += \"    VerticalRes:           \" + strconv.Itoa(desktop.VerticalRes()) + \" dpi\\n\\n\"\n" + 
    "\n" + 
    "    text += \"Screen Available :\\n\"\n" + 
    "    text += \"    ClientWidth:        \" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientWidth())) + \" px\\n\"\n" + 
    "    text += \"    ClientHeight:         \" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, desktop.ClientHeight())) + \" px\\n\\n\"\n" + 
    "\n" +     
    "    wX, wY := mainwin.Pos()\n" + 
    "    wWidth, wHeight := mainwin.Size()\n" + 
    "    text += \"Window Geometry :\\n\"\n" + 
    "    text += \"    WindowPos:     \" + \" (\" + strconv.Itoa(wX) + \", \" + strconv.Itoa(wY) + \")\" + \" dp\\n\"\n" + 
    "    text += \"    WindowSize:    \" + \" (\" + strconv.Itoa(wWidth) + \", \" + strconv.Itoa(wHeight) + \")\" + \" dp\\n\\n\"\n" + 
    "\n" +     
    "    cX, cY := mainwin.ClientPos()\n" + 
    "    cWidth, cHeight := mainwin.ClientSize()\n" + 
    "    text += \"Window Client Geometry :\\n\"\n" + 
    "    text += \"    ClientPos:     \" + \" (\" + strconv.Itoa(cX) + \", \" + strconv.Itoa(cY) + \")\" + \" dp\n\"\n" + 
    "    text += \"    ClientSize:    \" + \" (\" + strconv.Itoa(cWidth) + \", \" + strconv.Itoa(cHeight) + \")\" + \" dp\\n\\n\"\n" + 
    "\n" + 
    "    text += \"Window Geometry Screen Pixels:\\n\"\n" + 
    "    text += \"    WindowPos:     \" + \" (\" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wX)) + \", \" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wY)) + \")\" + \" px\\n\"\n" + 
    "    text += \"    WindowSize:    \" + \" (\" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wWidth)) + \", \" + strconv.Itoa(metrics.DpToPx(ui.GoDpr, wHeight)) + \")\" + \" px\\n\\n\"\n" + 
    "\n" + 
    "    return text\n" + 
    "}"*/

/*var codeMain string = "package main\n\n" + 
    "import (\n" + 
    "    \"strconv\"\n\n" + 
    "    ui \"github.com/utopiagio/utopia\"\n\n" + 
    "    \"github.com/utopiagio/utopia/desktop\"\n" +
    "    \"github.com/utopiagio/utopia/metrics\"\n" +
    ")\n\n" + 

    "var mainwin *ui.GoWindowObj\n" + 
    "var codewin *ui.GoWindowObj\n" + 
    "var popup *ui.GoPopupWindowObj\n" + 
    "var lblWindowProperties *ui.GoLabelObj\n\n" + 

    "func main() {\n" + 
    "    // create application instance before any other objects\n" + 
        "    app := ui.GoApplication(\"GoHelloDemo\")\n" + 
        "    // create application window\n" + 
        "    mainwin = ui.GoMainWindow(\"GoHello Demo - UtopiaGio Package\")\n" + 
        "    // set the window layout style to stack widgets vertically\n" + 
        "    mainwin.SetLayoutStyle(ui.VFlexBoxLayout)\n" + 
        "    mainwin.SetMargin(10,10,10,10)\n" + 
        "    mainwin.SetPadding(0,0,0,0)\n" + 
        "    mainwin.SetSize(640, 565)\n" + 
        "    mainwin.SetPos(100,100)\n\n" + 

        "    // setup the main window MenuBar\n" + 
        "    menuBar := mainwin.MenuBar()\n" + 
        "    menuBar.Show()\n" + 
        "    mnuCode := menuBar.AddMenu(\"Code\")\n" + 
        "    mnuCode.AddAction(\"Main code\", ShowMainCode_Clicked)\n" + 
        "    mnuCode.AddAction(\"GetWindowProperties code\", ShowGetWindowProperties_Clicked)\n\n" + 
        
        "    layoutWinProperties := ui.GoHFlexBoxLayout(mainwin.Layout())\n" + 
        "    layoutWinProperties.SetMargin(0,0,0,0)\n" + 
        "    layoutWinProperties.SetPadding(10,10,10,10)\n" + 
        "    layoutWinProperties.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)\n" + 

        "    lblWindowProperties = ui.GoLabel(layoutWinProperties, \"\")\n" + 
        "    lblWindowProperties.SetWrap(false)\n" + 
        "    lblWindowProperties.SetSizePolicy(ui.PreferredWidth, ui.ExpandingHeight)\n" + 
        "    lblWindowProperties.SetMinWidth(260)\n" + 
        "    lblWindowProperties.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightGray)\n" + 
        "    lblWindowProperties.SetMaxLines(23)\n" + 
        "    lblWindowProperties.SetPadding(8,8,8,8)\n\n" + 

        "    ui.GoSpacer(layoutWinProperties, 10)\n\n" + 

        "    layoutLblSizing := ui.GoVFlexBoxLayout(layoutWinProperties)\n\n" + 

        "    lblLabel0 := ui.GoLabel(layoutLblSizing, \"GoLabel...GoSizePolicy{Horiz: FixedWidth, Vert: FixedHeight}\")\n" + 
        "    lblLabel0.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)\n" + 
        "    lblLabel0.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)\n" + 
        "    lblLabel0.SetWidth(500)\n" + 
        "    lblLabel0.SetMaxLines(1)\n\n" + 
   
        "    ui.GoSpacer(layoutLblSizing, 10)\n\n" + 

        "    lblLabel1 := ui.GoLabel(layoutLblSizing, \"GoLabel...GoSizePolicy{Horiz: PreferredWidth, Vert: PreferredHeight}\")\n" + 
        "    lblLabel1.SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)\n" + 
        "    lblLabel1.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)\n" + 
        "    lblLabel1.SetMaxLines(0)\n\n" + 
 
        "    ui.GoSpacer(layoutLblSizing, 10)\n\n" + 
 
        "    lblLabel2 := ui.GoLabel(layoutLblSizing, \"GoLabel...GoSizePolicy{Horiz: ExpandingWidth, Vert: ExpandingHeight}\")\n" + 
        "    lblLabel2.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)\n" + 
        "    lblLabel2.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_LightBlue)\n" + 
        "    lblLabel2.SetMaxLines(0)\n\n" + 
    

        "    // Action Bar to contain button controls\n" + 
        "    layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())\n" + 
        "    layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)\n" + 
        "    layoutBottom.SetMargin(0,10,0,0)\n" + 
        "    layoutBottom.SetPadding(0,0,0,0)\n" + 
        "    layoutBottom.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)\n\n" + 

        "    leftpadding := ui.GoSpacer(layoutBottom, 0)\n" + 
        "    leftpadding.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)\n\n" + 

        "    btnMoveCode := ui.GoButton(layoutBottom, \"?\")\n" + 
        "    btnMoveCode.SetMargin(0,4,0,4)\n" + 
        "    btnMoveCode.SetPadding(4,4,2,4)\n" + 
        "    btnMoveCode.SetFaceColor(ui.Color_LightGrey)\n" + 
        "    btnMoveCode.SetOnClick(ShowMoveCode_Clicked)\n\n" + 

        "    btnMove := ui.GoButton(layoutBottom, \"Move\")\n" + 
        "    btnMove.SetMargin(0,4,4,4)\n" + 
        "    btnMove.SetPadding(2,4,4,4)\n" + 
        "    btnMove.SetOnClick(ActionMove_Clicked)\n\n" + 

        "    btnSize := ui.GoButton(layoutBottom, \"Size\")\n" + 
        "    btnSize.SetMargin(4,4,0,4)\n" + 
        "    btnSize.SetPadding(4,4,2,4)\n" + 
        "    btnSize.SetOnClick(ActionSize_Clicked)\n\n" + 

        "    btnSizeCode := ui.GoButton(layoutBottom, \"?\")\n" + 
        "    btnSizeCode.SetMargin(0,4,0,4)\n" + 
        "    btnSizeCode.SetPadding(2,4,4,4)\n" + 
        "    btnSizeCode.SetFaceColor(ui.Color_LightGrey)\n" + 
        "    btnSizeCode.SetOnClick(ShowSizeCode_Clicked)\n\n" + 

        "    padding := ui.GoSpacer(layoutBottom, 0)\n" + 
        "    padding.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)\n\n" + 

        "    btnClose := ui.GoButton(layoutBottom, \"Close\")\n" + 
        "    btnClose.SetWidth(260)\n" + 
        "    btnClose.SetHeight(160)\n" + 
        "    btnClose.SetMargin(4,4,4,4)\n" + 
        "    btnClose.SetPadding(4,4,4,4)\n" + 
        "    btnClose.SetOnClick(ActionExit_Clicked)\n\n" + 
    
        "    lblWindowProperties.SetText(\"WINDOW PROPERTIES>\")\n\n" + 

        "    // detect main window configuration changes\n" + 
        "    mainwin.SetOnConfig(UpdateWindowProperties)\n\n" + 

        "    // show the application window\n" + 
        "    mainwin.Show()\n\n" + 
    
        "    // run the application\n" + 
        "    app.Run()\n" + 
    "}\n\n" + 

    "func ActionExit_Clicked() {\n" + 
    "    mainwin.Close()\n" + 
    "}\n\n" + 

    "func ActionMove_Clicked() {\n" + 
    "    x, y := mainwin.Pos()\n" + 
    "    mainwin.SetPos(x + 10, y + 10)\n" + 
    "}\n\n" + 

    "func ActionSize_Clicked() {\n" + 
    "    width, height := mainwin.Size()\n" + 
    "    mainwin.SetSize(width + 10, height + 10)\n" + 
    "}\n\n" + 

    "func UpdateWindowProperties() {\n" + 
    "    lblWindowProperties.SetText(GetWindowProperties())\n" + 
   "}\n" */

/*var content string = "UtopiaGio is a Go framework library built on top of the Gio library module. Gio is a cross-platform immediate mode GUI.\n\n" + 
"The GoApplication class/structure maintains a list of GoWindows and manages the control of the GoWindows and their running threads.\n\n" +
"Each GoWindow runs it's own message loop, but it will be possible to send and receive communications over channels between windows.\n\n" +
"The framework allows the building of more complex programs without the necessity to access the Gio backend. In turn this means reduced calls to Gio directly, but the ability to write specific Gio routines still remains. It is also possible to use all of the Gio widget classes by encapsulating within the GioObject structure inside the Layout function.\n\n" +
"Inheritance is achieved using the new GioObject, and the user interface is provided by the new GioWidget." +
"New layout methods have been introduced requiring a very small change to the Gio package layout module. The Gio widget module is still used on some of the widgets, but the intention is to move any relevant code for GioWidgets to the internal/widget package.\n\n" +
"Access to the underlying OS Screen and Main Window has been provided through the desktop package, making it possible to retrieve position, size and scaling of gio windows. The Pos function has been added to the Gio package, which along with the Size function allows positioning and sizing of the gio window. Also available at run time using GoWindowObj SetPos() and SetSize() functions.\n\n" + 

"###### GoWindowObj\n\n" + 
"- [**AddPopupMenu()**](#addPopupMenu])  ( popupMenu **GoPopupMenuObj** )\n" +
"- [**Centre()**](#centre)\n" + 
"  - [**ClearPopupMenus()**](#clearpopupMenus)\n" + 
"  - [**Close()**](#close)\n" + 
"  - [**EscFullScreen()**](#escFullScreen)\n\n" + 
(**[GoWindowObj]) GoFullScreen()**  
(**[GoWindowObj]) IsMainWindow()** ( isMain [**bool**][bool] )  
(**[GoWindowObj]) IsModal()** ( isModal [**bool**][bool] )  
(**[GoWindowObj]) Layout()** ( layout [**\*GoLayoutObj**][GoLayoutObj] )  
(**[GoWindowObj]) Maximize()**  
(**[GoWindowObj]) Minimize()**  
(**[GoWindowObj]) MenuBar()** ( menuBar [**\*GoMenuBarObj**][GoMenuBarObj] )  
(**[GoWindowObj]) MenuPopup(** idx **[int][int] )** ( popupMenu [**\*GoPopupMenuObj**][GoPopupMenuObj] )  

---
(**[GoWindowObj]**) **AddPopupMenu()** ( popupMenu [**\*GoPopupMenuObj**][GoPopupMenuObj] )*/

/*"### GoButton\n" + 
"Add a **GoButtonObj** to the bottom layout and set its border and padding along with the onClick action function.\n" + 
"```\n" + 
"    btnClose := ui.GoButton(layoutBottom, \"Close\")\n" + 
"    btnClose.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_Blue)\n" + 
"    btnClose.SetPadding(4,4,4,4)\n" + 
"    btnClose.SetOnClick(ActionExit_Clicked)\n" + 
"```\n" + 
"Notice the parent object **layoutBottom**. This declaration renders the Button as a child of this layout. Also because the layout has a SizePolicy of ui.PreferredHeight, the layout will size to contain the button object and not expand.\n\n" + 

"The **GoButtonObj** also has the default SizePolicy of ui.PreferredWidth and ui.PreferredWidth resulting in a button just big enough to display the caption of the button plus some default padding.\n\n" + 

"The function ActionExit_Clicked() must be declared outside the package main() function as an external function.\n\n" + 
"```\n" + 
"    func ActionExit_Clicked() {\n" + 
"        log.Println(\"ActionExit_Clicked().......\")\n" + 
"        os.Exit(0)\n" + 
"    }\n" + 
"```"*/