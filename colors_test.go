package standard

import (
	"testing"

	"github.com/dscottboggs/attest"
)

func TestPercent2Color(t *testing.T) {
	t.Run("0% -- totally green", func(t *testing.T) {
		test := attest.New(t)
		col := Percent2Color(0)
		test.Equals(0, int(col.Red))
		test.Equals(255, int(col.Green))
		test.Equals(0, int(col.Blue))
	})
	t.Run("25% -- full green, half red", func(t *testing.T) {
		test := attest.New(t)
		col := Percent2Color(25)
		test.Equals(127, int(col.Red))
		test.Equals(255, int(col.Green))
		test.Equals(0, int(col.Blue))
	})
	t.Run("50% -- full gren and red", func(t *testing.T) {
		test := attest.New(t)
		col := Percent2Color(50)
		test.Equals(255, int(col.Red))
		test.Equals(255, int(col.Green))
		test.Equals(0, int(col.Blue))
	})
	t.Run("75% -- full red, half green", func(t *testing.T) {
		test := attest.New(t)
		col := Percent2Color(75)
		test.Equals(255, int(col.Red))
		test.Equals(127, int(col.Green))
		test.Equals(0, int(col.Blue))
	})
	t.Run("100% -- full red", func(t *testing.T) {
		test := attest.New(t)
		col := Percent2Color(100)
		test.Equals(255, int(col.Red))
		test.Equals(0, int(col.Green))
		test.Equals(0, int(col.Blue))
	})
}

func TestToHexString_Color(t *testing.T) {
	test := attest.New(t)
	color := Color{0x50, 0x9F, 0xFF}
	test.Equals("#509FFF", color.ToHexString())
}
