package kademlia

import (
	"net"
	"fmt"
	"bufio"
	//"log"
	//"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"log"
	"strings"
)

type Network struct {
	myContact *Contact
	myRoutingTable *RoutingTable
}


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_,err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func (network *Network) Listen() {
	// TODO
	//socket listening different events
	// PingMessage
	// FindContactMessage
	// FindDataMessage
	// StoreMessage
	ipAndPort := strings.Split(network.myContact.Address, ":")
	port := int(ipAndPort[1])

	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: port,
		IP: net.ParseIP(ipAndPort[0]),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error A %v\n", err)
		return
	}
	for {
		//data,remoteaddr,err := ser.ReadFromUDP(p)
		_,remoteaddr,err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err !=  nil {
			fmt.Printf("Some error B  %v", err)
			continue
		}
		//go unmarshallData(data)
		//go sendResponse(ser, remoteaddr)
	}

	//unserialize
}

func unmarshallData(data int) {
	//newTest := &ProtocolPackage{}
	//err = proto.Unmarshal(data, newTest)
	//if err != nil {
	//	log.Fatal("unmarshaling error: ", err)
	//}
}



func (network *Network) Sender (marshaledObject []byte, address string) (*ProtocolPackage){

	p :=  make([]byte, 2048)
	conn, err := net.DialTimeout("udp", address, 100)
	// //net.Dial("udp", "127.0.0.1:8080")

	if err != nil {
		fmt.Printf("Some error %v", err)
		return nil
	}
	fmt.Fprintf(conn, string(marshaledObject))
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)

		newTest := &ProtocolPackage{}
		err = proto.Unmarshal(p, newTest)

		//new contact and add it to bucket
		newContact := &Contact{
			ID: newTest.FindID,
			Address: address,
		}
		newContact.CalcDistance(network.myContact)
		network.myRoutingTable.AddContact(newContact)

		switch newTest.GetMessageSent() {
		case ProtocolPackage_PING:
			fmt.Printf("Ping")
			break;
		case ProtocolPackage_STORE:
			fmt.Printf("store")
			break;
		case ProtocolPackage_FINDNODE:
			fmt.Printf("find node")
			break;
		case ProtocolPackage_FINDVALUE:
			fmt.Printf("find value")
			break;

		}


		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		log.Printf("Unmarshalled to: %+v", newTest)
		conn.Close()
		return newTest
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()


	/*testingRep := make([]*ProtocolPackage_ContactInfo, 0)
	//testingRep := [3]kademlia.ProtocolPackage_ContactInfo{}
	t := ProtocolPackage_ContactInfo{
		ContactID: []byte("Client!!!!!"),
		Address: proto.String("localhost"),
		Distance: []byte("FAR!"),
	}
*/
	//testingRep = append(testingRep, &t)
	return nil
}

func proccessReceivedMessage () {

}

func proccessPong(){

}

func processFindConctactMessage()  {

}

func SendFindDataMessage()  {

}

func SendStoreMessage()  {

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

func (network *Network) marshalFindContact(findThisID *KademliaID, contacts *Contact) (*ProtocolPackage){
	typeOfMessage := ProtocolPackage_FINDNODE

	marshalPackage := &ProtocolPackage{
		ClientID: network.myContact.ID,
		Address: proto.String(network.myContact.Address),
		MessageSent: &typeOfMessage,
		FindID: &findThisID,

	}

	data, err := proto.Marshal(marshalPackage)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	return network.Sender(data, contacts.Address)
}

func (network *Network) SendFindContactMessage(contact *Contact, findThisID *KademliaID) (*ContactCandidates){
	// TODO
	// Serialize
	//
	result := network.marshalFindContact(findThisID, contact)
	var contactsReceived ContactCandidates = ContactCandidates{make([]Contact, 0, len(result.ContactsKNearest))}

	for i:=0 ; i < len(result.ContactsKNearest) ; i++{
		///var kademliaId *KademliaID = result.ContactsKNearest[i].ContactID
		newContact := Contact{ID: NewKademliaIDFromBytes(result.ContactsKNearest[i].ContactID), Address: *result.ContactsKNearest[i].Address}
		contactsReceived.contacts = append(contactsReceived.contacts, newContact)
	}
	return &contactsReceived
}


func (network *Network) SendFindDataMessage(hash string) {
	// TODO
	// Serialize

}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
	// Serialize
}
