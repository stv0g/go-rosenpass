// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

//go:build dragonfly || openbsd

package rosenpass

import (
	"fmt"
	"net"
)

// Open creates two UDP sockets, one for IPv4 and one for IPv6
//
// On DragonFly BSD and OpenBSD, listening on the
// "tcp" and "udp" networks does not listen for both IPv4 and IPv6
// connections. This is due to the fact that IPv4 traffic will not be
// routed to an IPv6 socket - two separate sockets are required if
// both address families are to be supported.
// See inet6(4) for details.
func (s *udpConn) Open() ([]receiveFunc, error) {
	networks := map[string]*net.UDPAddr{}

	for _, la := range s.listenAddrs {
		if network := networkFromAddr(la); network == "udp" {
			networks["udp4"] = la
			networks["udp6"] = la
		} else {
			networks[network] = la
		}
	}

	return s.open(networks)
}
