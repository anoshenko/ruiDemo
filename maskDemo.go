package main

import "github.com/anoshenko/rui"

const maskDemoText = `
GridLayout {
	style = demoPage,
	content = [
		ListLayout {
			width = 100%, height = 100%, padding = 32px,
			content = [
				ImageView {
					id = maskView, width = 100%, height = 150%, padding = 16px,
					src = "cat.jpg", fit = contain, background-color = red,
					horizontal-align = center, vertical-align = center,
					border = _{ style = dotted, width = 8px, color = blue },
					mask = image { src = star-mask.svg },
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
						DropDownList { row = 0, column = 1, id = maskImage1, current = 0, items = ["star-mask.svg", "Linear gradient", "Radial gradient", "Conic gradient"]},
						TextView { row = 1, text = "Fit" },
						DropDownList { row = 1, column = 1, id = maskFit1, current = 0, items = ["none", "contain", "cover"]},
						TextView { row = 2, text = "Horizontal align" },
						DropDownList { row = 2, column = 1, id = maskHAlign1, current = 0, items = ["left", "right", "center"]},
						TextView { row = 3, text = "Vertical align" },
						DropDownList { row = 3, column = 1, id = maskVAlign1, current = 0, items = ["top", "bottom", "center"]},
						TextView { row = 4, text = "Repeat" },
						DropDownList { row = 4, column = 1, id = maskRepeat1, current = 0, items = ["no-repeat", "repeat", "repeat-x", "repeat-y", "round", "space"]},
						TextView { row = 5, text = "mask-clip" },
						DropDownList { row = 5, column = 1, id = maskClip1, current = 0, items = ["border-box", "padding-box", "content-box"]},
						TextView { row = 6, text = "mask-origin" },
						DropDownList { row = 6, column = 1, id = maskOrigin1, current = 0, items = ["border-box", "padding-box", "content-box"]},
					]
				}
			]
		}
	]
}
`

func createMaskDemo(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, maskDemoText)
	if view == nil {
		return nil
	}

	updateMask1 := func(rui.DropDownList, int) {
		//images := []string{"cat.jpg", "winds.png", "gifsInEmail.gif", "mountain.svg"}
		image := rui.NewBackgroundImage(rui.Params{
			rui.Source:          "star-mask.svg", //images[rui.GetCurrent(view, "maskImage1")],
			rui.Fit:             rui.GetCurrent(view, "maskFit1"),
			rui.HorizontalAlign: rui.GetCurrent(view, "maskHAlign1"),
			rui.VerticalAlign:   rui.GetCurrent(view, "maskVAlign1"),
			rui.Repeat:          rui.GetCurrent(view, "maskRepeat1"),
			rui.Attachment:      rui.GetCurrent(view, "maskAttachment1"),
		})
		rui.Set(view, "maskView", rui.Mask, image)
	}

	for _, id := range []string{
		"maskFit1",
		"maskHAlign1",
		"maskVAlign1",
		"maskRepeat1",
	} {
		rui.Set(view, id, rui.DropDownEvent, updateMask1)
	}

	rui.Set(view, "maskImage1", rui.DropDownEvent, func(list rui.DropDownList, index int) {
		switch index {
		case 0:
			updateMask1(list, index)

		case 1:
			rui.Set(view, "maskView", rui.Mask, rui.NewBackgroundLinearGradient(rui.Params{
				rui.Direction: rui.ToBottomGradient,
				rui.Gradient: []rui.BackgroundGradientPoint{
					{Color: 0x00FFFFFF, Pos: rui.Percent(0)},
					{Color: 0xFFFFFFFF, Pos: rui.Percent(100)},
				},
			}))

		case 2:
			rui.Set(view, "maskView", rui.Mask, rui.NewBackgroundRadialGradient(rui.Params{
				rui.Gradient: []rui.BackgroundGradientPoint{
					{Color: 0xFF000000, Pos: rui.Percent(0)},
					{Color: 0x00000001, Pos: rui.Percent(100)},
				},
			}))

		case 3:
			rui.Set(view, "maskView", rui.Mask, rui.NewBackgroundConicGradient(rui.Params{
				rui.Repeating: true,
				rui.Gradient: []rui.BackgroundGradientAngle{
					{Color: 0xFFFFFFFF, Angle: 0},
					{Color: 0x00FFFFFF, Angle: rui.PiRad(1)},
				},
			}))
		}
	})

	rui.Set(view, "maskClip1", rui.DropDownEvent, func(_ rui.DropDownList, index int) {
		rui.Set(view, "maskView", rui.MaskClip, index)
	})

	rui.Set(view, "maskOrigin1", rui.DropDownEvent, func(_ rui.DropDownList, index int) {
		rui.Set(view, "maskView", rui.MaskOrigin, index)
	})

	return view
}
