# AODV Implementation in Golang

## Overview

This is a simplified implementation of the Ad hoc On-Demand Distance Vector (AODV) routing protocol in Golang. The AODV protocol is a reactive routing protocol that establishes routes only when needed.

## Features

- Sends Route Request (RREQ) messages.
- Propagates RREQ messages through intermediate nodes.
- Sends Route Reply (RREP) messages.
- Updates routing tables based on received RREP messages.

## Usage

### Prerequisites

- Golang installed on your machine.

### Running the Code

1. Clone the repository or copy the code to your local machine.
2. Save the code in a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the code using the following command:

```sh
go run main.go
```

### Code Explanation

#### `RouteRequest` Struct

Defines the structure for Route Request (RREQ) messages.

- `Source`: The source node of the request.
- `Destination`: The destination node of the request.
- `SeqNum`: Sequence number to uniquely identify the request.
- `TTL`: Time to live for the request, to limit the number of hops.

#### `RouteReply` Struct

Defines the structure for Route Reply (RREP) messages.

- `Source`: The source node of the reply (which is the destination of the original request).
- `Destination`: The destination node of the reply (which is the source of the original request).
- `SeqNum`: Sequence number to uniquely identify the reply.

#### Functions

- `sendRREQ(src, dest string)`: Initiates the sending of a Route Request.
- `propagateRREQ(rreq RouteRequest)`: Simulates the propagation of the RREQ message through intermediate nodes.
- `sendRREP(rreq RouteRequest)`: Sends a Route Reply in response to a RREQ message.
- `updateRoutingTable(rrep RouteReply)`: Updates the routing table based on the received RREP message.

## Example Output

When you run the code, you will see output indicating the sending and propagation of RREQ messages, as well as the sending of RREP messages and updates to the routing table.

```sh
Sending RREQ: {Source: NodeA, Destination: NodeB, SeqNum: 1, TTL: 5}
Propagating RREQ: {Source: NodeA, Destination: NodeB, SeqNum: 1, TTL: 4}
Sending RREP: {Source: NodeB, Destination: NodeA, SeqNum: 1}
Updated Routing Table: map[NodeA:NodeB]
```

## Notes

This is a simplified implementation and does not handle many practical aspects of AODV, such as collision handling, neighbor table maintenance, or route error messages. It is intended to understand the basic operation of the AODV protocol.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

**Eduardo Raider** - Software Engineer