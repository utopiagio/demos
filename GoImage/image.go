// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoImage/image.go */

package main

import (
    ui "github.com/utopiagio/utopia"
)

var mainwin *ui.GoWindowObj

func main() {
    // create application instance before any other objects
    app := ui.GoApplication("GoBoxLayoutDemo")
    // create application window
    mainwin = ui.GoMainWindow("GoBoxLayout Demo - UtopiaGio Package")
    // set the window layout style to stack widgets vertically
    mainwin.SetLayoutStyle(ui.HBoxLayout)
    mainwin.SetMargin(10,10,10,10)
    mainwin.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_Blue)
    mainwin.SetPadding(10,10,10,10)
    mainwin.SetSize(640, 240)
    mainwin.SetPos(100,100)
    
    imgNSU := ui.GoImage(mainwin.Layout(), "C:/Users/richa/OneDrive/Pictures/Saved Pictures/tony1/image1.jpeg")
    imgNSU.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)
    imgNSU.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_Blue)
    imgNSU.SetPadding(10,10,10,10)
    imgBenelli := ui.GoImage(mainwin.Layout(), "C:/Users/richa/OneDrive/Pictures/Saved Pictures/tony2/benelli750.jpeg")
    imgBenelli.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
    imgBenelli.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_Blue)
    imgBenelli.SetPadding(10,10,10,10)
    // show the application window
    mainwin.Show()
    
    // run the application
    app.Run()
}