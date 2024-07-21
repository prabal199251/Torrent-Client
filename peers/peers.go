package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

func Unmarshal(peersBin []byte) ([]Peer, error) {
	const PeerSize = 6 //4 for IP & 2 for Port

	numPeers := len(peersBin) / PeerSize

	if len(peersBin)%PeerSize != 0 {
		err := fmt.Errorf("received malformed peers")
		return nil, err
	}

	peers := make([]Peer, numPeers)

	for i := 0; i < numPeers; i++ {
		offset := i * PeerSize
		peers[i].IP = net.IP(peersBin[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16([]byte(peersBin[offset+4 : offset+6]))
	}

	return peers, nil
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
