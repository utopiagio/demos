// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/GoMenuBar/demomenubar.go */

package main

import (
	"log"
	"os"
	ui "github.com/utopiagio/utopia"
)

func main() {
	// create application instance before any other objects
	app := ui.GoApplication("MenuBarDemo")
	// create application window
	win := ui.GoWindow("MenuBar Demo - UtopiaGio Package")
	// set the window layout style to stack widgets vertically
	win.SetLayoutStyle(ui.VFlexBoxLayout)
	win.SetMargin(10,10,10,10)
	//win.SetMargin(15,15,15,15)
	//win.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)
	//win.SetPadding(10,10,10,10)
	win.SetPadding(0,0,0,0)


	menuBar := win.MenuBar()
	menuBar.Show()

	mnuFile := menuBar.AddMenu("File")
	mnuEdit := menuBar.AddMenu("Edit")
	mnuHelp := menuBar.AddMenu("Help")
	mnuFile.SetOnClick(ActionFile_Clicked)
	mnuEdit.SetOnClick(ActionEdit_Clicked)
	mnuHelp.SetOnClick(ActionHelp_Clicked)

	mnuFileNew := mnuFile.AddItem("New")
	mnuFileNew.SetOnClick(ActionFileNew_Clicked)
	mnuFile.AddAction("Open", ActionFileOpen_Clicked)
	mnuFile.AddAction("Save", ActionFileSave_Clicked)
	mnuFile.AddAction("Close", ActionFileClose_Clicked)
	mnuFile.AddAction("Exit", ActionExit_Clicked)

	mnuEdit.AddAction("Cut", ActionEditCut_Clicked)
	mnuEdit.AddAction("Copy", ActionEditCopy_Clicked)
	mnuEdit.AddAction("Paste", ActionEditPaste_Clicked)
	mnuEdit.AddAction("Undo", ActionEditUndo_Clicked)
	mnuEdit.AddAction("Redo", ActionEditRedo_Clicked)
	
	mnuHelp.AddAction("Documentation", ActionDocumentation_Clicked)
	mnuHelp.AddAction("About UtopiaGio", ActionAbout_Clicked)

	layoutTop := ui.GoHFlexBoxLayout(win.Layout())
	layoutTop.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
	layoutTop.SetMargin(0,0,0,0)
	layoutTop.SetPadding(10,10,10,10)
	layoutTop.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

	layoutBottom := ui.GoHFlexBoxLayout(win.Layout())
	layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
	layoutBottom.SetMargin(0,10,0,0)
	layoutBottom.SetPadding(0,0,0,0)
	layoutBottom.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

	expander := ui.GoSpacer(layoutBottom, 20)
	expander.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

	btnClose := ui.GoButton(layoutBottom, "Close")
	btnClose.SetWidth(260)
	btnClose.SetHeight(160)
	btnClose.SetMargin(10,10,10,10)
	//btnClose.SetBorder(ui.BorderSingleLine, 1, 2, ui.Color_Red)
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

func ActionFile_Clicked() {
	log.Println("ActionFile_Clicked().......")
	
}

func ActionEdit_Clicked() {
	log.Println("ActionEdit_Clicked().......")
}

func ActionHelp_Clicked() {
	log.Println("ActionHelp_Clicked().......")
}

func ActionDocumentation_Clicked() {
	log.Println("ActionDocumentationt_Clicked().......")
}

func ActionAbout_Clicked() {
	log.Println("ActionAbout_Clicked().......")
}

func ActionFileNew_Clicked() {
	log.Println("ActionFileNew_Clicked().......")
}

func ActionFileOpen_Clicked() {
	log.Println("ActionFileOpen_Clicked().......")
}

func ActionFileSave_Clicked() {
	log.Println("ActionFileSave_Clicked().......")
}

func ActionFileClose_Clicked() {
	log.Println("ActionFileClose_Clicked().......")
}

func ActionEditCopy_Clicked() {
	log.Println("ActionEditCopy_Clicked().......")
}

func ActionEditCut_Clicked() {
	log.Println("ActionEditCut_Clicked().......")
}

func ActionEditPaste_Clicked() {
	log.Println("ActionEditPaste_Clicked().......")
}

func ActionEditUndo_Clicked() {
	log.Println("ActionEditUndo_Clicked().......")
}

func ActionEditRedo_Clicked() {
	log.Println("ActionEditRedo_Clicked().......")
}

