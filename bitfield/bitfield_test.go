package bitfield

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPiece(t *testing.T) {
	bf := Bitfield{0b01010100, 0b01010100}
	outputs := []bool{false, true, false, true, false, true, false, false, false, true, false, true, false, true, false, false, false, false, false, false}

	for i:=0; i<len(outputs); i++ {
		assert.Equal(t, outputs[i], bf.HasPiece(i))
	}
}
