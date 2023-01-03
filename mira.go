package miraapi

import (
	"github.com/bearsh/hid"
	"golang.org/x/exp/constraints"
)

const (
	// Boox Mira USB vendor id
	Boox_Mira_VID uint16 = 0x0416

	// Boox Mira USB product id
	Boox_Mira_PID uint16 = 0x5020
)

// Command encodes wire device command codes
type Command byte

const (
	CmdRefresh        Command = 0x01
	CmdSetRefreshMode Command = 0x02
	CmdSetSpeed       Command = 0x04
	CmdSetContrast    Command = 0x05
	CmdSetColdLight   Command = 0x06
	CmdSetWarmLight   Command = 0x07
	CmdSetDitherMode  Command = 0x09
	CmdSetColorFilter Command = 0x11
)

// RefreshMode encodes wire device refresh modes
type RefreshMode byte

const (
	// RefreshDirect: black/white, fast
	RefreshDirect RefreshMode = 0x01

	// RefreshGray: gray scale, slow
	RefreshGray RefreshMode = 0x02

	// RefreshA2: fast
	RefreshA2 RefreshMode = 0x03
)

func SendCommand(device *hid.Device, command Command, data ...byte) error {
	_, err := device.SendFeatureReport(append([]byte{0, byte(command)}, data...))
	return err
}

func constraint[T constraints.Integer](value, min, max T) T {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

// Refresh refreshes the display
func Refresh(device *hid.Device) error {
	return SendCommand(device, CmdRefresh)
}

const (
	MinSpeed = 1
	MaxSpeed = 7
)

// SetSpeed sets a new speed value (1..7)
func SetSpeed(device *hid.Device, speed int) error {
	return SendCommand(device, CmdSetSpeed, byte(11-constraint(speed, MinSpeed, MaxSpeed)))
}

const (
	MinContrast = 0
	MaxContrast = 15
)

// SetContrast sets a new contrast value (0..15)
func SetContrast(device *hid.Device, contrast int) error {
	return SendCommand(device, CmdSetSpeed, byte(constraint(contrast, MinContrast, MaxContrast)))
}

// SetRefreshMode sets a new refresh mode
func SetRefreshMode(device *hid.Device, refreshMode RefreshMode) error {
	return SendCommand(device, CmdSetRefreshMode, byte(refreshMode))
}

const (
	MinDitherMode = 0
	MaxDitherMode = 3
)

// SetDitherMode sets a new dither mode (0..3)
func SetDitherMode(device *hid.Device, ditherMode int) error {
	return SendCommand(device, CmdSetDitherMode, byte(constraint(ditherMode, MinDitherMode, MaxDitherMode)))
}

// SetColorFilter sets new white + black filters (0..254)
func SetColorFilter(device *hid.Device, whiteFilter, blackFilter int) error {
	return SendCommand(device, CmdSetColorFilter,
		255-byte(constraint(whiteFilter, 0, 254)),
		byte(constraint(blackFilter, 0, 254)))
}

// SetColdLight sets a new cold light brightness (0..254)
func SetColdLight(device *hid.Device, brightness int) error {
	return SendCommand(device, CmdSetColdLight, byte(constraint(brightness, 0, 254)))
}

// SetWarmLight sets a new warm light brightness (0..254)
func SetWarmLight(device *hid.Device, brightness int) error {
	return SendCommand(device, CmdSetWarmLight, byte(constraint(brightness, 0, 254)))
}

var (
	ColdLab = []float64{95.34, -3.80, -22.55}
	WarmLab = []float64{92.46, 7.69, 51.66}
)
