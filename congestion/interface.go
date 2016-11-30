package congestion

import (
	"time"

	"github.com/lucas-clemente/quic-go/protocol"
)

// A SendAlgorithm performs congestion control and calculates the congestion window
type SendAlgorithm interface {
	TimeUntilSend(now time.Time, bytesInFlight protocol.ByteCount) time.Duration
	OnPacketSent(sentTime time.Time, bytesInFlight protocol.ByteCount, packetNumber protocol.PacketNumber, bytes protocol.ByteCount, isRetransmittable bool) bool
	GetCongestionWindow() protocol.ByteCount
	OnCongestionEvent(rttUpdated bool, priorInFlight protocol.ByteCount, bytesInFlight protocol.ByteCount, eventTime time.Time, ackedPackets PacketVector, lostPackets PacketVector, leastUnacked protocol.PacketNumber)
	SetNumEmulatedConnections(n int)
	OnRetransmissionTimeout(packetsRetransmitted bool)
	OnConnectionMigration()
	RetransmissionDelay() time.Duration

	// Experiments
	SetSlowStartLargeReduction(enabled bool)
}

// SendAlgorithmWithDebugInfo adds some debug functions to SendAlgorithm
type SendAlgorithmWithDebugInfo interface {
	SendAlgorithm
	BandwidthEstimate() protocol.Bandwidth

	// Stuff only used in testing

	HybridSlowStart() *HybridSlowStart
	SlowstartThreshold() protocol.PacketNumber
	RenoBeta() float32
	InRecovery() bool
}
