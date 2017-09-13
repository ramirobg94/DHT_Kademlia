package kademlia

import (
	"net"
	"fmt"
	"bufio"
)

type Network struct {
}


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_,err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func Listen(ip string, port int) {
	// TODO
	//socket listening different events
	// PingMessage
	// FindContactMessage
	// FindDataMessage
	// StoreMessage

	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: port,
		IP: net.ParseIP(ip),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error A %v\n", err)
		return
	}
	for {
		_,remoteaddr,err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err !=  nil {
			fmt.Printf("Some error B  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}

	//unserialize
}

func (network *Network) SendPingMessage(contact *Contact) {
	// TODO

	p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()

	// Serialize
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
	// Serialize
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
	// Serialize
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
	// Serialize
}