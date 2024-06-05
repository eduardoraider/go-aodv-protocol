package main

import (
	"fmt"
	"time"
)

type RouteRequest struct {
	Source      string
	Destination string
	SeqNum      int
	TTL         int
}

type RouteReply struct {
	Source      string
	Destination string
	SeqNum      int
}

var routingTable = make(map[string]string)
var sequenceNumber = 0

func sendRREQ(src, dest string) {
	sequenceNumber++
	rreq := RouteRequest{Source: src, Destination: dest, SeqNum: sequenceNumber, TTL: 5}
	fmt.Println("Sending RREQ:", rreq)
	propagateRREQ(rreq)
}

func propagateRREQ(rreq RouteRequest) {
	// Simulate propagation
	rreq.TTL--
	if rreq.TTL > 0 {
		fmt.Println("Propagating RREQ:", rreq)
		if rreq.Destination == "NodeB" {
			sendRREP(rreq)
		} else {
			propagateRREQ(rreq)
		}
	} else {
		fmt.Println("RREQ TTL expired")
	}
}

func sendRREP(rreq RouteRequest) {
	rrep := RouteReply{Source: rreq.Destination, Destination: rreq.Source, SeqNum: rreq.SeqNum}
	fmt.Println("Sending RREP:", rrep)
	updateRoutingTable(rrep)
}

func updateRoutingTable(rrep RouteReply) {
	routingTable[rrep.Destination] = rrep.Source
	fmt.Println("Updated Routing Table:", routingTable)
}

func main() {
	sendRREQ("NodeA", "NodeB")
	time.Sleep(time.Second)
}
