package main

import (
	"gomk/matrix"
	"gomk/oled"
	"gomk/rgb"

	"image/color"
	"machine"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/drivers/ws2812"

	"tinygo.org/x/drivers/mcp23017"
)

func init() {
	i2c := machine.I2C0
	err := i2c.Configure(machine.I2CConfig{
		SCL: machine.P0_29,
		SDA: machine.P1_13,
	})
	if err != nil {
		println("could not configure I2C: ", err)
		return
	}
	// Setup I2C Devices
	expander, err := mcp23017.NewI2C(i2c, 0)
	if err != nil {
		println("could not configure ioexpander:", err)
		return
	}
	oled.InitOLED(ssd1306.NewI2C(i2c))

	// Setup Keyboard Matrices
	matrix.MATRIX_RIGHT, err = matrix.GenerateRightMatrix(expander)
	if err != nil {
		println("could not configure right matrix: ", err)
		return
	}
	matrix.SetupLeftMatrix()
	// Setup neopixels
	neopixel_led := machine.P1_12
	neopixel_led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	rgb.LED_OBJECT = ws2812.New(neopixel_led)
	rgb.SetSingleColor(color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 0,
	})
	if err != nil {
		println("could not update LEDs: ", err)
		return
	}
}

func main() {
	for {
		matrix.CheckMatrixLeft(matrix.MATRIX_LEFT)
		matrix.CheckMatrixRight(matrix.MATRIX_RIGHT)
		if rgb.LED_ACTIVE {
			rgb.UpdateLeds()
		}
	}
}
