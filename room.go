package state

import "net"

// Room represents a room holding
// certaint state, and a set of
// clients to stream state to
type Room struct {
	state   *State
	clients map[string]*net.UDPConn
	// TODO
}
