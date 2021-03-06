package kademlia

import "time"
import (
	"math/rand"
	"fmt"
)


const TIME_TO_CHECK_LAST_VISITED = 1

/** StartRefreshManaging
*	Check to refresh the buckets if needed each minute
 */
func (kademlia *Kademlia) StartRefreshManaging(){
	var buck *bucket

	//Init date for all buckets

	for i:=0 ; i < IDLength * 8 ; i++ {
		buck = kademlia.network.myRoutingTable.buckets[i]
		buck.lastTimeVisited = time.Now().Local()
	}

	for{
		time.Sleep(70 * time.Second)

		fmt.Println("Check for refresh buckets")
		for i:=0 ; i < IDLength * 8 ; i++{
			buck = kademlia.network.myRoutingTable.buckets[i]

			if time.Since(buck.lastTimeVisited).Hours() > 1 && buck.list.Len() > 0{
				random := rand.Intn(buck.list.Len())
				for e := buck.list.Front() ; e != nil ; e = e.Next() {
					if random == 0{
						fmt.Println("\nProcess refresh of bucket \n", 160- i)
						kademlia.LookupContact(e.Value.(Contact).ID)
						break
					}
					random--
				}
			}
		}

	}
}
