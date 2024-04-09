package matrix

import (
	"machine"
	"machine/usb/hid/keyboard"
	"tinygo.org/x/drivers/mcp23017"
)

const (
	COLS_LEFT  = 6
	ROWS_LEFT  = 4
	COLS_RIGHT = 6
	ROWS_RIGHT = 4
)

type MatrixLeft struct {
	outputs [COLS_LEFT]machine.Pin
	inputs  [ROWS_LEFT]machine.Pin
}

type MatrixRight struct {
	outputs [COLS_RIGHT]mcp23017.Pin
	inputs  [ROWS_RIGHT]mcp23017.Pin
}

type MatrixRightIntermediate struct {
	outputs [COLS_RIGHT]int
	inputs  [ROWS_RIGHT]int
}

var MATRIX_LEFT = MatrixLeft{
	outputs: [COLS_LEFT]machine.Pin{
		machine.P0_13,
		machine.P0_14,
		machine.P0_15,
		machine.P0_16,
		machine.P0_24,
		machine.P0_25,
	},
	inputs: [ROWS_LEFT]machine.Pin{
		machine.P0_03,
		machine.P0_28,
		machine.P0_02,
		machine.P0_31,
	},
}

var MATRIX_RIGHT_INTERMEDIATE = MatrixRightIntermediate{
	outputs: [COLS_RIGHT]int{
		0, 1, 2, 3, 4, 5,
	},
	inputs: [ROWS_RIGHT]int{
		8, 9, 10, 11,
	},
}

var MATRIX_RIGHT MatrixRight

var MATRIX_LEFT_STATE [COLS_LEFT][ROWS_LEFT]bool
var MATRIX_RIGHT_STATE [COLS_RIGHT][COLS_RIGHT]bool

func GenerateRightMatrix(ioExpander *mcp23017.Device) (MatrixRight, error) {
	outputs := [COLS_RIGHT]mcp23017.Pin{}
	for i := range outputs {
		outputs[i] = ioExpander.Pin(MATRIX_RIGHT_INTERMEDIATE.outputs[i])
		err := outputs[i].SetMode(mcp23017.Output)
		if err != nil {
			return MatrixRight{}, err
		}
	}
	inputs := [ROWS_RIGHT]mcp23017.Pin{}
	for i := range inputs {
		inputs[i] = ioExpander.Pin(MATRIX_RIGHT_INTERMEDIATE.inputs[i])
		err := inputs[i].SetMode(mcp23017.Input)
		if err != nil {
			return MatrixRight{}, err
		}
	}
	return MatrixRight{
		outputs: outputs,
		inputs:  inputs,
	}, nil
}

func SetupLeftMatrix() {
	for _, output := range MATRIX_LEFT.outputs {
		output.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	for _, input := range MATRIX_LEFT.inputs {
		input.Configure(machine.PinConfig{Mode: machine.PinInput})
	}
}

func CheckMatrixLeft(left MatrixLeft) {
	for i, output := range left.outputs {
		output.High()
		for j, input := range left.inputs {
			keyState := input.Get()
			if keyState != MATRIX_LEFT_STATE[i][j] {
				handleLeft(i, j, keyState)
				MATRIX_LEFT_STATE[i][j] = keyState
			}
		}
		output.Low()
	}
}

func CheckMatrixRight(right MatrixRight) {
	for i, output := range right.outputs {
		err := output.High()
		if err != nil {
			println("Failed to set Pin high:", err)
			return
		}
		for j, input := range right.inputs {
			level, err := input.Get()
			if err != nil {
				println("Failed to read Pin:", err)
				return
			}
			if level != MATRIX_RIGHT_STATE[i][j] {
				handleRight(i, j, level)
				MATRIX_RIGHT_STATE[i][j] = level
			}
		}
		err = output.Low()
		if err != nil {
			println("Failed to set Pin low:", err)
			return
		}
	}
}

func handleLeft(col int, row int, keyState bool) {
	var keycode interface{} = NONE
	layerOffset := 0
	for keycode == PASSTHROUGH && CURRENT_LAYER-layerOffset >= 0 {
		keycode = KEYMAP.layers[CURRENT_LAYER-layerOffset].left[col+row*ROWS_LEFT]
		layerOffset++
	}
	handle(keyState, keycode)
}

func handleRight(col int, row int, keyState bool) {
	var keycode interface{} = NONE
	layerOffset := 0
	for keycode == PASSTHROUGH && CURRENT_LAYER-layerOffset >= 0 {
		keycode = KEYMAP.layers[CURRENT_LAYER-layerOffset].right[col+row*ROWS_RIGHT]
		layerOffset++
	}
	handle(keyState, keycode)
}

func handle(keyState bool, keycode interface{}) {
	kb := keyboard.Port()
	switch keycode {
	case ENCODER:
		break
	case LOWER:
		CURRENT_LAYER = 1
		break
	case UPPER:
		CURRENT_LAYER = 2
		break
	case NONE:
	case PASSTHROUGH:
		break
	default:
		var err error
		if keyState {
			err = kb.Down(keycode.(keyboard.Keycode))
		} else {
			err = kb.Up(keycode.(keyboard.Keycode))
		}
		if err != nil {
			println("Failed to send Keycode: ", err)
			return
		}
	}
}
