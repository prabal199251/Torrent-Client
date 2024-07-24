package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatRequest(t *testing.T) {
	msg := FormatRequest(4, 567, 4321)
	expected := &Message{
		ID: MsgRequest,
		PayLoad: []byte{
			0x00, 0x00, 0x00, 0x04, // Index
			0x00, 0x00, 0x02, 0x37, // Begin
			0x00, 0x00, 0x10, 0xe1, // Length
		},
	}
	assert.Equal(t, expected, msg)
}

func TestFormatHave(t *testing.T) {
	msg := FormatHave(4)
	expected := &Message{
		ID:      MsgHave,
		PayLoad: []byte{0x00, 0x00, 0x00, 0x04},
	}
	assert.Equal(t, expected, msg)
}
