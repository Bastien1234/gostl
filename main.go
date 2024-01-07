package main

import (
	"fmt"
	"gostl/hashmap"
	linkedlist "gostl/linked_list"
)

func main() {
	fmt.Println("coucou")

	fmt.Printf("Hash of Bastien : %d\n", hashmap.Hash("Bastien"))
	fmt.Printf("Hash of Misao : %d\n", hashmap.Hash("Misao"))
	fmt.Printf("Hash of Emma : %d\n", hashmap.Hash("Emma"))
	fmt.Printf("Hash of Luna : %d\n", hashmap.Hash("Bastien"))

	hm := hashmap.NewHashMap[string, string](50)

	fmt.Println(hm.Values[0])

	ll := linkedlist.NewLinkedList[string]()

	ll.Add("coucou")
	el, tr := ll.Get("coucou")
	ellll, fls := ll.Get("faux")

	fmt.Printf("true : %t false : %t\n", tr, fls)
	fmt.Println(el)
	fmt.Println(ellll)

}
