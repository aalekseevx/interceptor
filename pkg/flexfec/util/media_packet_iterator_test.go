package util

import (
	"testing"

	"github.com/pion/rtp"
	"github.com/stretchr/testify/assert"
)

func TestMediaPacketIterator_Simple(t *testing.T) {
	packets := []rtp.Packet{
		{Header: rtp.Header{SequenceNumber: 1000}},
		{Header: rtp.Header{SequenceNumber: 1001}},
		{Header: rtp.Header{SequenceNumber: 1002}},
		{Header: rtp.Header{SequenceNumber: 1003}},
		{Header: rtp.Header{SequenceNumber: 1004}},
		{Header: rtp.Header{SequenceNumber: 1005}},
	}

	evenIndices := []uint32{0, 2, 4}
	oddIndices := []uint32{1, 3, 5}

	evenIterator := NewMediaPacketIterator(packets, evenIndices)
	oddIterator := NewMediaPacketIterator(packets, oddIndices)

	assert.Equal(t, 1000, evenIterator.First().SequenceNumber)
	assert.Equal(t, 1001, oddIterator.First().SequenceNumber)
}
