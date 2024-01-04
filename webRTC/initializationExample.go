package main

import (
	"fmt"
	"time"

	"github.com/pion/webrtc/v3"
)

func main() {
	// Create a new WebRTC API
	api := webrtc.NewAPI(webrtc.WithMediaEngine(webrtc.MediaEngine{}))

	// Create a peer configuration
	config := webrtc.Configuration{}

	// Create a new peer connection with the configuration
	peerConnection, err := api.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// Create a data channel with the label "demo"
	dataChannel, err := peerConnection.CreateDataChannel("demo", nil)
	if err != nil {
		panic(err)
	}

	// Set up event handlers for the data channel
	dataChannel.OnOpen(func() {
		fmt.Println("Data Channel is open!")
		dataChannel.SendText("Hello, WebRTC!")
	})

	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Received message: %s\n", msg.Data)
	})

	// Create an offer to initiate the WebRTC connection
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	// Set the local description
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		panic(err)
	}

	// Print the generated offer in SDP format
	fmt.Println("Generated SDP offer:")
	fmt.Println(offer.SDP)

	// In a real application, you would send the offer to the remote peer
	// and receive the remote peer's answer. For simplicity, we'll simulate
	// this by waiting for a few seconds before generating an answer locally.

	time.Sleep(3 * time.Second)

	// Simulate receiving the offer and generating an answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}

	// Set the local description
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	// Print the generated answer in SDP format
	fmt.Println("Generated SDP answer:")
	fmt.Println(answer.SDP)

	// In a real application, you would send the answer back to the original peer.
	// For simplicity, we'll simulate this by waiting for a few more seconds
	// before closing the connection.

	time.Sleep(3 * time.Second)

	// Close the peer connection
	err = peerConnection.Close()
	if err != nil {
		panic(err)
	}
}
