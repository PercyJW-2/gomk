package matrix

import (
	"machine/usb/hid/keyboard"
)

type Code uint8

const (
	ENCODER Code = iota
	LOWER
	UPPER
	NONE
	PASSTHROUGH
)

const LAYERCOUNT = 3

type KeymapLayer struct {
	left  [COLS_LEFT * ROWS_LEFT]interface{}
	right [COLS_RIGHT * ROWS_RIGHT]interface{}
}

type Keymap struct {
	layers [LAYERCOUNT]KeymapLayer
}

var CURRENT_LAYER = 0

var KEYMAP = Keymap{
	layers: [LAYERCOUNT]KeymapLayer{
		{
			left: [COLS_LEFT * ROWS_LEFT]interface{}{
				keyboard.KeyTab, keyboard.KeyQ, keyboard.KeyW, keyboard.KeyE, keyboard.KeyR, keyboard.KeyT,
				keyboard.KeyLeftCtrl, keyboard.KeyA, keyboard.KeyS, keyboard.KeyD, keyboard.KeyF, keyboard.KeyG,
				keyboard.KeyLeftShift, keyboard.KeyY, keyboard.KeyX, keyboard.KeyC, keyboard.KeyV, keyboard.KeyB,
				NONE, keyboard.KeyLeftAlt, keyboard.KeyLeftGUI, LOWER, keyboard.KeySpace, ENCODER,
			},
			right: [COLS_RIGHT * ROWS_RIGHT]interface{}{
				keyboard.KeyZ, keyboard.KeyU, keyboard.KeyI, keyboard.KeyO, keyboard.KeyP, keyboard.KeySlash,
				keyboard.KeyH, keyboard.KeyJ, keyboard.KeyK, keyboard.KeyL, keyboard.KeySemicolon, keyboard.KeyQuote,
				keyboard.KeyN, keyboard.KeyM, keyboard.KeyComma, keyboard.KeyPeriod, keyboard.KeySlash, keyboard.KeyRightShift,
				keyboard.KeyEnter, UPPER, keyboard.KeyBackspace, keyboard.KeyDelete, NONE, NONE,
			},
		},
		{
			left: [COLS_LEFT * ROWS_LEFT]interface{}{
				keyboard.KeyEsc, keyboard.Key1, keyboard.Key2, keyboard.Key3, keyboard.Key4, keyboard.Key5,
				PASSTHROUGH, PASSTHROUGH, keyboard.KeyUpArrow, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
				PASSTHROUGH, keyboard.KeyLeftArrow, keyboard.KeyDownArrow, keyboard.KeyRightArrow, PASSTHROUGH, PASSTHROUGH,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
			},
			right: [COLS_RIGHT * ROWS_RIGHT]interface{}{
				keyboard.Key6, keyboard.Key7, keyboard.Key8, keyboard.Key9, keyboard.Key0, keyboard.KeyTilde,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
			},
		},
		{
			left: [COLS_LEFT * ROWS_LEFT]interface{}{
				keyboard.KeyF1, keyboard.KeyF2, keyboard.KeyF3, keyboard.KeyF4, keyboard.KeyF5, keyboard.KeyF6,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
			},
			right: [COLS_RIGHT * ROWS_RIGHT]interface{}{
				keyboard.KeyF7, keyboard.KeyF8, keyboard.KeyF9, keyboard.KeyF10, keyboard.KeyF11, keyboard.KeyF12,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
				PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH, PASSTHROUGH,
			},
		},
	},
}
