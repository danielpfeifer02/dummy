package wire

import (
	"github.com/quic-go-no-crypto/quic-go-no-crypto/internal/protocol"
)

// A Frame in QUIC
type Frame interface {
	Append(b []byte, version protocol.Version) ([]byte, error)
	Length(version protocol.Version) protocol.ByteCount
}
