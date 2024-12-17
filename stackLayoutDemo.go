package main

import (
	"fmt"

	"github.com/anoshenko/rui"
)

const stackLayoutDemoText = `
GridLayout {
	style = demoPage,
	content = [
		StackLayout {
			id = stackLayout, width = 100%, height = 100%
		},
		ListLayout {
			style = optionsPanel,
			content = [
				GridLayout {
					style = optionsTable,
					content = [
						Button { row = 0, column = 0:1, id = pushRed, content = "Push red view" },
						Button { row = 1, column = 0:1, id = pushGreen, content = "Push green view" },
						Button { row = 2, column = 0:1, id = pushBlue, content = "Push blue view" },
						Button { row = 3, column = 0:1, id = popView, content = "Pop view" },
						Button { row = 4, column = 0:1, id = firstToFront, content = "First view to front" },
						TextView { row = 5, text = "Animation" },
						DropDownList { row = 5, column = 1, id = pushAnimation, current = 0, items = ["end-to-start", "top-down", "down-up and rotate", "grow", "no animation"]},
						TextView { row = 6, text = "Timing" },
						DropDownList { row = 6, column = 1, id = pushTiming, current = 0, items = ["ease", "ease-in", "ease-out", "ease-in-out", "linear"]},
						TextView { row = 7, text = "Duration" },
						DropDownList { row = 7, column = 1, id = pushDuration, current = 1, items = ["0.5s", "1s", "2s"]},
						Checkbox { row = 8, column = 0:1, id = animatedToFront, content = "Move to front animation", checked = true },
					]
				}
			]
		}
	]
}
`

var stackLayoutPageCounter = 0

func createStackLayoutPage(session rui.Session, color rui.Color) rui.View {
	stackLayoutPageCounter++
	title := fmt.Sprintf("Page %d", stackLayoutPageCounter)

	return rui.NewGridLayout(session, rui.Params{
		rui.Title:           title,
		rui.BackgroundColor: color,
		rui.Margin:          rui.Px(8),
		rui.Radius:          rui.Px(16),
		rui.Border: rui.NewBorder(rui.Params{
			rui.Style:    rui.SolidLine,
			rui.ColorTag: rui.Black,
			rui.Width:    rui.Px(2),
		}),
		rui.CellVerticalAlign:   rui.CenterAlign,
		rui.CellHorizontalAlign: rui.CenterAlign,
		rui.Content:             title,
	})
}

func createStackLayoutDemo(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, stackLayoutDemoText)
	if view == nil {
		return nil
	}

	rui.Set(view, "pushTiming", rui.DropDownEvent, func(_ rui.DropDownList, index int) {
		timing := []string{rui.EaseTiming, rui.EaseInTiming, rui.EaseOutTiming, rui.EaseInOutTiming, rui.LinearTiming}
		if index >= 0 && index < len(timing) {
			rui.Set(view, "stackLayout", rui.PushTiming, timing[index])
		}
	})

	rui.Set(view, "pushDuration", rui.DropDownEvent, func(_ rui.DropDownList, index int) {
		duration := []float64{0.5, 1, 2}
		if index >= 0 && index < len(duration) {
			rui.Set(view, "stackLayout", rui.PushDuration, duration[index])
		}
	})

	rui.Set(view, "pushAnimation", rui.DropDownEvent, func(_ rui.DropDownList, index int) {
		switch index {
		case 0:
			if view.Session().TextDirection() == rui.RightToLeftDirection {
				rui.Set(view, "stackLayout", rui.PushTransform, rui.NewTransformProperty(rui.Params{
					rui.TranslateX: rui.Percent(-100),
				}))
			} else {
				rui.Set(view, "stackLayout", rui.PushTransform, rui.NewTransformProperty(rui.Params{
					rui.TranslateX: rui.Percent(100),
				}))
			}

		case 1:
			rui.Set(view, "stackLayout", rui.PushTransform, rui.NewTransformProperty(rui.Params{
				rui.TranslateY: rui.Percent(-100),
			}))

		case 2:
			rui.Set(view, "stackLayout", rui.PushTransform, rui.NewTransformProperty(rui.Params{
				rui.TranslateY: rui.Percent(100),
				rui.Rotate:     rui.PiRad(1),
			}))

		case 3:
			rui.Set(view, "stackLayout", rui.PushTransform, rui.NewTransformProperty(rui.Params{
				rui.ScaleX: 0.001,
				rui.ScaleY: 0.001,
			}))

		case 4:
			rui.Set(view, "stackLayout", rui.PushTransform, nil)
		}
	})

	rui.Set(view, "animatedToFront", rui.CheckboxChangedEvent, func(_ rui.Checkbox, checked bool) {
		rui.Set(view, "stackLayout", rui.MoveToFrontAnimation, checked)
	})

	rui.Set(view, "pushRed", rui.ClickEvent, func(_ rui.View) {
		if stackLayout := rui.StackLayoutByID(view, "stackLayout"); stackLayout != nil {
			stackLayout.Push(createStackLayoutPage(session, rui.Red))
		}
	})

	rui.Set(view, "pushGreen", rui.ClickEvent, func(_ rui.View) {
		if stackLayout := rui.StackLayoutByID(view, "stackLayout"); stackLayout != nil {
			stackLayout.Push(createStackLayoutPage(session, rui.Green))
		}
	})

	rui.Set(view, "pushBlue", rui.ClickEvent, func(_ rui.View) {
		if stackLayout := rui.StackLayoutByID(view, "stackLayout"); stackLayout != nil {
			stackLayout.Push(createStackLayoutPage(session, rui.Blue))
		}
	})

	rui.Set(view, "popView", rui.ClickEvent, func(_ rui.View) {
		if stackLayout := rui.StackLayoutByID(view, "stackLayout"); stackLayout != nil {
			stackLayout.Pop()
		}
	})

	rui.Set(view, "firstToFront", rui.ClickEvent, func(_ rui.View) {
		if stackLayout := rui.StackLayoutByID(view, "stackLayout"); stackLayout != nil {
			if views := stackLayout.Views(); len(views) > 1 {
				stackLayout.MoveToFront(views[0])
			}
		}
	})

	return view
}
