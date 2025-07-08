package main

import "github.com/anoshenko/rui"

const textStyleDemoText = `
GridLayout {
	style = demoPage,
	content = [
		GridLayout {
			width = 100%, height = 100%, cell-vertical-align = center, cell-horizontal-align = center,
			content = [
				TextView {
					id = textStyleText, padding = 16px, margin=16px, max-width = 80%, 
					border = { style = solid, width = 1px, color = darkgray },
					text = "Twenty years from now you will be more disappointed by the things that you didn't do than by the ones you did do. So throw off the bowlines. Sail away from the safe harbor. Catch the trade winds in your sails. Explore. Dream. Discover."
				}
			]
		},
		ListLayout {
			style = optionsPanel,
			content = [
				GridLayout {
					style = optionsTable,
					content = [
						TextView { row = 0, text = "Font name" },
						DropDownList { row = 0, column = 1, id = textStyleFont, current = 0, 
							items = ["default", "serif", "sans-serif", "Courier new",  "monospace", "cursive", "fantasy"],
							change-listeners = {
								current = TextStyleFont,
							},
						},
						TextView { row = 1, text = "Text size" },
						DropDownList { 
							row = 1, column = 1, id = textStyleSize, current = 0, 
							items = ["1em", "14pt", "12px", "1.5em"], drop-down-event = TextStyleSize,
						},
						TextView { row = 2, text = "Text color" },
						ColorPicker { 
							row = 2, column = 1, id = textStyleColor, color-changed = TextStyleColor,
						},
						TextView { row = 3, text = "Text weight" },
						DropDownList { 
							row = 3, column = 1, id = textStyleWeight, current = 0, drop-down-event = TextStyleWeight,
							items = ["default", "thin", "extra-light", "light", "normal", "medium", "semi-bold", "bold", "extra-bold", "black"],
						},
						Checkbox { 
							row = 4, column = 0:1, id = textStyleItalic, content = "Italic",
							checkbox-event = TextStyleItalic,
						},
						Checkbox { 
							row = 5, column = 0:1, id = textStyleSmallCaps, content = "Small-caps",
							checkbox-event = TextStyleSmallCaps,
						},
						Checkbox { 
							row = 6, column = 0:1, id = textStyleStrikethrough, content = "Strikethrough",
							checkbox-event = TextStyleStrikethrough,
						},
						Checkbox { 
							row = 7, column = 0:1, id = textStyleOverline, content = "Overline",
							checkbox-event = TextStyleOverline,
						},
						Checkbox { 
							row = 8, column = 0:1, id = textStyleUnderline, content = "Underline",
							checkbox-event = TextStyleUnderline,
						},
						TextView { row = 9, text = "Line style" },
						DropDownList { 
							row = 9, column = 1, id = textStyleLineStyle, current = 0, drop-down-event = TextStyleLineStyle,
							items = ["default", "solid", "dashed", "dotted", "double", "wavy"],
						},
						TextView { row = 10, text = "Line thickness" },
						DropDownList { 
							row = 10, column = 1, id = textStyleLineThickness, current = 0, drop-down-event = TextStyleLineThickness,
							items = ["default", "1px", "1.5px", "2px", "3px", "4px"],
						},
						TextView { row = 11, text = "Line color" },
						ColorPicker { 
							row = 11, column = 1, id = textStyleLineColor, color-changed = TextStyleLineColor,
						},
						TextView { row = 12, text = "Shadow" },
						DropDownList { 
							row = 12, column = 1, id = textStyleShadow, current = 0, drop-down-event = TextStyleShadow, 
							items = ["none", "gray, (x, y)=(1px, 1px), blur=0", "blue, (x, y)=(-2px, -2px), blur=1", "green, (x, y)=(0px, 0px), blur=3px"],
						},
						TextView { row = 13, text = "User select" },
						DropDownList { 
							row = 13, column = 1, id = textStyleUserSelect, current = 0, 
							items = ["false", "true"], drop-down-event = TextStyleUserSelect, 
						},
						TextView { row = 14, text = "Text align" },
						DropDownList { 
							row = 14, column = 1, id = textStyleAlign, current = 0, 
							items = ["left", "right", "center", "justify"], drop-down-event = TextStyleAlign, 
						},
						TextView { row = 15, text = "Text wrap" },
						DropDownList { 
							row = 15, column = 1, id = textStyleWrap, current = 0, 
							items = ["wrap", "nowrap", "balance"], drop-down-event = TextStyleWrap, 
						},
						TextView { row = 17, text = "writing-mode" },
						DropDownList { 
							row = 17, column = 1, id = textWritingMode, current = 0, drop-down-event = TextWritingMode,
							items = ["horizontal-top-to-bottom", "horizontal-bottom-to-top", "vertical-right-to-left", "vertical-left-to-right"],
						},
						TextView { row = 18, text = "vertical-text-orientation" },
						DropDownList { 
							row = 18, column = 1, id = textVerticalTextOrientation, current = 0, 
							items = ["mixed-text", "upright-text"], drop-down-event = TextVerticalTextOrientation,
						},
						Button { row = 19, column = 0:1, content = "Rui text", click-event = RuiTextClick },
					]
				}
			]
		}
	]
}
`

type textStyleDemo struct {
	view rui.View
}

func createTextStyleDemo(session rui.Session) rui.View {
	return rui.CreateViewFromText(session, textStyleDemoText, new(textStyleDemo))
}

func (demo *textStyleDemo) OnCreate(view rui.View) {
	demo.view = view
}

func (demo *textStyleDemo) TextStyleFont(view rui.View) {
	fonts := []string{"", "serif", "sans-serif", "Courier new", "monospace", "cursive", "fantasy"}
	if number := rui.GetCurrent(view); number > 0 && number < len(fonts) {
		rui.Set(demo.view, "textStyleText", rui.FontName, fonts[number])
	} else {
		rui.Set(demo.view, "textStyleText", rui.FontName, nil)
	}
}

func (demo *textStyleDemo) TextStyleSize(number int) {
	sizes := []string{"1em", "14pt", "12px", "1.5em"}
	if number >= 0 && number < len(sizes) {
		rui.Set(demo.view, "textStyleText", rui.TextSize, sizes[number])
	}
}

func (demo *textStyleDemo) TextStyleColor(color rui.Color) {
	rui.Set(demo.view, "textStyleText", rui.TextColor, color)
}

func (demo *textStyleDemo) RuiTextClick() {
	rui.ShowMessage("", "<pre>"+demo.view.String()+"</pre>", demo.view.Session())
}

func (demo *textStyleDemo) TextStyleWeight(number int) {
	rui.Set(demo.view, "textStyleText", rui.TextWeight, number)
}

func (demo *textStyleDemo) TextStyleItalic(state bool) {
	rui.Set(demo.view, "textStyleText", rui.Italic, state)
}

func (demo *textStyleDemo) TextStyleSmallCaps(state bool) {
	rui.Set(demo.view, "textStyleText", rui.SmallCaps, state)
}

func (demo *textStyleDemo) TextStyleStrikethrough(state bool) {
	rui.Set(demo.view, "textStyleText", rui.Strikethrough, state)
}

func (demo *textStyleDemo) TextStyleOverline(state bool) {
	rui.Set(demo.view, "textStyleText", rui.Overline, state)
}

func (demo *textStyleDemo) TextStyleUnderline(state bool) {
	rui.Set(demo.view, "textStyleText", rui.Underline, state)
}

func (demo *textStyleDemo) TextStyleLineStyle(number int) {
	styles := []string{"inherit", "solid", "dashed", "dotted", "double", "wavy"}
	if number > 0 && number < len(styles) {
		rui.Set(demo.view, "textStyleText", rui.TextLineStyle, styles[number])
	} else {
		rui.Set(demo.view, "textStyleText", rui.TextLineStyle, nil)
	}
}

func (demo *textStyleDemo) TextStyleLineThickness(number int) {
	sizes := []string{"", "1px", "1.5px", "2px", "3px", "4px"}
	if number > 0 && number < len(sizes) {
		rui.Set(demo.view, "textStyleText", rui.TextLineThickness, sizes[number])
	} else {
		rui.Set(demo.view, "textStyleText", rui.TextLineThickness, nil)
	}
}

func (demo *textStyleDemo) TextStyleLineColor(color rui.Color) {
	rui.Set(demo.view, "textStyleText", rui.TextLineColor, color)
}

func (demo *textStyleDemo) TextStyleShadow(number int) {
	switch number {
	case 0:
		rui.Set(demo.view, "textStyleText", rui.TextShadow, nil)

	case 1:
		rui.Set(demo.view, "textStyleText", rui.TextShadow, rui.NewTextShadow(rui.Px(1), rui.Px(1), rui.Px(0), rui.Gray))

	case 2:
		rui.Set(demo.view, "textStyleText", rui.TextShadow, rui.NewTextShadow(rui.Px(-2), rui.Px(-2), rui.Px(1), rui.Blue))

	case 3:
		rui.Set(demo.view, "textStyleText", rui.TextShadow, rui.NewTextShadow(rui.Px(0), rui.Px(0), rui.Px(3), rui.Green))
	}
}

func (demo *textStyleDemo) TextStyleUserSelect(number int) {
	switch number {
	case 0:
		rui.Set(demo.view, "textStyleText", rui.UserSelect, false)
		rui.Set(demo.view, "textStyleText", rui.Cursor, "default")

	case 1:
		rui.Set(demo.view, "textStyleText", rui.UserSelect, true)
		rui.Set(demo.view, "textStyleText", rui.Cursor, "text")
	}
}

func (demo *textStyleDemo) TextStyleAlign(number int) {
	rui.Set(demo.view, "textStyleText", rui.TextAlign, number)
}

func (demo *textStyleDemo) TextStyleWrap(number int) {
	rui.Set(demo.view, "textStyleText", rui.TextWrap, number)
}

func (demo *textStyleDemo) TextWritingMode(number int) {
	rui.Set(demo.view, "textStyleText", rui.WritingMode, number)
}

func (demo *textStyleDemo) TextVerticalTextOrientation(number int) {
	rui.Set(demo.view, "textStyleText", rui.VerticalTextOrientation, number)
}
