// Tempconv packet produce temperature conversion functions between ˚F and ˚C.
package ex_2_1_tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
	AbsoluteZeroK Kelvin     = 0
	FreezingK     Kelvin     = 273.15
	BoilingK      Kelvin     = 373.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g˚C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g˚F", f)
}

func (f Kelvin) String() string {
	return fmt.Sprintf("%gK", f)
}
