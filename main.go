package main

import (
	"./kademlia"
	//"time"
	/*"net"
	"fmt"
	"bufio"
	"github.com/golang/protobuf/proto"
	"log"
	"time"*/
	//"time"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	//kademlia.Listen("127.0.0.1", 8080);


	/*
	typeOfMessage := kademlia.ProtocolPackage_PING

	testingRep := make([]*kademlia.ProtocolPackage_ContactInfo, 0)
	//testingRep := [3]kademlia.ProtocolPackage_ContactInfo{}
	t := kademlia.ProtocolPackage_ContactInfo{
		ContactID: []byte("Client!!!!!"),
		Address: proto.String("localhost"),
		Distance: []byte("FAR!"),
	}

	testingRep = append(testingRep, &t)

	tx := kademlia.ProtocolPackage_ContactInfo{
		ContactID: []byte("Client11"),
		Address: proto.String("localhost1"),
		Distance: []byte("FAR1"),
	}
	testingRep = append(testingRep, &tx)
	testProto := &kademlia.ProtocolPackage{
		ClientID: []byte("String"),

		Ip: proto.String("localhost"),
		ListenPort: proto.Int32(1234),
		MessageSent: &typeOfMessage,
		ContactsKNearest: testingRep,

		//ContactsKNearest: testingRep,
		//ContactsKNearest: []byte("Client:123 ip:111"),
		//ContactsKNearest: []byte("Client:1234 ip:222"),
		//Then we can add both findID and findValue, they are optional
	}


	data, err := proto.Marshal(testProto)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}


	newTest := &kademlia.ProtocolPackage{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if testProto.GetIp() != newTest.GetIp() {
		log.Fatalf("data mismatch %q != %q", testProto.GetIp(), newTest.GetIp())
	}

	log.Printf("Unmarshalled to: %+v", newTest)

	for i := 0; i < len(newTest.ContactsKNearest); i++ {
		log.Print("ClientID: ", string(newTest.ContactsKNearest[i].ContactID[:]))
		log.Print("Address: ", *newTest.ContactsKNearest[i].Address)
		log.Print("Distance: ", string(newTest.ContactsKNearest[i].Distance[:]))

	}
	*/
	/*
	newNodes := kademlia.CreateRandomNetworks(100)

	kademlia.MakeMoreFriends(newNodes, 5)
	//newNodes[0].TestKademliaPing()
	newNodes[0].PrintNetwork()
	for i,_ := range newNodes {
		go newNodes[i].Listen()
	}*/



	/*for i,_ := range newNodes {
		go newNodes[i].Listen()
	}
	*/

	//go newNodes[1].TestKademliaPing(newNodes[0].GetMyContact())
	//go newNodes[1].TestKademliaPing(newNodes[2].GetMyContact())

	/*contacts := newNodes[33].GetMyRoutingTable().FindClosestContacts(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), 6)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
	fmt.Println("*****************")
	//go newNodes[33].TestKademliaPing(newNodes[99].GetMyContact())


	newNodes[33].SendFindContactMessage(newNodes[20].GetMyContact(),kademlia.NewKademliaID("2111111400000000000000000000000000000000"))

	contactsB := newNodes[33].GetMyRoutingTable().FindClosestContacts(kademlia.NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contactsB {
		fmt.Println(contactsB[i].String())
	}
	fmt.Println("*****************")

	for{
		time.Sleep(20 * time.Second)
		fmt.Println("hello")
	}
*/

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("hello, I am a Kademlia node.")
	printHelp()

	for scanner.Scan() {
		processText(scanner.Text())

	}

	/*for i, node := range newNodes {
		fmt.Println(i)
		node.PrintNetwork()
		node.Listen()
		//go runNode(&node)
	}*/

}

func processText(text string){
	//fmt.Println(text)
	//fmt.Println("hello")
	var words []string = strings.Split(text," ");

	var command = words[0]
	switch command {
	case "ping":
		processCommandPing(words)
		break
	case "info":
		processCommandInfo(words)
		break
	case "routingTable":
		processCommandRoutingTable(words)
		break
	case "help":
		break
	default:
		fmt.Println("command not found")
	}
	printHelp()
}

func processCommandPing(words []string){
	if  len(words) != 3 ||
		(words[1] != "--nodeID" && words[1] != "--nodeIP") {
		fmt.Println("error PING")
		return
	}
	fmt.Println("sending ping to ", words[2])
	fmt.Println("***********")
}

func processCommandInfo(words []string)  {
	if  len(words) != 1 || words[0] != "info" {
		fmt.Println("error INFO")
		return
	}
	fmt.Println("my info")
	fmt.Println("***********")
}

func processCommandRoutingTable(words []string)  {
	if  len(words) != 1 || words[0] != "routingTable" {
		fmt.Println("error routingTable")
		return
	}
	fmt.Println("my routing table")
	fmt.Println("***********")
}

func printHelp(){
	fmt.Println("*****************************************")
	fmt.Println("This is my help: (write option + enter)")
	fmt.Println("")
	fmt.Println("PING make a ping to the selected node")
	fmt.Println("$ ping --nodeID KademliaID")
	fmt.Println("$ ping --nodeIP KademliaID")
	fmt.Println("")
	fmt.Println("Return info about the current node")
	fmt.Println("$ info")
	fmt.Println("")
	fmt.Println("Return the routing table")
	fmt.Println("$ routingTable")
	fmt.Println("*****************************************")
}
/*func createRandomNode(){}

func SendPingMessageFake () {
	// TODO
	p :=  make([]byte, 2048)
	conn, err := net.DialTimeout("udp", "127.0.0.1:1234", 100)
	// //net.Dial("udp", "127.0.0.1:8080")

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

}*/