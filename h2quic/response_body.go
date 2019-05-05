package h2quic

import (
	"io"

	quic "github.com/attenberger/quic-go"
)

type responseBody struct {
	quic.Stream
}

var _ io.ReadCloser = &responseBody{}

func (rb *responseBody) Close() error {
	rb.Stream.CancelRead(0)
	return nil
}
