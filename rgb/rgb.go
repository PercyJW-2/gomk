package rgb

import (
	"gomk/matrix"
	"image/color"
	"tinygo.org/x/drivers/ws2812"
)

const LED_COUNT = ((matrix.COLS_LEFT + matrix.COLS_RIGHT) * matrix.ROWS_LEFT) - 4
const LED_ACTIVE = true

type LedMap struct {
	left  [matrix.ROWS_LEFT][matrix.COLS_LEFT]uint
	right [matrix.ROWS_RIGHT][matrix.COLS_RIGHT]uint
}

var KEY_TO_LEDNUMBER_MAP = LedMap{
	left: [matrix.ROWS_LEFT][matrix.COLS_LEFT]uint{
		{5, 4, 3, 2, 1, 0},
		{6, 7, 8, 9, 10, 11},
		{17, 16, 15, 14, 13, 12},
		{0, 18, 19, 20, 21, 0},
	},
	right: [matrix.ROWS_RIGHT][matrix.COLS_RIGHT]uint{
		{22, 23, 24, 25, 26, 27},
		{33, 32, 31, 30, 29, 28},
		{34, 35, 36, 37, 38, 39},
		{43, 42, 41, 40, 0, 0},
	},
}

var LED_OBJECT ws2812.Device
var LED_VALUES [LED_COUNT]color.RGBA
var LED_UPDATED = false

func SetSingleColor(color color.RGBA) {
	for i := range LED_VALUES {
		LED_VALUES[i] = color
	}
	LED_UPDATED = true
}

func setKeyColor(color color.RGBA, left bool, col int, row int) {
	var ledIndex uint
	if left {
		ledIndex = KEY_TO_LEDNUMBER_MAP.left[row][col]
	} else {
		ledIndex = KEY_TO_LEDNUMBER_MAP.right[row][col]
	}
	LED_VALUES[ledIndex] = color
	LED_UPDATED = true
}

func UpdateLeds() {
	if LED_UPDATED {
		err := LED_OBJECT.WriteColors(LED_VALUES[:])
		if err != nil {
			println("could not update LEDs: ", err)
		}
		LED_UPDATED = false
	}
}
