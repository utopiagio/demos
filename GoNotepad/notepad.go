/* notepad.go */

package main

import (
	//filepath_go "path/filepath"
	"log"
	"os"
	"strings"
	ui "github.com/utopiagio/utopia"
	dialog "github.com/utopiagio/utopia-x/filedialog"
	//icon "golang.org/x/exp/shiny/materialdesign/icons"	// eg: icon.FileFolder
)

var app *ui.GoApplicationObj
var win *ui.GoWindowObj
var txtPad *ui.GoTextEditObj
var openFilePath string
var filePath string
var fileSaved bool

func main() {
	// create application instance before any other objects
	app = ui.GoApplication("Notepad")
	// create application window
	win = ui.GoWindow("Notepad")
	// set the window layout style to stack widgets vertically
	win.SetLayoutStyle(ui.VFlexBoxLayout)
	win.SetMargin(10,10,10,10)
	//win.SetMargin(15,15,15,15)
	//win.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)
	//win.SetPadding(10,10,10,10)
	win.SetPadding(0,0,0,0)

	// setup menubar
	menuBar := win.MenuBar()
	menuBar.Show()

	mnuFile := menuBar.AddMenu("File")
	mnuEdit := menuBar.AddMenu("Edit")
	mnuFont := menuBar.AddMenu("Font")
	mnuSettings := menuBar.AddMenu("Settings")
	mnuHelp := menuBar.AddMenu("Help")
	//mnuFile.SetOnClick(ActionFile_Clicked)
	
	// add GoMenuItems to the menu dropdowns
	// GoMenuItems can be added with corresponding actions 
	mnuFile.AddAction("New", ActionFileNew_Clicked)
	mnuFile.AddAction("Open", ActionFileOpen_Clicked)
	mnuFile.AddAction("Save", ActionFileSave_Clicked)
	mnuFile.AddAction("SaveAs", ActionFileSaveAs_Clicked)
	mnuFile.AddAction("Close", ActionFileClose_Clicked)
	mnuFile.AddAction("Exit", ActionExit_Clicked)

	mnuEdit.AddAction("Cut", ActionEditCut_Clicked)
	mnuEdit.AddAction("Copy", ActionEditCopy_Clicked)
	mnuEdit.AddAction("Paste", ActionEditPaste_Clicked)
	mnuEdit.AddAction("Undo", ActionEditUndo_Clicked)
	mnuEdit.AddAction("Redo", ActionEditRedo_Clicked)
	
	mnuFontColor := mnuFont.AddItem("FontColour")
	mnuFontSize := mnuFont.AddItem("FontSize")	//, ActionFontSize_Clicked)
	/*mnuFontBold := */mnuFont.AddAction("FontBold", ActionFontBold_Clicked)
	mnuViewSize := mnuSettings.AddItem("ViewSize")

	mnuFontColor.AddAction("Red", ActionRedFontColor_Clicked)
	mnuFontColor.AddAction("Green", ActionGreenFontColor_Clicked)
	mnuFontColor.AddAction("Blue", ActionBlueFontColor_Clicked)
	mnuFontColor.AddAction("Orange", ActionOrangeFontColor_Clicked)

	mnuFontSize.AddAction("12pt", Action12ptFontSize_Clicked)
	mnuFontSize.AddAction("14pt", Action14ptFontSize_Clicked)
	mnuFontSize.AddAction("16pt", Action16ptFontSize_Clicked)

	mnuViewSize.AddAction("500 x 350", Action500ViewSizeSize_Clicked)
	mnuViewSize.AddAction("640 x 480", Action640ViewSizeSize_Clicked)
	mnuViewSize.AddAction("800 x 600", Action800ViewSizeSize_Clicked)

	mnuHelp.AddItem("Documentation")
	mnuHelp.AddItem("About UtopiaGio")

	//win.Layout().SetSpacing(10)

	// MainWindow has one edit layout
	layoutEdit := ui.GoHFlexBoxLayout(win.Layout())
	layoutEdit.SetMargin(0,0,0,0)
	layoutEdit.SetPadding(10,10,10,10)
	layoutEdit.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)

	txtPad = ui.GoTextEdit(layoutEdit, "Enter text here.")
	//txtPad.SetBorder(ui.BorderSingleLine, 1, 0, ui.Color_Blue)
	txtPad.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
	
	ui.GoSpacer(win.Layout(), 10)

	// MainWindow has one action layout 
	layoutAction := ui.GoHFlexBoxLayout(win.Layout())
	layoutAction.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)
	layoutAction.SetMargin(0,0,0,0)
	layoutAction.SetPadding(2,2,2,2)
	layoutAction.SetBorder(ui.BorderSingleLine, 2,10, ui.Color_Blue)
	layoutAction.SetHeight(200)

	// add actions  to actionLayout
	layoutMail := ui.GoHFlexBoxLayout(layoutAction)
	layoutMail.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)
	layoutMail.SetMargin(0,0,0,0)
	layoutMail.SetPadding(4,4,4,4)
	layoutMail.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_Blue)

	layoutMailAddress := ui.GoVFlexBoxLayout(layoutMail)
	layoutMailAddress.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
	layoutMailAddress.SetMargin(0,0,0,0)
	layoutMailAddress.SetPadding(0,0,0,0)
	//layoutMailAddress.SetBorder(ui.BorderSingleLine, 1, 0, ui.Color_Blue)

	lblMailAddress := ui.GoLabel(layoutMailAddress, "Mail Address")
	lblMailAddress.SetMargin(0,0,0,0)
	lblMailAddress.SetPadding(0,0,0,0)

	txtMailAddress := ui.GoTextEdit(layoutMailAddress, "richard.devel2go@gmail.com")
	txtMailAddress.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)
	txtMailAddress.SetMargin(10,0,0,0)
	txtMailAddress.SetBorder(ui.BorderSingleLine, 1, 0, ui.Color_Blue)
	txtMailAddress.SetPadding(0,0,0,0)

	btnSend := ui.GoButton(layoutMail, "Send")
	//btnClose.SetWidth(260)
	//btnClose.SetHeight(160)
	btnSend.SetMargin(20,0,0,0)
	//btnClose.SetBorder(ui.BorderSingleLine, 1, 2, ui.Color_Red)
	//btnSend.SetPadding(4,4,4,4)
	btnSend.SetOnClick(ActionSendMail_Clicked)

	ui.GoSpacer(layoutAction, 10)

	layoutPost := ui.GoVFlexBoxLayout(layoutAction)
	layoutPost.SetMargin(0,0,0,0)
	layoutPost.SetPadding(4,4,4,4)
	//layoutPost.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_Blue)
	layoutPost.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	//iconFileFolder := ui.GoIcon(icon.FileFolder)
	//ui.GoSpacer(layoutAction, 10)
	/*lblPost := */ui.GoLabel(layoutPost, "Publish")

	layoutPostTo := ui.GoHFlexBoxLayout(layoutPost)
	layoutPostTo.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

	ui.GoSpacer(layoutPostTo, 10)
	/*btnPost := ui.GoIconButton(layoutPost, iconFileFolder)*/
	facebookIcon := ui.GoIconPNG("../../icons/facebook-icon-48x48.png")

	facebook := ui.GoIconPNGButton(layoutPostTo, facebookIcon)
	facebook.SetHeight(24)
	facebook.SetWidth(24)
	//imagebox.SetMargin(5,0,0,0)
	//imagebox.SetBorder(ui.BorderSingleLine, 1, 5, ui.Color_Blue)
	//imagebox.SetPadding(4,4,4,4)
	facebook.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	ui.GoSpacer(layoutPostTo, 10)

	whatsappIcon := ui.GoIconPNG("../../icons/whatsapp-icon-48x48.png")

	whatsapp := ui.GoIconPNGButton(layoutPostTo, whatsappIcon)
	whatsapp.SetHeight(24)
	whatsapp.SetWidth(24)
	whatsapp.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	ui.GoSpacer(layoutPostTo, 10)

	signalappIcon := ui.GoIconPNG("../../icons/signalapp-icon-48x48.png")
	signal := ui.GoIconPNGButton(layoutPostTo, signalappIcon)
	signal.SetHeight(24)
	signal.SetWidth(24)
	signal.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	

	//ui.GoButton(layoutPost, "WhatsApp")
	//ui.GoButton(layoutPost, "Signal")

	expander := ui.GoSpacer(layoutAction, 20)
	expander.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	//expander := ui.GoLabel(layoutAction, "expander")
	//expander.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

	btnClose := ui.GoButton(layoutAction, "Close")
	//btnClose.SetWidth(260)
	//btnClose.SetHeight(160)
	//btnClose.SetMargin(20,0,0,0)
	btnClose.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_Blue)
	btnClose.SetPadding(4,4,4,4)
	btnClose.SetOnClick(ActionExit_Clicked)
	//btnClose.SetSizePolicy(ui.FixedWidth, ui.FixedHeight)

	//fileName = "UtopiaGio Unamed"
	openFilePath = ""
	filePath = "UtopiaGio Unamed"
	fileSaved = true
	// show the application window
	win.Show()
	// run the application
	app.Run()
}

func ActionEditCopy_Clicked() {
	log.Println("ActionEditCopy_Clicked().......")
	text := txtPad.SelectedText()
	app.ClipBoard().WriteText(text)
}

func ActionEditCut_Clicked() {
	log.Println("ActionEditCut_Clicked().......")
	text := txtPad.SelectedText()
	app.ClipBoard().WriteText(text)
	txtPad.Insert("")
}

func ActionEditPaste_Clicked() {
	log.Println("ActionEditPaste_Clicked().......")
	text := app.ClipBoard().ReadText()
	if text != "" {
		txtPad.Insert(text)
	}
}

func ActionEditUndo_Clicked() {
	log.Println("ActionEditUndo_Clicked().......")
}

func ActionEditRedo_Clicked() {
	log.Println("ActionEditRedo_Clicked().......")
}

func ActionExit_Clicked() {
	log.Println("ActionExit_Clicked().......")
	//win.Close()
	os.Exit(0)
}

func ActionFileNew_Clicked() {
	log.Println("ActionFileNew_Clicked().......")

	filePath = ""
	fileSaved = false
}

func ActionFileOpen_Clicked() {
	log.Println("ActionFile_Clicked().......")
	var resp int
	resp, filePath = dialog.GetOpenFileName(nil, openFilePath, "OpenFileDialog")
	log.Println("OpenFileDialog returned:", resp, filePath)
	if filePath == "" {filePath = "/UtopiaGio Unamed"}
	if resp > -1 {
		if _, err := os.Stat(filePath); err == nil {
			f, err := os.ReadFile(filePath)
			if err != nil {
		        log.Fatal(err)
		    }
		    text := string(f)
		    text = strings.ReplaceAll(text, "\r", "")
		    tabbedText := strings.ReplaceAll(text, "\t", "    ")
			txtPad.SetText(tabbedText)

			openFilePath = filePath
			fileSaved = true
		} else if os.IsNotExist(err) {
			log.Println("File or Path doesn't exist")
		} else {
			log.Println(err)
		}
	}
}

func ActionFileSave_Clicked() {
	log.Println("ActionFileSave_Clicked().......")
	//err = os.WriteFile(fileName, txtPad.Text(), 0644)
	if openFilePath == "" {openFilePath = "/UtopiaGio Unamed"}
	// Create creates or truncates the named file. If the file already exists, it is truncated.
    if f, errCreate := os.Create(openFilePath); errCreate == nil {
		defer f.Close()
		tabbedText := strings.Replace(txtPad.Text(), "    ", "\t", 1000)
		 _, errWrite := f.WriteString(tabbedText)
	    if errWrite == nil {
	    	fileSaved = true
	        log.Println("File written succesfully.")
	    } else {
			fileSaved = false
			log.Println("Failed to write file - errWrite", errWrite)
		}
	} else {
		fileSaved = false
		log.Println("Failed to create file - errCreate:", errCreate)
	}
}

func ActionFileSaveAs_Clicked() {
	log.Println("ActionFileSaveAs_Clicked().......")
	// show fileSave dialog
	var resp int
	resp, filePath = dialog.GetSaveFileName(nil, openFilePath, "SaveFileDialog")
	if resp > -1 {
		if f, errCreate := os.Create(filePath); errCreate == nil {
    		defer f.Close()
    		tabbedText := strings.Replace(txtPad.Text(), "    ", "\t", 1000)
    		 _, errWrite := f.WriteString(tabbedText)
		    if errWrite == nil {
		    	openFilePath = filePath
		    	fileSaved = true
		        log.Println("File written succesfully.")
		    } else {
				fileSaved = false
				log.Println("Failed to write file - errWrite", errWrite)
			}
		} else {
			fileSaved = false
			log.Println("Failed to create file - errCreate:", errCreate)
		}
	}
}

func ActionFileClose_Clicked() {
	log.Println("ActionFileClose_Clicked().......")
	openFilePath = ""
	txtPad.SetText("")
	fileSaved = true
}

func ActionFontBold_Clicked() {
	log.Println("ActionFontBold_Clicked().......")
	txtPad.SetFontBold(true)
}

func ActionRedFontColor_Clicked() {
	txtPad.SetFontColor(ui.Color_Red)
}

func ActionGreenFontColor_Clicked() {
	txtPad.SetFontColor(ui.Color_Green)
}

func ActionBlueFontColor_Clicked() {
	txtPad.SetFontColor(ui.Color_Blue)
}

func ActionOrangeFontColor_Clicked() {
	txtPad.SetFontColor(ui.Color_Orange)
}

func Action12ptFontSize_Clicked() {
	log.Println("Action12ptFontSize_Clicked().......")
	txtPad.SetFontSize(12)
}

func Action14ptFontSize_Clicked() {
	log.Println("Action14ptFontSize_Clicked().......")
	txtPad.SetFontSize(14)
}

func Action16ptFontSize_Clicked() {
	log.Println("Action16ptFontSize_Clicked().......")
	txtPad.SetFontSize(16)
}

func Action500ViewSizeSize_Clicked() {
	win.SetSize(500, 350)
}

func Action640ViewSizeSize_Clicked() {
	win.SetSize(640, 480)
}

func Action800ViewSizeSize_Clicked() {
	win.SetSize(800, 600)
}

func ActionSendMail_Clicked() {
	log.Println("ActionSendMail_Clicked().......")
}

func fileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		_, errRead := os.ReadFile(filePath)
		if errRead == nil {
			return true
		} else if os.IsNotExist(errRead) {
			log.Println("File or Path doesn't exist")
			return false
		} else {
			log.Println("File or Path read error:", errRead)
	        return false
	    }
	} else {
		return false
	}
}