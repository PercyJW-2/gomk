package oled

import (
	font "github.com/Nondzu/ssd1306_font"
	"tinygo.org/x/drivers/ssd1306"
)

var OLED_OBJECT ssd1306.Device
var OLED_DISPLAY font.Display

func InitOLED(device ssd1306.Device) {
	OLED_OBJECT = device
	OLED_OBJECT.Configure(ssd1306.Config{
		Width:    16,
		Height:   128,
		VccState: 0,
		Address:  0,
	})
	OLED_OBJECT.ClearBuffer()
	OLED_OBJECT.ClearDisplay()

	OLED_DISPLAY = font.NewDisplay(OLED_OBJECT)
	OLED_DISPLAY.Configure(font.Config{FontType: font.FONT_6x8})
	OLED_DISPLAY.XPos = 0
	OLED_DISPLAY.YPos = 0
	OLED_DISPLAY.PrintText("HELLO WORLD!")
}

func writeString(str string) {
	//var buffer = OLED_OBJECT.GetBuffer()

}
