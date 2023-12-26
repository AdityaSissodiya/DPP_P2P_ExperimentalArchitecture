package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"os"
	"time"
)

const (
	serverAddr = "127.0.0.1:12345"
	filePath   = "example.txt"
)

func main() {
	go startUDPServer()

	// Wait for the server to start
	<-time.After(1 * time.Second)

	sendFileOverUDP(filePath, serverAddr)
}

func startUDPServer() {
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server listening on", serverAddr)

	for {
		receiveFile(conn)
	}
}

func sendFileOverUDP(filePath, serverAddr string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Calculate CRC checksum for the file
	crc := crc32.NewIEEE()
	_, err = io.Copy(crc, file)
	if err != nil {
		fmt.Println("Error calculating CRC:", err)
		return
	}
	file.Seek(0, io.SeekStart) // Reset file pointer to the beginning

	// Send the file data along with the CRC checksum
	sendBuffer := make([]byte, 1024)
	binary.BigEndian.PutUint32(sendBuffer, crc.Sum32())

	for {
		n, err := file.Read(sendBuffer[4:])
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading from file:", err)
			return
		}

		_, err = conn.Write(sendBuffer[:n+4])
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
	}
}

func receiveFile(conn *net.UDPConn) {
	// Receive data along with CRC checksum
	receiveBuffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(receiveBuffer)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}

	// Extract CRC checksum from the received data
	receivedCRC := binary.BigEndian.Uint32(receiveBuffer[:4])
	receivedData := receiveBuffer[4:n]

	// Calculate CRC checksum for the received data
	crc := crc32.NewIEEE()
	_, err = crc.Write(receivedData)
	if err != nil {
		fmt.Println("Error calculating CRC:", err)
		return
	}

	// Compare the calculated CRC with the received CRC to verify integrity
	if crc.Sum32() != receivedCRC {
		fmt.Println("Checksum mismatch: Data may be corrupted.")
		return
	}
}
