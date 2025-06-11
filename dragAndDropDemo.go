package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anoshenko/rui"
)

const dragAndDropDemoText = `
GridLayout {
	cell-height = "auto, 1fr, auto", cell-width = "1fr, auto",  width = 100%, height = 100%,
	content = [
		ListLayout {
			row = 0, column = 0, orientation = start-to-end, margin = 8px, vertical-align = center,
			content = [
				"Objects to drag: ",
				ImageView {
					id = dragAndDropRed, padding = 4px, margin = 4px, src = red_icon.svg, 
					radius = 4px, border = _{ width = 1px, style = solid, color = gray },
					drag-data = "text/color:red", // drag-image = red_icon.svg,
				},
				ImageView {
					id = dragAndDropGreen, padding = 4px, margin = 4px, src = green_icon.svg, 
					radius = 4px, border = _{ width = 1px, style = solid, color = gray },
					drag-data = "text/color:green", // drag-image = green_icon.svg,
				},
				ImageView {
					id = dragAndDropBlue, padding = 4px, margin = 4px, src = blue_icon.svg, 
					radius = 4px, border = _{ width = 1px, style = solid, color = gray },
					drag-data = "text/color:blue", // drag-image = blue_icon.svg,
				},
			],
		},
		TextView {
			row = 1, column = 0, text = "Drop arrea ", padding-top = 8px, padding-left = 8px,
			cell-vertical-self-align = center, cell-horizontal-self-align = center,
		},
		ListLayout {
			row = 1, column = 0, id = dragAndDropTarget, 
			orientation = start-to-end, margin = 8px, padding = 8px, list-wrap = on,
			radius = 4px, border = _{ width = 1px, style = solid, color = gray },
		},
		Resizable {
			row = 2, column = 0:1, side = top, background-color = lightgrey, height = 200px,
			content = EditView {
				id = dragAndDropEventsLog, type = multiline, read-only = true, wrap = true,
			}
		},
		ListLayout {
			row = 0:1, column = 1, style = optionsPanel,
			content = [
				GridLayout {
					style = optionsTable,
					content = [
						Checkbox { 
							row = 0, column = 0:1, id = dragAndDropImage,
							content = "drag-image"
						},
						TextView { row = 1, text = "drag-image-x-offset" },
						NumberPicker { row = 1, column = 1, id = dragAndDropImageX, value = 0, min = -100, max = 100, step = 1},
						TextView { row = 2, text = "drag-image-y-offset" },
						NumberPicker { row = 2, column = 1, id = dragAndDropImageY, value = 0, min = -100, max = 100, step = 1},
						TextView { row = 3, text = "drag-effect" },
						DropDownList { row = 3, column = 1, id = dragAndDropEffect, current = 0, items = ["undefined", "copy", "move", "link"] },
						TextView { row = 4, text = "drag-effect-allowed" },
						DropDownList { row = 4, column = 1, id = dragAndDropEffectAllowed, current = 0, items = ["undefined", "copy", "move", "copy|move", "link", "copy|link", "move|link", "all"] },
					]
				}
			]
		},
	]
}`

func createDragAndDropDemo(session rui.Session) rui.View {
	rootView := rui.CreateViewFromText(session, dragAndDropDemoText)
	if rootView == nil {
		return nil
	}

	addToLog := func(tag string, event rui.DragAndDropEvent) {
		var buffer strings.Builder

		appendBool := func(name string, value bool) {
			buffer.WriteString(`, `)
			buffer.WriteString(name)
			if value {
				buffer.WriteString(` = true`)
			} else {
				buffer.WriteString(` = false`)
			}
		}

		appendInt := func(name string, value int) {
			buffer.WriteString(`, `)
			buffer.WriteString(name)
			buffer.WriteString(` = `)
			buffer.WriteString(strconv.Itoa(value))
		}

		appendPoint := func(name string, x, y float64) {
			buffer.WriteString(fmt.Sprintf(`, %s = (%g:%g)`, name, x, y))
		}

		buffer.WriteString(tag)
		buffer.WriteString(`: TimeStamp = `)
		buffer.WriteString(strconv.FormatUint(event.TimeStamp, 10))

		appendInt("Button", event.Button)
		appendInt("Buttons", event.Buttons)
		appendPoint("(X:Y)", event.X, event.Y)
		appendPoint("Client (X:Y)", event.ClientX, event.ClientY)
		appendPoint("Screen (X:Y)", event.ScreenX, event.ScreenY)
		appendBool("CtrlKey", event.CtrlKey)
		appendBool("ShiftKey", event.ShiftKey)
		appendBool("AltKey", event.AltKey)
		appendBool("MetaKey", event.MetaKey)
		if event.Target != nil {
			buffer.WriteString(`, Target = `)
			buffer.WriteString(event.Target.ID())
		} else {
			buffer.WriteString(`, Target = nil`)
		}
		if len(event.Data) > 0 {
			lead := `, Data = ["`
			for key, value := range event.Data {
				buffer.WriteString(lead)
				lead = `, "`
				buffer.WriteString(key)
				buffer.WriteString(`":"`)
				buffer.WriteString(value)
				buffer.WriteRune('"')
			}
			buffer.WriteRune(']')
		}

		switch event.DropEffect {
		case rui.DropEffectCopy:
			buffer.WriteString(`, DropEffect="copy"`)

		case rui.DropEffectMove:
			buffer.WriteString(`, DropEffect="move"`)

		case rui.DropEffectLink:
			buffer.WriteString(`, DropEffect="link"`)

		default:
			buffer.WriteString(`, DropEffect="none"`)
		}

		effects := []string{"undefined", "copy", "move", "copy|move", "link", "copy|link", "link|move", "all"}
		if event.EffectAllowed > 0 && event.EffectAllowed < len(effects) {
			buffer.WriteString(`, EffectAllowed="`)
			buffer.WriteString(effects[event.EffectAllowed])
			buffer.WriteString(`"`)
		}

		if event.Files != nil {
			buffer.WriteString(`, Files=[`)
			for i, file := range event.Files {
				if i > 0 {
					buffer.WriteString(", ")
				}
				buffer.WriteRune('"')
				buffer.WriteString(file.Name)
				buffer.WriteRune('"')
			}
			buffer.WriteRune(']')
		}

		buffer.WriteString(";\n\n")

		rui.AppendEditText(rootView, "dragAndDropEventsLog", buffer.String())
		rui.ScrollViewToEnd(rootView, "dragAndDropEventsLog")
	}

	list := rui.ListLayoutByID(rootView, "dragAndDropTarget")

	enterListener := func(view rui.View, event rui.DragAndDropEvent) {
		addToLog(view.ID()+" drag-enter", event)
		if event.Target != nil && event.Target.ID() == "dragAndDropTarget" && list != nil {
			list.Set(rui.Border, rui.NewBorder(rui.Params{
				rui.Style:    rui.SolidLine,
				rui.ColorTag: rui.Red,
				rui.Width:    rui.Px(1),
			}))
		}
	}

	leaveListener := func(view rui.View, event rui.DragAndDropEvent) {
		addToLog(view.ID()+" drag-leave", event)
		if event.Target != nil && event.Target.ID() == "dragAndDropTarget" && list != nil {
			list.Set(rui.Border, rui.NewBorder(rui.Params{
				rui.Style:    rui.SolidLine,
				rui.ColorTag: rui.Gray,
				rui.Width:    rui.Px(1),
			}))
		}
	}

	dragViews := []string{"dragAndDropRed", "dragAndDropGreen", "dragAndDropBlue"}

	for _, id := range dragViews {
		rui.SetParams(rootView, id, rui.Params{
			rui.DragStartEvent: func(view rui.View, event rui.DragAndDropEvent) {
				addToLog(view.ID()+" drag-start", event)
			},
			rui.DragEndEvent: func(view rui.View, event rui.DragAndDropEvent) {
				addToLog(view.ID()+" drag-end", event)
			},
			rui.DragEnterEvent: enterListener,
			rui.DragLeaveEvent: leaveListener,
		})
	}

	rui.SetParams(rootView, "dragAndDropTarget", rui.Params{
		rui.DropEvent: func(view rui.View, event rui.DragAndDropEvent) {
			addToLog(view.ID()+" drop", event)

			var image string
			switch event.Data["text/color"] {
			case "red":
				image = "red_icon.svg"

			case "green":
				image = "green_icon.svg"

			case "blue":
				image = "blue_icon.svg"

			default:
				if len(event.Files) > 0 {
					for _, file := range event.Files {
						switch file.MimeType {
						case "image/png", "image/jpeg", "image/gif", "image/svg+xml":
							list.LoadFile(file, func(file rui.FileInfo, data []byte) {
								if data != nil {
									list.Append(rui.NewImageView(session, rui.Params{
										rui.Source:  rui.InlineFileFromData(data, file.MimeType),
										rui.Padding: rui.Px(2),
									}))
								}
							})
						}
					}
				}
				return
			}

			if list != nil {
				list.Append(rui.NewImageView(session, rui.Params{
					rui.Source:  image,
					rui.Padding: rui.Px(2),
				}))

				list.Set(rui.Border, rui.NewBorder(rui.Params{
					rui.Style:    rui.SolidLine,
					rui.ColorTag: rui.Gray,
					rui.Width:    rui.Px(1),
				}))
			}
		},
		/*
			rui.DragOverEvent: func(_ rui.View, event rui.DragAndDropEvent) {
				addToLog("drag-over", event)
			},
		*/
		rui.DragStartEvent: func(view rui.View, event rui.DragAndDropEvent) {
			addToLog(view.ID()+" drag-start", event)
		},
		rui.DragEndEvent: func(view rui.View, event rui.DragAndDropEvent) {
			addToLog(view.ID()+" drag-end", event)
		},
		rui.DragEnterEvent: enterListener,
		rui.DragLeaveEvent: leaveListener,
	})

	rui.Set(rootView, "dragAndDropImage", rui.CheckboxChangedEvent, func(_ rui.Checkbox, checked bool) {
		if checked {
			dragImage := []string{"red_icon.svg", "green_icon.svg", "blue_icon.svg"}
			for i, id := range dragViews {
				rui.Set(rootView, id, rui.DragImage, dragImage[i])
			}
		} else {
			for _, id := range dragViews {
				rui.Set(rootView, id, rui.DragImage, nil)
			}
		}
	})

	rui.Set(rootView, "dragAndDropImageX", rui.NumberChangedEvent, func(_ rui.NumberPicker, value float64, _ float64) {
		for _, id := range dragViews {
			rui.Set(rootView, id, rui.DragImageXOffset, value)
		}
	})

	rui.Set(rootView, "dragAndDropImageY", rui.NumberChangedEvent, func(_ rui.NumberPicker, value float64, _ float64) {
		for _, id := range dragViews {
			rui.Set(rootView, id, rui.DragImageYOffset, value)
		}
	})

	rui.Set(rootView, "dragAndDropEffect", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		effect := 0

		switch number {
		case 0:
			effect = rui.DropEffectUndefined

		case 1:
			effect = rui.DropEffectCopy

		case 2:
			effect = rui.DropEffectMove

		case 3:
			effect = rui.DropEffectLink

		default:
			return
		}

		for _, id := range dragViews {
			rui.Set(rootView, id, rui.DropEffect, effect)
		}

		rui.Set(rootView, "dragAndDropTarget", rui.DropEffect, effect)
		/*
			switch number {
			case 0:
				rui.Set(rootView, "dragAndDropTarget", rui.DropEffect, nil)

			case 1:
				rui.Set(rootView, "dragAndDropTarget", rui.DropEffect, rui.DropEffectCopy)

			case 2:
				rui.Set(rootView, "dragAndDropTarget", rui.DropEffect, rui.DropEffectMove)

			case 3:
				rui.Set(rootView, "dragAndDropTarget", rui.DropEffect, rui.DropEffectLink)
			}
		*/
	})

	rui.Set(rootView, "dragAndDropEffectAllowed", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		for _, id := range dragViews {
			rui.Set(rootView, id, rui.DropEffectAllowed, number)
		}
	})

	return rootView
}
