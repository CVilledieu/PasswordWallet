package util

import "encoding/binary"

const (
	//HEADER is 2 bytes for type and 2 bytes for the number of keys
	HEADER         = 4
	MAX_KEY_SIZE   = 1000
	MAX_VALUE_SIZE = 3000
	MAX_NODE_SIZE  = HEADER + 8 + 2 + 4 + MAX_KEY_SIZE + MAX_VALUE_SIZE
)

type Node []byte

func newNode() Node {
	return make(Node, MAX_NODE_SIZE)
}

func (n Node) nkeys() uint16 {
	return binary.LittleEndian.Uint16(n[2:4])
}

func (n Node) nType() uint16 {
	return binary.LittleEndian.Uint16(n[0:2])
}

func (n Node) setHeader(nType, nKey uint16) {
	binary.LittleEndian.PutUint16(n[0:2], nType)
	binary.LittleEndian.PutUint16(n[2:4], nKey)
}
