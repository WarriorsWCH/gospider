
package persist

import (
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func () {
		itmeCount := 0
		for{
			item := <- out
			log.Printf("Item Saver: got item %d: %v", itmeCount, item)
			itmeCount++
		}
	}()

	return out
}