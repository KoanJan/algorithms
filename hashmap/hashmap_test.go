package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {

	// test
	m := NewHashMap()
	m.Put("Tony", "Hong")
	m.Put("Xie", "HIGH")
	if v, ok := m.Get("Tony"); ok {
		fmt.Printf("m[\"Tony\"] = %v\n", v)
	} else {
		fmt.Println("FOFF! Tony Hong was DEAD!!!")
	}
	if v, ok := m.Get("Xie"); ok {
		fmt.Printf("m[\"Xie\"] = %v\n", v)
	} else {
		fmt.Println("FOFF! Xie HIGH was DEAD!!!")
	}
}
