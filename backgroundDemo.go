package main

import "github.com/anoshenko/rui"

const backgroundDemoText = `
GridLayout {
	style = demoPage,
	content = [
		ListLayout {
			width = 100%, height = 100%, padding = 32px,
			content = [
				TextView {
					id = backgroundView, width = 100%, height = 150%, padding = 16px,
					text = "Sample text", text-size = 4em, background-color = red,
					border = _{ style = dotted, width = 8px, color = blue },
					background = image { src = cat.jpg },
				}
			]
		},		
		ListLayout {
			style = optionsPanel,
			content = [
				GridLayout {
					style = optionsTable,
					content = [
						TextView { row = 0, text = "Image" },
						DropDownList { row = 0, column = 1, id = backgroundImage1, current = 0, items = ["cat.jpg", "winds.png", "gifsInEmail.gif", "mountain.svg"]},
						TextView { row = 1, text = "Fit" },
						DropDownList { row = 1, column = 1, id = backgroundFit1, current = 0, items = ["none", "contain", "cover"]},
						TextView { row = 2, text = "Horizontal align" },
						DropDownList { row = 2, column = 1, id = backgroundHAlign1, current = 0, items = ["left", "right", "center"]},
						TextView { row = 3, text = "Vertical align" },
						DropDownList { row = 3, column = 1, id = backgroundVAlign1, current = 0, items = ["top", "bottom", "center"]},
						TextView { row = 4, text = "Repeat" },
						DropDownList { row = 4, column = 1, id = backgroundRepeat1, current = 0, items = ["no-repeat", "repeat", "repeat-x", "repeat-y", "round", "space"]},
						TextView { row = 5, text = "Clip" },
						DropDownList { row = 5, column = 1, id = backgroundClip1, current = 0, items = ["border-box", "padding-box", "content-box", "text"]},
						TextView { row = 6, text = "background-origin" },
						DropDownList { row = 6, column = 1, id = backgroundOrigin1, current = 1, items = ["border-box", "padding-box", "content-box"]},
						TextView { row = 7, text = "Attachment" },
						DropDownList { row = 7, column = 1, id = backgroundAttachment1, current = 0, items = ["scroll", "fixed", "local"]},
						TextView { row = 8, text = "Color" },
						ColorPicker { row = 8, column = 1, id = backgroundColor1, value = #FFF0F8FF },
						TextView { row = 9, text = "blend-mode" },
						DropDownList { row = 9, column = 1, id = backgroundBlendMode1, current = 0, items = ["normal", "multiply", "screen", "overlay", "darken", "lighten", "color-dodge", "color-burn", "hard-light", "soft-light", "difference", "exclusion", "hue", "saturation", "color", "luminosity"]},
					]
				}
			]
		}
	]
}
`

func createBackgroundDemo(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, backgroundDemoText)
	if view == nil {
		return nil
	}

	updateBackground1 := func(rui.DropDownList, int) {
		images := []string{"cat.jpg", "winds.png", "gifsInEmail.gif", "mountain.svg"}
		image := rui.NewBackgroundImage(rui.Params{
			rui.Source:          images[rui.GetCurrent(view, "backgroundImage1")],
			rui.Fit:             rui.GetCurrent(view, "backgroundFit1"),
			rui.HorizontalAlign: rui.GetCurrent(view, "backgroundHAlign1"),
			rui.VerticalAlign:   rui.GetCurrent(view, "backgroundVAlign1"),
			rui.Repeat:          rui.GetCurrent(view, "backgroundRepeat1"),
			rui.Attachment:      rui.GetCurrent(view, "backgroundAttachment1"),
		})
		rui.Set(view, "backgroundView", rui.Background, image)
	}

	for _, id := range []string{
		"backgroundImage1",
		"backgroundFit1",
		"backgroundHAlign1",
		"backgroundVAlign1",
		"backgroundRepeat1",
		"backgroundAttachment1",
	} {
		rui.Set(view, id, rui.DropDownEvent, updateBackground1)
	}

	rui.Set(view, "backgroundClip1", rui.DropDownEvent, func(index int) {
		rui.Set(view, "backgroundView", rui.BackgroundClip, index)
	})

	rui.Set(view, "backgroundOrigin1", rui.DropDownEvent, func(index int) {
		rui.Set(view, "backgroundView", rui.BackgroundOrigin, index)
	})

	rui.Set(view, "backgroundColor1", rui.ColorChangedEvent, func(_ rui.ColorPicker, color rui.Color) {
		rui.Set(view, "backgroundView", rui.BackgroundColor, color)
	})

	rui.Set(view, "backgroundBlendMode1", rui.DropDownEvent, func(_ rui.DropDownList, index int) {
		rui.Set(view, "backgroundView", rui.BackgroundBlendMode, index)
	})

	return view
}
