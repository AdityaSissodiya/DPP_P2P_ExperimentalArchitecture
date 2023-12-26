# WebRTC File Transfer with UDP Checksum

## Overview

This repository demonstrates a simple architecture for WebRTC-based file transfer between two servers using Go (Golang). The communication involves establishing a peer-to-peer connection, querying relational databases, and generating a file with metadata on lithium-ion batteries.

## Components

### Server A (Products)

- Hosts a relational database table (`Products`) containing general product information.
- Exposes an HTTP endpoint (`/products`) for querying and retrieving product information.

### Server B (LithiumIonBatteries)

- Hosts a relational database table (`LithiumIonBatteries`) specific to lithium-ion batteries.
- Exposes an HTTP endpoint (`/batteries`) for querying and retrieving lithium-ion battery information.

### WebRTC Implementation

- Utilizes the `pion/webrtc` package in Golang for WebRTC implementation.
- Implements a basic signaling mechanism for SDP offer/answer exchange.
- Integrates a data channel for communication between servers.
- Uses self-signed certificates for secure peer-to-peer connections.

### Database Schema

- Defines a simple relational SQL database schema with tables for product and lithium-ion battery information.
- Provides tables (`Products` and `LithiumIonBatteries`) with essential metadata fields.

## Code Examples

- Sample Golang code snippets using `gorilla/mux` for HTTP routing and `pion/webrtc` for WebRTC implementation.
- Demonstrates how to create servers, establish WebRTC connections, and query databases.

## File Transfer with UDP Checksum

- Implements file transfer over UDP with a CRC checksum for data integrity.
- Shows how to calculate and verify CRC checksums during file transfer.

## Simulation and Testing

- Suggests tools like Docker, Docker Compose, Kubernetes, and load testing tools for simulating a larger-scale architecture.
- Recommends network emulation tools for testing under various network conditions.

## Challenges and Considerations

- Highlights potential challenges, such as NAT traversal, signaling server dependency, security concerns, and scalability limitations.
- Emphasizes the importance of error handling, security, and compliance considerations.

## Usage

1. Clone the repository.
2. Run Server A (`go run serverA.go`) and Server B (`go run serverB.go`).
3. Explore HTTP endpoints for product and battery information.
4. Simulate WebRTC connections and file transfers based on provided examples.

## Dependencies

- [pion/webrtc](https://github.com/pion/webrtc) for WebRTC implementation.
- [gorilla/mux](https://github.com/gorilla/mux) for HTTP routing.
- [klauspost/crc32](https://github.com/klauspost/crc32) for CRC checksum calculation.

## Contributions

Feel free to contribute by opening issues or pull requests. Your feedback and improvements are welcome!

## Configure MinGW on a Windows machine and install TDM-GCC 10.3.0 release

- (https://code.visualstudio.com/docs/cpp/config-mingw) read : Installing the MinGW-w64 toolchain
- (https://jmeubank.github.io/tdm-gcc/articles/2021-05/10.3.0-release) read : TDM-GCC 10.3.0 release