package consistinghash

import (
	"fmt"
	"testing"
)

func TestConsistingHash(t *testing.T) {
	ring := NewHashRing()
	ring.AddNode(NewNode("127.0.0.1", "local1"))
	ring.AddNode(NewNode("127.0.0.1", "local2"))
	fmt.Println(ring)
	ring.Set("fork", "HongShite")
	fmt.Println(ring.Get("fork"))
	ring.Del("fork")
}
