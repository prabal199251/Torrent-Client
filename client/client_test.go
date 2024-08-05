package client

import (
	"net"
	"testing"

	"github.com/prabal199251/Torrent-Client/bitfield"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createClientAndServer(t *testing.T) (clientConn, serverConn net.Conn) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.Nil(t, err)

	done := make(chan struct{})

	go func() {
		defer ln.Close()
		serverConn, err = ln.Accept()
		require.Nil(t, err)
		done <- struct{}{}
	}()

	clientConn, err = net.Dial("tcp", ln.Addr().String())
	<-done

	return clientConn, serverConn
}

func TestClientServerConnection(t *testing.T) {
	clientConn, serverConn := createClientAndServer(t)
	defer clientConn.Close()
	defer serverConn.Close()

	testMessage := []byte("Hello, Server!")
	_, err := clientConn.Write(testMessage)
	require.Nil(t, err)

	buf := make([]byte, len(testMessage))
	_, err = serverConn.Read(buf)
	require.Nil(t, err)

	require.Equal(t, testMessage, buf)
}

func TestRecvBitfield(t *testing.T) {
	tests := map[string]struct {
		msg    []byte
		output bitfield.Bitfield
		fails  bool
	}{
		"successful bitfield" : {
			msg : []byte{0x00, 0x00, 0x00, 0x06, 5, 1, 2, 3, 4, 5},
			output: bitfield.Bitfield{1, 2, 3, 4, 5},
			fails: false,
		},

		"message is not a bitfield": {
			msg:    []byte{0x00, 0x00, 0x00, 0x06, 99, 1, 2, 3, 4, 5},
			output: nil,
			fails:  true,
		},
		"message is keep-alive": {
			msg:    []byte{0x00, 0x00, 0x00, 0x00},
			output: nil,
			fails:  true,
		},
	}

	for _, test := range tests {
		clientConn, serverConn := createClientAndServer(t)
		serverConn.Write(test.msg)

		bf, err := recvBitfiled(clientConn)

		if test.fails {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, bf, test.output)
		}
	}
}
