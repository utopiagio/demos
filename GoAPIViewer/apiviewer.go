// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/demos/gohello/apiviewer.go */

package main

import (
    "log"
    "strings"

    ui "github.com/utopiagio/utopia"
    "github.com/utopiagio/docs"
    //"github.com/utopiagio/utopia/docs/apireference"
    "github.com/utopiagio/utopia/history"
    "golang.org/x/exp/shiny/materialdesign/icons"

    "github.com/pkg/browser"
)

var docHistory *history.GoHistoryObj
var app *ui.GoApplicationObj
var mainwin *ui.GoWindowObj         // GoVFlexBoxLayout
//var popup *ui.GoPopupWindowObj
var navigator *ui.GoListViewObj     // Vertical layout
var hdrSection *ui.GoLabelObj
var mainview *ui.GoLayoutObj        // GoVBoxLayout
var pageview *ui.GoLayoutObj        // GoHFlexBoxLayout
var currentDoc *ui.GoRichTextObj

func main() {
    var startDoc string = "ove.Overview"
    docHistory = history.GoFileHistory(startDoc + "#")
    // create application instance before any other objects
    app = ui.GoApplication("APIViewer")
    // create application window
    mainwin = ui.GoMainWindow("UtopiaGio Package - API Viewer")
    // set the window layout style to stack widgets vertically
    mainwin.SetLayoutStyle(ui.VFlexBoxLayout)
    mainwin.SetMargin(10,10,10,10)
    mainwin.SetBorder(ui.BorderSingleLine, 1, 10, ui.Color_Red)
    mainwin.SetSize(1000, 600)
    mainwin.SetPos(100,100)

    // setup the main window MenuBar
    menuBar := mainwin.MenuBar()
    menuBar.Show()
    setupMenuBar()
    /*mnuFile := */menuBar.AddMenu("File")
    mnuPreferences := menuBar.AddMenu("Preferences")
    mnuColorScheme := mnuPreferences.AddItem("Color Scheme")
    mnuColorScheme.AddAction("Red", ActionColorScheme_Red)

    navBar := ui.GoHFlexBoxLayout(mainwin.Layout())
    navBar.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)
    navBar.SetHeight(36)
    navBar.SetBackgroundColor(ui.Color_LavenderBlush)

        ui.GoSpacer(navBar, 20)

        icoBack := ui.GoIconVG(icons.NavigationArrowBack)
        navBack := ui.GoIconVGButton(navBar, icoBack)
        navBack.SetOnClick(ActionBackHistory)
        navBack.SetMargin(4,4,4,4)
        navBack.SetMinWidth(40)
        navBack.SetFaceColor(ui.Color_LavenderBlush)

        icoForward := ui.GoIconVG(icons.NavigationArrowForward)
        navForward := ui.GoIconVGButton(navBar, icoForward)
        navForward.SetOnClick(ActionForwardHistory)
        navForward.SetMargin(4,4,4,4)
        navForward.SetMinWidth(40)
        navForward.SetFaceColor(ui.Color_LavenderBlush)

        pad := ui.GoSpacer(navBar, 0)
        pad.SetSizePolicy(ui.ExpandingWidth, ui.FixedHeight)

        //hint := ui.GoLabel(navBar, "Shortcut Keys. Page Back: '<' Pge Forward: '>'")
        //hint.SetMargin(0,7,0,5)
        ui.GoSpacer(navBar, 20)

    hdrLayout := ui.GoHFlexBoxLayout(mainwin.Layout())
    hdrLayout.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
    hdrLayout.SetBackgroundColor(ui.GoColor(0xFFFFF9FA))

        hdrLogo := ui.H2Label(hdrLayout, "UtopiaGio")
        hdrLogo.SetFontBold(true)
        hdrLogo.SetFontItalic(true)
        hdrLogo.SetSizePolicy(ui.FixedWidth, ui.PreferredHeight)
        hdrLogo.SetWidth(200)
        hdrLogo.SetPadding(14,14,10,10)

        hdrSection = ui.H4Label(hdrLayout, startDoc)
        hdrSection.SetFontBold(true)
        hdrSection.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
        hdrSection.SetPadding(10,20,10,10)
        hdrSection.SetBackgroundColor(ui.GoColor(0xFFFFF9FA))

    mainpanel := ui.GoHFlexBoxLayout(mainwin.Layout())
    mainpanel.SetBackgroundColor(ui.GoColor(0xFFFFFCFE))

        navpanel := ui.GoVBoxLayout(mainpanel)
        navpanel.SetSizePolicy(ui.FixedWidth, ui.ExpandingHeight)
        navpanel.SetWidth(200)
        navpanel.SetPadding(10,10,10,10)
        navpanel.SetBackgroundColor(ui.GoColor(0xFFFFF9FA))
    
            navigator = ui.GoListView(navpanel)
            navigator.SetLayoutMode(ui.Vertical)
            navigator.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
            setupNavigator()
            navigator.SetOnItemClicked(NavLink_Clicked)
            navigator.SetOnItemDoubleClicked(NavLink_DoubleClicked)

        mainview = ui.GoVBoxLayout(mainpanel)
        mainview.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)
            pageview = ui.GoHFlexBoxLayout(mainview)
            pageview.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)
            pageview.SetPadding(10,10,10,10)

    /*if content != "" {
        docsContent[0] = content
    } else {*/

    title, name, docContent := docs.Page(startDoc)
    hdrSection.SetText(title)

    for x := 0; x < len(docContent); x++ {
        currentDoc = ui.GoRichText(pageview, name)
        currentDoc.LoadMarkDown(docContent[x])
        currentDoc.SetOnLinkClick(Link_Clicked)
    }
    // show the main window
    mainwin.Show()
    // run the application
    app.Run()
}

func ActionBackHistory() {
    link := docHistory.Back()
    doc, anchor, found := strings.Cut(link, "#")
    if found {
        //log.Println("ActionBackHistory doc=", doc, "anchor=", anchor)
        // load document into viewer
        if doc != "" && doc != currentDoc.Name() {
            //log.Println("Load doc......................")
            Load(doc)
        }
        if anchor == "" {
            mainview.ScrollToOffset(0) 
        } else {
            offset := currentDoc.AnchorTable(anchor)
            //log.Println("mainview.ScrollToOffset......................", offset)
            mainview.ScrollToOffset(offset)
        }
        switchNavigatorFocus(doc)
    }
}

func ActionForwardHistory() {
   link := docHistory.Forward()
    doc, anchor, found := strings.Cut(link, "#")
    if found {
        //log.Println("ActionForwardHistory doc=", doc, "anchor=", anchor)
        // load document into viewer
        if doc != "" && doc != currentDoc.Name() {
            //log.Println("Load doc......................")
            Load(doc)
        }
        if anchor == "" {
            mainview.ScrollToOffset(0) 
        } else {
            offset := currentDoc.AnchorTable(anchor)
            //log.Println("mainview.ScrollToOffset......................", offset)
            mainview.ScrollToOffset(offset)
        }
        switchNavigatorFocus(doc)
    }
}

func ActionColorScheme_Red() {

}

func Link_Clicked(link string) {
    // link can be a web address or a golang repository.
    // a link to a golang repository is defined as 'package.page#anchor'
    // it can be just the package page eg. api.GoWindow#
    // or an anchor link in the page eg. api.GoWindow#Close
    // if the package page is already loaded then if an anchor link is
    // provided the page scrolls to the anchor link,
    // otherwise the package page is loaded first and then if an anchor link
    // is provided then the page scrolls to the anchor link.
    if strings.HasPrefix(link, "https://") {
        // Open link in default Web Browser
        err := browser.OpenURL(link)
        if err != nil {
            log.Println("Failed to open link in browser.")
        }
    } else {
        doc, anchor, found := strings.Cut(link, "#")
        if found {
            if doc == "" {
                doc = currentDoc.Name()
                link = doc + anchor
            }
            //log.Println("Link_Clicked:", link)
            //log.Println("Link doc =", doc, "anchor =", anchor)
            //log.Println("CurrentPath:",docHistory.CurrentPath())
            //log.Println("CurrentDoc:",currentDoc.Name())
            if link == docHistory.CurrentPath() {
                return
            }
            // load document into viewer
            if doc != currentDoc.Name() {
                //log.Println("Load doc......................")
                Load(doc)
            }
            if anchor != "" {
                offset := currentDoc.AnchorTable(anchor)
                //log.Println("mainview.ScrollToOffset......................", offset)
                mainview.ScrollToOffset(offset)
            }
            //log.Println("Add docHistory: (", link, " )")
            docHistory.Add(link)
        }
    }

    //**********************************************************************************************
    // RNW 28/03/24 For future development...................
    // Popup Window provides paste window to set reference on one sidr
    // Maybe this could be used as a copy and paste facility just in plain text
                /*docsContent := getDoc(doc)
                if len(docsContent) > 0 {
                    _, wanted, ok := strings.Cut(docsContent[0], "<a name=\"" + anchor + "\">")

                    if ok {
                        log.Println("wanted ok")
                        content, _, ok := strings.Cut(wanted, "</a>")
                        log.Println("content =", content)
                        if ok {
                            log.Println("content ok")
                            if ob.popup == nil {
                                ob.popup = GoPagePopup("Function Reference", content)
                                ob.popup.Window.SetOnClose(ob.PopupClosed)
                            } else {
                                ob.popup.AddContent(content)
                            }
                        }
                    }
                }*/
    // **********************************************************************************************
}

func Load(doc string) {
    var title string
    pageview.Clear()
    packg, docName, docContent := docs.Page(doc)
    if docName != "" {
        title = packg + " - " + docName
    } else {
        title = packg
        docName = packg
    }
    hdrSection.SetText(title)
    for x := 0; x < len(docContent); x++ {
        currentDoc = ui.GoRichText(pageview, doc)
        currentDoc.LoadMarkDown(docContent[x])
        currentDoc.SetOnLinkClick(Link_Clicked)
    }
    mainview.ScrollToOffset(0)
    //log.Println("Load doc.................EXIT.....")
}

/*func getDoc(doc string) (packg string, name string, content []string) {
    return packg, name, content = docs.Page(doc)
    
}*/

func NavLink_Clicked(nodeId []int) {
    var link string
    listItem := navigator.Item(nodeId)
    doc := listItem.Text()
    switch doc {
        case "Overview":
            link = "ove.Overview#"
        case "Reference":
            link = "ref.Index#"
        case "API Reference":
            link = "api.Index#"
        default: 
            if listItem.ParentControl().ObjectType() == "GoListViewItemObj" {
                //log.Println("listItem has ParentControl")
                //log.Println("listItem.Parent.Text =", listItem.ParentControl().(*ui.GoListViewItemObj).Text())
                switch listItem.ParentControl().(*ui.GoListViewItemObj).Text() {
                    case "Overview":
                        link = "ove." + doc + "#"
                    case "Reference":
                        link = "ref." + doc + "#"
                    case "API Reference":
                        link = "api." + doc + "#"
                    default:
                        // Show error page not available
                }
            }
    }
    //log.Println("Link_Clicked(", link, ")")
    Link_Clicked(link)
}

func NavLink_DoubleClicked(nodeId []int) {
    listItem := navigator.Item(nodeId)
    if listItem.IsExpanded() {
        listItem.SetExpanded(false)
    } else {
        listItem.SetExpanded(true)
    }
    navigator.SwitchFocus(listItem)
}

func switchNavigatorFocus(docRequest string) {
    var nodeLabel []string = nil
    packg, doc, found := strings.Cut(docRequest, ".")
    if found {
        switch packg {
            case "api":
                if doc == "Index" {
                    nodeLabel = []string{"API Reference"}
                } else {
                    nodeLabel = []string{"API Reference", doc}
                }
            case "ref":
                if doc == "Index" {
                    nodeLabel = []string{"Reference"}
                } else {
                    nodeLabel = []string{"Reference", doc}
                }
            case "ove":
                if doc == "Overview" {
                    nodeLabel = []string{"Overview"}
                } else {
                    nodeLabel = []string{"Overview", doc}
                }    
            default:  
        }
    }
    if nodeLabel != nil {
        item := navigator.ItemByLabel(nodeLabel)
        navigator.SwitchFocus(item)
    }
}

func setupMenuBar() {
}

func setupNavigator() {
    /*overview := */navigator.AddListItem(nil, "Overview")
    api_reference := navigator.AddListItem(nil, "API Reference")
    api_reference.AddListItem(nil, "GoApplication")
    api_reference.AddListItem(nil, "GoButton")
    api_reference.AddListItem(nil, "GoButtonGroup")
    api_reference.AddListItem(nil, "GoCanvas")
    api_reference.AddListItem(nil, "GoCheckBox")
    api_reference.AddListItem(nil, "GoImage")
    api_reference.AddListItem(nil, "GoLabel")
    api_reference.AddListItem(nil, "GoLayout")
    api_reference.AddListItem(nil, "GoList")
    api_reference.AddListItem(nil, "GoListView")
    api_reference.AddListItem(nil, "GoListViewItem")
    api_reference.AddListItem(nil, "GoMenu")
    api_reference.AddListItem(nil, "GoMenuBar")
    api_reference.AddListItem(nil, "GoMenuItem")
    api_reference.AddListItem(nil, "GoRadioButton")
    api_reference.AddListItem(nil, "GoRichText")
    api_reference.AddListItem(nil, "GoSlider")
    api_reference.AddListItem(nil, "GoSpacer")
    api_reference.AddListItem(nil, "GoWindow")
    api_reference.AddListItem(nil, "GioObject")
    api_reference.AddListItem(nil, "GioWidget")
    /*reference := */navigator.AddListItem(nil, "Reference")
}