// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"encoding/binary"
	"github.com/google/gopacket"
	"net"
	"strconv"
)

var (
	// We use two different endpoint types for IPv4 vs IPv6 addresses, so that
	// ordering with endpointA.LessThan(endpointB) sanely groups all IPv4
	// addresses and all IPv6 addresses, such that IPv6 > IPv4 for all addresses.
	EndpointIPv4 = gopacket.RegisterEndpointType(1, gopacket.EndpointTypeMetadata{"IPv4", func(b []byte) string {
		return net.IP(b).String()
	}})
	EndpointIPv6 = gopacket.RegisterEndpointType(2, gopacket.EndpointTypeMetadata{"IPv6", func(b []byte) string {
		return net.IP(b).String()
	}})

	EndpointMAC = gopacket.RegisterEndpointType(3, gopacket.EndpointTypeMetadata{"MAC", func(b []byte) string {
		return net.HardwareAddr(b).String()
	}})
	EndpointTCPPort = gopacket.RegisterEndpointType(4, gopacket.EndpointTypeMetadata{"TCP", func(b []byte) string {
		return strconv.Itoa(int(binary.BigEndian.Uint16(b)))
	}})
	EndpointUDPPort = gopacket.RegisterEndpointType(5, gopacket.EndpointTypeMetadata{"UDP", func(b []byte) string {
		return strconv.Itoa(int(binary.BigEndian.Uint16(b)))
	}})
	EndpointSCTPPort = gopacket.RegisterEndpointType(6, gopacket.EndpointTypeMetadata{"SCTP", func(b []byte) string {
		return strconv.Itoa(int(binary.BigEndian.Uint16(b)))
	}})
	EndpointRUDPPort = gopacket.RegisterEndpointType(7, gopacket.EndpointTypeMetadata{"RUDP", func(b []byte) string {
		return strconv.Itoa(int(b[0]))
	}})
	EndpointUDPLitePort = gopacket.RegisterEndpointType(8, gopacket.EndpointTypeMetadata{"UDPLite", func(b []byte) string {
		return strconv.Itoa(int(binary.BigEndian.Uint16(b)))
	}})
	EndpointPPP = gopacket.RegisterEndpointType(9, gopacket.EndpointTypeMetadata{"PPP", func([]byte) string {
		return "point"
	}})
)

// NewIPEndpoint creates a new IP (v4 or v6) endpoint from a net.IP address.
// It returns gopacket.InvalidEndpoint if the IP address is invalid.
func NewIPEndpoint(a net.IP) gopacket.Endpoint {
	switch len(a) {
	case 4:
		return gopacket.NewEndpoint(EndpointIPv4, []byte(a))
	case 16:
		return gopacket.NewEndpoint(EndpointIPv6, []byte(a))
	}
	return gopacket.InvalidEndpoint
}

// NewMACEndpoint returns a new MAC address endpoint.
func NewMACEndpoint(a net.HardwareAddr) gopacket.Endpoint {
	return gopacket.NewEndpoint(EndpointMAC, []byte(a))
}
func newPortEndpoint(t gopacket.EndpointType, p uint16) gopacket.Endpoint {
	return gopacket.NewEndpoint(t, []byte{byte(p >> 8), byte(p)})
}

// NewTCPPortEndpoint returns an endpoint based on a TCP port.
func NewTCPPortEndpoint(p TCPPort) gopacket.Endpoint {
	return newPortEndpoint(EndpointTCPPort, uint16(p))
}

// NewUDPPortEndpoint returns an endpoint based on a UDP port.
func NewUDPPortEndpoint(p UDPPort) gopacket.Endpoint {
	return newPortEndpoint(EndpointUDPPort, uint16(p))
}

// NewSCTPPortEndpoint returns an endpoint based on a SCTP port.
func NewSCTPPortEndpoint(p SCTPPort) gopacket.Endpoint {
	return newPortEndpoint(EndpointSCTPPort, uint16(p))
}

// NewRUDPPortEndpoint returns an endpoint based on a RUDP port.
func NewRUDPPortEndpoint(p RUDPPort) gopacket.Endpoint {
	return gopacket.NewEndpoint(EndpointRUDPPort, []byte{byte(p)})
}

// NewUDPLitePortEndpoint returns an endpoint based on a UDPLite port.
func NewUDPLitePortEndpoint(p UDPLitePort) gopacket.Endpoint {
	return newPortEndpoint(EndpointUDPLitePort, uint16(p))
}
