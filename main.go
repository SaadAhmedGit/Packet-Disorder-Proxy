package main

import (
	"container/heap"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"packet-shuffler/packetHeap"
)

const (
	PROXY_SERVER_HOST = "localhost"
	PROXY_SERVER_PORT = "9988"
	PROXY_SERVER_TYPE = "tcp"

	SERVER_HOST = "localhost"
	SERVER_PORT = "9989"
	SERVER_TYPE = "tcp"

	BUFFER_SIZE = 8096
)

func clientHandler(client net.Conn) {

	defer func() {
		client.Close()
		log.Printf("Connection closed with %v\n", client.RemoteAddr())
	}()

	client_buf := make([]byte, BUFFER_SIZE)

	// Connect to the server
	server, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		log.Println(err)
		return
	}

	defer server.Close()

	server_buf := make([]byte, BUFFER_SIZE)

	// Read image dimensions and forward them to the server
	n, err := client.Read(client_buf[:4])
	if err != nil {
		log.Println(err)
		return
	}

	rows := int(binary.LittleEndian.Uint32(client_buf[:n]))
	server.Write(client_buf[:n])

	n, err = client.Read(client_buf[:4])
	if err != nil {
		log.Println(err)
		return
	}

	cols := int(binary.LittleEndian.Uint32(client_buf[:n]))
	server.Write(client_buf[:n])

	log.Printf("Image size: %d x %d\n", cols, rows)

	// Create a min heap to shuffle packets
	h := &packetHeap.PacketHeap{}
	heap.Init(h)

PACKET_LOOP:
	for packet := 0; packet < rows; packet++ {
		n, err := client.Read(client_buf)

		// Send ack msg to client
		client.Write([]byte{1})

		// First 4 bytes of the packet is the intended packet id and this should be sent to the server.
		packet_id := int(binary.LittleEndian.Uint32(client_buf[:4]))

		// Create a deep copy to insert in the heap
		data := make([]byte, n)
		copy(data, client_buf[:n])

		switch err {
		case io.EOF:
			break PACKET_LOOP
		case nil:
			log.Printf("Received packet %d\n from %s\n", packet_id, client.RemoteAddr())

			// Push packet to the heap with a random priority
			heap.Push(h, packetHeap.Packet{Priority: rand.Int(), Content: data})

			// Send packet with the lowest priority to the server after every four packets received
			if h.Len() > 0 && (packet%4 == 0) {
				packet := heap.Pop(h).(packetHeap.Packet)
				server.Write(packet.Content)
				log.Printf("Forwarded %d bytes to server from %s\n", len(packet.Content), client.RemoteAddr())

			}
		default:
			log.Printf("Failed to receive the packet:%s", err)
			return
		}
	}

	//Empty the remaining packets
	for h.Len() > 0 {
		packet := heap.Pop(h).(packetHeap.Packet)
		server.Write(packet.Content)
		log.Printf("Forwarded %d bytes to server from %s\n", len(packet.Content), client.RemoteAddr())
	}

	//Forward server response back to the client
	n, err = server.Read(server_buf)
	if err != nil {
		log.Println(err)
	}
	client.Write(server_buf)
}

func main() {
	fmt.Printf("Listening on port %s...\n", PROXY_SERVER_PORT)
	proxy_server, err := net.Listen(PROXY_SERVER_TYPE, PROXY_SERVER_HOST+":"+PROXY_SERVER_PORT)
	if err != nil {
		log.Fatalf("Error listening on port: %s\n%s", PROXY_SERVER_PORT, err.Error())
	}
	defer proxy_server.Close()

	for {
		conn, err := proxy_server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go clientHandler(conn)
	}

}
