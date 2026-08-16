package main

import (
	"fmt"
	"math/rand"
)

// enums using const
type OrderStatus int

// zero indexed grouped values
const (
	Confirmed OrderStatus = iota
	Processing
	Shipped
	Delivered
	Cancelled
	Returned
	Refunded
	Hold
)

func chageOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

// Packet Analyzer
type PacketStatus int

const (
	Accepted PacketStatus = iota
	Rejected
	Suspicious
)

func (ps PacketStatus) String() string {
	// infer length ...
	// get a string from based on iota indexed value ps
	return [...]string{"Accepted", "Rejected", "Suspicious"}[ps]
}

type Protocol string

const (
	TCP   Protocol = "TCP"
	UDP   Protocol = "UDP"
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
)

type Packet struct {
	ID       int
	Protocol Protocol
	SrcIP    string
	DestIP   string
	Size     int
}
type PacketAnalyzer struct {
	TotalPackets      int
	AcceptedPackets   int
	RejectedPackets   int
	SuspiciousPackets int
}

func (pa *PacketAnalyzer) AnalyzePacket(p Packet) PacketStatus {
	pa.TotalPackets++

	if p.Protocol == HTTP { // insecure protocol
		pa.SuspiciousPackets++
		return Suspicious
	}

	if p.Size > 1500 {
		pa.RejectedPackets++
		return Rejected
	}
	pa.AcceptedPackets++
	return Accepted
}

// random generator
func generateIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}
func generatePacket() Packet {
	protocols := []Protocol{TCP, UDP, HTTP, HTTPS}
	return Packet{
		ID:       rand.Intn(10000),
		Protocol: protocols[rand.Intn(len(protocols))],
		SrcIP:    generateIP(),
		DestIP:   generateIP(),
		Size:     rand.Intn(2000),
	}
}

// enumerated data types
func main() {
	chageOrderStatus(Returned)
	analyzer := PacketAnalyzer{}
	for i := 0; i < 100; i++ {
		packet := generatePacket()
		status := analyzer.AnalyzePacket(packet)
		// for status if we use %d, it will give iota value, if %s it uses stringer interface, and calls String()
		fmt.Printf("packet #%d: protocol: %s, src: %s, dst: %s, size: %d bytes, status: %s\n",
			packet.ID, packet.Protocol, packet.SrcIP, packet.DestIP, packet.Size, status)
	}
	fmt.Printf("\nAnalysis Summary:\n")
	fmt.Printf("total packets: %d\n", analyzer.TotalPackets)
	fmt.Printf("accepted packets: %d\n", analyzer.AcceptedPackets)
	fmt.Printf("rejected packets: %d\n", analyzer.RejectedPackets)
	fmt.Printf("suspicious packets: %d\n", analyzer.SuspiciousPackets)
}
