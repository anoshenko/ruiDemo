package main

import (
	"fmt"

	"github.com/anoshenko/rui"
)

const rootViewText = `
GridLayout {
	id = rootLayout, width = 100%, height = 100%, cell-height = "auto, 1fr",
	content = [
		GridLayout {
			id = rootTitle, width = 100%, cell-width = "auto, 1fr", 
			cell-vertical-align = center, background-color = #ffc0ded9, text-color = black,
			content = [
				ImageView { 
					id = rootTitleButton, padding = 8px, src = menu_icon.svg,
					tooltip = "Select demo (alt-M)"
				},
				TextView { 
					id = rootTitleText, column = 1, padding-left = 8px, text = "Title",
				}
			],
		},
		StackLayout {
			id = rootViews, row = 1, move-to-front-animation = false,
		}
	]
}
`

type demoPage struct {
	title   string
	creator func(session rui.Session) rui.View
	view    rui.View
}

type demoSession struct {
	rootView rui.View
	pages    []demoPage
}

func (demo *demoSession) OnStart(session rui.Session) {
	rui.DebugLog("Session start")
}

func (demo *demoSession) OnFinish(session rui.Session) {
	rui.DebugLog("Session finish")
}

func (demo *demoSession) OnResume(session rui.Session) {
	rui.DebugLog("Session resume")
}

func (demo *demoSession) OnPause(session rui.Session) {
	rui.DebugLog("Session pause")
}

func (demo *demoSession) OnDisconnect(session rui.Session) {
	rui.DebugLog("Session disconnect")
}

func (demo *demoSession) OnReconnect(session rui.Session) {
	rui.DebugLog("Session reconnect")
}

func createDemo(_ rui.Session) rui.SessionContent {
	sessionContent := new(demoSession)
	sessionContent.pages = []demoPage{
		{"Text style", createTextStyleDemo, nil},
		{"View border", viewDemo, nil},
		{"Background image", createBackgroundDemo, nil},
		{"Mask image", createMaskDemo, nil},
		{"ListLayout", createListLayoutDemo, nil},
		{"GridLayout", createGridLayoutDemo, nil},
		{"ColumnLayout", createColumnLayoutDemo, nil},
		{"StackLayout", createStackLayoutDemo, nil},
		{"Tabs", createTabsDemo, nil},
		{"Resizable", createResizableDemo, nil},
		{"ListView", createListViewDemo, nil},
		{"Checkbox", createCheckboxDemo, nil},
		{"Controls", createControlsDemo, nil},
		{"FilePicker", createFilePickerDemo, nil},
		{"TableView", createTableViewDemo, nil},
		{"EditView", createEditDemo, nil},
		{"ImageView", createImageViewDemo, nil},
		{"Canvas", createCanvasDemo, nil},
		{"VideoPlayer", createVideoPlayerDemo, nil},
		{"AudioPlayer", createAudioPlayerDemo, nil},
		{"Popups", createPopupDemo, nil},
		{"Filter", createFilterDemo, nil},
		{"Clip", createClipDemo, nil},
		{"Transform", transformDemo, nil},
		{"Animation", createAnimationDemo, nil},
		{"Transition", createTransitionDemo, nil},
		{"Key events", createKeyEventsDemo, nil},
		{"Mouse events", createMouseEventsDemo, nil},
		{"Pointer events", createPointerEventsDemo, nil},
		{"Touch events", createTouchEventsDemo, nil},
		{"Drag and drop", createDragAndDropDemo, nil},
	}

	return sessionContent
}

func (demo *demoSession) CreateRootView(session rui.Session) rui.View {
	demo.rootView = rui.CreateViewFromText(session, rootViewText)
	if demo.rootView == nil {
		return nil
	}

	rui.Set(demo.rootView, "rootTitleButton", rui.ClickEvent, demo.clickMenuButton)
	session.SetHotKey(rui.KeyM, rui.AltKey, func(session rui.Session) {
		demo.clickMenuButton()
	})
	demo.showPage(0)
	return demo.rootView
}

func (demo *demoSession) clickMenuButton() {
	items := make([]string, len(demo.pages))
	for i, page := range demo.pages {
		items[i] = page.title
	}

	buttonFrame := rui.ViewByID(demo.rootView, "rootTitleButton").Frame()

	rui.ShowMenu(demo.rootView.Session(), rui.Params{
		rui.Items:           items,
		rui.OutsideClose:    true,
		rui.VerticalAlign:   rui.TopAlign,
		rui.HorizontalAlign: rui.LeftAlign,
		rui.MarginLeft:      rui.Px(buttonFrame.Bottom() / 2),
		rui.Arrow:           rui.LeftArrow,
		rui.ArrowAlign:      rui.LeftAlign,
		rui.ArrowSize:       rui.Px(12),
		rui.ArrowOffset:     rui.Px(buttonFrame.Left + (buttonFrame.Width-12)/2),
		rui.PopupMenuResult: func(n int) {
			demo.showPage(n)
		},
	})
}

func (demo *demoSession) showPage(index int) {
	if index < 0 || index >= len(demo.pages) {
		return
	}

	if stackLayout := rui.StackLayoutByID(demo.rootView, "rootViews"); stackLayout != nil {
		if demo.pages[index].view == nil {
			demo.pages[index].view = demo.pages[index].creator(demo.rootView.Session())
			stackLayout.Append(demo.pages[index].view)
		} else {
			stackLayout.MoveToFront(demo.pages[index].view)
			if view := rui.CanvasViewByID(demo.pages[index].view, "canvasDemoPage", "canvas"); view != nil {
				view.Redraw()
			}
		}
		rui.Set(demo.rootView, "rootTitleText", rui.Text, demo.pages[index].title)
		demo.rootView.Session().SetTitle(demo.pages[index].title)
	}
}

func main() {
	rui.ProtocolInDebugLog = true
	/*
		rui.SetDebugLog(func(text string) {
			if len(text) > 120 {
				text = text[:120] + "..."
			}
			log.Println(text)
		})
	*/
	rui.AddEmbedResources(&resources)

	//addr := rui.GetLocalIP() + ":8080"
	addr := "localhost:8000"
	fmt.Print(addr)
	rui.OpenBrowser("http://" + addr)
	rui.StartApp(addr, createDemo, rui.AppParams{
		Title:      "RUI demo",
		Icon:       "icon.svg",
		TitleColor: rui.Color(0xffc0ded9),
		//NoSocket:   true,
		//SocketAutoClose: 5,
	})
}
