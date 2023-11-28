package main

import (
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
	oledDisplay := ssd1306.NewI2C(i2c)
	oledDisplay.Configure(ssd1306.Config{
		Width:    16,
		Height:   128,
		VccState: 0,
		Address:  0,
	})
	// Setup Keyboard Matrices
	MATRIX_RIGHT, err = generateRightMatrix(expander)
	if err != nil {
		println("could not configure right matrix: ", err)
		return
	}
	setupLeftMatrix()
	// Setup neopixels
	leds := ws2812.New(machine.P0_03)
	err = leds.WriteColors([]color.RGBA{{A: 0, G: 0, B: 0, R: 0}})
	if err != nil {
		println("could not update LEDs: ", err)
		return
	}
}

func main() {
	for {
		checkMatrixLeft(MATRIX_LEFT)
		checkMatrixRight(MATRIX_RIGHT)
	}
}
