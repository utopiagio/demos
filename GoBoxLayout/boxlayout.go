// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoBoxLayout/boxlayout.go */

package main

import (
    //"log"
    "strconv"
    ui "github.com/utopiagio/utopia"

    //"github.com/utopiagio/utopia/desktop"
    //"github.com/utopiagio/utopia/metrics"

)

var mainwin *ui.GoWindowObj

func main() {
    // create application instance before any other objects
    app := ui.GoApplication("GoBoxLayoutDemo")
    // create application window
    mainwin = ui.GoWindow("GoBoxLayout Demo - UtopiaGio Package")
    // set the window layout style to stack widgets vertically
    mainwin.SetLayoutStyle(ui.VBoxLayout)
    mainwin.SetMargin(10,10,10,10)
    mainwin.SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_Blue)
    //mainwin.SetPadding(10,10,10,10)
    mainwin.SetSize(320, 240)
    mainwin.SetPos(100,100)   

    mainlayout := mainwin.Layout()
    var lblTest [12]*ui.GoLabelObj
    // Add some controls
    for i := 0; i < 12; i++ {
        lblTest[i] = ui.GoLabel(mainlayout, "Long String TestLabel" + strconv.Itoa(i + 1))
        lblTest[i].SetBorder(ui.BorderSingleLine, 2, 6, ui.Color_Blue)
        lblTest[i].SetSizePolicy(ui.PreferredWidth, ui.PreferredHeight)
        lblTest[i].SetMinWidth(200)
        lblTest[i].SetMaxWidth(400)
        //lblTest[i].SetMaxHeight(50)
        lblTest[i].SetMaxLines(1)
    }

    // show the application window
    mainwin.Show()
    // run the application
    app.Run()
}