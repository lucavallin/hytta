package io

import (
	"github.com/LuCavallin/yfirbord-grovepi/pkg/connections"
)

// Writer knows how to write from to the connection
type Writer struct {
	conn connections.ReadConnection
}

// NewWriter creates a new reader
func NewWriter(conn connections.ReadConnection) *Writer {
	return &Writer{
		conn: conn,
	}
}
