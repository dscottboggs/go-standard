package standard

import (
	"fmt"
)

// Color -- An 8-bit RGB color value
type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

// Percent2Color -- Convert a percent value to a Color on the gradient between
// green and red
func Percent2Color(percent int8) *Color {
	var r, g, b uint8 // b defaults to zero, don't need to set.
	red := uint16(float64(float64(percent)*2) / 100 * 255)
	green := uint16(float64(200-float64(percent)*2) / 100 * 255)
	if green > 255 {
		g = uint8(255)
	} else {
		g = uint8(green)
	}
	if red > 255 {
		r = uint8(255)
	} else {
		r = uint8(red)
	}
	c := Color{r, g, b}
	return &c
}

// ToHexString -- convert a color to a hex string: #XXXXXX
func (c *Color) ToHexString() string {
	return fmt.Sprintf("#%02X%02X%02X", c.Red, c.Green, c.Blue)
}
