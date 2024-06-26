package main

import (
	"github.com/anoshenko/rui"
)

const gridLayoutDemoText = `
GridLayout {
	style = demoPage,
	content = [
		GridLayout {
			id = gridLayout, width = 100%, height = 100%, 
			cell-width = "150px, 1fr, 30%", cell-height = "25%, 200px, 1fr",
			content = [
				TextView { row = 0, column = 0:1,
					text = "View 1", text-align = center, vertical-align = center, 
					background-color = #DDFF0000, radius = 8px, padding = 32px, 
					border = _{ style = solid, width = 1px, color = #FFA0A0A0 } 
				},
				TextView { row = 0:1, column = 2,
					text = "View 2", text-align = center, vertical-align = center, 
					background-color = #DD00FF00, radius = 8px, padding = 32px, 
					border = _{ style = solid, width = 1px, color = #FFA0A0A0 } 
				},
				TextView { row = 1:2, column = 0,
					text = "View 3", text-align = center, vertical-align = center, 
					background-color = #DD0000FF, radius = 8px, padding = 32px, 
					border = _{ style = solid, width = 1px, color = #FFA0A0A0 } 
				},
				TextView { row = 1, column = 1, 
					id = centralView,
					text = "View 4", text-align = center, vertical-align = center, 
					background-color = #DDFF00FF, radius = 8px, padding = 32px, 
					border = _{ style = solid, width = 1px, color = #FFA0A0A0 } 
				},
				TextView { row = 2, column = 1:2,
					text = "View 5", text-align = center, vertical-align = center, 
					background-color = #DD00FFFF, radius = 8px, padding = 32px, 
					border = _{ style = solid, width = 1px, color = #FFA0A0A0 } 
				},
			]
		},
		ListLayout {
			style = optionsPanel,
			content = [
				GridLayout {
					style = optionsTable,
					content = [
						TextView { row = 0, text = "Vertical align" },
						DropDownList { row = 0, column = 1, id = gridVAlign, current = 3,
							items = ["top", "bottom", "center", "stretch"]
						},
						TextView { row = 1, text = "Horizontal align" },
						DropDownList { row = 1, column = 1, id = gridHAlign, current = 3, 
							items = ["left", "right", "center", "stretch"]
						},
						TextView { row = 2, text = "Column gap" },
						DropDownList { row = 2, column = 1, id = gridColumnGap, current = 0, items = ["0", "8px"] },
						TextView { row = 3, text = "Row gap" },
						DropDownList { row = 3, column = 1, id = gridRowGap, current = 0, items = ["0", "8px"] },
						TextView { row = 4, text = "Vertical align of central view" },
						DropDownList { row = 4, column = 1, id = gridSelfVAlign, current = 0,
							items = ["default", "top", "bottom", "center", "stretch"]
						},
						TextView { row = 5, text = "Horizontal align of central view" },
						DropDownList { row = 5, column = 1, id = gridSelfHAlign, current = 0, 
							items = ["default", "left", "right", "center", "stretch"]
						},

					]
				}
			]
		}
	]
}
`

func createGridLayoutDemo(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, gridLayoutDemoText)
	if view == nil {
		return nil
	}

	rui.Set(view, "gridHAlign", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		rui.Set(view, "gridLayout", rui.CellHorizontalAlign, number)
	})

	rui.Set(view, "gridVAlign", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		rui.Set(view, "gridLayout", rui.CellVerticalAlign, number)
	})

	rui.Set(view, "gridSelfHAlign", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		if number == 0 {
			rui.Set(view, "centralView", rui.CellHorizontalSelfAlign, nil)
		} else {
			rui.Set(view, "centralView", rui.CellHorizontalSelfAlign, number-1)
		}
	})

	rui.Set(view, "gridSelfVAlign", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		if number == 0 {
			rui.Set(view, "centralView", rui.CellVerticalSelfAlign, nil)
		} else {
			rui.Set(view, "centralView", rui.CellVerticalSelfAlign, number-1)
		}
	})

	rui.Set(view, "gridColumnGap", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		switch number {
		case 0:
			rui.Set(view, "gridLayout", rui.GridColumnGap, rui.SizeUnit{Type: rui.Auto, Value: 0})

		case 1:
			rui.Set(view, "gridLayout", rui.GridColumnGap, rui.Px(8))
		}
	})

	rui.Set(view, "gridRowGap", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		switch number {
		case 0:
			rui.Set(view, "gridLayout", rui.GridRowGap, rui.SizeUnit{Type: rui.Auto, Value: 0})

		case 1:
			rui.Set(view, "gridLayout", rui.GridRowGap, rui.Px(8))
		}
	})

	return view
}
