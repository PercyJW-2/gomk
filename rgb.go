package main

import "image/color"

const LED_COUNT = ((COLS_LEFT + COLS_RIGHT) * ROWS_LEFT) - 4

var LED_VALUES [LED_COUNT]color.RGBA

func set_single_color(color color.RGBA) {
	for i := range LED_VALUES {
		LED_VALUES[i] = color
	}
}

func set_key_color(color color.RGBA, left bool, col int, row int) {
	if left {
		LED_VALUES[]
	} else {

	}
}
