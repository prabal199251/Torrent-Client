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

func TestParsePiece(t *testing.T) {
	tests := map[string]struct {
		inputIndex int
		inputBuf   []byte
		inputMsg   *Message
		outputN    int
		outputBuf  []byte
		fails      bool
	}{
		"parse valid piece": {
			inputIndex: 4,
			inputBuf:   make([]byte, 10),
			inputMsg: &Message{
				ID: MsgPiece,
				PayLoad: []byte{
					0x00, 0x00, 0x00, 0x04, // Index
					0x00, 0x00, 0x00, 0x02, // Begin
					0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, // Block
				},
			},
			outputBuf: []byte{0x00, 0x00, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, 0x00},
			outputN:   6,
			fails:     false,
		},

		"wrong message type": {
			inputIndex: 4,
			inputBuf:   make([]byte, 10),
			inputMsg: &Message{
				ID:      MsgChoke,
				PayLoad: []byte{},
			},
			outputBuf: make([]byte, 10),
			outputN:   0,
			fails:     true,
		},

		"payload too short": {
			inputIndex: 4,
			inputBuf:   make([]byte, 10),
			inputMsg: &Message{
				ID: MsgPiece,
				PayLoad: []byte{
					0x00, 0x00, 0x00, 0x04, // Index
					0x00, 0x00, 0x00, // Malformed offset
				},
			},
			outputBuf: make([]byte, 10),
			outputN:   0,
			fails:     true,
		},

		"wrong index": {
			inputIndex: 4,
			inputBuf:   make([]byte, 10),
			inputMsg: &Message{
				ID: MsgPiece,
				PayLoad: []byte{
					0x00, 0x00, 0x00, 0x06, // Index is 6, not 4
					0x00, 0x00, 0x00, 0x02, // Begin
					0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, // Block
				},
			},
			outputBuf: make([]byte, 10),
			outputN:   0,
			fails:     true,
		},

		"offset too high": {
			inputIndex: 4,
			inputBuf:   make([]byte, 10),
			inputMsg: &Message{
				ID: MsgPiece,
				PayLoad: []byte{
					0x00, 0x00, 0x00, 0x04, // Index is 4
					0x00, 0x00, 0x00, 0x0c, // Begin is 12 > 10
					0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, // Block
				},
			},
			outputBuf: make([]byte, 10),
			outputN:   0,
			fails:     true,
		},

		"expected offset but payload too long": {
			inputIndex: 4,
			inputBuf:   make([]byte, 10),
			inputMsg: &Message{
				ID: MsgPiece,
				PayLoad: []byte{
					0x00, 0x00, 0x00, 0x04, // Index is 4
					0x00, 0x00, 0x00, 0x02, // Begin is ok
					// Block is 10 long but begin=2; too long for input buffer
					0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x0a, 0x0b, 0x0c, 0x0d,
				},
			},
			outputBuf: make([]byte, 10),
			outputN:   0,
			fails:     true,
		},
	}

	for _, test := range tests {
		n ,err := ParsePiece(test.inputIndex, test.inputBuf, test.inputMsg)

		if test.fails {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, test.outputBuf, test.inputBuf)
		assert.Equal(t, test.outputN, n)
	}	
}
