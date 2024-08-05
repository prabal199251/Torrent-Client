package client

import (
	"net"
	"testing"

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

