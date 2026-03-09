package predict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTLE(t *testing.T) {
	t.Run("can construct value", func(t *testing.T) {
		tle, err := NewTLE([]string{
			"ISS (ZARYA)",
			"1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927",
			"2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537",
		})
		assert.Nil(t, err)
		assert.Equal(t, "ISS (ZARYA)", tle.Name)
		assert.Equal(t, 25544, tle.Catnum)
		assert.Equal(t, 292, tle.Setnum)
		assert.InDelta(t, 325.0288, tle.Meanan, 1.0e-4)
		assert.InDelta(t, 15.7212, tle.Meanmo, 1.0e-4)
		assert.InDelta(t, -.11606e-4, tle.Bstar, 1.0e-9)
		assert.InDelta(t, -.11606, tle.bstarFrac, 1.0e-5)
		assert.Equal(t, -4, tle.bstarExp)
	})

	t.Run("can parse positive Bstar exp", func(t *testing.T) {
		tle, err := NewTLE([]string{
			"ISS (ZARYA)",
			"1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606+4 0  2927",
			"2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537",
		})
		assert.Nil(t, err)
		assert.InDelta(t, -.11606, tle.bstarFrac, 1.0e-5)
		assert.Equal(t, 4, tle.bstarExp)
	})
}
