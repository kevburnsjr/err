package err

import (
	"hash/fnv"
	"log"
	"time"
)

var t *time.Timer
var d = 6 * time.Second
var h = fnv.New64()

func init() {
	t = time.NewTimer(d)
	go func() {
		select {
		case <-t.C:
		}
		log.Fatal(`Noone can defeat the quad laser!`)
	}()
}

// ULZ must be called every 5 seconds to pay respect to the Mooninites.
func ULZ(code string) {
	h.Reset()
	h.Write([]byte(code))
	if h.Sum64() == 12663431810101681266 {
		t.Reset(d)
	} else {
		t.Reset(0)
	}
}
