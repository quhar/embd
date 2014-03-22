package embd

import (
	"fmt"
	"strconv"
)

const (
	CapNormal int = 1 << iota
	CapI2C
	CapUART
	CapSPI
	CapGPMC
	CapLCD
	CapPWM
	CapAnalog
)

type PinDesc struct {
	ID      string
	Aliases []string
	Caps    int

	DigitalLogical int
	AnalogLogical  int
}

type PinMap []*PinDesc

func (m PinMap) Lookup(k interface{}, cap int) (*PinDesc, bool) {
	var ks string
	switch key := k.(type) {
	case int:
		ks = strconv.Itoa(key)
	case string:
		ks = key
	case fmt.Stringer:
		ks = key.String()
	default:
		return nil, false
	}

	for i := range m {
		pd := m[i]

		if pd.ID == ks {
			return pd, true
		}

		for j := range pd.Aliases {
			if pd.Aliases[j] == ks && pd.Caps&cap != 0 {
				return pd, true
			}
		}
	}

	return nil, false
}
