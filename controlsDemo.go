package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/anoshenko/rui"
)

const controlsDemoText = `
ListLayout {
	width = 100%, height = 100%, orientation = vertical, padding = 16px,
	content = [
		DetailsView {
			margin = 8px,
			summary = TextView { text = "Details title", background-color = #FFEEEEEE, padding = 4px },
			hide-summary-marker = true,
			content = "Details content"
		}
		ListLayout { orientation = horizontal, vertical-align = center, padding = 8px,
			border = _{ width = 1px, style = solid, color = #FF000000 }, radius = 4px,
			content = [
				Checkbox { id = controlsCheckbox, content = "Checkbox" },
				Button { id = controlsCheckboxButton, margin-left = 32px, content = "Check checkbox" },
			]
		},
		ListLayout { orientation = horizontal, margin-top = 16px, vertical-align = center,
			border = _{ width = 1px, style = solid, color = #FF000000 }, radius = 4px,
			content = [
				DropDownList { 
					id = controlsDropDownList, current = 0, margin = 16px,
					items = ["Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6"],
					disabled-items = [2, 4],
					item-separators = [1, 3],
				},
				"Selected: "
				TextView {
					id = controlsDropDownSelected, text = "Item 1", margin-left = 8px, 
				},
			]
		},
		ListLayout { orientation = horizontal, margin-top = 16px, padding = 8px, vertical-align = center,
			border = _{ width = 1px, style = solid, color = #FF000000 }, radius = 4px,
			content = [
				Button { id = controlsProgressDec, content = "<<" },
				Button { id = controlsProgressInc, content = ">>", margin-left = 12px },
				ProgressBar { id = controlsProgress, max = 100, value = 50, margin-left = 12px  },
				TextView { id = controlsProgressLabel, text = "50 / 100", margin-left = 12px },
			]
		},
		ListLayout { orientation = horizontal, margin-top = 16px, padding = 8px, vertical-align = center,
			border = _{ width = 1px, style = solid, color = #FF000000 }, radius = 4px,
			content = [
				"Enter number (-5...10)",
				NumberPicker { id = controlsNumberEditor, type = editor, width = 80px, 
					margin-left = 12px, min = -5, max = 10, step = 0.1, value = 0, precision = 1,
					data-list = [-2, 0, 2.5, 5] 
				},
				NumberPicker { id = controlsNumberSlider, type = slider, width = 150px, 
					margin-left = 12px, min = -5, max = 10, step = 0.1, value = 0, precision = 1,
					data-list = [-5, 0, 5, 10]  
				}
			]
		},
		ListLayout { orientation = horizontal, margin-top = 16px, padding = 8px, vertical-align = center,
			border = _{ width = 1px, style = solid, color = #FF000000 }, radius = 4px,
			content = [
				"Select color",
				ColorPicker { id = controlsColorPicker, value = #0000FF,
					margin = _{ left = 12px, right = 24px},
					data-list = [ #FF00FF00, yellow, black ]
				},
				"Result",
				View { id = controlsColorResult, width = 24px, height = 24px, margin-left = 12px, background-color = #0000FF }
			]
		},
		ListLayout { orientation = horizontal, margin-top = 16px, padding = 8px, vertical-align = center,
			border = _{ width = 1px, style = solid, color = #FF000000 }, radius = 4px,
			content = [
				"Select a time and date:",
				TimePicker { id = controlsTimePicker, min = "00:00", margin-left = 12px,
					data-list = [ "9:30", "12:00", "15:45" ] 
				},
				DatePicker { id = controlsDatePicker, min = "2001-01-01", margin-right = 24px,
					data-list = [ "2010-07-15", "2015-04-21" ] 
				},
				"Result:",
				TextView { id = controlsDateResult, margin-left = 12px }
			]
		},
		ListLayout { orientation = horizontal, margin-top = 16px, padding = 8px, vertical-align = center,
			border = _{ width = 1px, style = solid, color = #FF000000 }, radius = 4px,
			content = [
				View { 
					id = controlsHidden, width = 32px, height = 16px, margin-right = 16px, 
					background-color = red 
				},
				"Visibility",
				DropDownList { 
					id = controlsHiddenList, margin-left = 16px, current = 0, 
					items = ["visible", "invisible", "gone"]
				},
			],
		},
		ListLayout { orientation = horizontal, margin-top = 16px, padding = 8px, vertical-align = center,
			border = _{ width = 1px, style = solid, color = #FF000000 }, radius = 4px,
			content = [
				Button { 
					id = timerButton, content = "Start timer",
				},
				TextView { 
					id = timerText, margin-left = 16px,
				},
			]
		},
		Button {
			id = controlsMessage, margin-top = 16px, content = "Show message"
		}
	]
}
`

func createControlsDemo(session rui.Session) rui.View {
	view := rui.CreateViewFromText(session, controlsDemoText)
	if view == nil {
		return nil
	}

	rui.Set(view, "controlsCheckbox", rui.CheckboxChangedEvent, func(_ rui.Checkbox, checked bool) {
		if checked {
			rui.Set(view, "controlsCheckboxButton", rui.Content, "Uncheck checkbox")
		} else {
			rui.Set(view, "controlsCheckboxButton", rui.Content, "Check checkbox")
		}
	})

	rui.Set(view, "controlsCheckboxButton", rui.ClickEvent, func(rui.View) {
		checked := rui.IsCheckboxChecked(view, "controlsCheckbox")
		rui.Set(view, "controlsCheckbox", rui.Checked, !checked)
	})

	rui.Set(view, "controlsDropDownList", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		rui.Set(view, "controlsDropDownSelected", rui.Text, fmt.Sprintf("Item %d", number+1))
	})

	setProgressBar := func(dx float64) {
		if value := rui.GetProgressBarValue(view, "controlsProgress"); value >= 0 {
			max := rui.GetProgressBarMax(view, "controlsProgress")
			newValue := math.Min(math.Max(0, value+dx), max)
			if newValue != value {
				rui.Set(view, "controlsProgress", rui.Value, newValue)
				rui.Set(view, "controlsProgressLabel", rui.Text, fmt.Sprintf("%g / %g", newValue, max))
			}
		}
	}

	rui.Set(view, "controlsProgressDec", rui.ClickEvent, func(rui.View) {
		setProgressBar(-1)
	})

	rui.Set(view, "controlsProgressInc", rui.ClickEvent, func(rui.View) {
		setProgressBar(+1)
	})

	rui.Set(view, "controlsNumberEditor", rui.NumberChangedEvent, func(_ rui.NumberPicker, newValue float64) {
		rui.Set(view, "controlsNumberSlider", rui.Value, newValue)
	})

	rui.Set(view, "controlsNumberSlider", rui.NumberChangedEvent, func(_ rui.NumberPicker, newValue float64) {
		rui.Set(view, "controlsNumberEditor", rui.Value, newValue)
	})

	rui.Set(view, "controlsColorPicker", rui.ColorChangedEvent, func(_ rui.ColorPicker, newColor rui.Color) {
		rui.Set(view, "controlsColorResult", rui.BackgroundColor, newColor)
	})

	rui.Set(view, "controlsTimePicker", rui.Value, demoTime)
	rui.Set(view, "controlsDatePicker", rui.Value, demoTime)

	rui.Set(view, "controlsTimePicker", rui.TimeChangedEvent, func(_ rui.TimePicker, newDate time.Time) {
		demoTime = time.Date(demoTime.Year(), demoTime.Month(), demoTime.Day(), newDate.Hour(), newDate.Minute(),
			newDate.Second(), newDate.Nanosecond(), demoTime.Location())
		rui.Set(view, "controlsDateResult", rui.Text, demoTime.Format(time.RFC1123))
	})

	rui.Set(view, "controlsDatePicker", rui.DateChangedEvent, func(_ rui.DatePicker, newDate time.Time) {
		demoTime = time.Date(newDate.Year(), newDate.Month(), newDate.Day(), demoTime.Hour(), demoTime.Minute(),
			demoTime.Second(), demoTime.Nanosecond(), demoTime.Location())
		rui.Set(view, "controlsDateResult", rui.Text, demoTime.Format(time.RFC1123))
	})

	rui.Set(view, "controlsMessage", rui.ClickEvent, func(rui.View) {
		rui.ShowMessage("Hello", "Hello world!!!", session)
	})

	rui.Set(view, "controlsHiddenList", rui.DropDownEvent, func(_ rui.DropDownList, number int) {
		rui.Set(view, "controlsHidden", rui.Visibility, number)
	})

	timer := new(timerData)
	timer.button = rui.ButtonByID(view, "timerButton")
	timer.text = rui.TextViewByID(view, "timerText")

	timer.button.Set(rui.ClickEvent, timer.buttonClick)
	return view
}

type timerData struct {
	timerID int
	counter int
	button  rui.Button
	text    rui.TextView
}

func (timer *timerData) buttonClick(button rui.View) {
	session := button.Session()
	if timer.timerID == 0 {
		timer.timerID = session.StartTimer(1000, timer.timerFunc)
		timer.button.Set(rui.Content, "Stop timer")
	} else {
		session.StopTimer(timer.timerID)
		timer.timerID = 0
		timer.button.Set(rui.Content, "Start timer")
	}
}

func (timer *timerData) timerFunc(rui.Session) {
	timer.counter++
	timer.text.Set(rui.Text, strconv.Itoa(timer.counter))
}

var demoTime = time.Now()
