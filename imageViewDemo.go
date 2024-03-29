package main

import (
	"fmt"

	"github.com/anoshenko/rui"
)

const imageViewDemoText = `
GridLayout {
	style = demoPage,
	content = [
		GridLayout {
			id = imageViewGrid, 
			cell-height = "auto, 1fr",
			content = [
				TextView {
					id = imageViewInfo,
				},
				ImageView {
					id = imageView1, row = 1, width = 100%, height = 100%, src = "cat.jpg", srcset = "cat.jpg, cat@2x.jpg",
					border = _{ style = solid, width = 1px, color = #FF008800 } 
				},
			],
		},
		ListLayout {
			style = optionsPanel,
			content = [
				GridLayout {
					style = optionsTable,
					content = [
						TextView { row = 0, text = "Image" },
						DropDownList { row = 0, column = 1, id = imageViewImage, current = 0, items = ["cat.jpg", "winds.png", "gifsInEmail.gif", "mountain.svg"]},
						TextView { row = 1, text = "Fit" },
						DropDownList { row = 1, column = 1, id = imageViewFit, current = 0, items = ["none", "contain", "cover", "fill", "scale-down"]},
						TextView { row = 2, text = "Horizontal align" },
						DropDownList { row = 2, column = 1, id = imageViewHAlign, current = 0, items = ["left", "right", "center"]},
						TextView { row = 3, text = "Vertical align" },
						DropDownList { row = 3, column = 1, id = imageViewVAlign, current = 0, items = ["top", "bottom", "center"]},
						TextView { row = 4, text = "background-color" },
						ColorPicker { row = 4, column = 1, id = imageViewBackgroundColor, value = white },
						TextView { row = 5, text = "blend-mode" },
						DropDownList { row = 5, column = 1, id = imageViewBlendMode, current = 0, items = ["normal", "multiply", "screen", "overlay", "darken", "lighten", "color-dodge", "color-burn", "hard-light", "soft-light", "difference", "exclusion", "hue", "saturation", "color", "luminosity"]},

					]
				}
			]
		}
	]
}
`

func createImageViewDemo(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, imageViewDemoText)
	if view == nil {
		return nil
	}

	rui.Set(view, "imageView1", rui.LoadedEvent, func(imageView rui.ImageView) {
		w, h := imageView.NaturalSize()
		rui.Set(view, "imageViewInfo", rui.Text, fmt.Sprintf("Natural size: (%g, %g). Current URL: %s", w, h, imageView.CurrentSource()))
	})

	rui.Set(view, "imageView1", rui.ErrorEvent, func(_ rui.ImageView) {
		rui.Set(view, "imageViewInfo", rui.Text, "Image loading error")
	})

	rui.Set(view, "imageViewImage", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		images := []string{"cat.jpg", "winds.png", "gifsInEmail.gif", "mountain.svg"}
		srcset := []string{"cat.jpg, cat@2x.jpg", "", "", ""}
		if number < len(images) {
			rui.Set(view, "imageView1", rui.SrcSet, srcset[number])
			rui.Set(view, "imageView1", rui.Source, images[number])
		}
	})

	rui.Set(view, "imageViewFit", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		rui.Set(view, "imageView1", rui.Fit, number)
	})

	rui.Set(view, "imageViewHAlign", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		rui.Set(view, "imageView1", rui.ImageHorizontalAlign, number)
	})

	rui.Set(view, "imageViewVAlign", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		rui.Set(view, "imageView1", rui.ImageVerticalAlign, number)
	})

	rui.Set(view, "imageViewBackgroundColor", rui.ColorChangedEvent, func(_ rui.ColorPicker, color rui.Color) {
		rui.Set(view, "imageViewGrid", rui.BackgroundColor, color)
	})

	rui.Set(view, "imageViewBlendMode", rui.DropDownEvent, func(_ rui.DropDownList, index int) {
		rui.Set(view, "imageView1", rui.MixBlendMode, index)
	})

	return view
}
