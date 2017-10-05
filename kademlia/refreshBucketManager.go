package kademlia

import "time"
import (
	"math/rand"
	"fmt"
)




func (kademlia *Kademlia) StartRefreshManaging(){
	var buck bucket

	//Init date for all buckets
	for i:=0 ; i < IDLength * 8 ; i++ {
		buck = *kademlia.network.myRoutingTable.buckets[i]
		buck.lastTimeVisited = time.Now().Local()
	}

	for{
		time.Sleep(1 * time.Minute)
		for i:=0 ; i < IDLength * 8 ; i++{
			buck = *kademlia.network.myRoutingTable.buckets[i]
			duration := time.Since(buck.lastTimeVisited)

			if duration.Hours() > 1 && buck.list.Len() > 0{
				fmt.Println("\nProcess refresh of one bucket\n")
				random := rand.Intn(buck.list.Len())
				for e := buck.list.Front() ; e != nil ; e = e.Next() {
					if random == 0{
						kademlia.LookupContact(e.Value.(Contact).ID)
						break
					}
					random--
				}
			}
		}

	}
}
