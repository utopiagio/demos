// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoTextBox/textbox.go */

package main

import (
    //"log"
    
    ui "github.com/utopiagio/utopia"
)

var mainwin *ui.GoWindowObj

func main() {
    // create application instance before any other objects
    app := ui.GoApplication("GoTextBoxDemo")
    // create application window
    mainwin = ui.GoMainWindow("GoTextBox Demo - UtopiaGio Package")
    // set the window layout style to stack widgets vertically
    mainwin.SetLayoutStyle(ui.VFlexBoxLayout)
    mainwin.SetMargin(10,10,10,10)
    mainwin.SetPadding(0,0,0,0)
    //mainwin.SetMinSize(320, 480)
    mainwin.SetSize(640, 480)
    mainwin.SetPos(100,100)

    layoutTop := ui.GoVFlexBoxLayout(mainwin.Layout())
    layoutTop.SetMargin(0,0,0,0)
    layoutTop.SetBorder(ui.BorderSingleLine, 1, 10, ui.Color_Red)
    layoutTop.SetPadding(10,10,10,10)
    layoutTop.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)

    layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())
    layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
    //layoutTop.SetMargin(10,10,10,10)
    layoutBottom.SetMargin(0,0,0,0)
    layoutBottom.SetBorder(ui.BorderSingleLine, 1, 10, ui.Color_Red)
    layoutBottom.SetPadding(0,0,0,0)

    txtEdit := ui.GoTextEdit(layoutTop, "Hint.......")
    txtEdit.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
    txtEdit.SetWidth(200)
    txtEdit.SetMaxHeight(100)
    txtEdit.SetPadding(4,4,4,4)
    txtEdit.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Red)


    mainwin.Show()

    // run the application
    app.Run()
}