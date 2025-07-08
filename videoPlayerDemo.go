package main

import (
	"fmt"

	"github.com/anoshenko/rui"
)

const videoPlayerDemoText = `
GridLayout {
	cell-height = "auto, 1fr, auto, auto", width = 100%, height = 100%,
	content = [
		ListLayout {
			row = 0, orientation = start-to-end, padding = 4px,
			content = [
				Checkbox { 
					id = showVideoPlayerControls, content = "Controls", 
					checkbox-event = ShowControls,
				},
				Checkbox { 
					id = showVideoPlayerLoop, content = "Loop", checkbox-event = Loop,
				},
				Checkbox { 
					id = showVideoPlayerMuted, content = "Muted", checkbox-event = Muted,
				},
			],
		},
		GridLayout {
			row = 1, id = videoPlayerContainer,
			resize-event = ContainerResize,
			content = VideoPlayer {
				id = videoPlayer, src = testVideo.mp4, video-width = 320,
				play-event = PlayEvent, pause-event = PauseEvent,
				loaded-data-event = LoadedDataEvent,
				time-update-event = TimeUpdateEvent,
				player-error-event = PlayerError,
				volume-changed-event = VolumeChanged,
			},
		},
		ListLayout {
			row = 2, orientation = start-to-end, vertical-align = top, padding = 8px,
			content = [
				NumberPicker {
					id = videoPlayerSlider, width = 200px, type = slider
				}
				Button {
					id = videoPlayerPlay, content = "Play", margin-left = 16px, 
					click-event = PlayClick,
				}
			]
		},
		Resizable {
			row = 3, side = top, background-color = lightgrey, height = 200px,
			content = EditView {
				id = videoPlayerEventsLog, type = multiline, read-only = true, wrap = true,
			}
		},
	]
}`

type videoPlayerDemo struct {
	view  rui.View
	pause bool
}

func createVideoPlayerDemo(session rui.Session) rui.View {
	return rui.CreateViewFromText(session, videoPlayerDemoText, new(videoPlayerDemo))
}

func (demo *videoPlayerDemo) OnCreate(view rui.View) {
	demo.view = view
	demo.pause = true

	for _, event := range []rui.PropertyName{rui.AbortEvent, rui.CanPlayEvent, rui.CanPlayThroughEvent,
		rui.CompleteEvent, rui.EmptiedEvent, rui.EndedEvent, rui.LoadStartEvent,
		rui.LoadedMetadataEvent, rui.PlayingEvent, rui.SeekedEvent, rui.SeekingEvent,
		rui.StalledEvent, rui.SuspendEvent, rui.WaitingEvent} {

		rui.Set(view, "videoPlayer", event, func() {
			demo.writeToLog(string(event))
		})
	}

	for _, event := range []rui.PropertyName{rui.DurationChangedEvent, rui.RateChangedEvent} {
		rui.Set(view, "videoPlayer", event, func(value float64) {
			demo.writeToLog(fmt.Sprintf("%s: %g", string(event), value))
		})
	}

	rui.Set(view, "videoPlayerSlider", rui.NumberChangedEvent, func(value float64) {
		if demo.pause {
			rui.SetMediaPlayerCurrentTime(view, "videoPlayer", value)
		}
	})
}

func (demo *videoPlayerDemo) writeToLog(text string) {
	rui.AppendEditText(demo.view, "videoPlayerEventsLog", text+"\n")
	rui.ScrollViewToEnd(demo.view, "videoPlayerEventsLog")
}

func (demo *videoPlayerDemo) ContainerResize(frame rui.Frame) {
	rui.Set(demo.view, "videoPlayer", rui.VideoWidth, frame.Width)
	rui.Set(demo.view, "videoPlayer", rui.VideoHeight, frame.Height)
}

func (demo *videoPlayerDemo) ShowControls(state bool) {
	rui.Set(demo.view, "videoPlayer", rui.Controls, state)
}

func (demo *videoPlayerDemo) Loop(state bool) {
	rui.Set(demo.view, "videoPlayer", rui.Loop, state)
}

func (demo *videoPlayerDemo) Muted(state bool) {
	rui.Set(demo.view, "videoPlayer", rui.Muted, state)
}

func (demo *videoPlayerDemo) PlayClick() {
	if demo.pause {
		rui.MediaPlayerPlay(demo.view, "videoPlayer")
	} else {
		rui.MediaPlayerPause(demo.view, "videoPlayer")
	}
}

func (demo *videoPlayerDemo) PlayEvent() {
	demo.pause = false
	demo.writeToLog(string(rui.PlayEvent))
	rui.Set(demo.view, "videoPlayerPlay", rui.Content, "Pause")
}

func (demo *videoPlayerDemo) PauseEvent() {
	demo.pause = true
	demo.writeToLog(string(rui.PauseEvent))
	rui.Set(demo.view, "videoPlayerPlay", rui.Content, "Play")
}

func (demo *videoPlayerDemo) LoadedDataEvent() {
	demo.writeToLog("loaded-data-event")
	rui.Set(demo.view, "videoPlayerSlider", rui.Max, rui.MediaPlayerDuration(demo.view, "videoPlayer"))
}

func (demo *videoPlayerDemo) TimeUpdateEvent(time float64) {
	demo.writeToLog(fmt.Sprintf("time-update-event %gs", time))
	rui.Set(demo.view, "videoPlayerSlider", rui.Value, time)
}

func (demo *videoPlayerDemo) PlayerError(code int, message string) {
	demo.writeToLog(fmt.Sprintf(`player-error-event: code = %d, message = "%s"`, code, message))
}

func (demo *videoPlayerDemo) VolumeChanged(volume float64) {
	demo.writeToLog(fmt.Sprintf("%s: %g", string(rui.VolumeChangedEvent), volume))
	mutedFlag := rui.IsCheckboxChecked(demo.view, "showVideoPlayerMuted")
	if volume == 0 {
		if !mutedFlag {
			rui.Set(demo.view, "showVideoPlayerMuted", rui.Checked, true)
		}
	} else {
		if mutedFlag {
			rui.Set(demo.view, "showVideoPlayerMuted", rui.Checked, false)
		}
	}
}

//MAH00054.MP4
