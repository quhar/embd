package embd

func init() {
	Describers[HostBBB] = func(rev int) *Descriptor {
		return &Descriptor{
			GPIO: func() GPIO {
				return newGPIODriver(bbbPins)
			},
			I2C: newI2CDriver,
		}
	}
}

var bbbPins = PinMap{
	&PinDesc{ID: "P8_07", Aliases: []string{"66", "GPIO_66", "Caps: TIMER4"}, Caps: CapNormal | CapGPMC, DigitalLogical: 66},
	&PinDesc{ID: "P8_08", Aliases: []string{"67", "GPIO_67", "TIMER7"}, Caps: CapNormal | CapGPMC, DigitalLogical: 67},
	&PinDesc{ID: "P8_09", Aliases: []string{"69", "GPIO_69", "TIMER5"}, Caps: CapNormal | CapGPMC, DigitalLogical: 69},
	&PinDesc{ID: "P8_10", Aliases: []string{"68", "GPIO_68", "TIMER6"}, Caps: CapNormal | CapGPMC, DigitalLogical: 68},
	&PinDesc{ID: "P8_11", Aliases: []string{"45", "GPIO_45"}, Caps: CapNormal | CapGPMC, DigitalLogical: 45},
	&PinDesc{ID: "P8_12", Aliases: []string{"44", "GPIO_44"}, Caps: CapNormal | CapGPMC, DigitalLogical: 44},
	&PinDesc{ID: "P8_13", Aliases: []string{"23", "GPIO_23", "EHRPWM2B"}, Caps: CapNormal | CapGPMC, DigitalLogical: 23},
	&PinDesc{ID: "P8_14", Aliases: []string{"26", "GPIO_26"}, Caps: CapNormal | CapGPMC, DigitalLogical: 26},
	&PinDesc{ID: "P8_15", Aliases: []string{"47", "GPIO_47"}, Caps: CapNormal | CapGPMC, DigitalLogical: 47},
	&PinDesc{ID: "P8_16", Aliases: []string{"46", "GPIO_46"}, Caps: CapNormal | CapGPMC, DigitalLogical: 46},
	&PinDesc{ID: "P8_17", Aliases: []string{"27", "GPIO_27"}, Caps: CapNormal | CapGPMC, DigitalLogical: 27},
	&PinDesc{ID: "P8_18", Aliases: []string{"65", "GPIO_65"}, Caps: CapNormal | CapGPMC, DigitalLogical: 65},
	&PinDesc{ID: "P8_19", Aliases: []string{"22", "GPIO_22", "EHRPWM2A"}, Caps: CapNormal | CapGPMC, DigitalLogical: 22},
	&PinDesc{ID: "P8_26", Aliases: []string{"61", "GPIO_61"}, Caps: CapNormal | CapGPMC, DigitalLogical: 61},
	&PinDesc{ID: "P8_27", Aliases: []string{"86", "GPIO_86"}, Caps: CapNormal | CapLCD, DigitalLogical: 86},
	&PinDesc{ID: "P8_28", Aliases: []string{"88", "GPIO_88"}, Caps: CapNormal | CapLCD, DigitalLogical: 88},
	&PinDesc{ID: "P8_29", Aliases: []string{"87", "GPIO_87"}, Caps: CapNormal | CapLCD, DigitalLogical: 87},
	&PinDesc{ID: "P8_30", Aliases: []string{"89", "GPIO_89"}, Caps: CapNormal | CapLCD, DigitalLogical: 89},
	&PinDesc{ID: "P8_31", Aliases: []string{"10", "GPIO_10", "UART5_CTSN"}, Caps: CapNormal | CapLCD, DigitalLogical: 10},
	&PinDesc{ID: "P8_32", Aliases: []string{"11", "GPIO_11", "UART5_RTSN"}, Caps: CapNormal | CapLCD, DigitalLogical: 11},
	&PinDesc{ID: "P8_33", Aliases: []string{"9", "GPIO_9 ", "UART4_RTSN"}, Caps: CapNormal | CapLCD, DigitalLogical: 9},
	&PinDesc{ID: "P8_34", Aliases: []string{"81", "GPIO_81", "UART3_RTSN"}, Caps: CapNormal | CapLCD, DigitalLogical: 81},
	&PinDesc{ID: "P8_35", Aliases: []string{"8", "GPIO_8 ", "UART4_CTSN"}, Caps: CapNormal | CapLCD, DigitalLogical: 8},
	&PinDesc{ID: "P8_36", Aliases: []string{"80", "GPIO_80", "UART3_CTSN"}, Caps: CapNormal | CapLCD, DigitalLogical: 80},
	&PinDesc{ID: "P8_37", Aliases: []string{"78", "GPIO_78", "UART5_TXD"}, Caps: CapNormal | CapLCD, DigitalLogical: 78},
	&PinDesc{ID: "P8_38", Aliases: []string{"79", "GPIO_79", "UART5_RXD"}, Caps: CapNormal | CapLCD, DigitalLogical: 79},
	&PinDesc{ID: "P8_39", Aliases: []string{"76", "GPIO_76"}, Caps: CapNormal | CapLCD, DigitalLogical: 76},
	&PinDesc{ID: "P8_40", Aliases: []string{"77", "GPIO_77"}, Caps: CapNormal | CapLCD, DigitalLogical: 77},
	&PinDesc{ID: "P8_41", Aliases: []string{"74", "GPIO_74"}, Caps: CapNormal | CapLCD, DigitalLogical: 74},
	&PinDesc{ID: "P8_42", Aliases: []string{"75", "GPIO_75"}, Caps: CapNormal | CapLCD, DigitalLogical: 75},
	&PinDesc{ID: "P8_43", Aliases: []string{"72", "GPIO_72"}, Caps: CapNormal | CapLCD, DigitalLogical: 72},
	&PinDesc{ID: "P8_44", Aliases: []string{"73", "GPIO_73"}, Caps: CapNormal | CapLCD, DigitalLogical: 73},
	&PinDesc{ID: "P8_45", Aliases: []string{"70", "GPIO_70"}, Caps: CapNormal | CapLCD, DigitalLogical: 70},
	&PinDesc{ID: "P8_46", Aliases: []string{"71", "GPIO_71"}, Caps: CapNormal | CapLCD, DigitalLogical: 71},

	&PinDesc{ID: "P9_11", Aliases: []string{"30", "GPIO_30", "UART4_RXD"}, Caps: CapNormal | CapUART, DigitalLogical: 30},
	&PinDesc{ID: "P9_12", Aliases: []string{"60", "GPIO_60", "GPIO1_28"}, Caps: CapNormal, DigitalLogical: 60},
	&PinDesc{ID: "P9_13", Aliases: []string{"31", "GPIO_31", "UART4_TXD"}, Caps: CapNormal | CapUART, DigitalLogical: 31},
	&PinDesc{ID: "P9_14", Aliases: []string{"50", "GPIO_50", "EHRPWM1A"}, Caps: CapNormal | CapPWM, DigitalLogical: 50},
	&PinDesc{ID: "P9_15", Aliases: []string{"48", "GPIO_48", "GPIO1_16"}, Caps: CapNormal, DigitalLogical: 48},
	&PinDesc{ID: "P9_16", Aliases: []string{"51", "GPIO_51", "EHRPWM1B"}, Caps: CapNormal | CapPWM, DigitalLogical: 51},
	&PinDesc{ID: "P9_17", Aliases: []string{"5", "GPIO_5", "I2C1_SCL"}, Caps: CapNormal | CapI2C, DigitalLogical: 5},
	&PinDesc{ID: "P9_18", Aliases: []string{"4", "GPIO_4", "I2C1_SDA"}, Caps: CapNormal | CapI2C, DigitalLogical: 4},
	&PinDesc{ID: "P9_19", Aliases: []string{"13", "GPIO_13", "I2C2_SCL"}, Caps: CapNormal | CapI2C, DigitalLogical: 13},
	&PinDesc{ID: "P9_20", Aliases: []string{"12", "GPIO_12", "I2C2_SDA"}, Caps: CapNormal | CapI2C, DigitalLogical: 12},
	&PinDesc{ID: "P9_21", Aliases: []string{"3", "GPIO_3", "UART2_TXD"}, Caps: CapNormal | CapUART, DigitalLogical: 3},
	&PinDesc{ID: "P9_22", Aliases: []string{"2", "GPIO_2", "UART2_RXD"}, Caps: CapNormal | CapUART, DigitalLogical: 2},
	&PinDesc{ID: "P9_23", Aliases: []string{"49", "GPIO_49", "GPIO1_17"}, Caps: CapNormal, DigitalLogical: 49},
	&PinDesc{ID: "P9_24", Aliases: []string{"15", "GPIO_15", "UART1_TXD"}, Caps: CapNormal | CapUART, DigitalLogical: 15},
	&PinDesc{ID: "P9_25", Aliases: []string{"117", "GPIO_117", "GPIO3_21"}, Caps: CapNormal, DigitalLogical: 117},
	&PinDesc{ID: "P9_26", Aliases: []string{"14", "GPIO_14", "UART1_RXD"}, Caps: CapNormal | CapUART, DigitalLogical: 14},
	&PinDesc{ID: "P9_27", Aliases: []string{"115", "GPIO_115", "GPIO3_19"}, Caps: CapNormal, DigitalLogical: 115},
	&PinDesc{ID: "P9_28", Aliases: []string{"113", "GPIO_113", "SPI1_CS0"}, Caps: CapNormal | CapSPI, DigitalLogical: 113},
	&PinDesc{ID: "P9_29", Aliases: []string{"111", "GPIO_111", "SPI1_D0"}, Caps: CapNormal | CapSPI, DigitalLogical: 111},
	&PinDesc{ID: "P9_30", Aliases: []string{"112", "GPIO_112", "SPI1_D1"}, Caps: CapNormal | CapSPI, DigitalLogical: 112},
	&PinDesc{ID: "P9_31", Aliases: []string{"110", "GPIO_110", "SPI1_SCLK"}, Caps: CapNormal | CapSPI, DigitalLogical: 110},
	&PinDesc{ID: "P9_32", Aliases: []string{"VADC"}},
	&PinDesc{ID: "P9_33", Aliases: []string{"AIN4"}, Caps: CapAnalog, AnalogLogical: 4},
	&PinDesc{ID: "P9_34", Aliases: []string{"AGND"}},
	&PinDesc{ID: "P9_35", Aliases: []string{"AIN6"}, Caps: CapAnalog, AnalogLogical: 6},
	&PinDesc{ID: "P9_36", Aliases: []string{"AIN5"}, Caps: CapAnalog, AnalogLogical: 5},
	&PinDesc{ID: "P9_37", Aliases: []string{"AIN2"}, Caps: CapAnalog, AnalogLogical: 2},
	&PinDesc{ID: "P9_38", Aliases: []string{"AIN3"}, Caps: CapAnalog, AnalogLogical: 3},
	&PinDesc{ID: "P9_39", Aliases: []string{"AIN0"}, Caps: CapAnalog, AnalogLogical: 0},
	&PinDesc{ID: "P9_40", Aliases: []string{"AIN1"}, Caps: CapAnalog, AnalogLogical: 1},
}
