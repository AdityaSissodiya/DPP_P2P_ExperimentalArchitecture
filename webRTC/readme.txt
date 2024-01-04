1) We create a new WebRTC API and a peer connection.
2) We create a data channel named "demo" and set up event handlers for when the channel opens and receives messages.
3) We create an offer and print it in SDP format. In a real application, you would send this offer to the remote peer.
4) We simulate receiving the offer on the remote side, generate an answer, and print it in SDP format.
5) We simulate receiving the answer back on the original side.
6) We close the peer connection.